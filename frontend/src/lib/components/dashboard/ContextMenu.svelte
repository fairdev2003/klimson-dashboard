<script lang="ts">
	import { developerView } from '$lib/dashboard/stores/persist';
	import Icon from '@iconify/svelte';
	import { onMount } from 'svelte';
	import { fade } from 'svelte/transition';

	type Props = {
		x: number;
		y: number;
		context: ContextMenu[];
		close: () => void;
		title?: string;
	};

	let { x, y, context, close, title = 'Opcje sekcji' }: Props = $props();

	onMount(() => {
		const handleOutsideClick = () => close();

		setTimeout(() => {
			window.addEventListener('click', handleOutsideClick);
		}, 10);

		return () => window.removeEventListener('click', handleOutsideClick);
	});
</script>

{#if context.length > 0}
	<div
		class="absolute inset-0 z-[2000] flex h-screen w-screen items-center justify-center {$developerView
			? 'bg-blue-500/50'
			: ''}"
	>
		{#if $developerView}
			<p>development view</p>
		{/if}
	</div>
	<div
		in:fade={{ duration: 150 }}
		out:fade={{ duration: 150 }}
		class="fixed z-[2000] min-w-[240px] overflow-hidden rounded-xl border border-neutral-700 bg-neutral-800 p-1.5 shadow-2xl backdrop-blur-md"
		style="top: {y}px; left: {x}px;"
	>
		
		{#each context as items}
			<p class="p-2 text-sm text-neutral-400">{items.contextMenuName}</p>
			{#each items.items as item}
				<button
					onclick={(e) => {
						e.stopPropagation();
						item.action();
						close();
					}}
					class="group flex w-full items-center gap-3 rounded-lg px-3 py-2 text-left text-sm text-neutral-200 transition-all hover:bg-blue-700 hover:text-white"
				>
					{#if item.icon}
						<Icon
							icon={item.icon}
							class="text-lg {item.color
								? item.color
								: 'text-neutral-400'} transition-colors group-hover:text-white"
						/>
					{/if}
					<span class="flex-1">{item.label}</span>
				</button>
			{/each}
			
		{/each}
	</div>
{/if}
