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

<div class="flex items-center gap-2">
	<span class="justify-center flex items-center text-white">
		<TerminalPrefix {naming} />
	</span>
	<span class="text-gray-500 text-xs">
		<p class="ml-1 text-[13px]">
			{entry.metadata.command}
		</p>
	</span>
</div>
