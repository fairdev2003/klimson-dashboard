<script lang="ts">
	import { base_url } from '$lib/api/api.store';
	import Heading from '$lib/components/dashboard/typography/Heading.svelte';
	import { Dashboard } from '$lib/dashboard/logic';
	import { debug } from '$lib/dashboard/stores/debug';
	import Icon from '@iconify/svelte';
	import { onMount } from 'svelte';
	import { blur } from 'svelte/transition';
	import UsageChart from './components/UsageChart.svelte';
	let cpu = $state(0);
	let cpuUsage = $derived(Math.floor(cpu));
	let os: string = $state('OS_ID');

	let ws_connection_opened = $state(false);

	type CPUObjectType = {
		cpu: number;
		arch: string;
		os: string;
	};

	onMount(() => {
		const statement = $base_url == 'https://api.klimson.dev';
		const replace = statement ? 'https://' : 'http://';
		const websocket = statement ? 'wss://' : 'ws://';

		const socket = new WebSocket(`${$base_url.replace(replace, websocket)}/ws/stats/cpu`);

		socket.onmessage = (event) => {
			const data: CPUObjectType = JSON.parse(event.data);
			cpu = data.cpu;
		};

		socket.onopen = (event) => {
			ws_connection_opened = true;
		};

		return () => socket.close();
	});

	let opened = $state(false);
</script>

<button
	onclick={() => {
		opened = !opened;
	}}
	class="relative overflow-hidden cursor-pointer text-start group rounded-xl flex flex-col h-45 max-w-100 w-full md:w-70 lg:w-70 border gap-3 hover:ring-purple-400 hover:ring-2 border-widget-border bg-widget-background"
>
	<div class="absolute group-hover:flex hidden w-full h-full bg-purple-500/20"></div>
	{#if ws_connection_opened && cpu > 0}
		<div in:blur={{ duration: 300 }} class="flex flex-col gap-3 p-5">
			<div class="flex gap-4 items-center">
				<div class="flex gap-4 items-center">
					<Icon icon="solar:cpu-bold" width="40" height="40" />
					<Heading>CPU Usage</Heading>
				</div>
			</div>

			<div class="flex flex-col justify-between items-start gap-3">
				<div class="h-2 w-full mt-4 bg-foreground rounded-full overflow-hidden relative">
					<div
						class="h-full rounded-full transition-all duration-700 ease-out relative shadow-[0_0_10px_rgba(var(--color-bg),0.5)]"
						class:bg-cpu-good={cpuUsage <= 40}
						class:bg-cpu-warn={cpuUsage > 40 && cpuUsage <= 80}
						class:bg-cpu-bad={cpuUsage > 80}
						style="width: {cpuUsage}%"
					>
						<div class="absolute inset-0 bg-white/20 animate-pulse"></div>
					</div>
				</div>
				<h2 class="text-3xl font-bold">{cpuUsage}%</h2>
			</div>
			<div></div>
		</div>
	{/if}
</button>

<UsageChart bind:opened />
