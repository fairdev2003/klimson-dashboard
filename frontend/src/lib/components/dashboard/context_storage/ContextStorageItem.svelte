<script lang="ts">
	import Icon from '@iconify/svelte';
	import DatabaseModalInput from '../table/DatabaseModalInput.svelte';
	import Modal from '$lib/components/Modal.svelte';
	import Button from '$lib/components/Button.svelte';
	import { api } from '$lib/api/api';
	import type { ContextStorageType } from '$lib/api/requests/context_storage';
	import { context_storage } from './context_storage.store';
	import Check from '../Check.svelte';

	type Props = {
		context_record: ContextStorageType;
	};

	let { context_record }: Props = $props();

	let update_context_form = $state(context_record);
	let deletionModalOpened: boolean = $state(false);

	let updateLoading: boolean = $state(false);
	let deleteLoading: boolean = $state(false);

	let updateModalOpened: boolean = $state(false);
</script>

<div
	class="relative group bg-neutral-900 border border-neutral-700 items-center transition-colors w-55 h-30"
>
	<div class="flex justify-between group-hover:justify-end flex-col h-full p-3">
		<div class="flex gap-2 items-center group-hover:hidden">
			<Icon icon={context_record.icon} width="25" height="25" />
			<p class="font-bold text-lg">{context_record.key}</p>
		</div>

		<div class="flex select-text font-semibold justify-end p-3 text-sm">
			{context_record.value}
		</div>
	</div>

	<div
		class="absolute group-hover:flex gap-1 items-center right-0 top-0 w-full hidden h-13 justify-end p-3"
	>
		<button
			onclick={() => {
				updateModalOpened = !updateModalOpened;
			}}
			class="flex justify-center size-8 items-center transition-colors cursor-pointer text-green-500 hover:text-green-600 flex-col"
		>
			<Icon icon="mdi:edit" width="25" height="25" />
		</button>
		<button
			onclick={() => {
				deletionModalOpened = !deletionModalOpened;
			}}
			class="flex text-red-500 size-8 hover:text-red-600 cursor-pointer transition-colors justify-center items-center flex-col"
		>
			<Icon icon="mdi:trash" width="25" height="25" />
		</button>
	</div>
</div>

<Modal
	title="Editing context content"
	bind:opened={updateModalOpened}
	onClose={() => {
		updateModalOpened = false;
	}}
	className="w-100"
>
	<div class="flex flex-col gap-2">
		<DatabaseModalInput label="Key" bind:value={update_context_form.key} />
		<DatabaseModalInput label="Value" bind:value={update_context_form.value} />
		<Check label="Is Public?" bind:value={update_context_form.is_public} />
		<DatabaseModalInput label="Category" bind:value={update_context_form.category_name} />
		<DatabaseModalInput label="Type" bind:value={update_context_form.type} />
		<div class="mt-3 flex justify-end">
			<Button
				loading={updateLoading}
				onclick={async () => {
					updateLoading = true;

					const update_response = await api.context_storage.UpdateContextStorage(
						context_record.key,
						update_context_form
					);

					const response = await api.context_storage.GetPrivateContextStorage();

					$context_storage = response.data;

					updateLoading = false;
					updateModalOpened = false;
				}}
				theme="secondary"
				size="small">Update Context</Button
			>
		</div>
	</div>
</Modal>

<Modal
	title="Editing context content"
	bind:opened={deletionModalOpened}
	onClose={() => {
		deletionModalOpened = false;
	}}
	className="w-80"
>
	<div class="flex flex-col gap-2">
		<p>Are u sure u want delete this context storage?</p>
		<div class="mt-3 flex justify-end">
			<div class="flex gap-3">
				<Button
					onclick={() => {
						deletionModalOpened = false;
					}}
					theme="danger"
					size="small">No</Button
				>
				<Button
					loading={deleteLoading}
					onclick={async () => {
						if (!context_record.id) return;
						deleteLoading = true;
						const delete_response = await api.context_storage.DeleteContextStorage(
							context_record.id
						);

						const response = await api.context_storage.GetPrivateContextStorage();

						deleteLoading = false;
						$context_storage = response.data;
						deletionModalOpened = false;
					}}
					theme="correct"
					size="small">Yes</Button
				>
			</div>
		</div>
	</div>
</Modal>
