<script lang="ts">
	import { onMount } from 'svelte';
	import { api } from '$lib/api/api';
	import { routes } from '$lib/dashboard/stores/data.store';

	onMount(async () => {
		if ($routes) return;
		const response = await api.misc.GetRoutes();
		$routes = response.data;
		console.log('Routes', $routes);
	});
</script>

<div class="overflow-x-auto text-white">
	<div class="grid grid-cols-2 gap-5 p-5">
		{#each $routes as route, i}
			<div class="flex gap-1 border border-neutral-800/60 bg-neutral-900/60 p-2">
				{@render Method(route.method)}
				<a class="text-blue-500 hover:underline" href={api.api_config.host + route.path}
					>{route.path}</a
				>
			</div>
		{/each}
	</div>
</div>

{#snippet Method(method: string)}
	<div class="w-20">
		{#if method === 'GET'}
			<p class="text-green-500">{method}</p>
		{/if}
		{#if method === 'POST'}
			<p class="text-yellow-500">{method}</p>
		{/if}
		{#if method === 'PUT'}
			<p class="text-blue-500">{method}</p>
		{/if}
		{#if method === 'DELETE'}
			<p class="text-red-500">{method}</p>
		{/if}
	</div>
{/snippet}

<style>
	@import 'tailwindcss';
	@plugin '@tailwindcss/forms';
	@plugin '@tailwindcss/typography';

	label {
		color: white;
	}

	input,
	textarea {
		@apply mt-2 border-1 border-neutral-800/60 bg-neutral-900 text-white;
	}
</style>
