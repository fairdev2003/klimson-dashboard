<script lang="ts">
	import Campfire from '$lib/assets/campfire-loading.svg';
	import gsap from 'gsap';
	import { onMount, tick } from 'svelte';
	import Page from '../../routes/+page.svelte';

	let CampfireElement: HTMLImageElement | undefined = $state(undefined);
	let loadingOn: boolean = $state(true);

	onMount(async () => {
		await tick();

		gsap.to(CampfireElement || '', {
			scale: 1.5,
			duration: 1,
			repeat: -1,
			yoyo: true,
			ease: 'power1.inOut'
		});
	});
</script>

{#if loadingOn}
	<div
		class="z-100 fixed inset-0 flex flex-col items-center justify-center overflow-hidden bg-black/80"
	>
		<img
			bind:this={CampfireElement}
			src={Campfire}
			class="size-15 select-none"
			alt="campfire-loading"
		/>
	</div>
{/if}
