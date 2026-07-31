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
{#if content.child}{:else}
	<div
		onclick={() => {
			$route = content.route;
			goto(content.href);
		}}
		class:normal={$route !== content.route}
		class:selected={$route === content.route}
		class:disabled={content.disabled}
		class="flex overflow-hidden relative items-center px-3 cursor-pointer rounded-lg transition-colors h-10 gap-3"
	>
		<div class="flex gap-2 items-center">
			{#if content.icon}
				<Icon icon={content.icon} />
			{/if}

			<p class="text-neutral-300 text-sm">
				{content.name}
			</p>
		</div>
	</div>
{/if}

<style>
	@import 'tailwindcss';

	.link {
		@apply bg-blue-500;
	}

	.selected {
		@apply bg-blue-700/60 font-bold hover:bg-blue-700  text-white;
	}

	.normal {
		@apply bg-neutral-800 hover:bg-neutral-700 border-neutral-700 text-neutral-500;
	}

	.disabled {
		@apply opacity-50 hover:bg-neutral-800 hover:cursor-not-allowed;
	}

	.disabled:hover {
		background-color: none !important;
	}
</style>
