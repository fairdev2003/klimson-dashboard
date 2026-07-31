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
	import { route } from '$lib/dashboard/stores/persist';

	let contentRef: HTMLElement | null = $state(null);

	$effect(() => {
		if (!contentRef) return;

		if ($sidebar_open) {
			gsap.fromTo(contentRef, { x: -20 }, { x: 0, duration: 0.3, ease: 'power2.out' });
		}
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

<div
	style="width: {$sidebar_open ? '300px' : '40px'}; transition: width 0.3s ease;"
	class="sticky left-0 top-16.25 h-dvh z-500 overflow-hidden flex flex-col m-5"
>
	<div>
		<SidebarToggler bind:opened={$sidebar_open} onclick={handleToggle} />
	</div>

	{#if $sidebar_open}
		<div bind:this={contentRef}>
			<SidebarPill />

			<nav class="flex flex-col gap-2 mt-3">
				{#each Dashboard.constants.SidebarContents as content, i}
					<SidebarItem {content} />
				{/each}
			</nav>
		</div>
	{:else}
		<nav class="flex flex-col gap-2 mt-3">
			{#each Dashboard.constants.SidebarContents as content, i}
				<button
					onclick={() => {
						$route = content.route;

						goto(content.href);
					}}
					class:normal={$route !== content.route}
					class:selected={$route === content.route}
					class="p-2 focus:outline-none cursor-pointer rounded-xl mb-4"
				>
					<Icon icon={String(content.icon)} width="25" height="25"></Icon>
				</button>
			{/each}
		</nav>
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
		@apply bg-blue-700 hover:bg-blue-700  text-white;
	}

	.normal {
		@apply hover:bg-neutral-700 border-neutral-700;
	}
</style>
