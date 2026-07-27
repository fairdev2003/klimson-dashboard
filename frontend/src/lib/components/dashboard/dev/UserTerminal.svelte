<script lang="ts">
	import Dashboard from '$lib/dashboard/dashboard.svelte';
	import { debug, terminal } from '$lib/terminal/logic';
	import TerminalInput from './(components)/input/TerminalInput.svelte';
	import TerminalRecord from './(components)/TerminalRecord.svelte';
</script>

<div class="flex relative flex-col gap-1">
	{#each $debug as entry (entry.id)}
		<TerminalRecord naming={Dashboard.constants.TerminalNaming} {entry} />
	{/each}

	<TerminalInput
		bind:inputRef={terminal.inputRef}
		bind:commandLineValue={terminal.commandLineValue}
	/>
</div>
