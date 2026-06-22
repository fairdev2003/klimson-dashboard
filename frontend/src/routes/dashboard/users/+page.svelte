<script lang="ts">
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import Heading from '$lib/components/dashboard/typography/Heading.svelte';
	import Icon from '@iconify/svelte';
	import { fade } from 'svelte/transition';
	import AccountsTable from '$lib/dashboard/users/components/AccountsTable.svelte';
	import Button from '$lib/components/Button.svelte';
	import type { LabelName } from './helpers/user.types';
	import AccountController from './helpers/access.svelte';
	import account_controller from './helpers/access.svelte';

	let selectedLabel: LabelName = $derived(
		($page.url.searchParams.get('label') as LabelName) || 'acc'
	);

	function updateLabel(label: LabelName) {
		const newParams = new URLSearchParams($page.url.searchParams);
		newParams.set('label', label);

		goto(`?${newParams.toString()}`, { replaceState: true, keepFocus: true });
	}
</script>

<div in:fade={{ duration: 150 }} class="flex flex-col m-8 my-4 gap-4">
	<div class="flex justify-between items-center border-b border-neutral-700 pb-4">
		<div class="flex-col flex gap-1">
			<Heading>
				<div class="flex gap-2 items-center">
					<Icon icon="mdi:user-key" />
					<p>CMS Access</p>
				</div>
			</Heading>
			<span class="text-sm font-md text-neutral-400"
				>Control what users has permission to specific part of the dashboard.</span
			>
			<div class="flex mt-4 gap-2">
				<button
					onclick={() => updateLabel('acc')}
					class:selected-label-pill={selectedLabel === 'acc'}
					class:normal-label-pill={selectedLabel !== 'acc'}
					class="base-label-pill"
				>
					Registered accounts
				</button>
				<button
					onclick={() => updateLabel('roles')}
					class:selected-label-pill={selectedLabel === 'roles'}
					class:normal-label-pill={selectedLabel !== 'roles'}
					class="base-label-pill">Roles</button
				>
				<button
					onclick={() => updateLabel('perms')}
					class:selected-label-pill={selectedLabel === 'perms'}
					class:normal-label-pill={selectedLabel !== 'perms'}
					class="base-label-pill">Permission Registry</button
				>
			</div>
		</div>

		<div class="flex gap-4">
			<Button onclick={() => account_controller.DumpData()} theme="base">Dump data</Button>
			<Button theme="base">Implement role</Button>
			<Button theme="secondary">Add account</Button>
		</div>
	</div>

	<div>
		{#if selectedLabel === 'acc'}
			<AccountsTable />
		{/if}
	</div>
</div>

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
