<script lang="ts">
	import { contextMenuOptions } from '$lib/dashboard/stores/store';
	import { type Snippet } from 'svelte';
	import {
		dashboard_config,
		isMobile,
		mobile_sidebar_open,
		sidebar_open
	} from '$lib/dashboard/stores/persist';
	import { fade } from 'svelte/transition';
	import ContextMenu from './ContextMenu.svelte';

	import Sidebar from './sidebar/Sidebar.svelte';
	import CMSNavbar from './CMSNavbar.svelte';
	import DashboardDock from './dock/DashboardDock.svelte';

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
	}

	let { children }: Props = $props();
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
<div
	oncontextmenu={(e) => handleRightClick(e)}
	class="flex flex-col h-screen overflow-hidden bg-neutral-950 text-white"
>
	<header class="border-b h-[66px] border-neutral-700 flex items-center px-4 shrink-0">
		<CMSNavbar />
	</header>

	<div class="flex flex-1 overflow-hidden relative">
		{#if !$isMobile && $sidebar_open}
			<aside
				onclick={(e) => {
					e.stopPropagation();
				}}
				class="absolute lg:static z-20 w-75 bg-neutral-900 border-r border-neutral-700 transition-all"
			>
				<Sidebar />
			</aside>
		{/if}

		{#if !$sidebar_open}
			<div class="h-full w-75 lg:static hidden"></div>
		{/if}

		{#if $mobile_sidebar_open}
			<div class="lg:hidden absolute inset-0 bg-black/50 z-10">
				<aside
					class="absolute overflow-hidden lg:static z-20 w-75 bg-neutral-900 border-r border-neutral-700 transition-all"
				>
					<Sidebar />
				</aside>
			</div>
		{/if}

		<main in:fade={{ duration: 150 }} out:fade={{ duration: 150 }} class="flex-1">
			<DashboardDock />
			<div
				class:dock-invisible={!$dashboard_config.dock}
				class:dock-visible={$dashboard_config.dock}
				class=""
			>
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
		@apply h-[calc(100dvh-66px-100px)] overflow-y-auto;
	}

	.dock-invisible {
		@apply h-[calc(100dvh-66px)] overflow-y-auto;
	}
</style>
