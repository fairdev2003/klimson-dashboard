<script lang="ts">
	import type { TerminalEntry } from '$lib/dashboard/stores/debug';
	import type { TerminalNaming } from '../../console/terminal.svelte';
	import DatePrefix from '../helpers/DatePrefix.svelte';
	import TerminalPrefix from '../helpers/TerminalPrefix.svelte';

	type Props = {
		entry: TerminalEntry;
		naming: TerminalNaming | undefined;
	};

	let { entry, naming }: Props = $props();
</script>

<div class="flex gap-3 items-center leading-relaxed last:border-0">
	<DatePrefix color="bg-yellow-400/50" date={entry.date} {naming} />

	<span class="text-yellow-400">
		<p class="ml-1 text-[13px]">
			{entry.metadata.message}
		</p>
	</span>
</div>
