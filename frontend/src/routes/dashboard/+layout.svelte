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
	import { route } from '$lib/dashboard/stores/persist';

	import { page } from '$app/state';
	import { onMount } from 'svelte';

	$effect(() => {
		$route = page.url.pathname;
	});

	onMount(() => {
		const preventZoom = (e: any) => {
			if (e.touches.length > 1) {
				e.preventDefault();
			}
		};

		document.addEventListener('touchstart', preventZoom, { passive: false });
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
				<CMSNavbar />
				<DashboardLayout>
					{@render children()}
				</DashboardLayout>
			{:else}
				<div>Brak dostepu</div>
			{/if}
		{/await}
	</div>
</div>
