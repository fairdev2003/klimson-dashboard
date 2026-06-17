<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/state';
	import { api } from '$lib/api/api';
	import type { StorageRecord } from '$lib/api/requests/storage';
	import { storage_logic } from '$lib/dashboard/storage/storage.svelte';
	import { onMount } from 'svelte';
	import StorageRecordTile from '../../../storage/components/StorageRecordTile.svelte';
	import Button from '$lib/components/Button.svelte';
	import { toast } from '$lib/dashboard/stores/toast';

	let { params } = $props();
	let fileInput: HTMLInputElement | undefined = $state();
	let uploading: boolean = $state(false);

	$effect(async () => {
		await FetchStorageRecords();
	});

	async function FetchStorageRecords() {
		storage_logic.selected_path = params.path;

		const response = await api.storage_v2.GetStorageRecords(storage_logic.selected_path || '');

		storage_logic.storage_records = response.data;
	}

	async function UploadFile(event: Event) {
		const target = event.target as HTMLInputElement;
		const file = target.files?.[0];

		if (!file) return;

		uploading = true;
		try {
			const response = await api.storage_v2.SendImageData(file, storage_logic.selected_path || '');

			if (response.data.success) {
				toast.success('Success!');
				await FetchStorageRecords();
			}
		} catch (e) {
			console.error('Błąd uploadu:', e);
		} finally {
			uploading = false;
			target.value = '';
		}
	}
</script>

<div class="flex flex-col">
	{@render StorageHeader()}
	{storage_logic.storage_records}
</div>

{#snippet StorageHeader()}
	<p>{storage_logic.selected_path}</p>
	<input type="file" class="hidden" bind:this={fileInput} onchange={UploadFile} />
	<div
		class="p-4 bg-neutral-900 m-4 rounded-xl lg:flex-row flex flex-col gap-4 justify-between lg:items-center"
	>
		<div class="flex flex-col gap-2">
			<p>
				Listed: {(storage_logic.storage_records && storage_logic.storage_records.length) || '0'} files
			</p>
		</div>
		<div class="flex gap-3">
			<Button
				theme="base"
				onclick={async () => {
					const response = await api.storage_v2.CreateFolder('test', storage_logic.selected_path);

					console.log(response.data);
				}}>Create a folder</Button
			>
			<Button
				theme="secondary"
				onclick={() => {
					fileInput.click();
				}}>Upload a file</Button
			>
		</div>
	</div>
{/snippet}
