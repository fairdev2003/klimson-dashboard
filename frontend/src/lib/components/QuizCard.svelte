<script lang="ts">
	import Campfire from '$lib/assets/campfire-pfp.webp';
	import gsap from 'gsap';

	let contentToggled = false;
	let cardContent: HTMLDivElement;
	const topResults = [
		{ name: 'Anna', score: 95 },
		{ name: 'Marek', score: 90 },
		{ name: 'Ola', score: 85 },
		{ name: 'Piotr', score: 80 }
	];

	function toggleElement() {
		contentToggled = !contentToggled;

		if (contentToggled) {
			gsap.fromTo(
				cardContent,
				{ height: 0 },
				{ height: cardContent.scrollHeight, opacity: 1, duration: 0.3, ease: 'power2.out' }
			);
		} else {
			gsap.to(cardContent, {
				height: 0,

				duration: 0.3,
				ease: 'power2.out'
			});
		}
	}
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div
	onclick={toggleElement}
	class="w-md duration-50 flex h-auto cursor-pointer flex-col justify-center gap-5 rounded-2xl border-0 border-neutral-700 bg-neutral-800 p-5 transition-colors active:bg-neutral-500"
>
	<div class="flex">
		<img src={Campfire} alt="lorem-pfp" class="size-18 select-none rounded-2xl" />
		<div class="flex flex-col gap-1 p-3 text-white">
			<h3 class="text-lg font-bold">Plakietki Specjalności Drużyn</h3>
			<div class="flex items-center gap-2">
				<div
					class="w-15 flex h-6 select-none items-center justify-center rounded-full bg-green-500/50 text-sm font-bold text-green-400"
				>
					Łatwy
				</div>
				<div
					class="flex h-6 select-none items-center justify-center rounded-full bg-gray-700/50 px-3 text-sm font-bold text-gray-400"
				>
					Paweł Cyngot
				</div>
			</div>
		</div>
	</div>

	<div
		bind:this={cardContent}
		class="mt-0 overflow-hidden font-bold text-white"
		style="height: 0; opacity: 0;"
	>
		<div></div>
		<div
			class="flex h-40 w-full flex-col items-center justify-center gap-1 rounded-2xl bg-neutral-600/60"
		></div>
	</div>
</div>
