<script lang="ts">
	import { base_url } from '$lib/api/api.store';
	import Heading from '$lib/components/dashboard/typography/Heading.svelte';
	import Icon from '@iconify/svelte';
	import { onMount } from 'svelte';
	let cpu = $state(0);
	let cpuUsage = $derived(Math.floor(cpu));

	onMount(() => {
		const statement = base_url == 'https://api.klimson.dev';
		const replace = statement ? 'https://' : 'http://';
		const websocket = statement ? 'wss://' : 'ws://';

		const socket = new WebSocket(`${$base_url.replace(replace, websocket)}/ws/stats/cpu`);

		socket.onmessage = (event) => {
			const data = JSON.parse(event.data);
			cpu = data.cpu;
		};

		return () => socket.close();
	});
</script>

<button
	class="relative overflow-hidden cursor-pointer text-start group rounded-xl flex flex-col h-45 max-w-100 w-full md:w-70 lg:w-70 border gap-3 hover:ring-purple-400 hover:ring-2 border-neutral-700 bg-neutral-800/60"
>
	<div class="absolute group-hover:flex hidden w-full h-full bg-purple-500/20"></div>

	<div class="flex flex-col gap-3 p-5">
		<div class="flex gap-4 items-center">
			<div class="flex gap-4 items-center">
				<Icon icon="solar:cpu-bold" width="40" height="40" />
				<Heading>CPU Usage</Heading>
			</div>
		</div>

		<div class="flex justify-between gap-1 items-center">
			<h2 class="text-3xl font-bold">{cpuUsage}%</h2>
		</div>
		<div></div>
	</div>
</button>
