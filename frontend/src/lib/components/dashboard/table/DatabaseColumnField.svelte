<script lang="ts">
	import MovingTooltip from '../MovingTooltip.svelte';
	import type { DatabaseTableColumn } from './data_table.types';
	import TableTypeSpan from './TableTypeSpan.svelte';

	type Props = {
		column: DatabaseTableColumn;
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
			</div>
		{/snippet}

		<div class="flex gap-1 items-center">
			<TableTypeSpan type={column.type} />
			<p>{column.slug}</p>
		</div>
	</MovingTooltip>
</th>
