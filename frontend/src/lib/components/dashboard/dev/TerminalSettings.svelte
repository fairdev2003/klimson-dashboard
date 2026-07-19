<script lang="ts">
	import { slide } from 'svelte/transition';

	type Props = {
		opened: boolean;
	};

	let { opened = $bindable() }: Props = $props();
</script>

{#if opened}
	<div
		role="presentation"
		onmousedown={(e) => e.stopPropagation()}
		onmousemove={(e) => e.stopPropagation()}
		onclick={(e) => e.stopPropagation()}
		in:slide={{ duration: 300 }}
		out:slide={{ duration: 300 }}
		class="absolute p-2 right-0 active:cursor-auto top-10 z-10 w-3/5 h-50 bg-blue-800"
	>
		<p class="font-bold text-xs uppercase">Terminal window configuration</p>
	</div>
{/if}
