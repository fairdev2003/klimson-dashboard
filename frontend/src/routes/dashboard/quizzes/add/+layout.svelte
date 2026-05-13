<script lang="ts">
	import favicon from '$lib/assets/favicon.svg';
	import Button from '$lib/components/Button.svelte';
	import Navbar from '$lib/components/Navbar.svelte';
	import SectionList from '$lib/components/dashboard/SectionList.svelte';
	import Heading from '$lib/components/dashboard/typography/Heading.svelte';

	let { children } = $props();
</script>

<svelte:head>
	<title>Panel Admina | Lista Quizów</title>
	<meta name="description" content="Panel Admina HarcQuiz | Quizy" />
</svelte:head>
<div>
	{@render children?.()}
</div>
