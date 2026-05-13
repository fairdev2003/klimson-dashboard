<script lang="ts">
	type Props = {
		columns: string[];
		oncolumnclick: (e: string) => void;
	};

	let { columns, oncolumnclick }: Props = $props();
</script>

<thead>
	<tr class="rounded-xl border-b border-neutral-800 bg-neutral-900">
		{#each columns as column, i}
			<th
				onclick={() => {
					oncolumnclick(column);
				}}
				class="p-4 text-xs font-bold tracking-widest text-neutral-500 uppercase select-none"
				>{column}</th
			>
		{/each}
	</tr>
</thead>
