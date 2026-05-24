<script lang="ts">
	import { onMount } from 'svelte';
	import { gsap } from 'gsap';

	// Referencje do elementów dla GSAP
	let svgPathEl = $state<SVGInterfaceElement | null>(null);
	let flashBorderEl = $state<HTMLDivElement | null>(null);

	onMount(() => {
		if (!svgPathEl || !flashBorderEl) return;

		const pathLength = (svgPathEl as any).getTotalLength();

		gsap.set(svgPathEl, {
			strokeDasharray: pathLength,
			strokeDashoffset: pathLength
		});

		const tl = gsap.timeline();

		tl.to(svgPathEl, {
			strokeDashoffset: 0, // Wężyk płynnie okrąża kwadrat
			duration: 0.9,
			ease: 'linear'
		}).to(
			flashBorderEl,
			{
				opacity: 1,
				scale: 1.01, // Bardzo delikatne powiększenie podczas błysku
				duration: 0.1,
				repeat: 3, // Szybkie mignięcie 3 razy
				yoyo: true,
				ease: 'sine.inOut',
				onComplete: () => {
					gsap.set(flashBorderEl, { opacity: 0, scale: 1 });
					gsap.to(svgPathEl, { opacity: 0, duration: 0.3 });
				}
			},
			'-=0.1'
		); // Nakładamy błysk delikatnie na końcówkę animacji wężyka
	});
</script>

<div class="flex min-h-screen items-center justify-center bg-neutral-950">
	<div class="relative size-64 select-none rounded-3xl">
		<div
			class="absolute inset-0 z-10 flex flex-col items-center justify-center rounded-3xl bg-neutral-900 text-white shadow-xl"
		>
			<span class="text-4xl">🟩</span>
			<h3 class="mt-4 font-bold tracking-wide">Spotify Box</h3>
			<p class="text-sm text-neutral-400">GSAP Border Effect</p>
		</div>

		<svg
			class="absolute inset-0 z-20 size-full pointer-events-none overflow-visible"
			viewBox="0 0 256 256"
		>
			<rect
				x="1"
				y="1"
				width="254"
				height="254"
				rx="24"
				ry="24"
				class="fill-none stroke-neutral-700/60 stroke-[1px]"
			/>

			<rect
				bind:this={svgPathEl}
				x="1"
				y="1"
				width="254"
				height="254"
				rx="24"
				ry="24"
				class="fill-none stroke-green-500 stroke-[2px] [stroke-linecap:round] [filter:drop-shadow(0_0_4px_#22c55e)_drop-shadow(0_0_8px_#22c55e)]"
			/>
		</svg>

		<div
			bind:this={flashBorderEl}
			class="absolute inset-0 z-0 rounded-3xl border-2 border-green-500 opacity-0 pointer-events-none [box-shadow:0_0_20px_rgba(34,197,94,0.5)]"
		></div>
	</div>
</div>
