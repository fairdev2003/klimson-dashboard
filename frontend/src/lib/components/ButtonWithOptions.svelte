<script lang="ts">
	import { ChevronDown } from '@lucide/svelte';
	let optionsOpened: boolean = $state(false);
	let o: string[] = $state(['Siurak', 'Quiz', 'Blog', 'nakurwiam']);

	let boxRef: HTMLDivElement;

	type Option = string;

	type Actions = {
		onoptionclick?: (e: Option) => void;
		onclick?: () => void;
	};

	type Props = {
		options?: string[];
	};

	const { onoptionclick, options = o }: Actions & Props = $props();
</script>

<div class="relative flex w-[200px] items-center" bind:this={boxRef}>
	<button
		onclick={() => {
			onclick;
			optionsOpened = false;
		}}
		class="bg-secondary-add-button h-7 cursor-pointer px-2 text-sm hover:bg-blue-600/60"
		>Nowy projekt</button
	>
	<button
		onclick={() => {
			optionsOpened = !optionsOpened;
		}}
		class="bg-secondary-add-button flex h-7 w-7 cursor-pointer select-none items-center justify-center text-[10px] hover:bg-blue-600/60 focus:outline-0"
		><div class="text-sm"><ChevronDown /></div></button
	>
	{#if optionsOpened}
		<div class=" absolute top-7 z-10 flex w-full flex-col">
			{#if options}
				{#each options as option}
					<button
						class="bg-secondary-add-button hover:bg-primary-add-button cursor-pointer p-1 text-start text-sm"
						onclick={() => {
							if (!onoptionclick) return;
							onoptionclick(option);
							optionsOpened = false;
						}}>{option}</button
					>
				{/each}
			{/if}
		</div>
	{/if}
</div>

<svelte:document
	onmousedown={(e) => {
		if (!boxRef || !boxRef.contains(e.target as Node)) {
			optionsOpened = false;
		}
	}}
	onkeydown={(e) => {
		if (e.key === 'Escape') {
			optionsOpened = false;
		}
	}}
/>
