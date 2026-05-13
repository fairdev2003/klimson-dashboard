<script lang="ts">
	import type { Snippet } from 'svelte';
	import { tv } from 'tailwind-variants';

	const buttonStyles = tv({
		base: 'cursor-pointer p-3 rounded-md',

		variants: {
			theme: {
				base: 'bg-neutral-500'
			}
		}
	});

	type Props = {
		actionIcon: Snippet;
		tooltip?: string;
		theme: 'base';
	};

	const { actionIcon, tooltip, theme = 'base' }: Props = $props();
</script>

<button class={buttonStyles({ theme })}>
	{@render actionIcon()}
</button>
