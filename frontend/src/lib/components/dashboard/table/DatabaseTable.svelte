<script lang="ts">
	import type { ColumnData } from '$lib/api/requests/database.table';
	import Icon from '@iconify/svelte';
	import type { DatabaseTableProps } from './data_table.types';
	import DatabaseColumnField from './DatabaseColumnField.svelte';
	import DatabaseField from './DatabaseField.svelte';
	import { onMount } from 'svelte';
	import { sortColumnsOrder } from './data_table.helpers';
	import type { TableData } from '$lib/api/requests/misc';
	import { highlightedFields } from './data_table.store';
	import { sortOrderList } from './table_sorter';

	type Props = {
		data: any[];
		table_name: string;
		column: TableData[];
	};

	let schema: TableData[] = $state([]);
	let { data, table_name, column }: Props = $props();

	const blacklistedKeys = ['deleted_at', 'password', 'updated_at'];

	let filteredData = $derived(
		data
			? data.filter((row) => {
					const isDeleted =
						row.deleted_at !== null && row.deleted_at !== undefined && row.deleted_at !== '';
					return !isDeleted;
				})
			: []
	);

	onMount(() => {
		$highlightedFields = [];

		schema = sortColumnsOrder(column, sortOrderList);
	});
</script>

<div class="overflow-x-auto shadow-sm flex flex-col gap-4 max-w-7xl">
	<div class="flex flex-col gap-1">
		<a class="flex gap-1 items-center text-blue-500 hover:underline" href="/dashboard/database">
			<Icon icon="lets-icons:back" />
			<p>Back to database list</p>
		</a>
		<h1 class="font-bold text-blue-500 flex justify-between">
			<p>
				Viewing Table: <span class="font-normal text-white">{table_name}</span>
			</p>
			{#if $highlightedFields.length > 0}
				<p class="text-green-500">
					Changes: {$highlightedFields.length}
				</p>
			{/if}
		</h1>
	</div>

	<div class="w-full overflow-x-auto rounded-t-lg">
		<table
			class="w-full min-w-max border-collapse text-left text-sm border border-neutral-600 text-white"
		>
			<thead class="bg-neutral-800 text-xs tracking-wider">
				<tr>
					{#each schema as column}
						<DatabaseColumnField {column} />
					{/each}
				</tr>
			</thead>
			{#if filteredData && filteredData.length}
				<tbody class="bg-neutral-900 text-white">
					{#each filteredData as row}
						<tr class="transition-colors hover:bg-neutral-800/50">
							{#each schema as column}
								<DatabaseField {column} {row} />
							{/each}
						</tr>
					{/each}
				</tbody>
			{:else}
				<div class="w-full">
					<p>No records :c</p>
				</div>
			{/if}
		</table>
	</div>
</div>
