<script lang="ts">
	import type { TerminalEntry } from '$lib/dashboard/stores/debug';
	import ConsoleRecord from './records/ConsoleRecord.svelte';
	import ErrorRecord from './records/ErrorRecord.svelte';
	import MessageRecord from './records/MessageRecord.svelte';
	import SystemRecord from './records/SystemRecord.svelte';
	import WarningRecord from './records/WarningRecord.svelte';

	type Props = {
		entry: TerminalEntry;
	};

	let { entry }: Props = $props();
</script>

{#if entry.type === 'system'}
	<SystemRecord {entry} />
{/if}
{#if entry.type === 'warn'}
	<WarningRecord {entry} />
{/if}
{#if entry.type === 'error'}
	<ErrorRecord {entry} />
{/if}

{#if entry.type === 'message'}
	<MessageRecord {entry} />
{/if}
{#if entry.type === 'console'}
	<ConsoleRecord {entry} />
{/if}
