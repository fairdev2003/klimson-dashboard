<script lang="ts">
	import { scale } from 'svelte/transition';
	import { computePosition, flip, shift, offset, arrow } from '@floating-ui/dom';
	import type { Snippet } from 'svelte';

	type Props = {
		children: Snippet;
		content?: string;
		position?: 'top' | 'bottom' | 'left' | 'right';
		disabled?: boolean;
	};

	let {
		children,
		content = 'Tooltip content',
		position = 'top',
		disabled = false
	}: Props = $props();

	let hovered = $state(false);
	let referenceEl = $state<HTMLElement | null>(null);
	let floatingEl = $state<HTMLElement | null>(null);
	let arrowEl = $state<HTMLElement | null>(null);

	// Funkcja aktualizująca pozycję
	async function updatePosition() {
		if (!referenceEl || !floatingEl) return;

		const { x, y, placement, middlewareData } = await computePosition(referenceEl, floatingEl, {
			placement: position,
			middleware: [offset(8), flip(), shift({ padding: 5 }), arrow({ element: arrowEl! })]
		});

		Object.assign(floatingEl.style, {
			left: `${x}px`,
			top: `${y}px`
		});

		if (middlewareData.arrow && arrowEl) {
			const { x: ax, y: ay } = middlewareData.arrow;
			const side = placement.split('-')[0] as 'top' | 'bottom' | 'left' | 'right'; // Mapujemy stronę na przeciwną krawędź dla strzałki
			const staticSide = {
				top: 'bottom',
				right: 'left',
				bottom: 'top',
				left: 'right'
			}[side];
			Object.assign(arrowEl.style, {
				left: ax != null ? `${ax}px` : '',
				top: ay != null ? `${ay}px` : '',
				[staticSide!]: '-4px'
			});
		}
	}

	$effect(() => {
		if (hovered && referenceEl && floatingEl) {
			updatePosition();
		}
	});
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
<div
	bind:this={referenceEl}
	class="relative"
	onmouseenter={async () => {
		hovered = true;
	}}
	onmouseleave={() => (hovered = false)}
	onfocusin={() => (hovered = true)}
	onfocusout={() => (hovered = false)}
>
	{@render children()}

	{#if hovered && !disabled}
		<div
			bind:this={floatingEl}
			transition:scale={{ duration: 150, start: 0.95 }}
			class="pointer-events-none absolute top-0 left-0 z-2000 w-max max-w-xs"
		>
			<div
				class="relative rounded-lg border border-secondary bg-neutral-900 px-4 py-1.5 text-xs text-white shadow-xl"
			>
				{content}
			</div>
		</div>
	{/if}
</div>
