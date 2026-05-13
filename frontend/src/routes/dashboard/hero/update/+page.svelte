<script lang="ts">
	import { goto } from '$app/navigation';
	import { api } from '$lib/api/api';
	import Button from '$lib/components/Button.svelte';
	import ImageContainer from '$lib/components/dashboard/add/ImageContainer.svelte';
	import { blogForm, blogFormState } from '$lib/dashboard/stores/blog';
	import { heros } from '$lib/dashboard/stores/data.store';
	import { heroForm } from '$lib/dashboard/stores/hero';
	import { toast } from '$lib/dashboard/stores/toast';
	import { marked } from 'marked';
	import { onMount } from 'svelte';
	import { get } from 'svelte/store';

	function updateField(key: 'quote' | 'image_url' | 'author', value: string) {
		heroForm.update((form) => ({
			...form,
			[key]: value
		}));
	}
</script>

<div class="mx-auto mt-6 flex max-w-5xl flex-col gap-6">
	<!-- FORM -->
	<div class="flex flex-col gap-4 rounded-xl bg-neutral-900 p-6 shadow">
		<h2 class="text-lg font-semibold text-white">Nowy cytat</h2>

		<ImageContainer type="hero" mode="add" src={$heroForm.image_url} />

		<input
			class="rounded bg-neutral-800 px-3 py-2 text-white ring-1 ring-neutral-700 outline-none focus:ring-blue-500"
			placeholder="Cytat"
			value={$heroForm.quote}
			oninput={(e) => updateField('quote', (e.currentTarget as HTMLInputElement).value)}
		/>

		<input
			class="rounded bg-neutral-800 px-3 py-2 text-white ring-1 ring-neutral-700 outline-none focus:ring-blue-500"
			placeholder="Autor"
			value={$heroForm.author}
			oninput={(e) => updateField('author', (e.currentTarget as HTMLInputElement).value)}
		/>

		<Button
			theme="correct"
			className="mt-2"
			onclick={async () => {
				await api.hero.UpdateHero(get(heroForm));
				toast.show('Updated Hero', 'info', 4000);
				const r = await api.hero.FetchHeros();
				heros.set(r.data);
				goto('/dashboard/hero');
			}}>Opublikuj cytat</Button
		>
	</div>
</div>
