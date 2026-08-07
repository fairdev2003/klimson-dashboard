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
	title="Click 'Tab' to toggle sidebar"
	class:selected-toggler={opened}
	class:base-toggler={!opened}
	class="p-2 cursor-pointer focus:outline-none transition-colors text-text hover:bg-neutral-700 rounded-xl mb-4"
>
	<Icon icon="tabler:layout-sidebar" width="25" height="25"></Icon>
</button>

<style>
	@import 'tailwindcss';

	.selected-toggler {
		@apply bg-blue-700;
	}

	.selected-toggler {
		background-color: var(--color-primary);
	}

	.selected-toggler:hover {
		background-color: var(--color-primary);
	}

	.base-toggler {
		background-color: transparent;
	}

	.base-toggler:hover {
		background-color: var(--color-primary);
	}
</style>
