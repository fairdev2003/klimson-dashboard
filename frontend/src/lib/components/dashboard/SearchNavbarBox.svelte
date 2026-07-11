<script lang="ts">
	import { goto } from '$app/navigation';
	import { dashboard_config } from '$lib/dashboard/stores/persist';
	import Icon from '@iconify/svelte';
	import { fade } from 'svelte/transition';
	let value: string = $state('');
	let opened: boolean = $state(false);

	let boxRef: HTMLDivElement | null = null;
	let inputContainerRef: HTMLInputElement | null = null;

	// Logika wyszukiwania

	type Route = {
		name: string;
		url: string;
		icon: string;
		description: string;
	};

	const routes: Route[] = [
		{
			name: 'Hub',
			icon: 'mynaui:home',
			description: 'Main page',
			url: '/dashboard'
		},
		{
			name: 'Database Editor',
			icon: 'mynaui:database',
			description: 'Edit your database records',
			url: '/dashboard/database'
		},
		{
			name: 'Context Storage',
			icon: 'material-symbols:store',
			description: 'Manage persistent context',
			url: '/dashboard/context_storage'
		},
		{
			name: 'Spotify',
			icon: 'mynaui:music',
			description: 'Spotify integration settings',
			url: '/dashboard/spotify'
		},
		{
			name: 'Tools',
			icon: 'mynaui:cog',
			description: 'Utility tools',
			url: '/dashboard/tools'
		},
		{
			name: 'File Storage',
			icon: 'mynaui:folder',
			description: 'Manage uploaded files',
			url: '/dashboard/storage'
		},
		{
			name: 'PG3D',
			icon: 'mdi:controller',
			description: 'Pixel Gun 3D dashboard',
			url: '/dashboard/pg3d'
		},
		{
			name: 'API Routes',
			icon: 'material-symbols:api',
			description: 'Manage API endpoints',
			url: '/dashboard/routes'
		},
		{
			name: 'CMS Access',
			icon: 'mdi:user-key',
			description: 'Manage user content access',
			url: '/dashboard/users'
		},
		{
			name: 'Redis Storage',
			icon: 'devicon:redis',
			description: 'Redis persisted storage manager',
			url: '/dashboard/redis_writable'
		}
	];

	const searchedRoutes = $derived(
		value === '' ? routes : routes.filter((e) => e.name.toLowerCase().includes(value.toLowerCase()))
	);

	const searchedBookmarks = $derived(
		value === ''
			? $dashboard_config.bookmarks
			: $dashboard_config.bookmarks.filter((e) =>
					e.name.toLowerCase().includes(value.toLowerCase())
				)
	);

	$effect(() => {
		if (opened) {
			inputContainerRef?.focus();
		}
	});
</script>

<div class="relative hidden justify-center w-[500px] md:flex lg:flex z-100">
	<input
		class="h-[50px] w-full bg-neutral-800 border border-neutral-700/60 rounded-xl px-4 outline-none"
		bind:value
		onclick={() => {
			opened = true;
		}}
		placeholder="Click to search"
	/>

	{#if opened}
		<!-- svelte-ignore a11y_click_events_have_key_events -->
		<!-- svelte-ignore a11y_no_static_element_interactions -->
		<div
			class="bg-black/50 fixed inset-0 z-90"
			onclick={() => {
				opened = false;
			}}
		></div>

		<!-- svelte-ignore a11y_click_events_have_key_events -->
		<!-- svelte-ignore a11y_no_static_element_interactions -->
		<div
			in:fade={{ duration: 300 }}
			out:fade={{ duration: 300 }}
			class="fixed w-[500px] border border-neutral-700 p-3 z-200 bg-neutral-900 rounded-xl"
			onclick={(e) => e.stopPropagation()}
		>
			<input
				bind:this={inputContainerRef}
				bind:value
				class="w-full bg-neutral-800 p-3 rounded-lg border-0 outline-none"
				placeholder="Search..."
			/>
			<div class="mt-3 h-[300px] overflow-auto">
				<div class="flex gap-3 flex-col">
					{#if searchedRoutes.length > 0 || searchedBookmarks.length > 0}
						{@render RoutesTab()}
						{@render BookmarksTab()}
					{:else}
						<p class="text-neutral-500 p-3">No results found</p>
					{/if}
				</div>
			</div>
		</div>
	{/if}
</div>

{#snippet RoutesTab()}
	<p class="font-black">Dashboard routes</p>
	{#each searchedRoutes as search_route}
		<button
			onclick={() => {
				goto(search_route.url);
				opened = false;
			}}
			class="flex items-center gap-3 w-full h-15 rounded-lg cursor-pointer bg-neutral-800 hover:bg-neutral-700 px-4 transition-colors"
		>
			<div class="flex items-center justify-center size-8 bg-neutral-900 rounded-md shrink-0">
				<Icon icon={search_route.icon} width="20" height="20" class="text-neutral-300" />
			</div>
			<div class="flex flex-col">
				<p class="font-black text-start text-neutral-200">{search_route.name}</p>
				<p class="text-xs text-start text-neutral-400">{search_route.description}</p>
			</div>
		</button>
	{/each}
{/snippet}

{#snippet BookmarksTab()}
	<p class="font-black">Saved Bookmarks</p>
	{#each searchedBookmarks as search_bookmark}
		<button
			onclick={() => {
				window.open(search_bookmark.href, '_blank', 'noopener,noreferrer');

				opened = false;
			}}
			class="flex items-center gap-3 w-full h-15 rounded-lg cursor-pointer bg-neutral-800 hover:bg-neutral-700 px-4 transition-colors"
		>
			<div
				class="flex items-center justify-center size-8 {search_bookmark.color} rounded-md shrink-0"
			></div>
			<div class="flex flex-col">
				<p class="font-black text-start text-neutral-200">{search_bookmark.name}</p>
				<p class="text-xs text-start text-neutral-400">{search_bookmark.slug}</p>
			</div>
		</button>
	{/each}
{/snippet}

<svelte:document
	onkeydown={(e) => {
		if (e.key === 'Escape' && opened) {
			opened = false;
		}
	}}
/>
