<script lang="ts">
	import { goto } from '$app/navigation';
	import type { V2StorageRecord } from '$lib/api/requests/storage';
	import { storage_logic } from '$lib/dashboard/storage/storage.svelte';
	import { scale } from 'svelte/transition';

	type Props = {
		storage_record: V2StorageRecord;
		path: string;
	};

	let showInfo = $state(false);

	let { storage_record, path }: Props = $props();
</script>

<div
	onclick={() => {
		if (storage_record.file_type === 'folder') {
			console.log('Storage record: ', storage_record);

			storage_logic.selected_path = path;
			console.log('Path', storage_logic.selected_path);

			return;
		}
	}}
	class="w-40 h-30 rounded-lg cursor-pointer bg-neutral-800/60 text-white relative"
>
	{storage_record.file_name}
</div>
