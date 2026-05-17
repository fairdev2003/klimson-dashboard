<script lang="ts">
	import type { Snippet } from 'svelte';

	type Props = {
		className?: string;
		children: Snippet;
	};

	const tooltip = 'bottomsdwdadwadaw';

	const { className = '', children }: Props = $props();
</script>

<!-- svelte-ignore slot_element_deprecated -->
<span class={`group relative cursor-help ${className}`}>
	<!-- Tooltip -->
	{@render children()}
</span>
