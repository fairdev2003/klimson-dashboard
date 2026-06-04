<script lang="ts">
	import { goto } from '$app/navigation';
	import { api } from '$lib/api/api';
	import type { TableData } from '$lib/api/requests/misc';
	import Loader from '$lib/components/dashboard/Loader.svelte';
	import { toast } from '$lib/dashboard/stores/toast';
	import Icon from '@iconify/svelte';
	import { onMount } from 'svelte';

	let loading = $state(false);
	let index: number = $state(-1);

	let tablesLoading = $state(false);

	let tables: TableData[] = $state([]);

	onMount(async () => {
		tablesLoading = true;
		const response = await api.misc.GetTables();

		tables = response.data.tables;
		tablesLoading = response.status === 200;
	});

	function onDatabaseRecordClick(table_id: string, indexClicked: number) {
		loading = true;
		index = indexClicked;
		setTimeout(() => {
			goto(`/dashboard/database/${table_id}`);
			loading = false;
		}, 1000);
	}
</script>

<div class="m-5 flex flex-col items-center justify-center h-[calc(100vh-66px)] gap-1">
	<div
		class="bg-neutral-900 border flex flex-col gap-2 border-neutral-700 p-4 h-110 w-[95%] md:w-80 lg:w-80"
	>
		{#if tablesLoading}
			{#each tables as table, i}
				<button
					onclick={() => onDatabaseRecordClick(table.table, i)}
					class="bg-neutral-800 cursor-pointer hover:bg-neutral-700 active:bg-neutral-700 duration-150 transition-colors flex justify-between items-center gap-2 px-3 border border-neutral-700 h-14 w-full"
				>
					<div class="flex gap-2 items-center">
						<Icon icon={table.icon} width="30" height="30" />
						<div class="flex flex-col text-start">
							<p class="text-sm font-bold">{table.name}</p>
							<p class="text-xs text-neutral-400">{table.table}</p>
						</div>
					</div>
					{#if loading && index === i}
						<div class="loader pr-3">
							<Loader theme="regular" />
						</div>
					{/if}
				</button>
			{/each}
		{:else}
			<div class="flex mx-auto">
				<Loader />
			</div>
		{/if}
	</div>
</div>
