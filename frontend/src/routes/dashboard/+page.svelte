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
	import DropdownButton from '$lib/components/dashboard/settings/components/DropdownButton.svelte';
	import {
		settings_page_open,
		type SettingKey
	} from '$lib/components/dashboard/settings/store.svelte';
	import { toast } from '$lib/dashboard/stores/toast';
	import CPUWidget from './widgets/CPUWidget.svelte';
	import WidgetContainer from './widgets/containers/WidgetContainer.svelte';
	import Heading from '$lib/components/dashboard/typography/Heading.svelte';
	import {
		dashboard_load_date,
		settings_startup_modal
	} from '$lib/components/dashboard/stores/main';
	import { onMount } from 'svelte';
	import SecondHubContainer from './widgets/containers/SecondHubContainer.svelte';
	import { dockComponent } from './dashboard.svelte';
	import BaseDockComponent from '$lib/components/dashboard/dock/boxes/BaseDockComponent.svelte';

	let loading: boolean = $state(false);

	let BenchmarkModalOpened: boolean = $state(false);

	let benchmarkKey: TimeResponse = $state('blogResponseTime');

	onMount(() => {
		dockComponent.set(BaseDockComponent);
	});
</script>

<div class=" p-5 text-white flex flex-col gap-5">
	<div class="my-2 flex justify-between px-2">
		<div class="flex flex-col">
			<Heading>Hub</Heading>
			<span class="text-sm font-md text-neutral-400">System load at: {$dashboard_load_date}</span>
		</div>
		<button
			onclick={() => {
				$settings_startup_modal = 'widgets';
				$settings_page_open = 'customization';
				goto('/dashboard/settings');
			}}
			class="p-2 hover:bg-neutral-700 rounded-xl cursor-pointer"
		>
			<Icon icon="boxicons:edit-filled" width="30" height="30" />
		</button>
	</div>

	<WidgetContainer />

	<SecondHubContainer />
</div>

{#snippet GradientLeft()}
	<div
		class="absolute -left-5 w-50 h-50 rounded-full blur-[60px] pointer-events-none -z-10 bg-linear-to-tl bg-linear-to-tb from-blue-700 to-blue-500"
	></div>
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

<svelte:head>
	<title>Dashboard Hub</title>
	<meta name="description" content="Panel Admina HarcQuiz" />
</svelte:head>
