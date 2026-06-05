<script lang="ts">
	import tippy from 'tippy.js';
	import type { Attachment } from 'svelte/attachments';

	import {
		dashboardLoaded,
		requestTimes,
		type TimeResponse
	} from '$lib/dashboard/stores/data.store';

	import Modal from '$lib/components/Modal.svelte';
	import { dev } from '$app/environment';
	import { api } from '$lib/api/api';
	import { base_url } from '$lib/api/api.store';
	import DiskSpaceWidget from './widgets/DiskSpaceWidget.svelte';
	import Icon from '@iconify/svelte';
	import { goto } from '$app/navigation';
	import DatabaseWidget from './widgets/DatabaseWidget.svelte';
	import SpotifyWidget from './widgets/SpotifyWidget.svelte';

	// s

	function tooltip(content: string): Attachment {
		return (element) => {
			const tooltip = tippy(element, { content });
			return tooltip.destroy();
		};
	}

	let loading: boolean = $state(false);

	let BenchmarkModalOpened: boolean = $state(false);

	let benchmarkKey: TimeResponse = $state('blogResponseTime');
</script>

<div class=" p-5 text-white flex flex-col gap-5">
	<!-- {#if dev}
		<p>Panel na dev</p>
	{:else}
		<p>Panel na prod</p>
	{/if}
	<p>
		Obecny serwer: <a
			class="text-blue-500 hover:underline"
			href={api.api_config.dev_server + '/v1/api'}>{$base_url}</a
		>
	</p>
	{#if dev}
		<p>
			Panel na serwerze produkcyjnym: <a
				class="text-blue-500 hover:underline"
				href={api.api_config.prod_front + '/dashboard'}
				>{api.api_config.prod_front + '/dashboard'}</a
			>
		</p>
	{/if} -->

	<div
		class="flex justify-between gap-4 p-5 w-full bg-neutral-800/60 border border-neutral-700 rounded-xl items-center"
	>
		<div class="flex gap-4 items-center">
			<img
				width="40"
				height="40"
				class="rounded-full"
				src="https://api.klimson.dev/storage/interface/random/banana.webp"
				alt="me"
			/>
			<h1 class="text-xl font-bold">Witaj, Jakub</h1>
		</div>
		<div class="flex">
			<button
				onclick={() => {
					goto('/dashboard/settings');
				}}
				class="p-2 hover:bg-neutral-700 rounded-xl cursor-pointer"
			>
				<Icon icon="material-symbols:settings" width="30" height="30" />
			</button>
		</div>
	</div>

	<div class="flex gap-5 lg:flex-row flex-col">
		<DiskSpaceWidget />
		<DatabaseWidget />
		<SpotifyWidget />
	</div>
</div>

<div class="grid grid-cols-5 p-5 gap-5"></div>

{@render BenchmarkAll()}

{#snippet BenchmarkAll()}
	<div class="flex max-w-[550px] flex-col gap-4 rounded-lg p-5"></div>
{/snippet}

<Modal
	onClose={() => {
		BenchmarkModalOpened = false;
	}}
	title={benchmarkKey}
	bind:opened={BenchmarkModalOpened}
	className="w-[400px]"
>
	{@render Benchmark(benchmarkKey)}
</Modal>

{#snippet Benchmark(key: TimeResponse)}
	{@const entry: Record<TimeResponse, number> = $requestTimes}
	<div class="flex max-w-[550px] flex-col gap-4 rounded-lg">
		<div class="flex flex-col gap-1">
			<div class="flex justify-end text-xs font-medium">
				<span
					class={entry[key] < 300
						? 'text-green-500'
						: entry[key] < 1800
							? 'text-orange-500'
							: 'text-red-500'}
				>
					{$dashboardLoaded ? entry[key] + 'ms' : '-'}
				</span>
			</div>
			{#if $dashboardLoaded}
				<div class="h-2 w-full overflow-hidden rounded-full bg-neutral-800">
					<div
						class="h-full transition-all duration-500 ease-out {entry[key] < 300
							? 'bg-green-500'
							: entry[key] < 1800
								? 'bg-orange-500'
								: 'animate-pulse bg-red-500'}"
						style="width: {Math.min((entry[key] / 2000) * 100, 100)}%"
					></div>
				</div>
			{:else}
				<div class="h-2 w-full animate-pulse overflow-hidden rounded-full bg-green-800"></div>
			{/if}
		</div>
	</div>
{/snippet}
