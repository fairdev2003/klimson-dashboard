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
	<span class="text-white">
		<p class="ml-1 text-[13px]">
			{@html entry.metadata.html}
		</p>
	</span>
</div>
