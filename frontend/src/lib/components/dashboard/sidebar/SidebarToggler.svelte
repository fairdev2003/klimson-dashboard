<script lang="ts">
	import Icon from '@iconify/svelte';
	import { onMount } from 'svelte';

	type Props = { opened: boolean; onclick: () => void };

	let { opened = $bindable(), onclick }: Props = $props();
</script>

<button
	onclick={() => {
		onclick();
	}}
	class:selected-toggler={opened}
	class:base-toggler={!opened}
	class="p-2 cursor-pointer hover:bg-neutral-700 rounded-xl mb-4"
>
	<Icon icon="tabler:layout-sidebar" width="25" height="25"></Icon>
</button>

<style>
	@import 'tailwindcss';

	.selected-toggler {
		@apply bg-blue-700;
	}

	.base-toggler {
		@apply bg-transparent hover:bg-neutral-700;
	}
</style>
