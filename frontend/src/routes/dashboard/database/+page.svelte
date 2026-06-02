<script lang="ts">
	import { goto } from '$app/navigation';
	import { api } from '$lib/api/api';
	import type { TableData } from '$lib/api/requests/misc';
	import { toast } from '$lib/dashboard/stores/toast';
	import Icon from '@iconify/svelte';
	import { onMount } from 'svelte';

	let tables: TableData[] = $state([]);

	onMount(async () => {
		const response = await api.misc.GetTables();

		tables = response.data.tables;
	});

	function onDatabaseRecordClick(table_id: string) {
		setTimeout(() => {
			goto(`/dashboard/database/${table_id}`);
		}, 150);
	}
</script>

<div class="m-5 flex flex-col items-center justify-center h-[calc(100vh-66px)] gap-1">
	<div
		class="bg-neutral-900 border flex flex-col gap-2 border-neutral-700 p-4 h-110 w-[95%] md:w-80 lg:w-80"
	>
		{#each tables as table, i}
			<button
				onclick={() => onDatabaseRecordClick(table.table)}
				class="bg-neutral-800 cursor-pointer hover:bg-neutral-700 active:bg-neutral-700 duration-150 transition-colors flex items-center gap-2 px-3 border border-neutral-700 h-14 w-full"
			>
				<Icon icon={table.icon} width="30" height="30" />
				<div class="flex flex-col text-start">
					<p class="text-sm font-bold">{table.name}</p>
					<p class="text-xs text-neutral-400">{table.table}</p>
				</div>
			</button>
		{/each}
	</div>
</div>
