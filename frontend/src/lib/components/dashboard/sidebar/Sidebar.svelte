<script lang="ts">
	import { onMount } from 'svelte';
	import gsap from 'gsap';
	import SidebarToggler from './SidebarToggler.svelte';
	import SidebarPill from './SidebarPill.svelte';
	import PagesSelector from './PagesSelector.svelte';
	import SidebarItem from './SidebarItem.svelte';
	import { Dashboard } from '$lib/dashboard/logic';
	import { sidebar_open } from '$lib/dashboard/stores/store';
	import { goto } from '$app/navigation';
	import Icon from '@iconify/svelte';
	import { dashboard_config, route } from '$lib/dashboard/stores/persist';
	import { page } from '$app/state';
	import LatestRoutes from './LatestRoutes.svelte';
	import DashboardContents from './DashboardContents.svelte';
	import StarredContents from './StarredContents.svelte';
	import LatestUserContents from './LatestUserContents.svelte';
	import { toast } from '$lib/dashboard/stores/toast';

	let contentRef: HTMLElement | null = $state(null);

	$effect(() => {
		if (!contentRef) return;

		if ($sidebar_open) {
			gsap.fromTo(contentRef, { x: -20 }, { x: 0, duration: 0.3, ease: 'power2.out' });
		}
	});

	$effect(() => {
		console.log(page.route);
	});

	function handleToggle() {
		if ($sidebar_open && contentRef) {
			gsap.to(contentRef, {
				x: -20,
				duration: 0.2,
				ease: 'power2.in',
				onComplete: () => {
					$sidebar_open = false;
				}
			});
		} else {
			$sidebar_open = true;
		}
	}
</script>

{#if !$sidebar_open && $dashboard_config.sidebarBehavior === 'autoHide'}
	<div
		class="fixed inset-y-0 left-0 w-6 z-100 pointer-events-none"
		onmouseenter={() => {
			$sidebar_open = true;
		}}
	>
		<!-- Aktywny margines o szerokości 16px, który reaguje na mysz -->
		<div class="w-4 h-full pointer-events-auto"></div>
	</div>
{/if}

<div
	style="width: {$sidebar_open ? '300px' : '0px'}; transition: width 0.3s ease;"
	class="sticky left-0 top-16.25 h-dvh z-500 overflow-auto flex flex-col"
>
	{#if $sidebar_open}
		<!-- svelte-ignore a11y_no_static_element_interactions -->
		<div
			bind:this={contentRef}
			class="h-full bg-neutral-900/90 backdrop-blur-md border-r border-neutral-800"
			onmouseleave={() => {
				if ($dashboard_config.sidebarBehavior === 'alwaysOn') return;
				$sidebar_open = false;
			}}
		>
			<div class="m-5">
				<SidebarPill />
			</div>
			<nav>
				<StarredContents />
				<LatestUserContents />
				<DashboardContents />
			</nav>
		</div>
	{/if}
</div>

<svelte:document
	onkeydown={async (keyboard_event) => {
		if (keyboard_event.key === 'Tab') {
			keyboard_event.preventDefault();
			$sidebar_open = !$sidebar_open;
			return;
		}
	}}
/>

<style>
	@import 'tailwindcss';

	.selected {
		color: var(--color-text);
		background-color: var(--color-primary);
	}

	.selected:hover {
		background-color: var(--color-foreground);
	}

	.normal {
		background-color: transparent;
	}

	.normal:hover {
		background-color: var(--color-foreground);
	}
</style>
