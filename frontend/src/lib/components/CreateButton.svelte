<script lang="ts">
	import type { Snippet } from 'svelte';
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = {
		children?: Snippet;
	} & HTMLAttributes<HTMLButtonElement>;

	const { children, ...props }: Props = $props();
</script>

<button class="cursor-pointer bg-blue-700 p-2 px-4 hover:bg-blue-600" {...props}>
	{#if children}
		{@render children()}
	{:else}
		{@render PlaceholderText()}
	{/if}
</button>

{#snippet PlaceholderText()}
	<p>Button</p>
{/snippet}
