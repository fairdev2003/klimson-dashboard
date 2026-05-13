<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
<!-- svelte-ignore a11y_click_events_have_key_events -->
<script lang="ts">
	import Modal from '$lib/components/Modal.svelte';
	import { addFormContributor } from '$lib/dashboard/stores/store';
	import { onMount } from 'svelte';
	import { availableRoles, colorMap, contributors_loading, roles } from '../vars';
	import RoleMultipleSelect from './RoleMultipleSelect.svelte';
	import Button from '$lib/components/Button.svelte';
	import { api } from '$lib/api/api';
	import type { Permission, Role, RoleOption } from '../types';
	import { toast } from '$lib/dashboard/stores/toast';
	import { contributors, permissionList } from '$lib/dashboard/stores/data.store';
	import Icon from '@iconify/svelte';
	import { fade } from 'svelte/transition';
	import axios from 'axios';
	import { developerView } from '$lib/dashboard/stores/persist';
	import Tooltip from '$lib/components/dashboard/Tooltip.svelte';

	type Props = {
		onClose?: () => void;
		opened?: boolean;
	};

	let loading: boolean = $state(false);
	let permissions: string = $state('');

	$effect(() => {
		permissions = $addFormContributor.permissions;
	});

	let { onClose, opened = $bindable(true) }: Props = $props();
	function toggleRole(roleId: string) {
		let currentPerms = permissions
			? permissions
					.split(',')
					.map((p) => p.trim())
					.filter((p) => p !== '')
			: [];

		if (currentPerms.includes(roleId)) {
			currentPerms = currentPerms.filter((id) => id !== roleId);
		} else {
			currentPerms = [...currentPerms, roleId];
		}

		permissions = currentPerms.join(',');
	}

	function isRoleSelected(roleId: string): boolean {
		return (permissions || '')
			.split(',')
			.map((p) => p.trim())
			.includes(roleId);
	}

	async function UpdatePermissions() {
		loading = true;
		$contributors_loading = true;
		onClose?.();
		let message = '';
		try {
			if (permissions.length === 0) {
				permissions = ' ';
			} else {
				permissions.replace('no:permissions', '');
			}
			const response = await api.contributor.UpdateContributorPermissions(
				{
					id: $addFormContributor.id
				},
				permissions
			);

			if (response.data.message) {
				message = response.data.message;
			}
		} catch (error: unknown) {
			if (axios.isAxiosError(error)) {
				const message = error.response?.data?.message || 'Błąd serwera';
				toast.error(message);
			} else {
				toast.error('Wystąpił nieoczekiwany błąd');
			}
		} finally {
			const response = await api.contributor.GetContributors();
			$contributors = response.data;
			loading = false;
			$contributors_loading = false;
			toast.success(message);
		}
	}
</script>

<Modal
	title={`Edytuj uprawnienia dla ${$addFormContributor.name}`}
	className="w-[1200px] h-[600px] pb-5"
	onClose={() => onClose?.()}
	bind:opened
>
	<div class="grid grid-cols-2 gap-4 pt-5">
		{#each $permissionList as role}
			<div class="col-span-1 h-full">
				{@render PermissionRecord(role)}
			</div>
		{/each}
		{@render PermissionRecord({
			name: 'Dodaj limit quizów',
			icon: 'hugeicons:limit-order',
			color: 'red',
			description: 'Pozwala ustawić ile quizów może wykonać kontrybutor',
			tag: 'limit:$'
		})}
	</div>
	<div class="my-5 mb-7 flex justify-end">
		<Button {loading} onclick={UpdatePermissions} size="small" theme="secondary"
			>Zapisz uprawnienia</Button
		>
	</div>
</Modal>

{#snippet PermissionRecord(role: Permission)}
	{@const selected = isRoleSelected(role.tag)}
	<Tooltip position="top" content={role.tag}>
		<button
			onclick={() => toggleRole(role.tag)}
			class="relative flex w-full items-center gap-4 rounded-xl border p-4 text-left transition-all duration-300
                {selected
				? (colorMap[role.color] ?? 'border-blue-500 bg-blue-400/10 text-blue-500') +
					' border-opacity-100  shadow-lg'
				: 'border-neutral-800  text-neutral-500 hover:border-neutral-700'}"
		>
			<div class="flex flex-col justify-center rounded-md bg-neutral-800 p-2">
				<Icon icon={role.icon} class="text-current" width="20" />
			</div>
			<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
			<div class="flex flex-1 flex-col">
				<!-- svelte-ignore a11y_click_events_have_key_events -->
				<p
					onclick={() => {
						if ($developerView) {
							navigator.clipboard.writeText(role.tag);
							toast.info('Skopiowano id uprawnienia!');
						}
					}}
					class="truncate font-bold {selected ? 'text-white' : 'text-neutral-400'}"
				>
					{$developerView ? role.tag : role.name}
				</p>
				<p class="mt-1 text-xs text-neutral-500">{role.description}</p>
			</div>

			{#if selected}
				<div transition:fade={{ duration: 200 }}>
					<Icon icon="mdi:check-circle" class="items-center text-current" width="20" />
				</div>
			{/if}
		</button>
	</Tooltip>
{/snippet}
