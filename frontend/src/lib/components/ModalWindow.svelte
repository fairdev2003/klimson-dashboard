<script lang="ts">
	import type { Snippet } from 'svelte';
	import { tv } from 'tailwind-variants';

	const backdropVariants = tv({
		base: '',
		variants: {
			style: {
				normal: '',
				blur: ''
			}
		}
	});

	const modalVariants = tv({
		base: 'relative border flex max-w-2xl flex-col md:backdrop-blur-none bg-neutral-950',
		variants: {
			size: {
				small: '',
				popup: '',
				a4: 'lg:h-[70%] w-[98%] lg:w-7xl w-[98%]'
			},
			theme: {
				modern: 'border-neutral-800/60 bg-neutral-950/60 backdrop-blur-lg'
			}
		}
	});

	type Props = {
		opened?: boolean;
		size?: 'small' | 'popup' | 'a4';
		backdrop_style: 'normal' | 'blur';
		onClose?: () => void;
		title?: string;
		children?: Snippet;
		backgroundCloseDisabled?: boolean;
	};

	const {
		backdrop_style,
		children,
		onClose,
		opened,
		size = 'a4',
		title,
		backgroundCloseDisabled = false
	}: Props = $props();
</script>

{#if opened}
	<!-- svelte-ignore a11y_click_events_have_key_events -->
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<div
		onclick={(e) => {
			if (backgroundCloseDisabled) return;
			e.stopPropagation();
			if (onClose) {
				onClose();
			}
		}}
		class="fixed inset-0 z-50 flex items-center justify-center bg-black/50 text-white md:backdrop-blur-lg lg:backdrop-blur-lg"
	>
		<!-- svelte-ignore a11y_no_static_element_interactions -->
		<div
			onclick={(e) => {
				e.stopPropagation();
			}}
			class={modalVariants({ size })}
		>
			<!-- title -->
			<div
				class="mb-2 flex h-10 flex-shrink-0 items-center justify-between border-b border-neutral-700/60 bg-neutral-800/60 px-5"
			>
				<p>{title}</p>
			</div>
			<!-- content -->
			<div class="scrollable flex-1 overflow-y-auto p-6 pt-5">
				{#if children}
					{@render children()}
				{/if}
			</div>
		</div>
	</div>
{/if}
