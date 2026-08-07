<script lang="ts">
	import Icon from '@iconify/svelte';
	import { onMount, type Snippet } from 'svelte';
	import { slide } from 'svelte/transition';

	type Props = { opened: boolean; children: Snippet; name: string };

	let { opened = $bindable(), children, name }: Props = $props();

	onMount(() => {});
</script>

<div transition:slide={{ duration: 300 }} class="">
	<div
		onclick={() => {
			opened = !opened;
		}}
		class="flex gap-2 items-center cursor-pointer select-none"
	>
		<div class="text-text transition-transform duration-300 {opened ? 'rotate-180' : ''}">
			<Icon icon="mdi:chevron-down" width="25" height="25" />
		</div>

		<h3 class="text-text font-bold capitalize text-lg">{name}</h3>
	</div>
	{#if opened}
		<div class="">
			{@render children()}
		</div>
	{/if}
</div>

<style>
	@import 'tailwindcss';

	.selector {
		@apply;
	}
</style>
