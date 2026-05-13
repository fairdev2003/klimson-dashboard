<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->

<script lang="ts">
	import { developerView, sidebar_open } from '$lib/dashboard/stores/persist';
	import Icon from '@iconify/svelte';
	import EditPermissionsModal from './EditPermissionsModal.svelte';
	import { Edit } from '@lucide/svelte';
	import EditContributorsDetailsModal from './EditContributorsDetailsModal.svelte';
	import RowSettings from './RowSettings.svelte';
	import EditContributorCredentialsModal from './EditContributorCredentialsModal.svelte';
	import VeryfyPasswordModal from './VeryfyPasswordModal.svelte';

	import { toast } from '$lib/dashboard/stores/toast';
	import { get } from 'svelte/store';
	import type { Contributor, OptionType } from '../types';
	import { contributors } from '$lib/dashboard/stores/data.store';
	import {
		addFormContributor,
		contextMenuOptions,
		summary_open
	} from '$lib/dashboard/stores/store';
	import { contributors_loading, roles } from '../vars';
	import DetailsField from './fields/DetailsField.svelte';
	import PermissionsField from './fields/PermissionsField.svelte';
	import LastOnlineField from './fields/LastOnlineField.svelte';
	import SettingsField from './fields/SettingsField.svelte';
	import ContributorsTableHead from './ContributorsTableHead.svelte';
	import { onMount } from 'svelte';
	import { api } from '$lib/api/api';
	import { goto } from '$app/navigation';
	import { fade } from 'svelte/transition';

	let EditPermissionsModalOpen = $state(false);
	let EditConitrbutorDetailsModalOpen = $state(false);
	let EditConitrbutorCredentialsModalOpen = $state(false);
	let VerifyContributorPasswordModalOpen = $state(false);

	let contributorList: Contributor[] = $state([]);

	let columns: string[] = ['Kontrybutor', 'Uprawnienia', 'Ostatnie logowanie', 'Ustawienia'];
	type ContributorDictType = Record<(typeof columns)[number], string>;

	let ContributorDict: ContributorDictType = {
		Kontrybutor: 'name',
		Uprawnienia: 'permissions',
		'Ostatnie logowanie': 'last_login',
		Ustawienia: 'settings'
	};

	let getRowOptions = (contributor: Contributor): OptionType[] => [
		{
			label: 'Edytuj uprawnienia',
			description: `Zarządzaj dostępem dla ${contributor.name}`,
			icon: 'mdi:shield-key-outline',
			action: () => {
				EditPermissionsModalOpen = true;
				$addFormContributor = contributor;
			},
			color: 'text-blue-400',
			visible: true
		},
		{
			label: 'Zablokuj/Odblokuj kontrybutora',
			description: 'Tymczasowo zablokuj dostęp do panelu',
			icon: 'mdi:account-lock-outline',
			action: async () => {
				if (confirm(`Czy na pewno chcesz zawiesic/odwiesic użytkownika ${contributor.name}?`)) {
					$contributors_loading = true;
					try {
						const response = await api.contributor.SwitchContributorBan({ id: contributor.id });
						if (response.status === 200) {
							toast.show('Wykonano', 'success');
						}
					} catch (error) {
						toast.show(error, 'error');
					} finally {
						const response = await api.contributor.GetContributors();
						$contributors = response.data;
						$contributors_loading = false;
					}
				}
			},
			color: 'text-amber-400',
			visible: true
		},
		{
			label: 'Zmień hasło',
			description: `Bezpiecznie zmień hasło dla ${contributor.name}`,
			icon: 'material-symbols:lock',
			action: () => {
				EditConitrbutorCredentialsModalOpen = true;
			},
			color: 'text-green-500',
			visible: true
		},
		{
			label: 'Sprawdź poprawność hasła',
			description: 'Sprawdź czy pamiętasz hasło',
			icon: 'mdi:password-check',
			action: () => {
				VerifyContributorPasswordModalOpen = true;
			},
			color: 'text-indigo-400',
			visible: true
		},
		{
			label: 'Pobierz logi',
			description: 'Pobierz raport aktywności z bazy Go',
			icon: 'icon-park-solid:log',
			action: async () => {
				toast.show(`Generowanie logów dla ${contributor.name}...`);
				const t = setTimeout(() => {
					toast.show(`Logi dla ${contributor.name} zostały pobrane.`, 'success');
				}, 2000);
			},
			color: 'text-neutral-400',
			visible: true
		},
		{
			label: 'Usuń kontrybutora',
			description: 'Nieodwracalne usunięcie z zarządu',
			icon: 'mdi:trash-can-outline',
			action: async () => {
				let toastMessage = '';

				if (confirm(`Czy na pewno chcesz usunąć użytkownika ${contributor.name}?`)) {
					try {
						const response = await api.contributor.DeleteContributor({ id: contributor.id });
						if (response.status === 200) {
							toastMessage = response.data.message;
						}
					} catch (error) {
						toast.show(error, 'error');
					} finally {
						const response = await api.contributor.GetContributors();
						if (response.status === 200) {
							toast.show(toastMessage);
							$contributors = response.data;
						}
					}
				}
			},
			color: 'text-red-400',
			visible: true
		}

		// material-symbols:lock
	];

	$effect(() => {
		contributorList = $contributors;
	});

	onMount(() => {
		$contextMenuOptions = [
			{
				contextMenuName: 'Kontrybutorzy',
				items: [
					{
						label: 'Odśwież dane o kontrybutorach',
						action: async () => {
							const response = await api.contributor.GetContributors();
							$contributors = response.data;
							if (response.status === 200) {
								toast.show('Odświeżono rekordy!', 'success');
							}
						},
						icon: 'lucide:refresh-cw',
						color: 'text-green-500'
					},
					{
						label: 'Podsumowanie (Summary)',
						action: () => sidebar_open.set(!$sidebar_open),
						icon: 'ooui:text-summary-rtl',
						color: 'text-blue-500'
					},
					{
						label: 'Kopiuj ścieżkę API',
						action: () => {
							navigator.clipboard.writeText(window.location.pathname);
							toast.show('Skopiowano!', 'success');
						},
						icon: 'lucide:copy',
						color: ''
					},
					{
						label: 'Wyczyść cache panelu',
						action: () => {
							localStorage.clear();
							location.reload();
						},
						icon: 'material-symbols:cached',
						color: 'text-red-500'
					},
					{
						label: 'Wyloguj sesję',
						action: () => {
							localStorage.setItem('token', '');
							goto('/login');
						},
						icon: 'lucide:log-out',
						color: 'text-red-500'
					}
				]
			}
		];
	});

	let sortColumn = $state<string | null>(null);
	let sortDirection = $state<'asc' | 'desc'>('asc');

	// To zastępuje Twoje dotychczasowe contributorList = $contributors
	let sortedList = $derived.by(() => {
		const data = [...$contributors]; // Kopia danych ze store

		if (!sortColumn || !ContributorDict[sortColumn]) return data;

		const key = ContributorDict[sortColumn] as keyof Contributor;

		return data.sort((a, b) => {
			let valA = a[key] ?? '';
			let valB = b[key] ?? '';

			// Obsługa typów (string vs number)
			if (typeof valA === 'string') valA = valA.toLowerCase();
			if (typeof valB === 'string') valB = valB.toLowerCase();

			if (valA < valB) return sortDirection === 'asc' ? -1 : 1;
			if (valA > valB) return sortDirection === 'asc' ? 1 : -1;
			return 0;
		});
	});

	function handleSort(column: string) {
		if (sortColumn === column) {
			// Przełączanie: asc -> desc -> brak
			if (sortDirection === 'asc') {
				sortDirection = 'desc';
			} else {
				sortColumn = null;
				sortDirection = 'asc';
			}
		} else {
			sortColumn = column;
			sortDirection = 'asc';
		}
	}
</script>

<div class="w-full rounded-xl border border-neutral-800 bg-neutral-900/50 backdrop-blur-md">
	<table class="w-full border-collapse text-left">
		<ContributorsTableHead
			oncolumnclick={(column) => {
				handleSort(column);
			}}
			{columns}
		/>

		<tbody>
			{#each sortedList as person, i}
				<tr
					class={`group relative border-b border-neutral-800/50 transition-colors hover:bg-neutral-800/30 `}
				>
					<DetailsField
						{person}
						onclick={() => {
							$addFormContributor = person;
							$addFormContributor.password = '';
							EditConitrbutorDetailsModalOpen = true;
						}}
					/>
					<PermissionsField
						{person}
						onclick={() => {
							$addFormContributor = person;
							EditPermissionsModalOpen = true;
						}}
					/>
					<LastOnlineField />
					<SettingsField
						options={getRowOptions(person)}
						{person}
						onclick={() => {
							$addFormContributor = person;
						}}
					/>
				</tr>
			{/each}
		</tbody>
	</table>
	<div class="h-5 w-full"></div>
</div>

<EditPermissionsModal
	onClose={() => (EditPermissionsModalOpen = false)}
	bind:opened={EditPermissionsModalOpen}
/>

<EditContributorsDetailsModal
	onClose={() => (EditConitrbutorDetailsModalOpen = false)}
	bind:opened={EditConitrbutorDetailsModalOpen}
/>

<EditContributorCredentialsModal
	onClose={() => (EditConitrbutorCredentialsModalOpen = false)}
	bind:opened={EditConitrbutorCredentialsModalOpen}
/>

<VeryfyPasswordModal
	onClose={() => (VerifyContributorPasswordModalOpen = false)}
	bind:opened={VerifyContributorPasswordModalOpen}
/>
