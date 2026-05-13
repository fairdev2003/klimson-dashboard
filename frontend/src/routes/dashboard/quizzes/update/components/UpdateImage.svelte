<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<!-- svelte-ignore a11y_missing_attribute -->
<script lang="ts">
	import { api } from '$lib/api/api';
	import Button from '$lib/components/Button.svelte';
	import CreateFormInput from '$lib/components/dashboard/CreateFormInput.svelte';
	import ImageUpload from '$lib/components/dashboard/ImageUpload.svelte';
	import { updateFormQuiz } from '$lib/dashboard/stores/store';
	import { toast } from '$lib/dashboard/stores/toast';
	import { getContext, setContext } from 'svelte';
	import { imageFile, imageSrc } from './image.store';

	type Props = {
		onSave?: () => void;
	};

	let { onSave }: Props = $props();
	let loading: boolean = $state(false);

	async function handleSave() {
		if (!$imageSrc) {
			toast.show('Prześlij najpierw zdjęcie!', 'error');
			return;
		}

		loading = true;
		try {
			let formData = new FormData();
			if ($imageFile) {
				formData.append('image', $imageFile);
			}

			// Strzał do API, który zapisuje rekord w bazie
			await api.image.SendImage('quiz', formData);

			toast.show('Dane zapisane!', 'success');
			onSave?.();
		} catch (e) {
			toast.show('Błąd zapisu', 'error');
		} finally {
			loading = false;
		}
	}

	let preview: string | null = $state(null);

	let value: string = $state('');
	const type = 'quiz';

	$effect(() => {
		if ($imageFile) {
			console.log($imageFile);
		}
	});
</script>

<div class="flex flex-col gap-4">
	<div class="flex flex-1 flex-col gap-4 overflow-y-auto pt-5">
		<ImageUpload label="Fotka" />
		<CreateFormInput bind:value label="Link do zdjęcia"></CreateFormInput>
		<div class="flex justify-end">
			<Button
				{loading}
				theme="secondary"
				size="small"
				onclick={() => {
					console.log($imageFile);
				}}>Zapisz zmiany</Button
			>
		</div>
	</div>
</div>
