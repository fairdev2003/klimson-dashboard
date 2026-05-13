<script lang="ts">
	import favicon from '$lib/assets/favicon.svg';
	import Navbar from '$lib/components/Navbar.svelte';
	import SectionList from '$lib/components/dashboard/SectionList.svelte';

	let { children } = $props();
</script>

<svelte:head>
	<title>Panel Admina | Lista Pytań</title>
	<meta name="description" content="Panel Admina HarcQuiz" />
</svelte:head>

<div>
	{@render children?.()}
</div>
