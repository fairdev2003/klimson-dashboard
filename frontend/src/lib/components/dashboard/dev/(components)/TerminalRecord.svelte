<script lang="ts">
	import type { TerminalEntry } from '$lib/dashboard/stores/debug';
	import type { TerminalNaming } from '../console/terminal.svelte';

	import {
		ConsoleRecord,
		ErrorRecord,
		MessageRecord,
		RawRecord,
		SilentRecord,
		SystemRecord,
		WarningRecord,
		FormatRecord,
		ImageRecord
	} from '$lib/terminal/components';

	type Props = {
		entry: TerminalEntry;
		naming: TerminalNaming | undefined;
	};

	let { entry, naming }: Props = $props();

	const recordMap: Record<string, any> = {
		system: SystemRecord,
		warn: WarningRecord,
		error: ErrorRecord,
		message: MessageRecord,
		console: ConsoleRecord,
		silent: SilentRecord,
		raw: RawRecord,
		format: FormatRecord,
		image: ImageRecord
	};

	let Component = $derived(recordMap[entry.type]);
</script>

{#if Component}
	<Component {naming} {entry} />
{/if}
