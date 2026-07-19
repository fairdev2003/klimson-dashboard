<script lang="ts">
	import Loader from '$lib/components/dashboard/Loader.svelte';
	import { terminal } from '$lib/terminal/logic';
	import TerminalPrefix from '../helpers/TerminalPrefix.svelte';

	type Props = {
		inputRef: HTMLInputElement | undefined;
		commandLineValue: string;
	};

	let { inputRef = $bindable(), commandLineValue = $bindable() }: Props = $props();
</script>

<div class="flex gap-3 items-center border-b border-white/5 pb-1">
	<span class="text-neutral-600 flex whitespace-nowrap shrink-0 text-sm">
		<TerminalPrefix naming={terminal.terminal_naming} />
	</span>

	<div class="relative flex-1">
		<input
			bind:value={commandLineValue}
			bind:this={inputRef}
			class="bg-transparent w-full text-neutral-400 text-xs outline-none focus:ring-0 border-0 p-0"
		/>

		{#if !commandLineValue}
			<p class="absolute top-0 left-0 pointer-events-none text-neutral-600 text-xs select-none">
				Type '/' to proceed
			</p>
		{/if}
	</div>
</div>
