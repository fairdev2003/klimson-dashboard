<script lang="ts">
	import { fly } from 'svelte/transition';
	type Props = { mobileDockOpened: boolean };

	let { mobileDockOpened = $bindable() }: Props = $props();
</script>

<div
	onclick={() => {
		mobileDockOpened = !mobileDockOpened;
	}}
	transition:fly={{ duration: 450 }}
	class="absolute flex lg:hidden z-1 bg-foreground h-screen w-screen"
></div>
