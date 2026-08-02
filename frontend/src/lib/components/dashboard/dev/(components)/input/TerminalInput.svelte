<script lang="ts">
	import Loader from '$lib/components/dashboard/Loader.svelte';
	import { terminal } from '$lib/terminal/logic';
	import TerminalPrefix from '../helpers/TerminalPrefix.svelte';

	type Props = {
		inputRef: HTMLInputElement | undefined;
		commandLineValue: string;
	};

	let { inputRef = $bindable(), commandLineValue = $bindable() }: Props = $props();

	let suggestionGhost = $derived.by(() => {
		const fullKey = terminal.intelisense.intelisenseValue || '';
		if (
			fullKey.toLowerCase().startsWith(commandLineValue.toLowerCase()) &&
			commandLineValue.length > 0
		) {
			return fullKey.slice(commandLineValue.length);
		}
		return '';
	});
</script>

<div class="flex gap-3 items-center border-b border-white/5 pb-1">
	<span class="text-neutral-600 flex whitespace-nowrap shrink-0 text-sm">
		<TerminalPrefix naming={terminal.terminal_naming} />
	</span>

	<div class="relative flex-1 flex items-center">
		<div
			class="absolute inset-0 flex items-center text-xs pointer-events-none select-none overflow-hidden"
		>
			<span class="text-transparent">{commandLineValue}</span>
			<span class="text-neutral-600">{suggestionGhost}</span>

			{#if !commandLineValue && !terminal.intelisense.intelisenseValue}
				<span class="text-neutral-600">Type '/' to proceed</span>
			{/if}
		</div>

		<input
			onfocusin={() => {
				terminal.inputFocused = true;
			}}
			onfocusout={() => {
				terminal.inputFocused = false;
			}}
			onkeydown={async (keyboard_event) => {
				await terminal.keyboard_event.onArrowRightClicked(keyboard_event);
				await terminal.keyboard_event.handleTerminalInputEnclosure(keyboard_event);
			}}
			oninput={(e) => {
				terminal.intelisense.searchForSyntaxAndReturnCommand(terminal.commandLineValue);
			}}
			spellcheck="false"
			bind:value={commandLineValue}
			bind:this={inputRef}
			class="bg-transparent w-full text-neutral-400 text-xs outline-none focus:ring-0 border-0 p-0 relative z-10 caret-neutral-400"
		/>
	</div>
</div>
