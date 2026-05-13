<script lang="ts">
	import { onMount } from 'svelte';
	import gsap from 'gsap';
	import { goto } from '$app/navigation';

	let activeIndex = 0;

	let container: HTMLDivElement;
	let indicator: HTMLDivElement;

	function setActive(index: number, el: HTMLDivElement) {
		activeIndex = index;

		const rect = el.getBoundingClientRect();
		const parentRect = container.getBoundingClientRect();

		const x = rect.left - parentRect.left;
		const width = el.offsetWidth;

		gsap.to(indicator, { duration: 0.1, x, width, ease: 'power2.out' });
	}

	onMount(() => {
		const first = container.querySelector('div') as HTMLDivElement;
		if (first) setActive(0, first);

		const handleResize = () => {
			const activeEl = container.children[activeIndex] as HTMLDivElement;
			if (activeEl) setActive(activeIndex, activeEl);
		};

		window.addEventListener('resize', handleResize);
		return () => window.removeEventListener('resize', handleResize);
	});

	type Props = {
		items?: { name: string; path: string }[];
		onclick?: () => string;
	};

	const {
		items = [
			{ name: 'Start', path: '/dashboard' },
			{ name: 'Quizy', path: '/dashboard/quizzes' },
			{ name: 'Pytania', path: '/dashboard/questions' },
			{ name: 'Odpowiedzi', path: '/dashboard/answers' },
			{ name: 'Scieżki API', path: '/dashboard/routes' },
			{ name: 'Blog', path: '/dashboard/blog' }
		],
		onclick
	}: Props = $props();
</script>

<div
	bind:this={container}
	class="relative flex gap-4 overflow-x-auto whitespace-nowrap"
	style="scroll-behavior: smooth;"
>
	{#each items as item, i}
		<div
			class="inline-block cursor-pointer px-2 py-2 text-sm font-medium text-white"
			on:click={(e) => {
				setActive(i, e.currentTarget as HTMLDivElement);
			}}
		>
			{item.name}
		</div>
	{/each}

	<div
		bind:this={indicator}
		class="absolute bottom-0 h-[2px] bg-blue-500 transition-all duration-300"
		style="width:0; transform:translateX(0px);"
	></div>
</div>
