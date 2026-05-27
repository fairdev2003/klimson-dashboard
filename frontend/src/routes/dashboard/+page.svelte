<script lang="ts">
	import { dev } from '$app/environment';
	import { api } from '$lib/api/api';
	import { onMount } from 'svelte';
	import type { Quiz } from './quizzes/types';
	import type { ServerResponse } from '$lib/api/types';
	import tippy from 'tippy.js';
	import type { Attachment } from 'svelte/attachments';
	import { contributors_loading, latest_requests } from './contributors/vars';
	import { fade } from 'svelte/transition';
	import {
		dashboardLoaded,
		requestTimes,
		type TimeResponse
	} from '$lib/dashboard/stores/data.store';
	import Button from '$lib/components/Button.svelte';

	import { contextMenuOptions } from '$lib/dashboard/stores/store';
	import Modal from '$lib/components/Modal.svelte';
	import Dashboard from './dashboard.svelte';
	import { debug } from '$lib/dashboard/stores/debug';
	import { goto } from '$app/navigation';
	import TotalQuizzes from './widgets/TotalQuizzes.svelte';
	import TotalContributors from './widgets/TotalContributors.svelte';
	import TotalApiRoutes from './widgets/TotalApiRoutes.svelte';
	import DevWidget from './widgets/DevWidget.svelte';
	import ProdWidget from './widgets/ProdWidget.svelte';
	// s

	function tooltip(content: string): Attachment {
		return (element) => {
			const tooltip = tippy(element, { content });
			return tooltip.destroy();
		};
	}

	let loading: boolean = $state(false);

	onMount(() => {
		contextMenuOptions.update((context) => [
			{
				contextMenuName: 'Opcje benchamrku',
				items: [
					{
						label: 'Benchmark serwera',
						action: RefreshBenchmark,
						icon: 'catppuccin:benchmark',
						color: 'text-blue-500'
					}
				]
			},
			{
				contextMenuName: 'Inne',
				items: [
					{
						label: 'Przejdz do edycji quizów',
						action: () => goto('/dashboard/quizzes'),
						icon: 'tdesign:icon',
						color: 'text-blue-500'
					}
				]
			}
		]);
	});

	const toggleBenchmarkStatus = (isRunning: boolean) => {
		contextMenuOptions.update((options) =>
			options.map((group) => ({
				...group,
				items: group.items.map((item) => {
					if (item.label.includes('Benchmark')) {
						return {
							...item,
							label: isRunning ? 'Benchmarkuje...' : 'Benchmark serwera',
							color: isRunning ? 'text-orange-500' : 'text-blue-500',
							icon: isRunning ? 'line-md:loading-twotone-loop' : 'catppuccin:benchmark'
						};
					}
					return item;
				})
			}))
		);
	};

	async function RefreshBenchmark() {
		loading = true;
		debug.success('Git');
		toggleBenchmarkStatus(true);
		const response = await Dashboard.Load();
		if (response) {
			loading = false;
		}
		toggleBenchmarkStatus(false);
	}

	let BenchmarkModalOpened: boolean = $state(false);
	let benchmarkKey: TimeResponse = $state('blogResponseTime');
</script>

<!-- <div class=" p-5 text-white">
	{#if dev}
		<p>Panel na dev</p>
	{:else}
		<p>Panel na prod</p>
	{/if}
	<p>
		Obecny serwer: <a
			class="text-blue-500 hover:underline"
			href={api.api_config.dev_server + '/v1/api'}>{api.api_config.dev_server}</a
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
	{/if}
</div> -->

<div class="grid grid-cols-5 p-5 gap-5"></div>

{@render BenchmarkAll()}

{#snippet BenchmarkAll()}
	<div class="flex max-w-[550px] flex-col gap-4 rounded-lg p-5">
		<Button
			loadingText="Wykonuje benchmark..."
			size="small"
			theme="secondary"
			{loading}
			onclick={RefreshBenchmark}>Odśwież</Button
		>
	</div>
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
