<script lang="ts">
	import RDBInput from '$lib/components/modal/RDBInput.svelte';
	import RDBModal from '$lib/components/modal/RDBModal.svelte';
	import type { Role } from '$lib/types/user';
	import { onMount } from 'svelte';
	import BindedUsers from './BindedUsers.svelte';
	import RoleMainImplemetation from './RoleMainImplemetation.svelte';
	import RolePermSection from './RolePermSection.svelte';
	import { debug } from '$lib/terminal/logic';

	type Props = {
		implementRoleModalOpened?: boolean;
	};

	type LabelName = 'main-page' | 'perms' | 'users';
	let selectedLabel: LabelName = $state('main-page');
	let initialForm = {
		name: '',
		color: '',
		permissions: [],
		icon: ''
	};

	let { implementRoleModalOpened = $bindable(false) }: Props = $props();

	function updateLabel(label: LabelName) {
		selectedLabel = label;
	}

	let role: Role = $state(initialForm);

	let isExitLocked = $derived(
		role.name !== initialForm.name ||
			role.color !== initialForm.color ||
			role.icon !== initialForm.icon ||
			role.permissions.length > 0
	);

	$effect(() => {
		isExitLocked = JSON.stringify(initialForm) !== JSON.stringify(role);
	});

	onMount(() => {
		isExitLocked = false;
	});
</script>

<RDBModal
	bind:opened={implementRoleModalOpened}
	border="borderless"
	title="Adding new role"
	size="form_preset"
	form_config={{
		onLog: () => {
			debug.log('Debug:', role);
		},
		onSubmit: () => {}
	}}
	bind:backgroundExitLocked={isExitLocked}
>
	<div class="flex flex-col gap-4">
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
		<div>
			{#if selectedLabel === 'main-page'}
				<RoleMainImplemetation bind:role />
			{/if}
			{#if selectedLabel === 'perms'}
				<RolePermSection bind:role />
			{/if}
			{#if selectedLabel === 'users'}
				<BindedUsers bind:role />
			{/if}
		</div>
	</div>
</RDBModal>

<style>
	@import 'tailwindcss';

	.selected-label-pill {
		@apply text-blue-400 bg-blue-500/20 hover:bg-blue-500/40;
	}

	.normal-label-pill {
		@apply bg-neutral-800 text-neutral-400 hover:bg-neutral-700;
	}

	.base-label-pill {
		@apply select-none px-4 p-2 rounded-full text-sm transition-all cursor-pointer;
	}
</style>
