<script lang="ts">
	import { addFormQuiz, contextMenuOptions, updateFormQuiz } from '$lib/dashboard/stores/store';
	import { onMount, tick } from 'svelte';
	import { goto } from '$app/navigation';
	import ResiableList from '../add/component/ResiableList.svelte';
	import Heading from '$lib/components/dashboard/typography/Heading.svelte';
	import UpdateNav from './components/FormNav.svelte';
	import UpdateQuizSettings from './components/UpdateQuizSettings.svelte';
	import UpdateQuestions from './components/UpdateQuestions.svelte';
	import UpdateBasicInfo from './components/UpdateBasicInfo.svelte';
	import type { Quiz } from '../types';
	import Button from '$lib/components/Button.svelte';
	import Icon from '@iconify/svelte';
	import { toast } from '$lib/dashboard/stores/toast';
	import ImageUpload from '$lib/components/dashboard/ImageUpload.svelte';
	import { api } from '$lib/api/api';
	import { imageFile } from './components/image.store';
	import { debug } from '$lib/dashboard/stores/debug';
	import { quizzes } from '$lib/dashboard/stores/data.store';
	import gsap from 'gsap';

	onMount(async () => {
		await tick();

		gsap.fromTo('.whole', { opacity: 0, y: -60, delay: 400, ease: 'power3' }, { opacity: 1, y: 0 });

		if (!$addFormQuiz.id) {
			goto('/dashboard/quizzes');
		}
	});
	let quiz: Quiz = $state($addFormQuiz);

	let imageFormOpen: boolean = $state(false);
	let basicInfoFormOpen: boolean = $state(false);
	let questionFormOpen: boolean = $state(false);
	let settingsFormOpen: boolean = $state(false);
	let src: string = $state('');
	let file: File | undefined = $state();

	onMount(() => {
		if ($addFormQuiz.image_url) {
			src = api.image.quiz($addFormQuiz.image_url);
		}
	});

	async function SendImage() {
		toast.info('cos sie dzieje');
		console.log(file);
		const formData = new FormData();
		if (file) {
			formData.append('image', file);
		}
		const id = crypto.randomUUID();
		formData.append('id', id);
		const response = await api.image.SendImage('quiz', formData);

		updateFormQuiz('image_url', response.file_name);

		const response2 = await api.quiz.UpdateQuizImage(
			{ id: $addFormQuiz.id },
			{ image_url: response.file_name }
		);
		console.log(response2.data);

		src = api.image.quiz(response.file_name);
		if (response.file_name) {
			debug.log('Przesłano zdjęcie');
			toast.show('Przesłano zdjęcie', 'success', 4000);
			imageFormOpen = false;
			const r = await api.quiz.GetAll();
			$quizzes = r.data;
		}
		console.log(src);
	}
</script>

<div class="">
	<UpdateNav
		><Heading>Aktualizacja quizu: <span class="text-neutral-500">{quiz?.title}</span></Heading
		></UpdateNav
	>
	<section class="whole mx-auto mt-5 flex w-6/10 flex-col gap-5">
		<ResiableList bind:isOpen={basicInfoFormOpen} section_name="Podstawowe informacje">
			{#snippet nav_icon()}
				<Icon icon="material-symbols:chat-info" height="30" width="30" />
			{/snippet}
			<UpdateBasicInfo
				onSave={() => {
					quiz.title = $addFormQuiz.title;
					basicInfoFormOpen = !basicInfoFormOpen;
				}}
			/>
		</ResiableList>
		<ResiableList bind:isOpen={imageFormOpen} section_name="Multimedia">
			{#snippet nav_icon()}
				<Icon icon="material-symbols:image" height="30" width="30" />
			{/snippet}
			<div class="flex flex-col gap-4">
				<ImageUpload bind:src bind:file />
				<div class="flex justify-end">
					<Button
						theme="secondary"
						size="small"
						onclick={async () => {
							await SendImage();
						}}>Zapisz obrazek</Button
					>
				</div>
			</div>
		</ResiableList>

		<ResiableList bind:isOpen={questionFormOpen} section_name="Pytania">
			{#snippet nav_icon()}
				<Icon icon="ix:question-filled" height="30" width="30" />
			{/snippet}
			<UpdateQuestions
				onSave={() => {
					questionFormOpen = !questionFormOpen;
				}}
			/>
		</ResiableList>
		<ResiableList bind:isOpen={settingsFormOpen} section_name="Statystyki i ustawienia">
			{#snippet nav_icon()}
				<Icon icon="material-symbols:settings-rounded" height="30" width="30" />
			{/snippet}
			<UpdateQuizSettings
				onSave={() => {
					settingsFormOpen = !settingsFormOpen;
				}}
			/>
		</ResiableList>
	</section>
	<div class="mx-auto mt-5 flex w-6/10 justify-between text-white">
		<!-- <Button
			theme="danger"
			size="medium"
			onclick={() => {
				toast.show('Funkcja usuwania quizu nie jest jeszcze zaimplementowana.', 'info');
			}}
		>
			<div class="flex items-center justify-center">
				Usuń quiz <span class="ms-2"><Icon icon="mdi:trash-can-outline" /></span>
			</div>
		</Button> -->
	</div>
</div>
