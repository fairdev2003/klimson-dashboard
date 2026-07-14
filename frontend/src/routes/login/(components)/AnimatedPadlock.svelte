<script lang="ts">
	import { onMount } from 'svelte';
	import { blur } from 'svelte/transition';
	import { gsap } from 'gsap';
	import FancyLoader from '../../dashboard/redis_writable/(components)/FancyLoader.svelte';

	let {
		padlock_loading = $bindable(),
		opened = $bindable(false)
	}: { padlock_loading: boolean; opened?: boolean } = $props();

	let key = $state(0);
	let shackleRef: HTMLDivElement;

	$effect(() => {
		const tl = gsap.timeline();

		if (opened) {
			// Animacja otwieraniax
			tl.to(shackleRef, {
				y: -10,
				duration: 0.2,
				ease: 'power2.out'
			}).to(shackleRef, {
				y: -12,
				x: 4,
				rotation: 15,
				transformOrigin: 'bottom left',
				duration: 0.3,
				ease: 'back.out(2)'
			});
		} else {
			tl.to(shackleRef, {
				y: 0,
				x: 0,
				rotation: 0,
				duration: 0.3,
				ease: 'power2.inOut'
			});
		}
	});
</script>

<div
	in:blur={{ duration: 300 }}
	class="flex flex-col items-center justify-center w-8 h-10 relative"
>
	<div
		bind:this={shackleRef}
		class="w-4 h-4 border-[2px] border-white border-b-0 rounded-t-md -mb-1 top-0 absolute"
	></div>

	<div class="w-5 h-5 bg-white rounded-sm flex items-center justify-center relative z-10"></div>
</div>
