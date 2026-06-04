<script lang="ts">
	import Icon from '@iconify/svelte';
	import { blur } from 'svelte/transition';

	type DropdownOption = {
		key: string;
		value: string;
	};

	type Props = {
		options: DropdownOption[];
		onchoose?: (e: DropdownOption) => void;
		current_value: any;
		label?: string;
	};

	function handleClickOutside(event: MouseEvent) {
		const target = event.target as HTMLElement;

		const clickedButton = target.closest('.dropdown-button');
		const clickedDropdown = target.closest('.dropdown-container');

		if (!clickedDropdown && !clickedButton) {
			opened = false;
		}
	}

	let opened: boolean = $state(false);

	let { options, onchoose, current_value = $bindable(), label }: Props = $props();

	let l = $derived(options.find((opt) => opt.value === current_value)?.key || 'Wybierz opcję');
</script>

<div class="relative">
	<button
		class="bg-neutral-800 dropdown-button flex items-center gap-1 justify-between min-w-50 px-4 p-2 border cursor-pointer hover:bg-neutral-700 border-neutral-700 rounded-xl"
		onclick={() => (opened = !opened)}
	>
		<p>{label || l}</p>
		{#if opened}
			<Icon icon="mynaui:chevron-up" width="20" height="20" />
		{:else}
			<Icon icon="mynaui:chevron-down" width="20" height="20" />
		{/if}
	</button>

	{#if opened}
		<div
			in:blur={{ duration: 300 }}
			out:blur={{ duration: 300 }}
			class="absolute dropdown-container overflow-hidden top-full left-1/2 -translate-x-1/2 mt-2 bg-neutral-800 border border-neutral-700 w-48 z-50 rounded-md"
		>
			<div class="flex flex-col">
				{#each options as option}
					<button
						onclick={() => {
							current_value = option.value;

							onchoose?.(option);

							opened = false;
						}}
						class:selected={current_value === option.value}
						class="px-4 py-2 cursor-pointer hover:bg-neutral-700 text-left transition-colors"
					>
						{option.key}
					</button>
				{/each}
			</div>
		</div>
	{/if}
</div>

<svelte:document onclick={(e) => handleClickOutside(e)} />

<style>
	@import 'tailwindcss';

	.selected {
		@apply bg-blue-500 hover:bg-blue-500;
	}
</style>
