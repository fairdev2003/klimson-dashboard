<script lang="ts">
	import { routes } from '$lib/dashboard/stores/data.store';
</script>

<div class="flex gap-2">
	<p class="font-black">Routes Explorer</p>
	{#if $routes}
		<div class="bg-blue-500/50 hidden lg:flex p-1 rounded-md px-3 text-xs text-blue-300">
			{$routes.length} endpoints
		</div>
	{/if}
</div>
