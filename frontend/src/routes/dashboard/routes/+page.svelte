<script lang="ts">
	import { onMount } from 'svelte';
	import { api } from '$lib/api/api';
	import { routes } from '$lib/dashboard/stores/data.store';
	import type { RoutesResponse } from './types';
	import { blur } from 'svelte/transition';
	import { goto } from '$app/navigation';
	import { base_url } from '$lib/api/api.store';
	import RoutesDocky from '$lib/components/dashboard/dock/boxes/RoutesDocky.svelte';
	import DashboardDock from '$lib/components/dashboard/dock/DashboardDock.svelte';
	import { dockComponent } from '../dashboard.svelte';

	onMount(async () => {
		dockComponent.set(RoutesDocky);
		if ($routes) return;
		const response = await api.misc.GetRoutes();
		$routes = response.data;
		console.log('Routes', $routes);
	});

	let searchBoxValue = $state('');
	let category = $state('');

	const filtered_routes = $derived.by(() => {
		let result =
			category === ''
				? $routes
				: $routes.filter((e) => e.path.toLowerCase().includes(category.toLowerCase()));

		if (searchBoxValue !== '') {
			const query = searchBoxValue.toLowerCase();
			result = result.filter(
				(e) => e.method.toLowerCase().includes(query) || e.path.toLowerCase().includes(query)
			);
		}

		return result;
	});
	type FastSearchItem = {
		name: string;
		path: string;
		color: string;
	};

	const fastSearchList: FastSearchItem[] = [
		{
			name: 'Storage',
			path: '/storage',
			color: 'text-violet-500 bg-violet-800/40 hover:bg-violet-800/60'
		},
		{
			name: 'Database',
			path: '/database',
			color: 'text-pink-500 bg-pink-800/40 hover:bg-pink-800/60'
		},
		{
			name: 'Pixel Gun 3D',
			path: '/pg3d',
			color: 'text-purple-500 bg-purple-800/40 hover:bg-purple-800/60'
		},
		{
			name: 'Spotify',
			path: '/spotify',
			color: 'text-green-500 bg-green-800/40 hover:bg-green-800/60'
		},
		{
			name: 'Context Storage',
			path: '/context_storage',
			color: 'text-indigo-500 bg-indigo-800/40 hover:bg-indigo-800/60'
		},
		{
			name: 'Private',
			path: '/admin',
			color: 'text-slate-400 bg-slate-400/40 hover:bg-slate-400/60'
		}
	];
</script>

<div class="overflow-x-auto text-white">
	<div class="flex flex-col lg:w-4xl mx-auto gap-5 p-5">
		<input
			bind:value={searchBoxValue}
			placeholder="Search Server Api routes..."
			class="flex gap-1 rounded-xl bg-neutral-900 p-4"
		/>

		<div class="flex flex-wrap gap-1.5 items-center">
			{#each fastSearchList as { color, name, path }}
				<!-- svelte-ignore a11y_click_events_have_key_events -->
				<!-- svelte-ignore a11y_no_static_element_interactions -->
				<span
					onclick={() => {
						if (category === path) {
							category = '';
							return;
						}

						category = path;
					}}
					class:opacity-30={category != path && category != ''}
					class="p-1 px-2 rounded-full {color} text-xs cursor-pointer truncate">{name}</span
				>
			{/each}
		</div>

		{#each filtered_routes as route, i}
			<button
				onclick={() => {
					window.open($base_url + route.path, '_blank', 'noopener,noreferrer');
				}}
				in:blur={{ duration: 300 }}
				class="flex overflow-hidden w-auto gap-4 rounded-xl truncate items-center bg-neutral-900 hover:bg-neutral-800 transition-colors cursor-pointer p-4"
			>
				{@render Method(route.method)}
				<div class="flex truncate">
					<span class="text-purple-600 font-black"> {`${$base_url}`}</span>
					<p class="text-blue-500 hover:underline">
						{route.path}
					</p>
				</div>
			</button>
		{/each}
	</div>
</div>

{#snippet Method(method: string)}
	<div class="w-20">
		{#if method === 'GET'}
			<div class="bg-green-500/50 p-1 rounded-lg flex justify-center">
				<p class="text-green-500 font-black">{method}</p>
			</div>
		{/if}
		{#if method === 'POST'}
			<div class="bg-yellow-500/50 p-1 rounded-lg flex justify-center">
				<p class="text-yellow-500 font-black">{method}</p>
			</div>
		{/if}
		{#if method === 'PUT'}
			<div class="bg-blue-500/50 p-1 rounded-lg flex justify-center">
				<p class="text-blue-500 font-black">{method}</p>
			</div>
		{/if}
		{#if method === 'DELETE'}
			<div class="bg-red-500/50 p-1 rounded-lg flex justify-center">
				<p class="text-red-500 font-black">{method}</p>
			</div>
		{/if}
		{#if method === 'HEAD'}
			<div class="bg-purple-700/20 p-1 rounded-lg flex justify-center">
				<p class="text-purple-700 font-black">{method}</p>
			</div>
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
