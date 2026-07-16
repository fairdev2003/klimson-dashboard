<script lang="ts">
	import type { TerminalEntry } from '$lib/dashboard/stores/debug';
	import TerminalPrefix from '../TerminalPrefix.svelte';

	type Props = {
		entry: TerminalEntry;
	};

	let { entry }: Props = $props();
</script>

<div class="flex gap-3 items-center leading-relaxed last:border-0">
	<span class="text-neutral-600 whitespace-nowrap shrink-0">
		{#if entry.metadata.command}
			<TerminalPrefix />
		{/if}

		{#if entry.metadata.message}
			[{new Date(entry.date).toLocaleTimeString()}]
		{/if}
	</span>
	<span class="text-blue-400">
		<p class="ml-1 text-[13px]">
			{entry.metadata.message}

			{entry.metadata.command}
		</p>
	</span>
</div>

{#snippet Old()}
	<div class="flex gap-3 items-center pb-1 leading-relaxed last:border-0">
		<span class="text-neutral-600 whitespace-nowrap w-20 shrink-0">
			{#if entry.metadata.command}
				<TerminalPrefix />
			{/if}

			{#if entry.metadata.message}
				[{new Date(entry.date).toLocaleTimeString()}]
			{/if}
		</span>
		<span
			class={entry.type === 'error'
				? 'text-red-400'
				: entry.type === 'warn'
					? 'text-yellow-400'
					: entry.type === 'success'
						? 'text-green-400 '
						: entry.type === 'system'
							? 'text-blue-400'
							: 'text-neutral-300'}
		>
			<p class="ml-1 text-[13px]">
				{entry.metadata.message}

				{entry.metadata.command}
			</p>
		</span>
	</div>
{/snippet}
