<script lang="ts">
	import { goto } from '$app/navigation';
	import Dashboard from '$lib/dashboard/dashboard.svelte';
	import { cubicOut } from 'svelte/easing';
	import { fly, slide } from 'svelte/transition';
	import DockMenuItem from './DockMenuItem.svelte';
	import SidebarPill from '../sidebar/SidebarPill.svelte';
	import { route } from '$lib/dashboard/stores/persist';

	type Props = { mobileDockOpened: boolean };

	let routesSectionOpened = $state(true);
	let { mobileDockOpened = $bindable() }: Props = $props();
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div
	onclick={(e) => {
		e.stopPropagation();
	}}
	transition:fly={{ y: -40, duration: 250, easing: cubicOut }}
	class="lg:hidden p-4 flex flex-col gap-4 h-dvh shadow-2xl w-full bg-background absolute z-20 border-b border-neutral-800"
>
	<div class="flex flex-col gap-4 w-full overflow-auto scroll-class">
		<SidebarPill />
		<DockMenuItem bind:opened={routesSectionOpened} name="Dashboard Routes">
			<div class="flex flex-wrap gap-2 mt-4" transition:slide={{ duration: 300 }}>
				{#each Dashboard.constants.SidebarContents as item}
					<button
						class:normal={item.route !== $route}
						class:selected={item.route === $route}
						onclick={() => {
							goto(item.href);
							$route = item.route;
							mobileDockOpened = false;
						}}
						class="bg-foreground flex items-center justify-center p-2 rounded-xl"
					>
						<p class="mx-2">
							{item.name}
						</p>
					</button>
				{/each}
			</div>
		</DockMenuItem>
	</div>
</div>

<style>
	.selected {
		background-color: color-mix(in srgb, var(--color-primary) 60%, transparent);
		font-weight: bold;
		color: var(--color-primary-foreground, #ffffff);
	}
	.selected:hover {
		background-color: var(--color-primary);
	}

	.normal {
		background-color: var(--color-background);
		border-color: var(--color-border);
		color: var(--color-secondary-text);
	}
	.normal:hover {
		background-color: var(--color-foreground);
	}

	.disabled {
		opacity: 0.5;
		cursor: not-allowed;
	}

	.disabled:hover {
		background-color: transparent !important;
	}
</style>
