<script lang="ts">
	import type { TerminalEntry } from '$lib/dashboard/stores/debug';
	import TerminalPrefix from '../TerminalPrefix.svelte';

	type Props = {
		entry: TerminalEntry;
	};

	let { entry }: Props = $props();
</script>

<div class="flex items-center gap-2">
	<span class="justify-center flex items-center text-white">
		<TerminalPrefix />
	</span>
	<span class="text-white">
		<p class="ml-1 text-[13px]">
			{entry.metadata.command}
		</p>
	</span>
</div>
