<script lang="ts">
	import Dashboard from '$lib/dashboard/dashboard.svelte';
	import SidebarItem from './SidebarItem.svelte';
</script>

<div class="px-5">
	<p>Latest endpoint i have spot you in:</p>
	<div class="flex flex-col gap-2 px-5">
		{#each Dashboard.state.latestRoutes.slice(Dashboard.state.latestRoutes.length - 5, Dashboard.state.latestRoutes.length) as latestRoute}
			<SidebarItem content={latestRoute} />
		{/each}
	</div>
</div>
