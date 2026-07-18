<script lang="ts">
	import { onMount } from 'svelte';
	import gsap from 'gsap';

	onMount(() => {
		animate();
	});

	function animate() {
		gsap.fromTo(
			'#orbiter',
			{ rotation: 160, svgOrigin: '250 250' },
			{
				rotation: 160 + 360,
				duration: 1,
				ease: 'power2.inOut',
				svgOrigin: '250 250',
				overwrite: true
			}
		);
	}

	type Props = {
		onclick?: () => boolean;
	};

	let { onclick }: Props = $props();
</script>

<svg
	xmlns="http://www.w3.org/2000/svg"
	viewBox="0 0 500 500"
	width="50"
	height="50"
	class="mr-5 cursor-pointer"
	onclick={() => {
		const bool = onclick?.();
		if (bool) {
			animate();
		}
	}}
>
	<defs>
		<filter
			id="pixelete-filter-1"
			color-interpolation-filters="sRGB"
			x="0"
			y="0"
			width="100%"
			height="100%"
		>
			<feFlood height="2" width="2" />
			<feComposite width="12" height="12" />
			<feTile result="a" />
			<feComposite in="SourceGraphic" in2="a" operator="in" />
			<feMorphology operator="dilate" radius="6" />
		</filter>
	</defs>
	<ellipse
		style="fill: none; fill-rule: nonzero; stroke-width: 20px"
		class="stroke-blue-800 flex justify-center items-center"
		cx="250"
		cy="250"
		rx="210"
		ry="210"
	/>
	<text
		x="250"
		y="250"
		font-family="Arial, sans-serif"
		font-size="200"
		font-weight="bold"
		text-anchor="middle"
		dominant-baseline="middle"
		class="underline fill-blue-800 pointer-events-none"
	>
		K
	</text>

	<g id="orbiter">
		<ellipse
			style="stroke: rgb(26, 114, 181)"
			class="fill-blue-500"
			cx="250"
			cy="40"
			rx="40"
			ry="40"
		/>
	</g>
</svg>
