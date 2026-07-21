<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import gsap from 'gsap';
	import RDBInput from '$lib/components/modal/RDBInput.svelte';
	import { blur, slide } from 'svelte/transition';
	import { userLogin } from '$lib/dashboard/stores/persist';
	import { goto } from '$app/navigation';
	import type { AxiosResponse } from 'axios';
	import { api } from '$lib/api/api';
	import Loader from '$lib/components/dashboard/Loader.svelte';
	import FancyLoader from '../../dashboard/redis/(components)/FancyLoader.svelte';
	import { load } from '../../dashboard/content_manager/+page';
	import Button from '$lib/components/Button.svelte';
	import { spotifyApp } from '$lib/components/spotify/spotify.svelte';
	import Icon from '@iconify/svelte';

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
	let spotifyContainer: HTMLDivElement | undefined;
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

	onMount(async () => {
		await spotifyApp.zamontujKurwe();
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

	const slugs = ['background1', 'background2', 'background3', 'background4'] as const;
	type BackgroundSlug = (typeof slugs)[number];
	type Background = { name: string; slug: string; url: string };

	const backgrounds: Background[] = [
		{
			name: 'Archipelac',
			slug: 'background1',
			url: 'https://api.klimson.dev/interface/bucket/backgrounds/background1.webp'
		},
		{
			name: 'Meteor Rain',
			slug: 'background2',
			url: 'https://api.klimson.dev/interface/bucket/backgrounds/background2.svg'
		},

		{
			name: 'Arch linux',
			slug: 'background3',
			url: 'https://api.klimson.dev/interface/bucket/backgrounds/background3.webp'
		},

		{
			name: 'Modded MC Base',
			slug: 'background4',
			url: 'https://api.klimson.dev/interface/bucket/backgrounds/background4.webp'
		},
		{
			name: 'Sernik',
			slug: 'background5',
			url: 'https://api.klimson.dev/interface/bucket/backgrounds/background5.webp'
		}
	];
	let currentWallpaper: Background | undefined = $state(
		backgrounds.find((e) => e.slug === 'background1')
	);
	let nextWallpaper: Background | undefined = $state();

	type SingleBackground = (typeof backgrounds)[number];
	let selected_background: BackgroundSlug = $state('background3');

	let currentBgDiv: HTMLDivElement | undefined;
	let nextBgDiv: HTMLDivElement | undefined;

	function changeBackground(backgroundSlug: string) {
		const targetBg = backgrounds.find((e) => e.slug === backgroundSlug);
		if (!targetBg || targetBg.slug === currentWallpaper?.slug) return;

		selected_background = targetBg.slug;
		nextWallpaper = targetBg;

		if (!nextBgDiv || !currentBgDiv) return;

		gsap.fromTo(
			nextBgDiv,
			{ opacity: 0 },
			{
				opacity: 1,
				duration: 1,
				ease: 'power2.inOut',
				onComplete: () => {
					currentWallpaper = targetBg;
					nextWallpaper = undefined;
					gsap.set(nextBgDiv, { opacity: 0 });
				}
			}
		);
	}

	let backgroundSelectionMenuOpened = $state(false);
</script>

<div class="relative flex justify-center items-center w-screen h-screen overflow-hidden">
	<div
		bind:this={currentBgDiv}
		class="absolute background inset-0 -z-10"
		style="background-image: url('{currentWallpaper?.url}'); opacity: 1;"
	></div>

	<div
		bind:this={nextBgDiv}
		class="absolute background inset-0 -z-10"
		style="background-image: url('{nextWallpaper?.url}'); opacity: 0;"
	></div>

	<div class="flex flex-col justify-between items-center mb-70 select-none">
		<div bind:this={dateContainer} class="flex flex-col justify-center items-center">
			<h1 class="text-white text-[130px] font-black lilita-one-regular">{timeString}</h1>
			<h1 class="text-white text-[20px] font-black lilita-one-regular">{dateString}</h1>
		</div>
		<div
			class="flex flex-col justify-center gap-y-5 items-center opacity-0"
			bind:this={loginContainer}
			style="border-radius: 8px;"
		>
			<!-- <img
				class="size-30 rounded-full"
				src="https://klimson.dev/_app/immutable/assets/explore_more.DVfVQDvY.png"
				alt="xD"
			/> -->
			<div
				class="size-30 border-2 justify-center items-center flex rounded-full bg-white/40 border-white/50"
			>
				<Icon icon="mdi:user" class="size-22 text-white/60" />
			</div>
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

	<div class="fixed hidden m-4 bottom-0 left-0 flex flex-col gap-4">
		{#each backgrounds as _bg}
			<button
				class="bg-white/10 p-2 px-4 rounded-xl lilita text-white hover:bg-white/20 transition-colors"
				onclick={(e) => {
					e.stopPropagation();
					changeBackground(_bg.slug);
				}}>{_bg.name}</button
			>
		{/each}
	</div>
	<div class="group fixed m-4 bottom-0 left-0 flex flex-col gap-4">
		{#if backgroundSelectionMenuOpened}
			<!-- svelte-ignore a11y_click_events_have_key_events -->
			<div
				in:slide={{ duration: 300 }}
				class="h-120 w-100 bg-selection-container bg-white/20 rounded-2xl items-center gap-2 flex flex-col p-3 overflow-y-auto"
				onclick={(e) => {
					e.stopPropagation();
				}}
			>
				{#each backgrounds as _bg}
					<img
						onclick={() => {
							changeBackground(_bg.slug);
						}}
						src={_bg.url}
						class="w-1/2 aspect-video object-cover h-1/2 hover:opacity-90 cursor-pointer"
					/>
				{/each}
			</div>
		{/if}
		<button
			class="bg-white/20 hover:bg-white/30 cursor-pointer size-12 rounded-full items-center flex justify-center text-white/90"
			><Icon
				icon="mdi-light:chevron-up"
				width="50"
				height="50"
				onclick={(e) => {
					e.stopPropagation();
					backgroundSelectionMenuOpened = !backgroundSelectionMenuOpened;
				}}
				onmouseup={() => {
					console.log('MOUSE IS DOW');
				}}
				class="mb-1.5 ml-0.5"
			/></button
		>
	</div>

	{#if spotifyApp.spotify?.item?.name}
		<div
			bind:this={spotifyContainer}
			class="spotify fixed bottom-15 left-1/2 px-6 p-4 flex-col -translate-x-1/2 bg-white/40 rounded-2xl gap-2 flex min-w-100"
		>
			<div
				class="flex gap-4 items-center justify-between"
				onclick={(e) => {
					e.stopPropagation();
				}}
			>
				<div class="flex gap-4">
					<img src={spotifyApp.getAlbumCover()} class="rounded-xl size-15" alt="album_cover" />
					<div class="flex gap-1 justify-center flex-col">
						<p class="lilita text-white/90 truncate max-w-65">{spotifyApp.getSong()}</p>
						<p class="lilita text-xs text-white/70 truncate max-w-65">{spotifyApp.getArtist()}</p>
					</div>
				</div>

				<div class="controls flex gap-2 text-white">
					<Icon icon="raphael:arrowleft" class="size-10" />

					<Icon icon="material-symbols:pause" class="size-10" />
					<Icon icon="raphael:arrowright" class="size-10" />
				</div>
			</div>
		</div>
	{/if}
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
			gsap.to(spotifyContainer, {
				y: 60,
				opacity: 0,

				duration: 0.25,
				ease: 'power2.out'
			});

			gsap.to(spotifyContainer, {
				scale: 0.6,
				duration: 0.25,
				ease: 'power2.out'
			});
			gsap.fromTo(
				loginContainer,
				{
					y: 60,
					opacity: 0,
					scale: 0.6,
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

			gsap.to(spotifyContainer, {
				scale: 1,
				duration: 0.25,
				ease: 'power2.out'
			});
			gsap.to(spotifyContainer, {
				y: 0,
				opacity: 1,

				duration: 0.25,
				ease: 'power2.out'
			});
			gsap.to(loginContainer, {
				opacity: 0,
				duration: 0.35,
				y: 60,
				scale: 0.6,
				ease: 'power2.out'
			});
		}
	}}
/>

<style>
	@import 'tailwindcss';

	/* Stylowanie scrollbara dla przeglądarek opartych na Chromium (Chrome, Edge, Safari) */
	.bg-selection-container::-webkit-scrollbar {
		width: 8px; /* szerokość scrollbara */
		display: none;
	}

	.background4 {
		@apply bg-[url('../../../lib/assets/background4.webp')];
	}

	.background3 {
		@apply bg-[url('../../../lib/assets/background3.webp')];
	}

	.background2 {
		@apply bg-[url('../../../lib/assets/background2.svg')];
	}

	.background1 {
		@apply bg-[url('../../../lib/assets/background1.webp')];
	}

	.background {
		@apply bg-neutral-950 blur-md bg-no-repeat bg-cover;
	}

	.lilita-one-regular,
	.lilita {
		font-family: 'Lilita One', sans-serif;
		font-weight: 900;
		font-style: normal;
	}
</style>
