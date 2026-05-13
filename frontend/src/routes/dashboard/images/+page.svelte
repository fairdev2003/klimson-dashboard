<script lang="ts">
	import { api } from '$lib/api/api';
	import type { ImageList } from '$lib/api/types';

	import { onMount, type Component } from 'svelte';
	import { images } from '$lib/dashboard/stores/images';
	import Heading from '$lib/components/dashboard/typography/Heading.svelte';
	import Button from '$lib/components/Button.svelte';
	import QuizImagesList from './components/QuizImagesList.svelte';
	import QuestionImagesList from './components/QuestionImagesList.svelte';
	import HeroImagesList from './components/HeroImagesList.svelte';
	import StaticImagesList from './components/StaticImagesList.svelte';

	async function FetchImages(): Promise<ImageList> {
		const key = 'static';
		const response = await api.image.ListImages(key);
		console.log('Images');
		console.log(response);
		return response;
	}

	onMount(async () => {
		$images = await FetchImages();
	});

	type Folder = { tag: string; name: string; component: Component };

	const folders: Folder[] = [
		{ tag: 'quiz', name: 'Quiz', component: QuizImagesList },
		{ tag: 'question', name: 'Pytanie', component: QuestionImagesList },
		{ tag: 'hero', name: 'Hero', component: HeroImagesList },
		{ tag: 'static', name: 'Statyczne', component: StaticImagesList }
	];
	let selectedFolder: Folder | undefined = $state(folders[0]);
</script>

<div class="p-5 text-white">
	<Heading>Zdjęcia</Heading>
	<div class="flex justify-between">
		<div class="flex gap-2 mt-5">
			{#each folders as folder}
				<Button
					onclick={() => {
						selectedFolder = folder;
					}}
					theme={selectedFolder && selectedFolder.tag === folder.tag ? 'secondary' : 'base'}
					>{folder.name}</Button
				>
			{/each}
		</div>
		<Button theme="secondary">Dodaj</Button>
	</div>
	{#if selectedFolder}
		<selectedFolder.component />
	{/if}
</div>
