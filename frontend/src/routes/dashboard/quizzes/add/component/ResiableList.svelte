<script lang="ts">
	import type { Snippet } from 'svelte';
	import { slide } from 'svelte/transition';
	import { imageForm } from './form/form.store';
	import { toast } from '$lib/dashboard/stores/toast';

	let {
		section_name,
		children,
		isOpen = $bindable(),
		onclick,
		nav_icon,
		disabled,
		disabledMessage = 'error'
	} = $props<{
		section_name: string;
		children?: Snippet;
		isOpen?: boolean;
		nav_icon?: Snippet;
		onclick?: () => void;
		disabled?: boolean;
		disabledMessage?: string;
	}>();
</script>

<div class="overflow-visible z-100 border border-neutral-700 bg-neutral-900">
	<button
		type="button"
		class="flex w-full {disabled
			? 'opacity-30'
			: 'hover:bg-neutral-800 cursor-pointer'} items-center justify-between border-b p-5 text-left text-white transition-all"
		style="border-color: {isOpen ? 'rgb(64, 64, 64)' : 'transparent'}"
		onclick={() => {
			if (disabled) {
				toast.warning(disabledMessage);
				return;
			}
			onclick ? onclick() : (isOpen = !isOpen);
		}}
		aria-expanded={isOpen}
	>
		<div class="flex items-center gap-4">
			{#if nav_icon}
				{@render nav_icon()}
			{/if}
			<span class="text-lg font-medium">{section_name}</span>
		</div>
		<svg
			class="h-5 w-5 transform transition-transform duration-300 {isOpen ? 'rotate-180' : ''}"
			fill="none"
			viewBox="0 0 24 24"
			stroke="currentColor"
		>
			<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
		</svg>
	</button>

	{#if isOpen}
		<div transition:slide={{ duration: 300 }} class="p-5">
			<div class="leading-relaxed text-neutral-400">
				{@render children?.()}
			</div>
		</div>
	{/if}
</div>
