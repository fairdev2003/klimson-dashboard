<script lang="ts">
	import favicon from '$lib/assets/favicon.svg';
	import Navbar from '$lib/components/Navbar.svelte';
	import SectionList from '$lib/components/dashboard/SectionList.svelte';
	import { AppWindowIcon } from '@lucide/svelte';

	let { children } = $props();
</script>

<svelte:head>
	<title>Panel Admina | Scieżki API</title>
	<meta name="description" content="Panel Admina HarcQuiz" />
</svelte:head>

<div>
	{@render children?.()}
</div>
