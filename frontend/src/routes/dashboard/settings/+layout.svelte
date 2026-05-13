<script lang="ts">
	import favicon from '$lib/assets/favicon.svg';
	import Navbar from '$lib/components/Navbar.svelte';
	import SectionList from '$lib/components/dashboard/SectionList.svelte';

	let { children } = $props();
</script>

<div>
	{@render children?.()}
</div>
