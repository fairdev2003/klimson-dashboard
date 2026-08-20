<script lang="ts">
	import Dashboard from '$lib/dashboard/dashboard.svelte';
	import { type SidebarItemType } from './sidebar.types';
	import SidebarItem from './SidebarItem.svelte';
</script>

<div class="flex flex-col mt-3 overflow-auto">
	<p class="text-xs font-bold uppercase text-blue-400 px-4 pb-4">Dashboard contents</p>
	{#each Dashboard.constants.SidebarContents as content, i}
		<SidebarItem {content} />
	{/each}
</div>
