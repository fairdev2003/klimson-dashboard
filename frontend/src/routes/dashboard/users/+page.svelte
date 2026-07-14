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
	import Roles from './(components)/Roles.svelte';
	import PermissionsRegistry from './(components)/PermissionsRegistry.svelte';
	import Modal from '$lib/components/Modal.svelte';
	import InputSettingsRecord from '$lib/components/dashboard/settings/records/InputSettingsRecord.svelte';
	import Input from '$lib/components/dashboard/settings/components/Input.svelte';
	import DatabaseModalInput from '$lib/components/dashboard/table/DatabaseModalInput.svelte';
	import MultipleDropdown from '$lib/components/dashboard/settings/components/MultipleDropdown.svelte';

	let selectedLabel: LabelName = $derived(
		($page.url.searchParams.get('label') as LabelName) || 'acc'
	);

	function updateLabel(label: LabelName) {
		const newParams = new URLSearchParams($page.url.searchParams);
		newParams.set('label', label);

		goto(`?${newParams.toString()}`, { replaceState: true, keepFocus: true });
	}

	let roleModalOpened = $state(false);
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
			<Button
				theme="base"
				onclick={() => {
					roleModalOpened = true;
				}}>Implement role</Button
			>
			<Button theme="secondary">Add account</Button>
		</div>
	</div>

	<div>
		{#if selectedLabel === 'acc'}
			<div class="flex flex-col gap-2 w-2xl mx-auto">
				<div class="bg-neutral-800 justify-between flex rounded-lg p-3 px-6 items-center">
					<div class="flex items-center gap-3">
						<img
							src="https://api.klimson.dev/interface/bucket/random/pixelgunicon.png"
							class="size-10 rounded-full"
						/>
						<div class="flex flex-col">
							<span class="flex gap-0.5 items-center text-white">
								<p
									onclick={() => {
										goto(`/dashboard/redis_writable/${rdb}/info`);
									}}
									class="font-black hover:underline cursor-pointer"
								>
									Jew hunter
								</p>
							</span>
							<p class="hover:underline cursor-pointer">
								<!-- pill -->
								<span class="rounded-full px-2 text-xs p-0.5 bg-black">Nigga</span>
							</p>
						</div>
					</div>
					<div class="flex items-center gap-2">
						<button
							class="p-2 hover:bg-neutral-700/50 hover:text-blue-400 rounded-xl cursor-pointer"
						>
							<Icon icon="boxicons:edit-filled" width="20" height="20" />
						</button>
						<button
							class="p-2 hover:bg-neutral-700/50 hover:text-red-400 rounded-xl cursor-pointer"
						>
							<Icon icon="boxicons:trash-filled" width="20" height="20" />
						</button>
					</div>
				</div>
			</div>
		{/if}

		{#if selectedLabel === 'roles'}
			<Roles />
		{/if}

		{#if selectedLabel === 'perms'}
			<PermissionsRegistry />
		{/if}
	</div>
</div>

<Modal
	className="w-100 h-100"
	onClose={() => (roleModalOpened = false)}
	bind:opened={roleModalOpened}
>
	<DatabaseModalInput label="Role name" />
</Modal>

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
