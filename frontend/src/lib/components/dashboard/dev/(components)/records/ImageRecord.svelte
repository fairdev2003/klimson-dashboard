<script lang="ts">
	import { goto } from '$app/navigation';
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
	<DatePrefix color="bg-gray-700/50" date={entry.date} {naming} />

	<!-- svelte-ignore a11y_click_events_have_key_events -->
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->

	<span class="text-neutral-300">
		<img
			onclick={() => {
				if (!entry.metadata.src) return;
				window.open(entry.metadata.src, '_blank', 'noopener,noreferrer');
			}}
			src={entry.metadata.src}
			class="size-1/3 cursor-pointer hover:opacity-90"
			alt={`image-${entry.id}`}
		/>
	</span>
</div>
