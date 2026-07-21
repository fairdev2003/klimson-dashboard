<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import gsap from 'gsap';
	import RDBInput from '$lib/components/modal/RDBInput.svelte';
	import { blur } from 'svelte/transition';
	import { userLogin } from '$lib/dashboard/stores/persist';
	import { goto } from '$app/navigation';
	import type { AxiosResponse } from 'axios';
	import { api } from '$lib/api/api';
	import Loader from '$lib/components/dashboard/Loader.svelte';
	import FancyLoader from '../../dashboard/redis/(components)/FancyLoader.svelte';
	import { load } from '../../dashboard/content_manager/+page';

	let timeString = $state('');
	let dateString = $state('');

	let loginPage: boolean = $state(false);

	function updateDateTime() {
		const now = new Date();

		const hh = now.getHours().toString().padStart(2, '0');
		const mm = now.getMinutes().toString().padStart(2, '0');
		const ss = now.getSeconds().toString().padStart(2, '0');
		timeString = `${hh}:${mm}:${ss}`;

		const dd = now.getDate().toString().padStart(2, '0');
		const mo = (now.getMonth() + 1).toString().padStart(2, '0');
		const yyyy = now.toLocaleDateString('pl-PL', {
			day: '2-digit',

			month: 'long',

			weekday: 'long'
		});
		dateString = yyyy;
	}

	let interval: ReturnType<typeof setInterval>;
	let dateContainer: HTMLDivElement | undefined;
	let loginContainer: HTMLDivElement | undefined;
	let inputRef: HTMLInputElement | undefined;

	let pass: string = $state('');
	let error: string = $state('');
	let loading: boolean = $state(false);

	async function handleLogin() {
		if (!$userLogin || !pass) return;

		try {
			const response = await api.api.post(
				'/login',
				{ password: pass, login: $userLogin },
				{ withCredentials: true }
			);

			if (response.status === 200) {
				loading = false;
				goto('/dashboard');
			}
		} catch (e: any) {
			loading = false;
			console.log('Pełny obiekt błędu:', e);

			if (e.response && e.response.data) {
				error = e.response.data.message || 'Wystąpił nieznany błąd';
			} else {
				error = 'Błąd połączenia z serwerem';
			}
		} finally {
			loading = false;
		}
	}

	onMount(() => {
		if (!$userLogin) {
			goto('/login');
		}

		if (!dateContainer || !loginContainer) return;

		gsap.fromTo(dateContainer, { opacity: 0.5 }, { opacity: 1, duration: 0.5 });

		updateDateTime();
		interval = setInterval(updateDateTime, 1000);
	});

	onDestroy(() => {
		clearInterval(interval);
	});
</script>

<div class="relative flex justify-center items-center w-screen h-screen overflow-hidden">
	<div class="absolute inset-0 background -z-10"></div>

	<div class="flex flex-col justify-between items-center mb-70 select-none">
		<div bind:this={dateContainer} class="flex flex-col justify-center items-center">
			<h1 class="text-white text-[130px] font-black lilita-one-regular">{timeString}</h1>
			<h1 class="text-white text-[20px] font-black lilita-one-regular">{dateString}</h1>
		</div>
		<div
			class="flex flex-col justify-center gap-y-5 items-center opacity-0"
			bind:this={loginContainer}
		>
			<img
				class="size-30 rounded-full"
				src="https://klimson.dev/_app/immutable/assets/explore_more.DVfVQDvY.png"
				alt="xD"
			/>
			<p class="font-thin text-2xl lilita-one-regular text-white">{$userLogin}</p>
			<div class="relative">
				<input
					bind:this={inputRef}
					bind:value={pass}
					placeholder="Password"
					onkeydown={(e) => {
						if (e.key === 'Enter') {
							loading = true;
							setTimeout(async () => {
								await handleLogin();
							}, 2000);
						}
					}}
					type="password"
					class="text-white lilita-one-regular bg-white/20 border-2 p-1 px-3 rounded-lg border-white/30"
				/>
				{#if loading}
					<div class="absolute -right-7 top-2 translate-y-1/2">
						<FancyLoader size="medium" color="white" />
					</div>
				{/if}
			</div>
		</div>
	</div>

	<div lang="flex justify-center items-center"></div>
</div>

<svelte:document
	onclick={() => {
		if (!dateContainer || !loginContainer) return;

		if (!loginPage) {
			loginPage = true;

			gsap.to(dateContainer, {
				y: -60,
				opacity: 0,

				duration: 0.25,
				ease: 'power2.out'
			});

			gsap.to(dateContainer, {
				scale: 0.6,
				duration: 0.25,
				ease: 'power2.out'
			});
			gsap.fromTo(
				loginContainer,
				{
					y: 60,
					opacity: 0,
					scale: 0,
					duration: 0.25,
					ease: 'power2.out'
				},
				{
					y: 0,
					opacity: 1,
					scale: 1,
					duration: 0.25,
					ease: 'power2.out',
					onComplete: () => {
						if (!inputRef) return;

						inputRef.focus();
					}
				}
			);
		}
	}}
	onkeydown={async (e) => {
		if (e.key === 'Escape') {
			loginPage = false;
			if (!dateContainer || !loginContainer) return;

			gsap.to(dateContainer, {
				scale: 1,
				duration: 0.25,
				ease: 'power2.out'
			});
			gsap.to(dateContainer, {
				y: 0,
				opacity: 1,

				duration: 0.25,
				ease: 'power2.out'
			});

			gsap.to(loginContainer, { opacity: 0, y: 60, scale: 0, ease: 'power2.out' });
		}
	}}
/>

<style>
	@import 'tailwindcss';

	.background {
		@apply bg-neutral-950 blur-sm bg-[url('../../../lib/assets/background2.svg')] bg-no-repeat bg-cover;
	}

	.lilita-one-regular {
		font-family: 'Lilita One', sans-serif;
		font-weight: 900;
		font-style: normal;
	}
</style>
