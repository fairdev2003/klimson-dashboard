<script lang="ts">
	import { debug } from '$lib/terminal/logic';
	import { blur } from 'svelte/transition';
	import TerminalInput from './(components)/input/TerminalInput.svelte';
	import TerminalRecord from './(components)/TerminalRecord.svelte';
	import { terminal } from './console/terminal.svelte';
	import HttpTerminalLogger from './HttpTerminalLogger.svelte';
	import TerminalSettings from './TerminalSettings.svelte';
	import UserTerminal from './UserTerminal.svelte';
	import { tick } from 'svelte';
	import LogTerminal from './LogTerminal.svelte';

	$effect(() => {
		if ($debug && terminal.debugContainer) {
			tick().then(() => {
				terminal.debugContainer!.scrollTo({ top: terminal.debugContainer!.scrollHeight });
			});
		}
	});
</script>

<div
	bind:this={terminal.debugContainer}
	class:w-300={terminal.fullscreen}
	class:h-150={terminal.fullscreen}
	class:w-200={!terminal.fullscreen}
	class:h-100={!terminal.fullscreen}
	class="flex relative flex-col gap-1 overflow-y-auto bg-neutral-900/95 p-4 font-mono text-[11px] backdrop-blur-md"
>
	{#if terminal.terminalPage === 'user'}
		<UserTerminal />
	{/if}
	{#if terminal.terminalPage === 'only-logs'}
		<LogTerminal />
	{/if}
	{#if terminal.terminalPage === 'http'}
		<HttpTerminalLogger />
	{/if}
</div>
