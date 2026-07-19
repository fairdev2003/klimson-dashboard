<script lang="ts">
	import { debug, terminal } from '$lib/terminal/logic';
	import TerminalInput from './(components)/input/TerminalInput.svelte';
	import TerminalRecord from './(components)/TerminalRecord.svelte';
</script>

<div class="flex relative flex-col gap-1">
	{#each $debug as entry (entry.id)}
		<TerminalRecord naming={{ name: 'klimson-dashboard', path: 'main' }} {entry} />
	{/each}
</div>
