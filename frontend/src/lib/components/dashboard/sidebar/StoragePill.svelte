<script lang="ts">
	import { api } from '$lib/api/api';
	import { base_url } from '$lib/api/api.store';
	import type { DiskData } from '$lib/api/requests/misc';
	import { debug } from '$lib/dashboard/stores/debug';
	import Icon from '@iconify/svelte';
	import { onMount } from 'svelte';

	let percent: number = $state(0);
	let power = $derived(Math.floor(percent * 10));
	let disk: DiskData | undefined = $state();
	let usedGb = $derived(formatBytes(disk?.used ?? 0));
	let totalGb = $derived(formatBytes(disk?.total ?? 0));
	let loadingRoute = $state(false);

	onMount(async () => {
		const response = await api.misc.GetDisk();
		disk = response.data;
		percent = Number(disk.percentage) / 100;

		gsap.to('.blob-item', {
			backgroundColor: (i) => (i < power ? '#22c55e' : '#525252'),
			duration: 1,
			stagger: 0.1,
			ease: 'power1.inOut'
		});
	});

	function formatBytes(bytes: number | string, decimals = 2) {
		const b = Number(bytes);
		if (b === 0) return '0 GB';

		const gb = b / (1024 * 1024 * 1024);

		return `${gb.toFixed(decimals)} GB`;
	}
</script>

<div
	class="bg-neutral-800 rounded-lg gap-2 flex flex-col w-full p-3 hover:bg-neutral-700 transition-colors cursor-pointer"
>
	<div class="flex gap-2 items-center">
		<Icon icon="entypo:drive" width="17" height="17" />
		<p class="text-sm font-black">
			{$base_url === 'https://api.klimson.dev' ? 'Mikrus Storage' : 'Local Storage'}
		</p>
	</div>
	<div class="w-full h-0.5 bg-neutral-500">
		<div
			style="width: {Number(percent) * 100}%;"
			class="h-0.5 bg-green-500 transition-all delay-150 duration-500"
		></div>
	</div>
	<div>
		<p class="text-xs text-neutral-400">
			{usedGb} / {totalGb}
		</p>
	</div>
</div>
