<script lang="ts">
	import { gsap } from 'gsap';
	import { onMount, tick } from 'svelte';
	import { spotifyApp } from './spotify.svelte';
	import Icon from '@iconify/svelte';
	import Button from '../Button.svelte';

	onMount(async () => {
		await spotifyApp.zamontujKurwe();
	});

	$effect(() => {
		const spotifySong = spotifyApp.spotify;
		if (!spotifySong && modalEl) {
			closeModal();
		}
	});

	let svgPathEl = $state<HTMLOrSVGElement | null>(null);
	let flashBorderEl = $state<HTMLDivElement | null>(null);

	let currentTrackId = $state<string | null>(null);

	function startGlowAnimation() {
		requestAnimationFrame(() => {
			if (!svgPathEl || !flashBorderEl) return;

			const pathLength = (svgPathEl as any).getTotalLength();

			gsap.killTweensOf([svgPathEl, flashBorderEl]);

			gsap.set(svgPathEl, {
				strokeDasharray: pathLength,
				strokeDashoffset: pathLength,
				opacity: 1
			});
			gsap.set(flashBorderEl, { opacity: 0, scale: 1 });

			const tl = gsap.timeline();

			tl.to(svgPathEl, {
				strokeDashoffset: 0,
				duration: 0.9,
				ease: 'linear'
			})
				.to(
					flashBorderEl,
					{
						opacity: 1,
						scale: 1.005,
						duration: 0.1,
						repeat: 3,
						yoyo: true,
						ease: 'sine.inOut'
					},
					'-=0.1'
				)
				.to(svgPathEl, {
					opacity: 0,
					duration: 0.3,
					onComplete: () => {
						gsap.set(flashBorderEl, { opacity: 0, scale: 1 });
					}
				});
		});
	}

	$effect(() => {
		const trackId = spotifyApp.spotify?.item?.id || spotifyApp.getSong();

		if (trackId && trackId !== currentTrackId && svgPathEl) {
			currentTrackId = trackId;
			startGlowAnimation();
		}
	});

	type Props = {
		responsiveState?: 'desktop' | 'mobile';
	};

	let musicInfoModalOpenState: boolean = $state(false);
	let modalEl: HTMLDivElement;
	let modalContentEl: HTMLDivElement;

	function closeModal() {
		gsap.to(modalContentEl, {
			opacity: 0,
			scale: 0.9,
			duration: 0.25,
			ease: 'power2.in',
			onComplete: () => {
				musicInfoModalOpenState = false;
				document.body.style.overflow = 'auto';
			}
		});

		gsap.to(modalEl, {
			opacity: 0,
			duration: 0.25,
			ease: 'power2.in'
		});
	}

	async function openModal() {
		document.body.style.overflow = 'hidden';
		musicInfoModalOpenState = true;

		await tick();

		gsap.fromTo(modalEl, { opacity: 0 }, { opacity: 1, duration: 0.3, ease: 'power2.in' });

		gsap.fromTo(
			modalContentEl,
			{ opacity: 0, scale: 0.9 },
			{
				opacity: 1,
				scale: 1,
				delay: 0.3,
				duration: 0.4,
				ease: 'back.out(1.7)'
			}
		);
	}

	const { responsiveState = 'desktop' }: Props = $props();

	let containerEl: HTMLDivElement | undefined = $state();
	let textEl: HTMLDivElement | undefined = $state();
	let currentSong: string | null = $state('');

	let containerWidth = $derived(containerEl?.clientWidth);
	let textWidth = $derived(textEl?.clientWidth);

	let tween: gsap.core.Tween | null = null;

	onMount(() => {
		checkOverflowAndAnimate();
	});

	$effect(() => {
		const song = spotifyApp.getSong();
		if (song !== currentSong) {
			currentSong = song;
			requestAnimationFrame(() => {
				checkOverflowAndAnimate();
			});
		}
	});

	function checkOverflowAndAnimate() {
		if (!containerEl || !textEl) return;

		if (tween) {
			tween.kill();
			tween = null;
		}
		gsap.set(textEl, { x: 0 });

		const containerWidth = containerEl.clientWidth;
		const textWidth = textEl.scrollWidth;

		if (textWidth > containerWidth) {
			const distance = textWidth - containerWidth + 12;

			tween = gsap.to(textEl, {
				x: -distance,
				duration: distance / 35,
				ease: 'sine.inOut',
				repeat: -1,
				yoyo: true,
				delay: 1.2,
				repeatDelay: 1.2
			});
		}
	}
</script>

{#if spotifyApp.spotify && spotifyApp.getSong()}
	<div class="flex flex-col gap-2">
		<div
			class="relative justify-between bg-widget-background flex mx-auto p-2 items-center gap-2 overflow-visible rounded-lg shadow-lg object-center transition-all select-none w-full max-w-110"
		>
			<div class="flex gap-2 items-center">
				<span class="text-spotify flex items-center justify-center shrink-0">
					<Icon icon="mdi:spotify" width="20" height="20" class="mb-1" />
				</span>
				<p class="text-neutral-200 font-medium leading-none">Im listening to spotify</p>
			</div>
			<div
				class="p-1 px-2 bg-neutral-800 items-center text-spotify transition-colors rounded-full shadow-lg flex gap-1"
			>
				<Icon icon="material-symbols:devices" width="14" height="14" />
				<p class="text-xs">{spotifyApp.spotify.device.name}</p>
			</div>
		</div>
		<!-- svelte-ignore a11y_click_events_have_key_events -->
		<!-- svelte-ignore a11y_no_static_element_interactions -->
		<div
			onclick={() => {
				openModal();
			}}
			class="relative group bg-widget-background flex flex-col mx-auto p-4 cursor-pointer items-center gap-2 overflow-visible rounded-lg shadow-lg object-center transition-all select-none hover:bg-neutral-800 active:bg-neutral-700/70 w-full m-4"
		>
			<svg
				class="absolute inset-0 size-full pointer-events-none overflow-visible z-5"
				width="100%"
				height="100%"
			>
				<rect
					x="1"
					y="1"
					width="calc(100% - 2px)"
					height="calc(100% - 2px)"
					rx="8"
					ry="8"
					class="fill-none stroke-[1px]"
				/>

				<rect
					bind:this={svgPathEl}
					x="1"
					y="1"
					width="calc(100% - 2px)"
					height="calc(100% - 2px)"
					rx="8"
					ry="8"
					class="fill-none stroke-spotify stroke-[2px] opacity-0 [stroke-linecap:round] [filter:drop-shadow(0_0_4px_var(--color-spotify))_drop-shadow(0_0_8px_var(--color-spotify))]"
				/>
			</svg>

			<div
				bind:this={flashBorderEl}
				class="absolute inset-0 z-6 rounded-lg border-2 border-spotify opacity-0 pointer-events-none [box-shadow:0_0_20px_rgba(34,197,94,0.5)]"
			></div>
			<div class="flex gap-2 w-full items-center">
				<div class="relative size-12 shrink-0 shadow-lg">
					<span class="text-spotify rounded-full absolute bg-neutral-800 -bottom-1 -right-1 z-10">
						<Icon icon="mdi:spotify" width="24" height="24" />
					</span>
					<img
						src={spotifyApp.getAlbumCover()}
						alt="Album cover"
						class="rounded-lg border-3 border-neutral-800 size-12 object-cover"
					/>
				</div>

				<div class="relative justify-between flex flex-col w-full gap-2">
					<div>
						<div
							bind:this={containerEl}
							class="w-70 lg:w-full overflow-hidden whitespace-nowrap relative"
						>
							<div
								bind:this={textEl}
								class="inline-block font-black will-change-transform text-white"
							>
								{spotifyApp.getSong()}
							</div>
						</div>
						<div class="text-xs">{spotifyApp.getArtist()}</div>
					</div>
				</div>
			</div>
			<div class="flex items-center gap-4 px-2 w-full">
				<p class="text-xs shrink-0 text-neutral-400">{spotifyApp.formatMs(spotifyApp.progress)}</p>
				<div class="flex-1">
					{@render ProgressBar()}
				</div>
				<p class="text-xs shrink-0 text-neutral-400">{spotifyApp.formatMs(spotifyApp.duration)}</p>
			</div>
		</div>
	</div>
{/if}

{#if musicInfoModalOpenState}
	{@render Modal()}
{/if}

{#snippet ProgressBar(className?: string)}
	{#if spotifyApp.duration > 0}
		<div class={`h-1 w-full rounded-full bg-white/30 ${className}`}>
			<div
				class="h-full rounded-full bg-spotify transition-all duration-500 ease-linear"
				style="width: {(spotifyApp.progress / spotifyApp.duration) * 100}%"
			></div>
		</div>
	{/if}
{/snippet}

{#snippet Modal()}
	<!-- svelte-ignore a11y_click_events_have_key_events -->
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<div
		onclick={() => closeModal()}
		class={`fixed inset-0 z-2000 flex items-center justify-center overflow-hidden bg-neutral-950/90 backdrop-blur-md ${responsiveState === 'mobile' && 'pb-10'}`}
	>
		<div
			onclick={(e) => {
				e.stopPropagation();
			}}
			bind:this={modalEl}
			class="white lg:w-150 w-9/10 border border-neutral-700/60 bg-neutral-800/60 bg-gradient-to-r shadow-2xl backdrop-blur-3xl"
		>
			<div
				bind:this={modalContentEl}
				class="modal-content relative flex w-full gap-5 p-6 text-white"
			>
				{#if spotifyApp.getSong() && spotifyApp.getAlbumCover()}
					<img class="size-20 rounded-lg" src={spotifyApp.getAlbumCover()!} alt="Album cover" />
				{/if}
				<div class="relative flex w-full flex-col">
					{#if spotifyApp.getSong()}{/if}
					{#if spotifyApp.getSong()}
						<!-- svelte-ignore a11y_missing_attribute -->
						<a class="flex cursor-pointer items-center gap-1 font-bold">
							{spotifyApp.getSong()}
						</a>
					{/if}

					{#if spotifyApp.getSong()}
						<p class="font-sm mb-1 text-sm text-gray-400">
							{spotifyApp.getArtist()?.replaceAll(';', ',')}
						</p>
					{/if}
					{#if spotifyApp.getSong()}
						<div>
							{@render ProgressBar()}
						</div>
					{/if}
					{#if spotifyApp.getSong()}
						<div class="mt-1 flex justify-between font-semibold">
							<p class="text-[11px]">{spotifyApp.formatMs(spotifyApp.progress)}</p>
							<p class="text-[11px]">{spotifyApp.formatMs(spotifyApp.duration)}</p>
						</div>
					{/if}
				</div>
			</div>
		</div>
	</div>
{/snippet}

<svelte:window
	onkeydown={(a) => {
		if (a.key === 'Escape') {
			closeModal();
		}
	}}
/>

{#snippet AudioVisualizer()}
	<div class="visualizer flex gap-1">
		<div class="bar"></div>
		<div class="bar"></div>
		<div class="bar"></div>
		<div class="bar"></div>
		<div class="bar"></div>
		<div class="bar"></div>
	</div>
{/snippet}

<style>
	a {
		color: white;
		background-color: transparent;
		padding: 0;
	}
	.visualizer {
		display: flex;
		align-items: flex-end;
		height: 10px;
		scale: 60%;
	}

	.bar {
		width: 3px;
		background-color: var(--color-spotify);
		border-radius: 2px;

		transform-origin: bottom;

		animation: dance 0.5s infinite alternate ease-in-out;
	}

	.bar:nth-child(1) {
		height: 25px;
		animation-delay: 0.1s;
	}
	.bar:nth-child(2) {
		height: 25px;
		animation-delay: 0.2s;
	}
	.bar:nth-child(3) {
		height: 15px;
		animation-delay: 0.3s;
	}
	.bar:nth-child(4) {
		height: 10px;
		animation-delay: 0.1s;
	}
	.bar:nth-child(5) {
		height: 30px;
		animation-delay: 0.2s;
	}
	.bar:nth-child(6) {
		height: 20px;
		animation-delay: 0.3s;
	}

	@keyframes dance {
		from {
			transform: scaleY(0.3);
		}
		to {
			transform: scaleY(1);
		}
	}

	.mask-gradient {
		mask-image: linear-gradient(to right, transparent 0%, black 5%, black 95%, transparent 100%);
		-webkit-mask-image: linear-gradient(
			to right,
			transparent 0%,
			black 5%,
			black 95%,
			transparent 100%
		);
	}
</style>
