<script lang="ts">
	import Heading from '$lib/components/dashboard/typography/Heading.svelte';
	import Icon from '@iconify/svelte';
	import { onMount } from 'svelte';
	import { derived } from 'svelte/store';
	import gsap from 'gsap';
	import { api } from '$lib/api/api';
	import type { DiskData } from '$lib/api/requests/misc';
	import { debug } from '$lib/dashboard/stores/debug';
	import { goto, preloadData } from '$app/navigation';
	import { base_url } from '$lib/api/api.store';
	import { blur } from 'svelte/transition';
	import Loader from '$lib/components/dashboard/Loader.svelte';
	import { dashboard_config } from '$lib/dashboard/stores/persist';

	let percent: number = $state(0);
	let power = $derived(Math.floor(percent * 10));
	let disk: DiskData | undefined = $state();
	let usedGb = $derived(formatBytes(disk?.used ?? 0));
	let totalGb = $derived(formatBytes(disk?.total ?? 0));
	let loadingRoute = $state(false);
	let os = $state('');

	onMount(async () => {
		const response = await api.misc.GetDisk();
		disk = response.data;
		os = response.data.os;
		percent = Number(disk.percentage) / 100;

		const rootStyles = getComputedStyle(document.documentElement);

		const primaryColor = rootStyles.getPropertyValue('--color-primary').trim();
		const mutedColor = rootStyles.getPropertyValue('--color-foreground').trim();

		gsap.to('.blob-item', {
			backgroundColor: (i) => (i < power ? primaryColor : mutedColor),
			duration: 1,
			stagger: 0.1,
			ease: 'power1.inOut'
		});
	});

	$effect(() => {
		let config = $dashboard_config ?? $dashboard_config;

		const rootStyles = getComputedStyle(document.documentElement);

		const primaryColor = rootStyles.getPropertyValue('--color-primary').trim();
		const mutedColor = rootStyles.getPropertyValue('--color-foreground').trim();

		gsap.to('.blob-item', {
			backgroundColor: (i) => (i < power ? primaryColor : mutedColor),
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

<button
	onclick={async () => {
		const route = '/dashboard/storage';
		loadingRoute = true;
		await preloadData(route);

		setTimeout(() => {
			loadingRoute = false;

			goto(route);
		}, 100);
	}}
	class="relative overflow-hidden cursor-pointer text-start group rounded-xl flex flex-col h-45 max-w-100 w-full md:w-70 lg:w-70 border gap-3 hover:ring-green-400 hover:ring-2 border-widget-border bg-widget-background"
>
	<div class="absolute group-hover:flex hidden w-full h-full bg-green-500/20"></div>
	<div class="flex flex-col gap-3 p-5">
		<div class:justify-between={loadingRoute} class="flex gap-4 items-center">
			<div class="flex gap-4 items-center">
				<Icon icon="entypo:drive" width="40" height="40" />
				<Heading>Disk</Heading>
			</div>
			{#if loadingRoute}
				<Loader theme="regular" />
			{:else}
				<p class="text-xs mt-1 font-semibold text-primary">
					{os}
				</p>
			{/if}
		</div>

		<div class="flex justify-between gap-1 items-center">
			<div class="grid grid-cols-10 h-10 gap-1 w-50 lg:w-3/5">
				{#each Array(10) as _, i}
					<div class="blob-item bg-neutral-600 rounded-full col-span-1 h-10"></div>
				{/each}
			</div>
		</div>
		<div>
			{#key usedGb}
				<p in:blur={{ duration: 450 }} class="text-secondary-text text-md font-semibold">
					{usedGb} / {totalGb}
				</p>
			{/key}
		</div>
	</div>
</button>

<style>
	@import 'tailwindcss';

	.powered {
		@apply bg-green-500 transition-colors duration-300;
	}
</style>
