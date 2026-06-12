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
		current_value?: string[];
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

	let l = $derived(options.find((opt) => current_value?.includes(opt.key))?.key || 'Wybierz opcję');
</script>

<div class="relative">
	<button
		class:min-w-50={set_w}
		class="bg-neutral-800 dropdown-button flex items-center gap-1 justify-between px-4 p-2 border cursor-pointer hover:bg-neutral-700 border-neutral-700 rounded-xl"
		onclick={() => (opened = !opened)}
	>
		{#if children}
			{@render children?.()}
		{:else if label}
			<p>{label}</p>
		{:else}
			<p>
				{current_value?.length > 0 ? `${current_value?.length} items selected` : 'Select options'}
			</p>
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
			class="absolute dropdown-container w-50 overflow-hidden top-full left-1/2 -translate-x-1/2 mt-2 bg-neutral-800 border border-neutral-700 w-48 z-50 rounded-md"
		>
			<div class="flex flex-col">
				{#each options as option}
					<button
						onclick={() => {
							const list = current_value || [];
							if (list.includes(option.value)) {
								current_value = list.filter((val) => val !== option.value);
								return;
							}
							current_value = [option.value, ...list];
							onchoose?.(option);
						}}
						class:selected={current_value?.includes(option.value)}
						class="px-2 py-2 cursor-pointer hover:bg-neutral-700/60 text-left transition-colors"
					>
						<div class="flex gap-3 items-center">
							<div
								class:selected-checkmark={current_value?.includes(option.value)}
								class="size-5 flex bg-neutral-700 justify-center items-center rounded-sm"
							>
								{#if current_value?.includes(option.value)}
									<Icon icon="material-symbols:check" />
								{/if}
							</div>
							<p>
								{option.key}
							</p>
						</div>
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
		@apply bg-blue-700/40 hover:bg-blue-600/40;
	}

	.selected-checkmark {
		@apply bg-blue-700;
	}
</style>
