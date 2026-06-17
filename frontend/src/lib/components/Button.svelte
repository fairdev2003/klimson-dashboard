<script lang="ts">
	import type { Snippet } from 'svelte';
	import { tv } from 'tailwind-variants';
	import Loader from './dashboard/Loader.svelte';

	const buttonStyles = tv({
		base: 'cursor-pointer px-4 py-2 self-start flex gap-2 items-center text-white ',
		variants: {
			theme: {
				base: 'bg-primary border-1 border-secondary hover:bg-secondary',
				danger: 'bg-red-500 hover:bg-red-500/80',
				secondary: 'bg-blue-700 hover:bg-blue-600',
				correct: 'bg-green-500 hover:bg-green-600'
			},
			loading: {
				true: 'opacity-50',
				false: 'opacity-100 hover:bg'
			},
			disabled: {
				true: 'bg-neutral-600 hover:bg-neutral-600 border-neutral-700 text-neutral-400 cursor-not-allowed grayscale-[0.5] opacity-60',
				false: 'cursor-pointer'
			},
			size: {
				small: 'text-sm px-3 py-1 rounded-lg',
				medium: 'text-base px-4 py-2 rounded-lg',
				large: 'text-lg px-6 py-3'
			}
		},
		defaultVariants: {
			size: 'medium'
		}
	});

	type Props = {
		theme?: 'base' | 'danger' | 'secondary' | 'correct';
		children?: Snippet;
		className?: string;
		onclick?: () => void;
		loading?: boolean;
		disabled?: boolean; // Nowy prop
		loadingText?: string;
		size?: 'small' | 'medium' | 'large';
	};
	const {
		theme = 'base',
		children = PlainText,
		className = '',
		onclick,
		loading,
		disabled = false,
		loadingText,
		size
	}: Props = $props();
</script>

{#snippet PlainText()}
	<p>Click me</p>
{/snippet}

<button
	class={buttonStyles({ theme, className, loading, size, disabled })}
	onclick={() => {
		if (disabled) return;

		if (!loading) {
			if (onclick) {
				onclick();
			}
		}
	}}
>
	{#if loading}
		<Loader theme="regular" />
	{/if}

	{#if loadingText && loading}
		{loadingText}
	{:else}
		{@render children()}
	{/if}
</button>

<!-- <div
	class="border-1 w-50 flex h-8 items-center justify-center border-neutral-700/60 bg-neutral-800/60 text-white"
>
	{#if timeLeft > 0}
		Sesja wygaśnie w {clock}
	{/if}
</div> -->
