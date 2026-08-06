<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<script lang="ts">
	import CMSNavbar from '$lib/components/dashboard/CMSNavbar.svelte';
	import DashboardLayout from '$lib/components/dashboard/DashboardLayout.svelte';
	import Loader from '$lib/components/dashboard/Loader.svelte';
	import type { AxiosError } from 'axios';
	import { dashboardLoadState } from '$lib/dashboard/stores/data.store';
	import Toast from '$lib/components/dashboard/Toast.svelte';
	import { goto } from '$app/navigation';
	import { dashboard_config, route } from '$lib/dashboard/stores/persist';

	import { page } from '$app/state';
	import { onMount } from 'svelte';
	import { toast } from '$lib/dashboard/stores/toast';
	import FancyLoader from './redis/(components)/FancyLoader.svelte';
	import Console from '$lib/components/dashboard/dev/Terminal.svelte';
	import Dashboard from '$lib/dashboard/dashboard.svelte';
	import { useDebounce } from '@ariefsn/svelte-use';
	import { debug } from '$lib/terminal/logic';
	import axios from 'axios';
	import { api } from '$lib/api/api';

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

	$effect(() => {
		const currentConfig = $dashboard_config;

		if (!currentConfig) return;

		const timer = setTimeout(async () => {
			try {
				const api_resp = await api.api.post(
					`/admin/redis/user-config/set?user_id=root`,
					currentConfig
				);

				if (api_resp.status === 200) {
					toast.success(api_resp.data.message || 'Config saved');
					debug.system(api_resp);
				}
			} catch (error) {
				if (axios.isAxiosError(error)) {
					toast.error(error.message);
				} else {
					toast.error('An unexpected error occurred');
				}
			}
		}, 500);

		return () => clearTimeout(timer);
	});
</script>

<svelte:head>
	<link
		rel="icon"
		href="https://api.klimson.dev/interface/bucket/blackout_bot/8107-admin-badge-blue.favicon"
	/>
	<meta name="description" content="Panel Admina HarcQuiz" />
</svelte:head>

<div class="h-full">
	<Toast />
	<Console />
	<div class="flex h-auto flex-col text-white">
		{#await Dashboard.Load()}
			<div class="mx-auto mt-10 flex items-center gap-3">
				<FancyLoader color="blue" />
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
