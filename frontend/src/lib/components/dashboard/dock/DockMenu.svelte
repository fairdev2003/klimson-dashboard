<script lang="ts">
	import { goto } from '$app/navigation';
	import Button from '$lib/components/Button.svelte';
	import Dashboard from '$lib/dashboard/dashboard.svelte';
	import { onMount } from 'svelte';
	import { slide } from 'svelte/transition';
</script>

<div
	transition:slide={{ duration: 300 }}
	class="lg:hidden p-4 flex h-150 shadow-2xl w-full bg-neutral-900 absolute z-100 border-b border-neutral-800"
>
	<div class="flex flex-col gap-2 w-full">
		{#each Dashboard.constants.SidebarContents as item}
			<button
				onclick={() => {
					goto(item.href);
				}}
				class="text-start bg-neutral-800 p-2 rounded-xl"
			>
				{item.name}
			</button>
		{/each}
	</div>
</div>
