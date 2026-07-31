<script lang="ts">
	import { goto } from '$app/navigation';
	import Button from '$lib/components/Button.svelte';
	import Dashboard from '$lib/dashboard/dashboard.svelte';
	import { onMount } from 'svelte';
	import { cubicOut } from 'svelte/easing';
	import { fly, slide } from 'svelte/transition';
	import DockMenuItem from './DockMenuItem.svelte';
	import SidebarPill from '../sidebar/SidebarPill.svelte';

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
	class="lg:hidden p-4 flex flex-col gap-4 h-dvh shadow-2xl w-full bg-neutral-900 absolute z-20 border-b border-neutral-800"
>
	<div class="flex flex-col gap-4 w-full overflow-auto scroll-class">
		<SidebarPill />
		<DockMenuItem bind:opened={routesSectionOpened} name="Dashboard Routes">
			<div class="flex flex-wrap gap-2 mt-4" transition:slide={{ duration: 300 }}>
				{#each Dashboard.constants.SidebarContents as item}
					<button
						onclick={() => {
							goto(item.href);
							mobileDockOpened = false;
						}}
						class="bg-neutral-800 p-2 rounded-xl"
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
