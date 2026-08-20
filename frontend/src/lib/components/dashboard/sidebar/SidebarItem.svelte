<script lang="ts">
	import { goto } from '$app/navigation';
	import Icon from '@iconify/svelte';
	import type { SidebarItemType } from './sidebar.types';
	import { route } from '$lib/dashboard/stores/persist';
	import Dashboard from '$lib/dashboard/dashboard.svelte';
	import { slide } from 'svelte/transition';

	type Props = { content: SidebarItemType };
	let { content }: Props = $props();

	let opened: boolean = $state(false);

	function redirectTo(routeRecord: SidebarItemType) {
		$route = content.route;

		Dashboard.state.latestRoutes.push(routeRecord);

		if (routeRecord.child && routeRecord.child.length > 0) {
			opened = !opened;
		}

		if (routeRecord.disabled) return;

		goto(routeRecord.href);
	}
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->

<div>
	<button
		onclick={() => redirectTo(content)}
		class:selected={$route === content.route}
		class:normal={$route !== content.route}
		class="h-8 cursor-pointer relative flex items-center group overflow-hidden w-full"
	>
		{@render RouteRecordIndicator($route === content.route)}

		{@render ChildComponentsDropdown(content.child ? true : false)}

		<div class="flex items-center gap-2 px-4">
			{#if content && content.icon}
				<Icon icon={content.icon} />
				<p>{content.name}</p>
			{/if}
		</div>
	</button>
	{#if opened}
		<div transition:slide={{ duration: 150 }} class="w-full flex flex-col gap-1 pl-6">
			{#each content.child as child}
				<button
					onclick={() => redirectTo(child)}
					class:selected={$route === child.route}
					class:normal={$route !== child.route}
				>
					{@render RouteRecordIndicator($route === child.route)}

					<div class=" flex items-center gap-2 px-4">
						{#if child && child.icon}
							<Icon icon={child.icon} />
							<p>{child.name}</p>
						{/if}
					</div>
				</button>
			{/each}
		</div>
	{/if}
</div>

{#snippet RouteRecordIndicator(selected: boolean)}
	<div
		class:bg-blue-700={selected}
		class:bg-transparaent={!selected}
		class="w-1 h-full left-0 z-2 top-0 group group-hover:bg-blue-700 absolute"
	></div>
{/snippet}

{#snippet ChildComponentsDropdown(childExists?: boolean)}
	{#if childExists}
		<div class="absolute z-1 right-2 text-neutral-300">
			<Icon
				icon="mdi:chevron-down"
				class="w-4 h-4 transition-transform duration-250 ease-in-out {opened
					? 'rotate-180'
					: 'rotate-0'}"
			/>
		</div>
	{/if}
{/snippet}

<style>
	@import 'tailwindcss';

	.selected {
		@apply bg-blue-700/40;
	}

	.normal {
		@apply hover:bg-white/10 hover:border-blue-800/60 text-neutral-400;
	}
</style>
