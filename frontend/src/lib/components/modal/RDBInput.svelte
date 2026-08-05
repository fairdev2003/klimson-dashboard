<script lang="ts">
	import { debug } from '$lib/terminal/logic';

	type Props = {
		value?: string;
		label?: string;
		oninput?: (e: string) => void;
		placeholder?: string;
	};

	let { value = $bindable(), label, oninput, placeholder }: Props = $props();
</script>

<div class="flex flex-col gap-1">
	<span class="text-neutral-400 uppercase justify-between items-center flex font-bold text-xs">
		<p>{label}</p>
	</span>

	<input
		bind:value
		{placeholder}
		class="rounded-lg placeholder:text-neutral-600 border-0 bg-neutral-800 p-2"
	/>
</div>
