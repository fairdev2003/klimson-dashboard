<script lang="ts">
	import Icon from '@iconify/svelte';
	import { slide } from 'svelte/transition';
	import { terminal, type TerminalPage } from './console/terminal.svelte';

	let opened = $state(false);
	type OptionRecord = { name: string; value: TerminalPage };
	const options: OptionRecord[] = [
		{
			name: 'Dashboard Terminal',
			value: 'user'
		},
		{
			name: 'Logs',
			value: 'only-logs'
		},
		{
			name: 'Http logger',
			value: 'http'
		}
	];
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div
	onmousedown={(e) => e.stopPropagation()}
	onmousemove={(e) => e.stopPropagation()}
	onclick={(e) => e.stopPropagation()}
	class="relative"
>
	{@render Button()}
	{#if opened}
		<div
			in:slide={{ duration: 300 }}
			out:slide={{ duration: 300 }}
			class="absolute top-0 left-0 -right-10 z-10 w-60 p-3 bg-neutral-800 rounded-xl flex gap-2 flex-col"
		>
			{#each options as option}
				{@render OptionButton(option)}
			{/each}
		</div>
	{/if}
</div>

{#snippet OptionButton(option: OptionRecord)}
	<button
		onclick={() => {
			terminal.terminalPage = option.value;
			opened = false;
		}}
		class:option-base={option.value !== terminal.terminalPage}
		class:option-selected={option.value === terminal.terminalPage}
		class="flex items-center gap-2 h-8 w-full px-4 rounded-xl cursor-pointer"
	>
		<p class="text-[10px] font-bold tracking-widest text-white uppercase">{option.name}</p>
	</button>
{/snippet}

{#snippet Button()}
	<button
		onclick={() => (opened = !opened)}
		class="flex items-center justify-between gap-2 h-8 w-50 px-4 rounded-full hover:bg-white/10 cursor-pointer"
	>
		<p class="text-[10px] font-bold tracking-widest text-white uppercase">
			{options.find((e) => e.value === terminal.terminalPage)?.name}
		</p>
		<Icon icon="mdi:chevron-down" />
	</button>
{/snippet}

<style>
	@import 'tailwindcss';

	.option-selected {
		@apply bg-blue-800 hover:bg-blue-600;
	}

	.option-base {
		@apply hover:bg-white/10 bg-transparent;
	}
</style>
