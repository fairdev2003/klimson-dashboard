<script lang="ts">
	import { api } from '$lib/api/api';
	import type { ImageKey } from '$lib/api/types';
	import { debug } from '$lib/dashboard/stores/debug';
	import { heroForm } from '$lib/dashboard/stores/hero';
	import { selectedQuestionToEdit, updateFormQuiz } from '$lib/dashboard/stores/store';
	import { toast } from '$lib/dashboard/stores/toast';
	import { onMount } from 'svelte';

	let preview: string | null = $state(null);
	let fileInput: HTMLInputElement | undefined = $state();
	let imageURL: string = $state('');
	let imageName: string | undefined = $state();

	type Props = {
		type?: ImageKey;
		src?: string;
	};

	let { type = 'quiz', src = $bindable(imageName) }: Props = $props();

	const handleFile = async (e: Event) => {
		const target = e.target as HTMLInputElement;
		const file = target.files?.[0];
		if (!file) return;

		preview = URL.createObjectURL(file);

		const id = crypto.randomUUID();

		const formData = new FormData();
		formData.append('image', file);
		formData.append('id', id);

		const response = await api.image.SendImage(type, formData);
		if (type === 'quiz') {
			updateFormQuiz('image_url', response.file_name);
		}
		if (type === 'question') {
			$selectedQuestionToEdit.image_url = response.file_name;
		}
		if (type === 'hero') {
			heroForm.update((hero) => {
				return { ...hero, image_url: response.file_name };
			});
		}

		src = response.file_name;
		if (response.file_name) {
			debug.log('Przesłano zdjęcie');
			toast.show('Przesłano zdjęcie', 'success', 4000);
		}
		console.log(src);
	};

	onMount(() => {
		if (type === 'question') {
			imageURL = api.image.getImage(type, $selectedQuestionToEdit.image_url);
		}
		if (src) {
			imageURL = src;
		}
	});
</script>

<div
	class="bg-primary border-1 border-secondary relative mx-auto flex h-full w-full max-w-[400px] items-center justify-center p-5"
>
	<!-- chcialbym aby ten element prosil o wybor zdjecia z komputera -->
	<button
		onclick={() => {
			if (!fileInput) return;
			fileInput.click();
		}}
		class="absolute left-0 top-0 z-10 aspect-video h-full w-full cursor-pointer p-5 transition-colors hover:bg-white/20"
	></button>
	<!-- ukryty input do wyboru pliku -->
	<input bind:this={fileInput} type="file" accept="image/*" class="hidden" onchange={handleFile} />
	<div class="flex aspect-video w-full items-center justify-center">
		{#if src}
			<img src={api.image.getImage(type, src)} alt={src} class="h-full w-full object-contain" />
		{:else}
			<p class="text-neutral-500">Brak zdjęcia. Kliknij aby dodać</p>
		{/if}
	</div>
</div>
