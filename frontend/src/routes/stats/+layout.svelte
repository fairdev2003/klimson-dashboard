<script lang="ts">
	let { children } = $props();
</script>

<svelte:head>
	<title>Statystyki | Quizy</title>
	<meta name="description" content="Panel Admina HarcQuiz" />
</svelte:head>

<div>
	{@render children?.()}
</div>
