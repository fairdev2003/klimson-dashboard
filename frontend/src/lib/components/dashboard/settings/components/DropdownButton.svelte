<script lang="ts">
	import Icon from '@iconify/svelte';
	import type { Snippet } from 'svelte';
	import { blur } from 'svelte/transition';

	type DropdownOption = {
		key: string;
		value: string;
	};

	type Props = {
		options: DropdownOption[];
		onchoose?: (e: DropdownOption) => void;
		current_value?: any;
		error_text?: string;
		children?: Snippet;
		set_w: boolean;
		label: string;
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

	let {
		options,
		onchoose,
		set_w = false,
		current_value = $bindable(),
		children,
		error_text,
		label
	}: Props = $props();

	let l = $derived(options.find((opt) => opt.value === current_value)?.key || 'Select an option');
</script>

<div class="relative">
	<button
		class:min-w-50={set_w}
		class="bg-foreground text-text dropdown-button flex items-center gap-1 justify-between px-4 p-2 border cursor-pointer hover:bg-primary border-border rounded-xl"
		onclick={() => (opened = !opened)}
	>
		{#if children}
			{@render children?.()}
		{:else if label}
			<p class="text-text">{label}</p>
		{:else}
			<p>{l}</p>
		{/if}

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
			class="absolute dropdown-container w-50 overflow-hidden top-full left-1/2 -translate-x-1/2 mt-2 bg-foreground border border-border w-48 z-50 rounded-md"
		>
			<div class="flex flex-col">
				{#each options as option}
					<button
						onclick={() => {
							if (current_value === option.value) {
								opened = false;

								return;
							}
							current_value = option.value;

							onchoose?.(option);

							opened = false;
						}}
						class:selected={current_value === option.value}
						class="px-4 py-2 cursor-pointer hover:bg-primary text-secondary-text text-left transition-colors"
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
		background-color: var(--color-primary);
		color: var(--color-text);
	}

	.selected:hover {
		background-color: var(--color-primary);
	}
</style>
