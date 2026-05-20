<script lang="ts">
	import type { ColumnData, TableDataType } from '$lib/api/requests/database.table';
	import type { TableData } from '$lib/api/requests/misc';
	import MovingTooltip from '../MovingTooltip.svelte';
	import type { DatabaseTableColumn } from './data_table.types';
	import TableTypeSpan from './TableTypeSpan.svelte';

	type Props = {
		column: TableData;
	};

	let { column }: Props = $props();
</script>

<th scope="col" class="p-3 font-semibold border border-neutral-700 cursor-pointer">
	<MovingTooltip>
		{#snippet tooltipContent()}
			<div class="flex flex-col gap-3 p-1">
				<p class="text-xs">
					type: {column.type}
				</p>
				<p class="text-xs">
					{#if column.foreign_key === 'true'}
						Foreign Key Column
					{/if}
				</p>
			</div>
		{/snippet}

		<div class="flex gap-1 items-center">
			<TableTypeSpan {column} />
			<p>{column.name}</p>
		</div>
	</MovingTooltip>
</th>
