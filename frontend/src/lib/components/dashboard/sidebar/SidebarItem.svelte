<script lang="ts">
	import { goto } from '$app/navigation';
	import Icon from '@iconify/svelte';
	import type { SidebarItemType } from './sidebar.types';
	import { route } from '$lib/dashboard/stores/persist';

	type Props = { content: SidebarItemType };
	let { content }: Props = $props();

	function redirectTo(routeRecord: SidebarItemType) {
		$route = content.route;

		if (routeRecord.disabled) return;

		goto(routeRecord.href);
	}
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->

<div
	onclick={() => redirectTo(content)}
	class="h-10 cursor-pointer relative border border-neutral-800 items-center h-full hover:bg-white/10 rounded-lg group hover:border-blue-700 overflow-hidden"
>
	{@render RouteRecordIndicator()}

	{@render ChildComponentsDropdown(content.child ? true : false)}

	{@render CurrentRouteSelection($route === content.route)}

	<div class="flex h-full">
		<p class="px-4 text-neutral-300">{content.name}</p>
	</div>
</div>

{#snippet RouteRecordIndicator()}
	<div
		class="w-1 h-full left-0 z-2 top-0 bg-neutral-800 group group-hover:bg-blue-700 absolute"
	></div>
{/snippet}

{#snippet ChildComponentsDropdown(childExists?: boolean)}
	{#if childExists}
		<div class="absolute z-1 right-2 text-neutral-300 top-1.5">
			<Icon icon="mdi:chevron-down" />
		</div>
	{/if}
{/snippet}

{#snippet CurrentRouteSelection(selected: boolean)}
	{#if selected}
		<div class="absolute z-1 w-full h-full bg-blue-700/40"></div>
	{/if}
{/snippet}

<style>
	@import 'tailwindcss';
	.link {
		background-color: var(--color-primary);
	}

	.selected {
		@apply bg-neutral-700 hover:bg-neutral-600;
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
