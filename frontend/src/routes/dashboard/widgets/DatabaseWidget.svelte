<script lang="ts">
	import { goto, preloadData } from '$app/navigation';
	import { api } from '$lib/api/api';
	import Loader from '$lib/components/dashboard/Loader.svelte';
	import Heading from '$lib/components/dashboard/typography/Heading.svelte';
	import Icon from '@iconify/svelte';
	import { onMount } from 'svelte';
	import { blur } from 'svelte/transition';

	let loadingRoute = $state(false);
	let databaseCount = $state(0);

	onMount(async () => {
		const response = await api.misc.GetTables();

		databaseCount = response.data.tables.length;
	});
</script>

<button
	onclick={async () => {
		const route = '/dashboard/database';

		loadingRoute = true;
		await preloadData(route);

		setTimeout(() => {
			loadingRoute = false;

			goto(route);
		}, 100);
	}}
	class="relative overflow-hidden cursor-pointer text-start group rounded-xl flex flex-col h-45 max-w-100 w-full md:w-70 lg:w-70 border gap-3 hover:ring-blue-400 hover:ring-2 border-widget-border bg-widget-background"
>
	<div class="absolute group-hover:flex hidden w-full h-full bg-blue-500/20"></div>
	<div class:justify-between={loadingRoute} class="flex flex-col gap-3 p-5">
		<div class="flex gap-4 items-center">
			<div class="flex gap-4 items-center">
				<Icon icon="material-symbols:database" width="40" height="40" />
				<Heading>Databases</Heading>
			</div>
			{#if loadingRoute}
				<Loader theme="regular" />
			{/if}
		</div>

		<div class="flex text-start justify-between gap-1 items-center">
			{#key databaseCount}
				<h2 in:blur={{ duration: 300, delay: 300 }} class="text-3xl font-bold">{databaseCount}</h2>
			{/key}
		</div>
		<div></div>
	</div>
</button>
