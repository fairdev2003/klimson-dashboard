<script lang="ts">
	import { contextMenuOptions } from '$lib/dashboard/stores/store';
	import { type Snippet } from 'svelte';
	import { sidebar_open } from '$lib/dashboard/stores/persist';
	import { fade } from 'svelte/transition';
	import ContextMenu from './ContextMenu.svelte';

	import Sidebar from './sidebar/Sidebar.svelte';

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
	class="relative mt-[65px] flex min-h-[calc(100vh-66px)] bg-neutral-950 text-white"
>
	<div
		class:hidden={!$sidebar_open}
		class="shrink-0 self-stretch w-75 bg-neutral-900 border-t border-neutral-700 z-10
           
           /* 1. Domyślnie na małych ekranach: */
           absolute inset-y-0 left-0 h-full
           
           /* 2. Na dużych ekranach (lg+): */
           lg:static lg:flex lg:flex-col"
	>
		<Sidebar />
	</div>

	<!-- Główna zawartość -->
	<div
		in:fade={{ duration: 150 }}
		out:fade={{ duration: 150 }}
		class="flex-1 flex-col overflow-x-auto duration-300"
	>
		<!-- <DashboardDock /> -->
		<div class="mb-5 flex flex-col gap-5">
			{@render children()}
		</div>
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
