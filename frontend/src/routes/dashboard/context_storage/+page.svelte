<script lang="ts">
	import { api } from '$lib/api/api';
	import type { ContextStorageType } from '$lib/api/requests/context_storage';
	import Button from '$lib/components/Button.svelte';
	import Check from '$lib/components/dashboard/Check.svelte';
	import { context_storage } from '$lib/components/dashboard/context_storage/context_storage.store';
	import ContextStorageItem from '$lib/components/dashboard/context_storage/ContextStorageItem.svelte';
	import HarcCheckBox from '$lib/components/dashboard/HarcCheckBox.svelte';
	import DatabaseModalInput from '$lib/components/dashboard/table/DatabaseModalInput.svelte';
	import FieldEditModal from '$lib/components/dashboard/table/FieldEditModal.svelte';
	import Heading from '$lib/components/dashboard/typography/Heading.svelte';
	import Modal from '$lib/components/Modal.svelte';
	import { onMount } from 'svelte';

	let opened: boolean = $state(false);

	let context_storages: ContextStorageType[] = $state([]);

	let new_context_form = $state<ContextStorageType>({
		key: '',
		value: '',
		category_name: '',
		type: '',
		icon: '',
		is_public: false
	});

	$effect(() => {
		console.log(new_context_form);
	});

	let loading: boolean = $state(false);

	onMount(async () => {
		const response = await api.context_storage.GetPrivateContextStorages();

		$context_storage = response.data;
	});
</script>

<div>
	<div class="flex justify-between m-4 p-2 rounded-xl items-center">
		<Heading>Context Storage</Heading>
		<Button
			onclick={() => {
				opened = !opened;
			}}
			theme="secondary"
			size="medium">New Context</Button
		>
	</div>

	<div class="m-5 flex flex-wrap gap-5">
		{#each $context_storage as context_record}
			<ContextStorageItem {context_record} />
		{/each}
	</div>
</div>

<Modal
	title="Editing context content"
	bind:opened
	onClose={() => {
		opened = false;
	}}
	className="w-100"
>
	<div class="flex flex-col gap-2">
		<DatabaseModalInput label="Key" bind:value={new_context_form.key} />
		<DatabaseModalInput label="Value" bind:value={new_context_form.value} />
		<HarcCheckBox label="Public" bind:checked={new_context_form.is_public} />
		<DatabaseModalInput label="Category" bind:value={new_context_form.category_name} />
		<DatabaseModalInput label="Type" bind:value={new_context_form.type} />
		<div class="mt-3 flex justify-between">
			<Button
				{loading}
				onclick={() => {
					console.log(new_context_form);
				}}
				theme="base"
				size="small">Inpsect context</Button
			>
			<Button
				{loading}
				onclick={async () => {
					loading = true;
					const create_response = await api.context_storage.CreateContextStorage(new_context_form);

					const response = await api.context_storage.GetPrivateContextStorages();

					$context_storage = response.data;

					loading = false;
					opened = false;
				}}
				theme="secondary"
				size="small">Add Context</Button
			>
		</div>
	</div>
</Modal>

<svelte:head>
	<title>Context Storage</title>
	<meta name="description" content="View databases" />
</svelte:head>
