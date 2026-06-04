<script lang="ts">
	import Icon from '@iconify/svelte';
	import { blur } from 'svelte/transition';

	type Props = {
		checked: boolean;
	};

	let { checked = $bindable() }: Props = $props();
</script>

<button
	onclick={() => {
		checked = !checked;
	}}
	class="border-neutral-700 hover:border-white/80 transition-colors cursor-pointer grid grid-cols-2 p-0.5 gap-0.5 w-16 h-8 border-2"
>
	{#if checked}
		<div class="col-span-1"></div>
		<div
			in:blur={{ duration: 300 }}
			class="bg-green-500 col-span-1 flex items-center justify-center"
		>
			<Icon icon="icon-park-solid:correct" width="20" height="20" />
		</div>
	{:else}
		<div in:blur={{ duration: 300 }} class="bg-red-500 col-span-1 flex items-center justify-center">
			<Icon icon="mingcute:close-fill" width="20" height="20" />
		</div>
		<div class="col-span-1"></div>
	{/if}
</button>
