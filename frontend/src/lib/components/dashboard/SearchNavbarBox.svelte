<script lang="ts">
	import gsap from 'gsap';
	let value: string = $state('');
	let opened: boolean = $state(false);

	let boxRef: HTMLDivElement | null = null;
	let inputContainerRef: HTMLDivElement | null = null;

	function handleClickOutside(e: MouseEvent) {
		if (!opened) return;

		if (inputContainerRef && !inputContainerRef.contains(e.target as Node)) {
			gsap.to(boxRef, {
				height: '50px',
				opacity: 0,
				duration: 0.3,
				onComplete: () => {
					opened = false;
				}
			});
		}
	}

	$effect(() => {
		if (opened) {
			gsap.fromTo(
				boxRef,
				{ height: '70px', opacity: 0 },
				{
					height: '400px',
					opacity: 1,
					duration: 0.3,
					onComplete: () => {
						inputContainerRef?.focus();
					}
				}
			);
		}
	});
</script>

<div class="relative hidden h-[50px] justify-center w-[500px] md:flex lg:flex">
	{#if !opened}
		<input
			class="h-full w-full bg-transparent border-1 border-neutral-700/60 rounded-xl placeholder-neutral-500 ring-0 outline-none select-none focus:outline-none"
			bind:value
			onclick={() => {
				opened = true;
			}}
			placeholder="Kliknij aby wyszukać"
		/>
	{/if}

	{#if opened}
		<div
			class="bg-black/50 fixed inset-1 w-full h-full z-90"
			onclick={() => {
				gsap.to(boxRef, {
					height: '70px',
					duration: 0.3,
					onComplete: () => {
						opened = false;
					}
				});
			}}
		></div>
		<div
			bind:this={boxRef}
			onclick={(e) => {
				e.stopPropagation();
			}}
			class="absolute w-[500px] border border-neutral-700 p-3 z-100 h-100 bg-neutral-900 rounded-xl overflow-hidden"
		>
			<input
				bind:this={inputContainerRef}
				class="w-full bg-neutral-800 outline-0 p-3 focus:outline-0 active:outline-0 rounded-lg focus:border-0 active:border-0 border-0"
			/>
		</div>
	{/if}
</div>

<style>
	input:invalid {
		outline: none;
		box-shadow: none;
	}
</style>
