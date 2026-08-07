<script lang="ts">
	import { contextMenuOptions } from '$lib/dashboard/stores/store';
	import { type Snippet } from 'svelte';
	import { dashboard_config, isMobile, sidebar_open } from '$lib/dashboard/stores/persist';
	import { fade, slide } from 'svelte/transition';
	import ContextMenu from './ContextMenu.svelte';

	import Sidebar from './sidebar/Sidebar.svelte';
	import CMSNavbar from './CMSNavbar.svelte';
	import DashboardDock from './dock/DashboardDock.svelte';
	import { storage_logic } from '$lib/dashboard/storage/storage.svelte';
	import Icon from '@iconify/svelte';
	import { page } from '$app/state';
	import { debug } from '$lib/terminal/logic';

	type Props = {
		children: Snippet;
	};

	let menuPos: { x: number; y: number } = $state({ x: 0, y: 0 });
	let showMenu: boolean = $state(false);
	let currentTarget: HTMLDivElement | null = $state(null);

	function handleRightClick(e: MouseEvent) {
		e.preventDefault();
		menuPos = { x: e.clientX, y: e.clientY };
		// showMenu = true;
		debug.log(page.route);
	}

	let { children }: Props = $props();
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
<div
	oncontextmenu={(e) => handleRightClick(e)}
	class="flex flex-col h-screen overflow-hidden bg-primary-background text-white"
>
	<header class="border-b h-20 border-neutral-700 flex items-center px-4 shrink-0">
		<CMSNavbar />
	</header>

	<div class="flex flex-1 relative">
		<aside class="hidden lg:flex z-20 bg-background border-r border-border transition-all">
			<Sidebar />
		</aside>

		<main in:fade={{ duration: 150 }} out:fade={{ duration: 150 }} class="flex-1 text-text">
			<DashboardDock />
			<div class="">
				{@render children()}
			</div>
		</main>
	</div>

	{#if showMenu}
		<ContextMenu
			{...menuPos}
			title="Opcje sektora"
			context={$contextMenuOptions}
			close={() => (showMenu = false)}
		/>
	{/if}
</div>

<style>
	@import 'tailwindcss';

	.dock-visible {
		@apply h-[calc(100dvh-80px-100px)] overflow-y-auto;
	}

	.dock-invisible {
		@apply h-[calc(100dvh)] overflow-y-auto;
	}
</style>
