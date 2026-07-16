<script lang="ts">
	import type { TerminalEntry } from '$lib/dashboard/stores/debug';
	import type { TerminalNaming } from '../console/terminal.svelte';
	import ConsoleRecord from './records/ConsoleRecord.svelte';
	import ErrorRecord from './records/ErrorRecord.svelte';
	import MessageRecord from './records/MessageRecord.svelte';
	import SilentRecord from './records/SilentRecord.svelte';
	import SystemRecord from './records/SystemRecord.svelte';
	import WarningRecord from './records/WarningRecord.svelte';

	type Props = {
		entry: TerminalEntry;
		naming: TerminalNaming | undefined;
	};

	let { entry, naming }: Props = $props();
</script>

{#if entry.type === 'system'}
	<SystemRecord {naming} {entry} />
{/if}
{#if entry.type === 'warn'}
	<WarningRecord {naming} {entry} />
{/if}
{#if entry.type === 'error'}
	<ErrorRecord {naming} {entry} />
{/if}

{#if entry.type === 'message'}
	<MessageRecord {naming} {entry} />
{/if}
{#if entry.type === 'console'}
	<ConsoleRecord {naming} {entry} />
{/if}
{#if entry.type === 'silent'}
	<SilentRecord {naming} {entry} />
{/if}
