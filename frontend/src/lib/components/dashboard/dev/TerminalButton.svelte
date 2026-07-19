<script lang="ts">
	import { debug } from '$lib/terminal/logic';
	import { Terminal } from '@lucide/svelte';

	type Props = {
		onclick?: () => void;
	};

	let {
		onclick = () => {
			debug.warn("Terminal doesnt have 'onclick' event");
		}
	}: Props = $props();
</script>

<button
	class="group flex items-center gap-2 rounded-full border border-neutral-800 bg-neutral-950 p-2 pr-4 shadow-xl transition-all hover:border-blue-500/50 active:scale-95"
	{onclick}
>
	<div class="rounded-full bg-blue-500/20 p-2 transition-colors group-hover:bg-blue-500">
		<Terminal class="h-4 w-4 text-blue-500 group-hover:text-white" />
	</div>
	<span class="text-xs font-bold tracking-tighter text-neutral-400 uppercase">Terminal</span>
</button>
