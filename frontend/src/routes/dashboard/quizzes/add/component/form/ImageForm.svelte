<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<!-- svelte-ignore a11y_missing_attribute -->

<script lang="ts">
	import { api } from '$lib/api/api';
	import Button from '$lib/components/Button.svelte';
	import CreateFormInput from '$lib/components/dashboard/CreateFormInput.svelte';
	import ImageUpload from '$lib/components/dashboard/ImageUpload.svelte';
	import { debug } from '$lib/dashboard/stores/debug';
	import { heroForm } from '$lib/dashboard/stores/hero';
	import { selectedQuestionToEdit, updateFormQuiz } from '$lib/dashboard/stores/store';
	import { toast } from '$lib/dashboard/stores/toast';
	import Icon from '@iconify/svelte';

	type Props = {
		onSave?: () => void;
	};

	let { onSave }: Props = $props();
	let loading: boolean = $state(false);

	async function handleSave() {
		loading = true;
		await new Promise((resolve) => setTimeout(resolve, 2000));
		loading = false;
		toast.show('Zmiany zostały zapisane pomyślnie!', 'success');
		onSave?.();
	}

	let preview: string | null = $state(null);
	let fileInput: HTMLInputElement | undefined = $state();
	let imageURL: string = $state('');
	let imageName: string | undefined = $state();
	let modalOpen: boolean = $state(false);
	let value: string = $state('');
	let isDragging = $state(false);
	const type = 'quiz';
	let src: string = $state('');

	function onDragOver(e: DragEvent) {
		e.preventDefault();
		isDragging = true;
	}

	function onDragLeave() {
		isDragging = false;
	}

	async function onDrop(e: DragEvent) {
		e.preventDefault();
		isDragging = false;

		const file = e.dataTransfer?.files?.[0];
		if (!file) return;
		preview = URL.createObjectURL(file);

		await uploadFile(file);
	}

	const handleFile = async (e: Event) => {
		const target = e.target as HTMLInputElement;
		const file = target.files?.[0];
		if (!file) return;
		preview = URL.createObjectURL(file);

		await uploadFile(file);
	};

	async function uploadFile(file: File) {
		preview = URL.createObjectURL(file);

		const id = crypto.randomUUID();
		const formData = new FormData();
		formData.append('image', file);
		formData.append('id', id);

		const response = await api.image.SendImage(type, formData);

		if (type === 'quiz') {
			updateFormQuiz('image_url', response.file_name);
		}

		src = response.file_name;

		toast.show('Przesłano zdjęcie', 'success', 4000);
	}

	let imageReference: string = $state('');
	$effect(() => {
		toast.show(imageReference);
		debug.log(imageReference);
	});
</script>

<div class="flex flex-col gap-4">
	<div class="flex flex-1 flex-col gap-4 overflow-y-auto pt-5">
		<ImageUpload bind:src={imageReference} label="Fotka" />
		<CreateFormInput bind:value label="Link do zdjęcia"></CreateFormInput>
		<div class="flex justify-end">
			<Button
				{loading}
				theme="secondary"
				size="small"
				onclick={function () {
					handleSave ? handleSave() : null;
				}}>Zapisz zmiany</Button
			>
		</div>
	</div>
</div>

{#snippet S()}
	<div class="text-neutral-400">
		<h1 class="font-bold text-white">Prześlij zdjęcie</h1>
		<p>Przeslij zdjecie albo wklej link do zdjecia</p>
	</div>
	{#if preview}
		<div>
			<div class="flex aspect-video h-30 w-full items-center justify-center">
				<img src={preview} class="h-full w-full object-contain" />
			</div>
		</div>
	{/if}

	<div
		ondragover={onDragOver}
		ondragleave={onDragLeave}
		ondrop={onDrop}
		class="
		mx-auto flex w-6/10 flex-col justify-center gap-2 rounded-lg border-1 border-dashed p-5 text-center
		transition-colors
		{isDragging
			? 'border-blue-400 bg-blue-500/20 text-white'
			: 'border-blue-700 bg-blue-700/10 text-neutral-400'}
	"
	>
		<Icon icon="uil:image-upload" width="35" height="35" class="mx-auto  " />
		<div>
			<!-- svelte-ignore a11y_no_static_element_interactions -->
			<p class="flex justify-center gap-1">
				<a
					class="cursor-pointer text-blue-500 hover:underline"
					onclick={() => {
						if (!fileInput) return;
						fileInput.click();
					}}>Kliknij tutaj aby przesłać zdjecie</a
				>albo przeciągnij i upuść tutaj
			</p>
		</div>
		<div>
			<p class="flex justify-center gap-1">Max file size is 15MB</p>
		</div>
	</div>
{/snippet}
