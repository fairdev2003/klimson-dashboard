<script lang="ts">
	import favicon from '$lib/assets/favicon.svg';
	import './../app.css';
	import './../scrollbar.css';

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<div class="">
	{@render children?.()}
</div>
