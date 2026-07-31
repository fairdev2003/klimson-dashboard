<script lang="ts">
	import { goto } from '$app/navigation';
	import Button from '$lib/components/Button.svelte';
	import Dashboard from '$lib/dashboard/dashboard.svelte';
	import { onMount } from 'svelte';
	import { cubicOut } from 'svelte/easing';
	import { fly, slide } from 'svelte/transition';

	type Props = { mobileDockOpened: boolean };

	let { mobileDockOpened = $bindable() }: Props = $props();

	function whizz(node: HTMLElement, { from, to }: { from: DOMRect; to: DOMRect }, params: any) {
		const dx = from.left - to.left;
		const dy = from.top - to.top;

		const d = Math.sqrt(dx * dx + dy * dy);

		return {
			delay: 0,
			duration: Math.sqrt(d) * 120,
			easing: cubicOut
		};
	}
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div
	onclick={(e) => {
		e.stopPropagation();
	}}
	transition:fly={{ y: -40, duration: 250, easing: cubicOut }}
	class="lg:hidden p-4 flex h-150 shadow-2xl w-full bg-neutral-900 absolute z-20 border-b border-neutral-800"
>
	<div class="flex flex-col gap-2 w-full">
		{#each Dashboard.constants.SidebarContents as item}
			<button
				onclick={() => {
					goto(item.href);
					mobileDockOpened = false;
				}}
				class="text-start bg-neutral-800 p-2 rounded-xl"
			>
				{item.name}
			</button>
		{/each}
	</div>
</div>
