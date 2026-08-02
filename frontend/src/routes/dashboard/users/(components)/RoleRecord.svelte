<script lang="ts">
	import RDBModal from '$lib/components/modal/RDBModal.svelte';
	import type { Role } from '$lib/types/user';
	import Icon from '@iconify/svelte';
	import { onMount } from 'svelte';
	import account_controller from '../helpers/access.svelte';
	import { debug } from '$lib/terminal/logic';
	import RoleMainImplemetation from './(sections)/RoleMainImplemetation.svelte';
	import RolePermSection from './(sections)/RolePermSection.svelte';
	import BindedUsers from './(sections)/BindedUsers.svelte';

	type Props = { role: Role };

	let { role }: Props = $props();

	type LabelName = 'main-page' | 'perms' | 'users';

	let deleteRoleModalOpened = $state(false);
	let updateRoleModalOpened = $state(false);
	let selectedLabel: LabelName = $state('main-page');
	function updateLabel(label: LabelName) {
		selectedLabel = label;
	}
</script>

<RDBModal
	bind:opened={deleteRoleModalOpened}
	border="borderless"
	title="Delete the role"
	size="accept_preset"
	form_config={{
		onDelete: async () => {
			if (!role.id) {
				debug.warn('Role id is missing but required to delete the record.');
				return;
			}

			const close = await account_controller.DeleteRoleAndFetchNew(role.id);
			if (close) {
				deleteRoleModalOpened = false;
			}
		}
	}}
>
	<p class="text-red-400">Do you really delete this role from existance?</p>
</RDBModal>

<RDBModal
	bind:opened={updateRoleModalOpened}
	border="borderless"
	title="Delete the role"
	size="form_preset"
	form_config={{
		onSubmit: async () => {}
	}}
>
	{#snippet stickyBar()}
		<div class="flex flex-wrap p-4 px-0 top-0 gap-2 sticky mb-4 bg-neutral-900 w-full h-10">
			<button
				onclick={() => updateLabel('main-page')}
				class:selected-label-pill={selectedLabel === 'main-page'}
				class:normal-label-pill={selectedLabel !== 'main-page'}
				class="base-label-pill"
			>
				Main Info
			</button>
			<button
				onclick={() => updateLabel('perms')}
				class:selected-label-pill={selectedLabel === 'perms'}
				class:normal-label-pill={selectedLabel !== 'perms'}
				class="base-label-pill">Permissions</button
			>
			<button
				onclick={() => updateLabel('users')}
				class:selected-label-pill={selectedLabel === 'users'}
				class:normal-label-pill={selectedLabel !== 'users'}
				class="base-label-pill">Binded Users</button
			>
		</div>
	{/snippet}
	<div class="flex flex-col gap-4">
		<div>
			{#if selectedLabel === 'main-page'}
				<RoleMainImplemetation bind:role={account_controller.role} />
			{/if}
			{#if selectedLabel === 'perms'}
				<RolePermSection />
			{/if}
			{#if selectedLabel === 'users'}
				<BindedUsers bind:role={account_controller.role} />
			{/if}
		</div>
	</div></RDBModal
>

<div class="bg-neutral-800 justify-between flex rounded-lg p-3 px-6 items-center">
	<div class="flex items-center gap-3">
		<div class="flex flex-col">
			<span class="flex gap-0.5 items-center text-white">
				<p class="font-black hover:underline cursor-pointer">
					{role.name}
				</p>
			</span>
			<p class="hover:underline cursor-pointer"></p>
		</div>
	</div>
	<div class="flex items-center gap-2">
		<button
			onclick={() => {}}
			class="p-2 hover:bg-neutral-700/50 hover:text-blue-400 rounded-xl cursor-pointer"
		>
			<Icon icon="eos-icons:role-binding" width="20" height="20" />
		</button>
		<button
			onclick={() => {
				selectedLabel = 'main-page';
				account_controller.role = role;
				updateRoleModalOpened = !updateRoleModalOpened;
			}}
			class="p-2 hover:bg-neutral-700/50 hover:text-blue-400 rounded-xl cursor-pointer"
		>
			<Icon icon="boxicons:edit-filled" width="20" height="20" />
		</button>
		<button
			onclick={() => {
				deleteRoleModalOpened = !deleteRoleModalOpened;
			}}
			class="p-2 hover:bg-neutral-700/50 hover:text-red-400 rounded-xl cursor-pointer"
		>
			<Icon icon="boxicons:trash-filled" width="20" height="20" />
		</button>
	</div>
</div>

<style>
	@import 'tailwindcss';

	.selected-label-pill {
		@apply text-blue-400 bg-blue-500/20 hover:bg-blue-500/40;
	}

	.normal-label-pill {
		@apply bg-neutral-800 text-neutral-400 hover:bg-neutral-700 cursor-pointer;
	}

	.base-label-pill {
		@apply px-4 p-2 rounded-full text-sm transition-all;
	}

	.disabled-label-pill {
		@apply cursor-not-allowed bg-neutral-800 opacity-30;
	}
</style>
