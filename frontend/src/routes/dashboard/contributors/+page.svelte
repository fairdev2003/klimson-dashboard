<script lang="ts">
	import Icon from '@iconify/svelte';

	import ContributorsTable from './component/ContributorsTable.svelte';
	import Button from '$lib/components/Button.svelte';
	import AddContributorModal from './component/AddContributorModal.svelte';
	import { contextMenuOptions, resetContributor } from '$lib/dashboard/stores/store';
	import { onMount } from 'svelte';
	import { contributors, permissionList } from '$lib/dashboard/stores/data.store';
	import { contributors_loading } from './vars';
	import { fade } from 'svelte/transition';
	import { api } from '$lib/api/api';
	import type { AxiosResponse } from 'axios';
	import type { ServerResponse } from '$lib/api/types';
	import type { Permission } from './types';

	let openedAddContributorModal: boolean = $state(false);

	onMount(() => {
		const currentOptions = [...$contextMenuOptions];

		contextMenuOptions.set([
			currentOptions[0],
			{
				label: 'Dodaj nowy rekord',
				action: () => {
					openedAddContributorModal = true;
				},
				icon: 'mingcute:user-add-fill',
				color: 'text-blue-500'
			},
			...currentOptions.slice(1)
		]);
	});

	$effect(() => {
		if ($contributors_loading) {
			document.body.style.overflow = 'hidden';
			return;
		}
		document.body.style.overflow = 'auto';
	});

	let height: number = $state(0);
	function calculateHeight(e: HTMLDivElement) {
		height = e.clientHeight;
	}

	let permissions: Permission[] = $state([]);
	onMount(async () => {
		const response = await api.contributor.GetPermissionList();
		permissions = response.data;
	});
</script>

{#if $contributors_loading}
	<div
		in:fade={{ duration: 150 }}
		out:fade={{ duration: 150 }}
		class="absolute inset-x-0 top-0 z-[140] h-screen bg-black/50"
	></div>
{/if}
<div use:calculateHeight class="relative flex h-[1000px] flex-col gap-5 p-5 text-white">
	<div class="mt-5 flex justify-between">
		<div class="flex flex-col gap-1">
			<div class="mt-5 flex items-center gap-2">
				<Icon icon="mdi:account-group" class="me-2 inline-block" height="24" width="24" />
				<h1 class="inline-block text-2xl font-bold">Kontrybutorzy</h1>
			</div>
			<p class="max-w-[600px] text-sm text-neutral-400">
				W tym miejscu możesz zarządzać współtwórcami aplikacji HarcQuiz. Nadaj im odpowiednie
				uprawnienia aby mogli edytować quizy i inne treści. Jako root aplikacji masz pełną kontrolę
				nad uprawnieniami współtwórców. Pamiętaj aby przydzielać uprawnienia ostrożnie! Kliknij na
				dana w sekcje w tabeli aby edytować szczegóły lub uprawnienia konkretnego współtwórcy.
			</p>
		</div>
		<div>
			<Button
				onclick={() => {
					resetContributor();
					openedAddContributorModal = true;
				}}
				size="medium"
				theme="secondary">Dodaj nowego kontrybutora</Button
			>
		</div>
	</div>
	<ContributorsTable />
</div>

<AddContributorModal
	onClose={() => {
		openedAddContributorModal = false;
	}}
	bind:opened={openedAddContributorModal}
/>
