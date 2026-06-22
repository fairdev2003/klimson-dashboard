<script lang="ts">
	import { onMount } from 'svelte';
	import gsap from 'gsap';
	import Icon from '@iconify/svelte';

	let pages = ['Major Section', 'Minor Section'];
	let pageIndex = $state(0);

	let textRef: HTMLParagraphElement;

	function animateText(direction: 'next' | 'prev') {
		const tl = gsap.timeline();

		// 1. Wyjazd starego tekstu
		tl.to(textRef, {
			x: direction === 'next' ? -50 : 50,
			opacity: 0,
			duration: 0.2,
			ease: 'power2.in'
		});

		// 2. Zmiana indeksu (wewnątrz timelinii, żeby idealnie trafić w moment niewidoczności)
		tl.add(() => {
			if (direction === 'next') {
				pageIndex = (pageIndex + 1) % pages.length;
			} else {
				pageIndex = (pageIndex - 1 + pages.length) % pages.length;
			}
		});

		// 3. Wjazd nowego tekstu
		tl.fromTo(
			textRef,
			{ x: direction === 'next' ? 50 : -50, opacity: 0 },
			{ x: 0, opacity: 1, duration: 0.3, ease: 'power2.out' }
		);
	}

	function nextPage() {
		animateText('next');
	}
	function prevPage() {
		animateText('prev');
	}
</script>

<div
	class="relative px-2 items-center flex justify-between w-full h-10 mt-4 rounded-lg text-blue-400 bg-blue-500/20"
>
	<button onclick={prevPage} class="p-2 rounded-lg cursor-pointer hover:bg-blue-500/40">
		<Icon icon="material-symbols:chevron-left" />
	</button>
	<p bind:this={textRef} class="uppercase text-xs font-black">{pages[pageIndex]}</p>
	<button onclick={nextPage} class="p-2 rounded-lg cursor-pointer hover:bg-blue-500/40">
		<Icon icon="material-symbols:chevron-right" />
	</button>
</div>
