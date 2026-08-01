<script lang="ts">
	import { debug, type TerminalEntry } from '$lib/dashboard/stores/debug';
	import { onMount } from 'svelte';
	import type { TerminalNaming } from '../../console/terminal.svelte';
	import TerminalPrefix from '../helpers/TerminalPrefix.svelte';
	import type { KlimsonFetchType } from '$lib/types/stats';
	import { api } from '$lib/api/api';

	type Props = {
		entry: TerminalEntry;
	};

	let { entry }: Props = $props();
</script>

<div class="flex items-center gap-2">
	<div class="text-gray-500 text-xs flex gap-4">
		<div class="">
			<img
				aria-hidden="true"
				src="https://api.klimson.dev/interface/bucket/random/banana.webp"
				alt="fastfetch-image"
				class="size-50"
			/>
		</div>
		{#if entry.metadata.fastfetch}
			<div class="flex flex-col">
				<span class="text-blue-500 font-bold">Server Info</span>
				<div class="w-full h-4"></div>

				{#each Object.keys(entry.metadata.fastfetch) as name}
					{@render FastfetchRecord(name, entry.metadata.fastfetch[name])}
				{/each}
			</div>
		{/if}
	</div>
</div>

{#snippet FastfetchRecord(key: string, value: any)}
	<div class="flex gap-1 text-sm">
		<span class="text-blue-500 font-bold">{key}: </span>

		{#if key.includes('time') && key !== 'uptime'}
			<p>
				{new Date(value).toLocaleDateString('en-US', {
					year: 'numeric',
					month: 'numeric',
					day: 'numeric',
					minute: '2-digit',
					hour: '2-digit'
				})}
			</p>
		{:else if key.includes('uptime')}
			<p>
				{(() => {
					const totalSeconds = Number(value);
					const hours = Math.floor(totalSeconds / 3600);
					const minutes = Math.floor((totalSeconds % 3600) / 60);
					return `${hours}h ${minutes}m`;
				})()}
			</p>
		{:else}
			<p>{value}</p>
		{/if}
	</div>
{/snippet}
