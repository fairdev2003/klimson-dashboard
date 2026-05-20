<script lang="ts">
	import { goto } from '$app/navigation';
	import { route } from '$lib/dashboard/stores/persist';
	import Icon from '@iconify/svelte';
	import type { SidebarItemType } from './sidebar.types';

	type Props = { content: SidebarItemType };
	let { content }: Props = $props();

	let statement = $derived($route === content.href);
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
{#if content.child}{:else}
	<div
		onclick={() => {
			goto(content.href);
		}}
		class:normal={!statement}
		class:selected={statement}
		class="flex items-center px-3 cursor-pointer transition-colors border-x border-b h-10 gap-3"
	>
		{#if content.icon}
			<Icon icon={content.icon} />
		{/if}
		<p class="text-neutral-300 text-sm">
			{content.name}
		</p>
	</div>
{/if}

<style>
	@import 'tailwindcss';

	.link {
		@apply bg-blue-500;
	}

	.selected {
		@apply bg-blue-700/60 font-bold hover:bg-blue-700 border-blue-500 border text-white;
	}

	.normal {
		@apply bg-neutral-800 hover:bg-neutral-700 border-neutral-700 text-neutral-500;
	}
</style>
