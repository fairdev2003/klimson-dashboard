<script lang="ts">
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
</script>

<div class="m-5 flex flex-col gap-1">
	<div class="mb-3">
		<h1 class="font-bold text-blue-500">
			Listed tables: <span class="font-normal text-white">{tables.length}</span>
		</h1>
	</div>
	{#each tables as table, i}
		<a
			href="/dashboard/database/{table.table}"
			class="text-blue-500 items-center flex gap-1 hover:underline"
		>
			<Icon icon={table.icon} />
			<p>{table.name}</p>
		</a>
	{/each}
</div>
