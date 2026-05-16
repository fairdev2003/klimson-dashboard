<script lang="ts">
	import { onMount } from 'svelte';
	import DatabaseColumn from './DatabaseColumn.svelte';
	import MovingTooltip from '../MovingTooltip.svelte';

	type Props = {
		row: any;
		column: DatabaseColumn;
	};

	let { row, column }: Props = $props();
	let field_id: `${string}-${string}-${string}-${string}-${string}` | undefined = $state();

	onMount(() => {
		field_id = crypto.randomUUID();
	});
</script>

<td
	onclick={() => {}}
	class="whitespace-nowrap cursor-pointer border text-neutral-300 font-medium border-neutral-700 hover:bg-neutral-700"
>
	<MovingTooltip>
		{#snippet tooltipContent()}
			<div class="flex flex-col gap-3 p-1">
				<p class="text-xs">
					{row[column.slug]}
				</p>
				<p class="text-xs">
					type: {column.type}
				</p>
			</div>
		{/snippet}
		<div class="p-3">
			{row[column.slug]}
		</div>
	</MovingTooltip>
</td>
