<script lang="ts">
	import Icon from '@iconify/svelte';

	type Props = {
		checked: boolean;
	};

	let { checked = $bindable() }: Props = $props();
</script>

<button
	onclick={() => {
		checked = !checked;
	}}
	class="border-neutral-700 hover:border-blue-500/80 hover:bg-blue-500/30 transition-colors duration-300 cursor-pointer
           flex items-center rounded-lg w-16 h-8 border-2 bg-neutral-900 p-0.5 relative"
>
	<div
		class="h-full w-6 rounded-md flex items-center justify-center
               transition-all duration-300 ease-in-out shadow-sm
               {checked ? 'translate-x-8 bg-green-500' : 'translate-x-0 bg-neutral-500'}"
	>
		{#if checked}
			<Icon icon="icon-park-solid:correct" width="18" height="18" />
		{:else}
			<Icon icon="mingcute:close-fill" width="18" height="18" />
		{/if}
	</div>
</button>
