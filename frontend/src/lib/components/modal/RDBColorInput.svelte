<script lang="ts">
	import { debug } from '$lib/terminal/logic';

	type Props = {
		value?: string;
		label?: string;
	};

	let { value = $bindable(), label }: Props = $props();
</script>

<div class="flex flex-col gap-1">
	<span class="text-neutral-400 uppercase justify-between items-center flex font-bold text-xs">
		<p>{label}: <input class="text-white text-xs p-0 bg-transparent w-14 border-0" bind:value /></p>
	</span>

	<input
		type="color"
		bind:value
		class="rounded-lg hover:bg-neutral-700 transition-colors cursor-pointer border-0 bg-neutral-800 p-1"
	/>
</div>
