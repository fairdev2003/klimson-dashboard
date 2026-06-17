<script lang="ts">
	import type { Component } from 'svelte';
	import type { PageProps } from './$types';
	import IpView from '$lib/components/dashboard/tools/buttons/IpView.svelte';

	let { data }: PageProps = $props();

	let searchBoxValue = $state('');

	type ToolRecord = {
		component: Component;
		search_alliases: string;
	};

	const tool_records: ToolRecord[] = [{ component: IpView, search_alliases: 'ip' }];

	const filtered_records = $derived.by(() => {
		let result: ToolRecord[] = [];
		const query = searchBoxValue.toLowerCase();
		result = tool_records.filter((e) => e.search_alliases.toLowerCase().includes(query));
		return result;
	});
</script>

<div class="flex flex-col w-2xl mx-auto gap-5 p-5">
	<input
		bind:value={searchBoxValue}
		placeholder="Search Server Api routes..."
		class="flex border-neutral-700 gap-1 rounded-xl bg-neutral-900 p-4"
	/>

	<div class="flex flex-col gap-4">
		{#each filtered_records as tool_record}
			<tool_record.component />
		{/each}
	</div>
</div>
