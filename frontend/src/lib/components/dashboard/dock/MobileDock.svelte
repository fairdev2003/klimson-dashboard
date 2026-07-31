<script lang="ts">
	import DockBackground from './DockBackground.svelte';
	import DockMenu from './DockMenu.svelte';

	type Props = { mobileDockOpened: boolean };

	let { mobileDockOpened = $bindable() }: Props = $props();
</script>

<div>
	{#if mobileDockOpened}
		<DockBackground bind:mobileDockOpened />
		<DockMenu bind:mobileDockOpened />
	{/if}
</div>
