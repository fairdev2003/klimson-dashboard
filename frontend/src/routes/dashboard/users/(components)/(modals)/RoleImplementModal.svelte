<script lang="ts">
	import RDBInput from '$lib/components/modal/RDBInput.svelte';
	import RDBModal from '$lib/components/modal/RDBModal.svelte';
	import type { Role } from '$lib/types/user';
	import { onMount } from 'svelte';
	import BindedUsers from '../(sections)/BindedUsers.svelte';
	import RoleMainImplemetation from '../(sections)/RoleMainImplemetation.svelte';
	import RolePermSection from '../(sections)/RolePermSection.svelte';
	import { debug } from '$lib/terminal/logic';
	import account_controller from '../../helpers/access.svelte';
	import MovingTooltip from '$lib/components/dashboard/MovingTooltip.svelte';

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
</script>

<RDBModal
	bind:opened={implementRoleModalOpened}
	border="borderless"
	title="Adding new role"
	size="form_preset"
	form_config={{
		onLog: () => {
			debug.log('Debug:', account_controller.role);
		},
		onSubmit: async () => {
			account_controller.CreateNewRole();
		}
	}}
>
	{#snippet stickyBar()}
		<div class="flex flex-wrap p-4 px-0 top-0 gap-2 sticky mb-4 bg-neutral-900 w-full h-10">
			<MovingTooltip>
				{#snippet tooltipContent()}
					<p class="text-xs max-w-40">Manage role info</p>
				{/snippet}
				<button
					onclick={() => updateLabel('main-page')}
					class:selected-label-pill={selectedLabel === 'main-page'}
					class:normal-label-pill={selectedLabel !== 'main-page'}
					class="base-label-pill"
				>
					Main Info
				</button>
			</MovingTooltip>

			<MovingTooltip>
				{#snippet tooltipContent()}
					<p class="text-xs max-w-40">Manage permissions</p>
				{/snippet}
				<button
					onclick={() => updateLabel('perms')}
					class:selected-label-pill={selectedLabel === 'perms'}
					class:normal-label-pill={selectedLabel !== 'perms'}
					class="base-label-pill">Permissions</button
				>
			</MovingTooltip>

			<MovingTooltip>
				{#snippet tooltipContent()}
					<p class="text-xs max-w-40">Section not available in create view</p>
				{/snippet}
				<button class="base-label-pill disabled-label-pill">Binded Users</button>
			</MovingTooltip>
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
		</div>
	</div>
</RDBModal>

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
