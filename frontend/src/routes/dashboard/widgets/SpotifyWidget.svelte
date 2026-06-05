<script lang="ts">
	import { goto } from '$app/navigation';
	import { spotifyApp } from '$lib/components/spotify/spotify.svelte';
	import Icon from '@iconify/svelte';
	import { onMount } from 'svelte';
	import { blur } from 'svelte/transition';

	onMount(async () => {
		await spotifyApp.zamontujKurwe();
	});
</script>

{#key spotifyApp.spotify?.item?.name}
	<button
		in:blur={{ duration: 300 }}
		onclick={() => {
			goto('/dashboard/spotify');
		}}
		class="relative overflow-hidden cursor-pointer text-start group rounded-xl flex flex-col h-45 max-w-100 w-full md:w-70 lg:w-70 border gap-3 hover:ring-green-500 hover:ring-2 border-neutral-700 bg-neutral-800/60"
	>
		<div class="absolute group-hover:flex hidden w-full h-full bg-green-500/20"></div>
		<div class="p-5">
			<div class="flex gap-3">
				<img
					src={spotifyApp.getAlbumCover()}
					width="50"
					height="50"
					class="rounded-full"
					alt="album_cover"
				/>
				<div class="flex relative flex-col text-start gap3 justify-center w-full">
					<h3 class="text-sm font-semibold truncate">{spotifyApp.getSong()}</h3>
					<p class="text-xs text-neutral-400 font-semibold">{spotifyApp.getArtist()}</p>
					<div class="absolute right-5 top-1/2 translate-[50%]">
						{@render AudioVisualizer()}
					</div>
				</div>
			</div>
			<div>
				<div class="flex justify-between mt-4 mx-12">
					<Icon icon="fluent:previous-20-filled" width="30" height="3s0" />

					<Icon icon="material-symbols:pause" width="30" height="30" />
					<Icon icon="fluent:next-20-filled" width="30" height="30" />
				</div>
				<div class={`h-1 rounded-full bg-white/30 w-full mt-4`}>
					<div
						class="h-full rounded-full bg-green-400 transition-all duration-500 ease-linear"
						style="width: {(spotifyApp.progress / spotifyApp.duration) * 100}%"
					></div>
				</div>

				<div class="flex justify-between mt-2 relative">
					<p class="text-[11px] text-white">{spotifyApp.formatMs(spotifyApp.progress)}</p>
					<p class="text-[11px] text-white">{spotifyApp.formatMs(spotifyApp.duration)}</p>
				</div>
				<div
					class="absolute group-hover:flex hidden bottom-0 px-5 border-b-0 right-1/2 translate-x-1/2 rounded-t-lg border-2 p-1 border-neutral-700 bg-neutral-800"
				>
					<p class="text-xs text-white font-semibold">{spotifyApp.spotify?.device.name}</p>
				</div>
			</div>
		</div>
	</button>
{/key}

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
	.visualizer {
		display: flex;
		align-items: flex-end;
		height: 10px;
		scale: 60%;
	}

	.bar {
		width: 3px;
		background-color: #1ed760;
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
</style>
