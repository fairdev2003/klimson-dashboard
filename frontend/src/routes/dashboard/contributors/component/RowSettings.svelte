<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<script lang="ts">
	import { addFormContributor } from '$lib/dashboard/stores/store';
	import Icon from '@iconify/svelte';
	import { fade } from 'svelte/transition';
	let btn: HTMLButtonElement | null = null;
	let boxRef: HTMLDivElement | null = null;

	type OptionType = {
		label: string;
		description: string;
		icon: string;
		action: () => void;
		color?: string;
	};

	function handleClickOutside(e: MouseEvent) {
		if (!boxRef || !boxRef.contains(e.target as Node)) {
			if (!btn || !btn.contains(e.target as Node)) {
				open = false;
			}
		}
	}

	type Props = {
		options?: OptionType[];
		onclick?: () => void;
	};

	let { options, onclick }: Props = $props();
	let open = $state(false);
</script>

<div class="">
	<button
		onclick={() => {
			open = !open;

			onclick?.();
		}}
		bind:this={btn}
		class="p-2 text-neutral-500 transition-colors hover:text-white"
	>
		<Icon icon="fluent:ios-arrow-right-24-filled" />
	</button>

	{#if open}
		<div
			in:fade={{ duration: 150 }}
			out:fade={{ duration: 150 }}
			bind:this={boxRef}
			class="border-1 z-100 absolute right-0 top-12 w-[400px] border-neutral-700/60 bg-neutral-800"
		>
			<p class="p-1 text-start text-neutral-400">{$addFormContributor.name}</p>
			{#each options as option}
				{@render Option(option)}
			{/each}
		</div>
	{/if}
</div>

{#snippet Option(action: OptionType)}
	<button
		class="w-full text-start"
		onclick={() => {
			action.action();
			open = false;
		}}
	>
		<div class="flex cursor-pointer items-center gap-3 p-4 hover:bg-neutral-700">
			<Icon icon={action.icon} width="24" height="24" class={action.color} />
			<div>
				<h3 class="font-semibold">{action.label}</h3>
				<p class="text-sm text-neutral-400">{action.description}</p>
			</div>
		</div>
	</button>
{/snippet}

<svelte:document
	onmousedown={(e) => {
		handleClickOutside(e);
	}}
/>
