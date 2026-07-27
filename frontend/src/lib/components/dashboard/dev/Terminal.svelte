<script lang="ts">
	import { debug } from '$lib/dashboard/stores/debug';
	import { onMount, tick } from 'svelte';
	import { debugOn } from '$lib/dashboard/stores/persist';

	import { terminal } from '$lib/terminal/logic';
	import TerminalButton from './TerminalButton.svelte';
	import TerminalContent from './TerminalContent.svelte';
	import TerminalHeader from './TerminalHeader.svelte';
	import { blur } from 'svelte/transition';

	$effect(() => {
		if ($debug && terminal.debugContainer) {
			tick().then(() => {
				terminal.debugContainer!.scrollTo({ top: terminal.debugContainer!.scrollHeight });
			});
		}
	});

	onMount(() => {
		terminal.loaded = true;
	});

	type Props = {
		standalone?: boolean;
	};

	let { standalone = false }: Props = $props();
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
{#if $debugOn && terminal.loaded}
	<div
		bind:this={terminal.boxRef}
		class={`${standalone ? 'flex' : 'fixed'} z-2000 flex ${terminal.isDragging ? '' : 'transition-all'} flex-col shadow-2xl`}
		style="left: {terminal.pos.x}px; bottom: {20 - terminal.pos.y}px; {terminal.isDragging
			? 'z-index: 1000'
			: ''}"
	>
		{#if terminal.terminalOpened}
			<TerminalHeader />
			<TerminalContent />
		{:else}
			<!-- sssssh -->
		{/if}
	</div>
{/if}

<svelte:document
	onkeydown={async (keyboard_event) => {
		if (terminal.commandLineValue.length === 0 && keyboard_event.key === '/') {
			await terminal.keyboard_event.onSlashKeyClick(keyboard_event);
		}
		if (keyboard_event.key === 'F2') {
			await terminal.keyboard_event.onF2KeyClicked(keyboard_event);
		}
		if (keyboard_event.key === 'F3') {
			await terminal.keyboard_event.onF3KeyClicked(keyboard_event);
		}
		if (keyboard_event.key === 'ArrowDown') {
			await terminal.keyboard_event.onArrowDownClicked(keyboard_event);
		}
		if (keyboard_event.key === 'ArrowUp') {
			await terminal.keyboard_event.onArrowUpClicked(keyboard_event);
		}
		if (keyboard_event.ctrlKey && keyboard_event.key === 'x') {
			await terminal.keyboard_event.onConsoleClearClicked();
		}
		if (keyboard_event.key === 'Enter') {
			await terminal.keyboard_event.onEnter();
		}
	}}
/>

<style>
	@import 'tailwindcss';

	div::-webkit-scrollbar {
		width: 4px;
	}
	div::-webkit-scrollbar-track {
		background: transparent;
	}
	div::-webkit-scrollbar-thumb {
		background: #262626;
		border-radius: 10px;
	}
	div::-webkit-scrollbar-thumb:hover {
		background: #3b82f6;
	}
</style>
