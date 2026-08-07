<script lang="ts">
	import { goto } from '$app/navigation';
	import type { Snippet } from 'svelte';
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = {
		children: Snippet;
		padding?: boolean;
		idTrackingDisabled?: boolean;
	} & HTMLAttributes<HTMLHeadingElement>;

	const { padding = false, idTrackingDisabled = false, children, ...props }: Props = $props();
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<h1
	{...props}
	id={props.id}
	class={`heading text-primary group relative text-2xl font-bold ${props.class} ${padding && 'pt-5 '} font-bold`}
>
	<span>{@render children()}</span>
</h1>

<style>
</style>
