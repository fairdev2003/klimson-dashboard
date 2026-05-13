<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/state';
	import { debugOn, route, sidebar_open } from '$lib/dashboard/stores/persist';
	import {
		Album,
		SquareCode,
		Wallpaper,
		Newspaper,
		Menu,
		Settings,
		Image,
		ListStart,
		Info
	} from '@lucide/svelte';
	import type { Component } from 'svelte';
	import tippy from 'tippy.js';
	import Tooltip from './Tooltip.svelte';
	import { get } from 'svelte/store';
	import { toast } from '$lib/dashboard/stores/toast';

	type SContent = {
		image: Component;
		href: string;
		name: string;
	};

	let selected_route: string = $state(page.url.pathname);

	const contents: SContent[] = [
		{ image: ListStart, href: '/dashboard', name: 'Start' },
		{ image: Album, href: '/dashboard/quizzes', name: 'Quizy' },
		{ image: Album, href: '/dashboard/contributors', name: 'Kontrybutorzy' },
		{ image: Newspaper, href: '/dashboard/blog', name: 'Blog' },
		{ image: Wallpaper, href: '/dashboard/hero', name: 'Hero' },
		{ image: SquareCode, href: '/dashboard/routes', name: 'API' },
		{ image: Image, href: '/dashboard/images', name: 'Zdjęcia' },
		{ image: Settings, href: '/dashboard/settings', name: 'Ustawienia' },
		{ image: Info, href: '/dashboard/info', name: 'Informacje' }
	];
</script>

<div
	style:width={$sidebar_open ? '300px' : '72px'}
	oncontextmenu={(e) => {
		toast.info('s');
	}}
	class="shrink-0 self-stretch border-r-1 border-neutral-700 bg-neutral-900 transition-all duration-300 ease-in-out"
>
	<div class="sticky top-[65px] flex h-[calc(100vh-65px)] flex-col justify-start overflow-hidden">
		<div class="flex w-full items-center px-4 py-2">
			<button
				class="flex cursor-pointer items-center justify-center rounded-md p-2 transition-colors hover:bg-blue-900/50 hover:text-white"
				onclick={() => {
					$sidebar_open = !$sidebar_open;
				}}
			>
				<Menu size={24} />
			</button>
		</div>

		<nav class="mt-4 flex flex-col gap-1 px-2">
			{#each contents as content}
				{@render Content(content)}
			{/each}
		</nav>

		<div class="mt-auto mb-4"></div>
	</div>
</div>

{#snippet Content(content: SContent)}
	{@const statement = $route === content.href}
	<Tooltip position="top" content={content.name} disabled={$sidebar_open}>
		<button
			class="group relative flex w-full cursor-pointer items-center overflow-hidden rounded-lg px-4 py-3 transition-all hover:bg-neutral-700/50"
			class:bg-neutral-700={statement}
			class:text-blue-400={statement}
			onclick={() => {
				selected_route = content.href;
				goto(content.href);
			}}
		>
			<div class="flex min-w-[24px] shrink-0 items-center justify-center">
				<content.image />
			</div>

			{#if $sidebar_open}
				<span class="ml-4 text-sm font-medium whitespace-nowrap transition-opacity duration-300">
					{content.name}
				</span>
			{/if}
		</button>
	</Tooltip>
{/snippet}
