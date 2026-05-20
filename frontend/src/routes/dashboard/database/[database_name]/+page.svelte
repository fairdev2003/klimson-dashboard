<script lang="ts">
	import DatabaseTable from '$lib/components/dashboard/table/DatabaseTable.svelte';
	import { onMount } from 'svelte';
	import type { PageProps } from './$types';
	import { api } from '$lib/api/api';
	import type { TableDataType } from '$lib/api/requests/database.table';
	import Loader from '$lib/components/dashboard/Loader.svelte';
	import { route } from '$lib/dashboard/stores/persist';

	let { data, params }: PageProps = $props();

	let table_data: TableDataType | undefined = $state();

	onMount(async () => {
		const response = await api.database.table.GetTableData(params.database_name);
		route.set('/dashboard/database');
		table_data = response.data;
		console.log(table_data);
	});
</script>

<div class="p-6">
	{#if table_data}
		<DatabaseTable
			data={table_data.data}
			column={table_data.columns}
			table_name={params.database_name}
		/>
	{:else}
		<div class="flex justify-center items-center">
			<Loader />
		</div>
	{/if}
</div>
