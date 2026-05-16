<script lang="ts">
	import DatabaseTable from '$lib/components/dashboard/table/DatabaseTable.svelte';
	import type { PageProps } from './$types';

	let { data, params }: PageProps = $props();
</script>

<div class="p-6">
	<DatabaseTable />
</div>
