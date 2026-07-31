<script lang="ts">
	import { type Snippet } from 'svelte';
	import { fade } from 'svelte/transition';

	let {
		children,
		tooltipContent,
		delay = 800
	}: {
		children: Snippet;
		tooltipContent: Snippet;
		delay?: number;
	} = $props();

	let mouseX = $state(0);
	let mouseY = $state(0);
	let isVisible = $state(false);

	function portal(node: HTMLElement) {
		document.body.appendChild(node);
		return {
			destroy() {
				if (node.parentNode) node.parentNode.removeChild(node);
			}
		};
	}

	function handleMouseMove(e: MouseEvent) {
		mouseX = e.clientX + 13;
		mouseY = e.clientY + 13;
	}
</script>

<div
	class="relative inline-block"
	onmousemove={handleMouseMove}
	onmouseenter={() => (isVisible = true)}
	onmouseleave={() => (isVisible = false)}
>
	{@render children()}

	{#if isVisible}
		<div
			use:portal
			in:fade={{ duration: 150, delay }}
			class="text-white pointer-events-none fixed z-[9999] bg-neutral-900"
			style="left: {mouseX}px; top: {mouseY}px; min-width: max-content;"
		>
			<div class="p-2">
				{@render tooltipContent()}
			</div>
		</div>
	{/if}
</div>

<style>
	/* Ważne: style stąd nie dosięgną elementu w body, 
       dlatego klasy Tailwindowe w use:portal są bezpieczniejsze */
</style>
