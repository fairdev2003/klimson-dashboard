<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<script lang="ts">
	import f from '$lib/assets/dashboard.png';
	import CMSNavbar from '$lib/components/dashboard/CMSNavbar.svelte';
	import DashboardLayout from '$lib/components/dashboard/DashboardLayout.svelte';
	import Loader from '$lib/components/dashboard/Loader.svelte';
	import type { AxiosError } from 'axios';
	import Dashboard from './dashboard.svelte';
	import { dashboardLoadState } from '$lib/dashboard/stores/data.store';
	import Toast from '$lib/components/dashboard/Toast.svelte';
	import Debug from '$lib/components/dashboard/Debug.svelte';
	import { goto } from '$app/navigation';
	import { fade, scale } from 'svelte/transition';
	import { latest_requests } from './contributors/vars';
	import { developerView, route } from '$lib/dashboard/stores/persist';
	import { toast } from '$lib/dashboard/stores/toast';
	import { api } from '$lib/api/api';

	import { page } from '$app/state';
	import { onMount } from 'svelte';

	onMount(() => {
		goto($route);
	});

	$effect(() => {
		$route = page.url.pathname;
	});

	let { children } = $props();

	function redirectTo(e: HTMLDivElement) {
		setTimeout(() => {
			goto('/login');
			console.log('Przeniesiono');
		}, 10000);
	}

	let pos = $state({ x: 20, y: 20 });
	let isDragging = $state(false);

	// Funkcja obsługująca ruch
	function handleMouseDown(e: MouseEvent) {
		isDragging = true;

		const onMouseMove = (moveEvent: MouseEvent) => {
			if (isDragging) {
				// Obliczamy nową pozycję (odejmujemy trochę, by kursor był na środku nagłówka)
				pos.x += moveEvent.movementX;
				pos.y += moveEvent.movementY;
			}
		};

		const onMouseUp = () => {
			isDragging = false;
			window.removeEventListener('mousemove', onMouseMove);
			window.removeEventListener('mouseup', onMouseUp);
		};

		window.addEventListener('mousemove', onMouseMove);
		window.addEventListener('mouseup', onMouseUp);
	}
</script>

<svelte:head>
	<link rel="icon" href={f} />
	<title>Panel Admina</title>
	<meta name="description" content="Panel Admina HarcQuiz" />
</svelte:head>

<div class="h-full">
	<Toast />
	<Debug />
	<div class="flex h-auto flex-col text-white">
		{#await Dashboard.Load()}
			<div class="mx-auto mt-10 flex items-center gap-3">
				<Loader />
				<p>{$dashboardLoadState}</p>
			</div>
		{:then access}
			{#if access}
				{#if $developerView}
					{@render ApiDebug()}
				{/if}
				<CMSNavbar />
				<DashboardLayout>
					{@render children()}
				</DashboardLayout>
			{:else}
				<div>Brak dostepu</div>
			{/if}
		{:catch error: AxiosError}
			<div in:fade={{ duration: 150 }} use:redirectTo class="flex flex-col gap-2 bg-red-500/60 p-2">
				{#if error.code === 'ERR_BAD_REQUEST'}
					<p class="text-[13px]">ODPOWIEDZ Z SERWERA - przeniesnie nastapi w 10 sekund</p>
					<p class="text-[13px]">Błąd: {error.response?.data.error as string}</p>
					{#if error.response?.data.required}
						<p class="text-[13px]">Brak: {error.response?.data.required as string}</p>
					{/if}
				{/if}
			</div>
		{/await}
	</div>
</div>

{#snippet ApiDebug()}
	<div
		class="fixed z-[2000] w-72 overflow-hidden rounded-lg border border-neutral-800 bg-neutral-900 p-0 shadow-2xl select-none"
		style="left: {pos.x}px; top: {pos.y}px; transition: {isDragging ? 'none' : 'all 0.1s'}"
	>
		<div
			role="button"
			tabindex="0"
			class="flex cursor-grab items-center justify-between bg-neutral-800 px-4 py-2 active:cursor-grabbing"
			onmousedown={handleMouseDown}
		>
			<h3 class="text-[10px] font-bold tracking-widest text-neutral-400 uppercase">API Monitor</h3>
			<div class="flex gap-1">
				<div class="h-2 w-2 animate-pulse rounded-full bg-green-500"></div>
			</div>
		</div>

		<div class="flex max-h-64 flex-col gap-2 overflow-y-auto bg-black/20 p-4">
			{#each Object.entries($latest_requests) as [url, time]}
				<div class="flex flex-col border-b border-neutral-800 pb-2 last:border-0">
					<span
						class="cursor-pointer truncate font-mono text-[10px] text-blue-400 hover:underline"
						onclick={() => {
							navigator.clipboard.writeText(api.api_config.baseURL + url.replace('/', ''));
							toast.success('Skopiowano do schowka');
						}}>{url}</span
					>
					<div class="mt-1 flex items-center justify-between">
						<span class="text-[9px] text-neutral-500 italic">Duration:</span>
						<span
							class="text-[10px] font-bold {parseInt(time) > 500
								? 'text-red-500'
								: 'text-orange-500'}"
						>
							{time}
						</span>
					</div>
				</div>
			{/each}

			{#if Object.keys($latest_requests).length === 0}
				<span class="text-center text-[10px] text-neutral-600 italic">Oczekiwanie na ruch...</span>
			{/if}
		</div>
	</div>
{/snippet}
