<script lang="ts">
	import favicon from '$lib/assets/favicon.svg';
	import Navbar from '$lib/components/Navbar.svelte';

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
	<title>HarcQuiz | Quiz</title>
	<meta name="description" content="Panel Admina HarcQuiz" />
</svelte:head>

<div>
	{@render children?.()}
</div>
