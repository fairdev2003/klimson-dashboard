<script lang="ts">
	import Icon from '@iconify/svelte';

	type Props = {
		onButtonClick: () => void;
		title?: string;
	};

	let { onButtonClick, title }: Props = $props();
</script>

<div class="flex justify-between items-center">
	<p class="uppercase text-xs font-black text-white">{title}</p>
	<button
		onclick={onButtonClick}
		class="p-2 hover:bg-neutral-700/50 hover:text-red-400 rounded-xl cursor-pointer"
	>
		<Icon icon="material-symbols:close" />
	</button>
</div>
