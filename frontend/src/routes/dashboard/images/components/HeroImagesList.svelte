<script lang="ts">
	import { api } from '$lib/api/api';
	import { onMount } from 'svelte';
	import ListedImage from './ListedImage.svelte';

	let list: string[] = $state([]);

	onMount(async () => {
		list = await api.image.ListImages('hero');
		console.log(list);
	});
</script>

<div class="flex flex-wrap gap-4 p-5">
	{#each list as src}
		<ListedImage {src} />
	{/each}
</div>
