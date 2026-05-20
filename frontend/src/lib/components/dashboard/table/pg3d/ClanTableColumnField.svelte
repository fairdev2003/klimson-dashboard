<script lang="ts">
	import type { ColumnData, TableDataType } from '$lib/api/requests/database.table';
	import type { TableData } from '$lib/api/requests/misc';
	import MovingTooltip from '../../MovingTooltip.svelte';

	type Props = {
		column: string;
	};

	let { column }: Props = $props();
</script>

<th scope="col" class="p-3 font-semibold border border-neutral-700 cursor-pointer">
	<MovingTooltip>
		{#snippet tooltipContent()}
			<p class="text-xs">
				{column}
			</p>
		{/snippet}

		<div class="flex gap-1 items-center">
			<p>{column}</p>
		</div>
	</MovingTooltip>
</th>
