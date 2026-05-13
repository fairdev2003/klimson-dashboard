<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/state'; // Svelte 5 style
	import Button from '$lib/components/Button.svelte';
</script>

<div class="flex min-h-[60vh] flex-col items-center justify-center gap-4">
	<h1 class="text-6xl font-bold text-white">{page.status}</h1>

	{#if page.status === 404}
		<p class="text-neutral-400">Ups! Ta część panelu sterowania nie istnieje.</p>
	{:else}
		<p class="text-neutral-400">Coś poszło nie tak w dashboardzie.</p>
	{/if}

	<Button
		onclick={() => {
			goto('/dashboard');
		}}
		theme="secondary">Wróć do pulpitu</Button
	>
</div>
