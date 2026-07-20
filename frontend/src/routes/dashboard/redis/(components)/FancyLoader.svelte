<script lang="ts">
	import { onMount } from 'svelte';
	import gsap from 'gsap';
	import { tv } from 'tailwind-variants';

	let firstsquareRef: HTMLDivElement;
	let secondsquareRef: HTMLDivElement;
	let state: boolean = $state(false);

	type FancyLoaderColor = 'red' | 'orange' | 'blue' | 'white' | 'gray';

	type Props = {
		color?: FancyLoaderColor;
		centerAnimationEnabled?: boolean;
		rotate?: 360 | 720;
		size?: 'big' | 'medium' | 'minimalistic';
	};
	const { color = 'red', centerAnimationEnabled, rotate = 720, size = 'big' }: Props = $props();

	const loaderStyles = tv({
		base: 'h-10 top-0 w-10  absolute z-2',
		variants: {
			color: {
				red: 'border-red-500',
				orange: 'border-orange-500',
				blue: 'border-blue-500',
				white: 'border-white',
				gray: 'border-neutral-600'
			},
			size: {
				big: 'h-10 w-10 border-[5px]',
				medium: 'h-6 w-6 border-[5px]',
				minimalistic: 'w-3 h-3 border-[2px]'
			}
		},
		defaultVariants: {
			color: 'red',
			size: 'big'
		}
	});

	const centerSquareStyles = tv({
		base: 'w-full h-full overflow-hidden',
		variants: {
			color: {
				red: 'bg-red-500',
				orange: 'bg-orange-500',
				blue: 'bg-blue-500',
				white: 'bg-white',
				gray: 'bg-neutral-600'
			}
		},
		defaultVariants: {
			color: 'red'
		}
	});

	onMount(() => {
		const tl = gsap.timeline({ repeat: -1 });
		const tr = gsap.timeline({ repeat: -1 });

		tl.to(firstsquareRef, {
			rotation: rotate,
			duration: 1.5,
			ease: 'power2.inOut'
		})
			.to(secondsquareRef, {
				height: '0%',
				duration: 0.5,
				ease: 'power2.inOut'
			})
			.to(firstsquareRef, {
				rotation: rotate * 2,
				duration: 1.5,
				ease: 'power2.inOut'
			})
			.to(secondsquareRef, {
				height: '100%',
				duration: 0.5,
				ease: 'power2.inOut'
			});
	});
</script>

<div class="relative flex items-center justify-center">
	<div bind:this={firstsquareRef} class={loaderStyles({ color, size })}>
		<div bind:this={secondsquareRef} class={centerSquareStyles({ color })}></div>
	</div>
</div>
