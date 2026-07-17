<script lang="ts">
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import Heading from '$lib/components/dashboard/typography/Heading.svelte';
	import Icon from '@iconify/svelte';
	import { blur, fade, slide } from 'svelte/transition';
	import Button from '$lib/components/Button.svelte';
	import type { LabelName } from './helpers/user.types';
	import account_controller from './helpers/access.svelte';
	import Roles from './(components)/Roles.svelte';
	import PermissionsRegistry from './(components)/PermissionsRegistry.svelte';
	import Modal from '$lib/components/Modal.svelte';
	import DatabaseModalInput from '$lib/components/dashboard/table/DatabaseModalInput.svelte';
	import RDBModal from '$lib/components/modal/RDBModal.svelte';
	import type { User } from '$lib/api/types';
	import { debug } from '$lib/dashboard/stores/debug';
	import { api } from '$lib/api/api';
	import UserListing from './(components)/UserListing.svelte';
	import { onMount } from 'svelte';
	import type { AxiosError } from 'axios';
	import UserPreview from './(components)/UserPreview.svelte';
	import RDBInput from '$lib/components/modal/RDBInput.svelte';
	import { terminal } from '$lib/terminal/logic';
	import { bold } from '$lib/terminal/style';

	let selectedLabel: LabelName = $derived(
		($page.url.searchParams.get('label') as LabelName) || 'acc'
	);

	let editUserModalOpened = $state(false);
	let addUserModalOpened = $state(false);
	let implementRoleModalOpened = $state(false);

	function updateLabel(label: LabelName) {
		const newParams = new URLSearchParams($page.url.searchParams);
		newParams.set('label', label);

		goto(`?${newParams.toString()}`, { replaceState: true, keepFocus: true });
	}

	let users: User[] = $state([]);
	let usersLoading: boolean = $state(true);

	let currentUser: User | undefined = $state();

	onMount(async () => {
		await FetchUsers();
	});

	async function FetchUsers() {
		try {
			const response = await api.user.List();

			if (response.status === 200) {
				users = response.data.users;
				debug.log(users);
			}
		} catch (error) {
			debug.error(error);
		} finally {
			debug.log('Listing rule has been ended');
		}
	}

	async function refetchUsers() {
		await FetchUsers();
	}

	const createEmptyUser = {
		first_name: '',
		last_name: '',
		nickname: '',
		pfp: '',
		blocked: false
	};

	let userForm: User = $state(createEmptyUser);

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
					implementRoleModalOpened = true;
				}}>Implement role</Button
			>
			<Button
				onclick={() => {
					userForm = createEmptyUser;

					addUserModalOpened = true;
				}}
				theme="secondary">Add account</Button
			>
		</div>
	</div>

	<div>
		{#if selectedLabel === 'acc'}
			<UserListing
				{users}
				{usersLoading}
				onEditButtonClick={(user) => {
					currentUser = user;

					userForm = {
						first_name: user.first_name,
						last_name: user.last_name,
						nickname: user.nickname,
						pfp: user.pfp,
						blocked: user
					};

					editUserModalOpened = true;
				}}
			/>
		{/if}

		{#if selectedLabel === 'roles'}
			<Roles />
		{/if}

		{#if selectedLabel === 'perms'}
			<PermissionsRegistry />
		{/if}
	</div>
</div>

<RDBModal
	bind:opened={implementRoleModalOpened}
	border="borderless"
	title="Adding new role"
	size="form_preset"
	form_config={{
		onLog: () => {}
	}}
>
	<RDBInput label="ROLE NAME" />
</RDBModal>

<RDBModal
	border="borderless"
	form_config={{
		onSubmit: async () => {
			try {
				const response = await api.user.Create(userForm);

				if (response.status === 200) {
					debug.log('Success');
				}
			} catch (error) {
				debug.log(error);
			} finally {
				addUserModalOpened = false;
			}
		},
		onCancel: () => {
			addUserModalOpened = false;
		},
		onLog: () => {
			debug.log(userForm);
		}
	}}
	title="Adding new account"
	size="form_preset"
	bind:opened={addUserModalOpened}
>
	<div class="flex flex-col gap-2">
		{#if userForm.nickname}
			<div in:slide={{ duration: 300 }} out:slide={{ duration: 300 }}>
				<UserPreview user={userForm} />
			</div>
		{/if}
		<div class="flex flex-col gap-1">
			<span class="text-neutral-400 uppercase justify-between items-center flex font-bold text-xs">
				<p>NICKNAME</p>
			</span>

			<input bind:value={userForm.nickname} class="rounded-lg border-0 bg-neutral-800 p-2" />
		</div>
		<div class="flex flex-col gap-1">
			<span class="text-neutral-400 uppercase justify-between items-center flex font-bold text-xs">
				<p>FIRST NAME</p>
			</span>

			<input bind:value={userForm.first_name} class="rounded-lg border-0 bg-neutral-800 p-2" />
		</div>
		<div class="flex flex-col gap-1">
			<span class="text-neutral-400 uppercase justify-between items-center flex font-bold text-xs">
				<p>LAST NAME</p>
			</span>

			<input bind:value={userForm.last_name} class="rounded-lg border-0 bg-neutral-800 p-2" />
		</div>

		<div class="flex flex-col gap-1">
			<span class="text-neutral-400 uppercase justify-between items-center flex font-bold text-xs">
				<p>PFP HREF</p>
			</span>

			<input bind:value={userForm.pfp} class="rounded-lg border-0 bg-neutral-800 p-2" />
		</div>
	</div>
</RDBModal>

<RDBModal
	border="borderless"
	bind:form={userForm}
	form_config={{
		onSubmit: async () => {
			if (!currentUser) {
				debug.warn('No user!');
				return;
			}
			try {
				const response = await api.user.Update(currentUser.id as string | number, userForm);

				if (response.status === 200) {
					debug.success('Success');
					await refetchUsers();
				}
			} catch (error) {
				debug.log(error);
			} finally {
				editUserModalOpened = false;
			}
		},
		onCancel: () => {
			editUserModalOpened = false;
		},
		onLog: () => {
			debug.log(userForm);
		}
	}}
	title={`Editing "${currentUser?.nickname}"`}
	size="form_preset"
	bind:opened={editUserModalOpened}
>
	<div class="flex flex-col gap-2">
		<UserPreview user={currentUser} />
		<span class="text-neutral-400 uppercase justify-between items-center flex font-bold text-xs">
			<p>ID: {currentUser?.id}</p>
		</span>

		<div class="flex flex-col gap-1">
			<span class="text-neutral-400 uppercase justify-between items-center flex font-bold text-xs">
				<p>NICKNAME</p>
			</span>

			<input bind:value={userForm.nickname} class="rounded-lg border-0 bg-neutral-800 p-2" />
		</div>
		<div class="flex flex-col gap-1">
			<span class="text-neutral-400 uppercase justify-between items-center flex font-bold text-xs">
				<p>FIRST NAME</p>
			</span>

			<input bind:value={userForm.first_name} class="rounded-lg border-0 bg-neutral-800 p-2" />
		</div>
		<div class="flex flex-col gap-1">
			<span class="text-neutral-400 uppercase justify-between items-center flex font-bold text-xs">
				<p>LAST NAME</p>
			</span>

			<input bind:value={userForm.last_name} class="rounded-lg border-0 bg-neutral-800 p-2" />
		</div>

		<div class="flex flex-col gap-1">
			<span class="text-neutral-400 uppercase justify-between items-center flex font-bold text-xs">
				<p>PFP HREF</p>
			</span>

			<input bind:value={userForm.pfp} class="rounded-lg border-0 bg-neutral-800 p-2" />
		</div>
	</div>
</RDBModal>

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
