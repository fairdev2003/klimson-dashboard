<script lang="ts">
	import type { DatabaseTableColumn } from './data_table.types';
	import TableTypeSpan from './TableTypeSpan.svelte';

	type Props = {
		column: DatabaseTableColumn;
	};

	let { column }: Props = $props();
</script>

<th scope="col" class="p-3 font-semibold border border-neutral-700">
	<div class="flex gap-1">
		<TableTypeSpan type={column.type} />
		<p>{column.slug}</p>
	</div>
</th>
