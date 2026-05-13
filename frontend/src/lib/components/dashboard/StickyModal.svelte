<script lang="ts">
	import { onMount, tick, type Snippet } from 'svelte';
	import gsap from 'gsap';

	type Props = {
		opened: boolean;
		onClose: () => void;
		className?: string;
		children?: Snippet;
		onopen?: () => void;
	};
	let modalEl: HTMLDivElement;
	let backgroundEl: HTMLDivElement;

	onMount(() => onopen?.());

	$effect(async () => {
		if (opened) {
			await tick();

			gsap.fromTo(
				backgroundEl,
				{ backgroundColor: 'rgba(0, 0, 0, 0)' },
				{ backgroundColor: 'rgba(0, 0, 0, 0.3)' }
			);
			gsap.fromTo(modalEl, { y: -200 }, { y: 0, duration: 0.4 });
		} else {
			handleClose();
		}
	});

	let { opened = $bindable(), onClose, className, children, onopen }: Props = $props();

	async function handleClose() {
		if (!modalEl || !backgroundEl) return;

		const tl = gsap.timeline({
			onComplete: () => {
				document.body.style.overflow = 'auto';
				opened = false;
				onClose?.();
			}
		});

		tl.to(modalEl, { y: -200, opacity: 0, duration: 0.3, ease: 'power2.in' }).to(
			backgroundEl,
			{ backgroundColor: 'rgba(0, 0, 0, 0)', duration: 0.3 },
			'-=0.2'
		);
	}
</script>

{#if opened}
	<!-- svelte-ignore a11y_click_events_have_key_events -->
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<div
		bind:this={backgroundEl}
		class="fixed inset-0 z-2000 flex items-center justify-center"
		onclick={handleClose}
	>
		<div
			onclick={(e) => {
				e.stopPropagation();
			}}
			bind:this={modalEl}
			class={`absolute p-5 min-w-150 top-4 z-2000 ${className} rounded-lg border border-neutral-800 bg-neutral-950 text-white`}
		>
			{#if children}
				{@render children()}
			{/if}
		</div>
	</div>
{/if}
