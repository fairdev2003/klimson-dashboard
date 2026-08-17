<script lang="ts">
	import { goto } from '$app/navigation';
	import { route } from '$lib/dashboard/stores/persist';
	import Icon from '@iconify/svelte';
	import type { SidebarItemType } from './sidebar.types';
	import { toast } from '$lib/dashboard/stores/toast';
	import { fade } from 'svelte/transition';

	type Props = { content: SidebarItemType };
	let { content }: Props = $props();
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->

<div>
	<button
		onclick={() => {
			$route = content.route;

			goto(content.href);
		}}
		class:normal={$route !== content.route}
		class:selected={$route === content.route}
		class="p-2 focus:outline-none cursor-pointer w-full text-neutral-200 rounded-lg flex gap-2 items-start"
	>
		<Icon icon={String(content.icon)} width="25" height="25"></Icon>
		<div>
			<p class="text-neutral-200">
				{content.name}
			</p>
		</div>
	</button>
	<div class="flex gap-2">
		{#if content.child}
			<div class="w-0.5 bg-neutral-700 shadow-lg h-full"></div>
			<div class="ml-5 mt-2 flex gap-2 flex-col">
				{#each content.child as child}
					<p>{child.name}</p>
				{/each}
			</div>
		{/if}
	</div>
</div>

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
