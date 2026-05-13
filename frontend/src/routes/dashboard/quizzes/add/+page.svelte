<script lang="ts">
	import Button from '$lib/components/Button.svelte';
	import Heading from '$lib/components/dashboard/typography/Heading.svelte';
	import Icon from '@iconify/svelte';
	import ResiableList from './component/ResiableList.svelte';
	import ImageForm from './component/form/ImageForm.svelte';
	import BasicInformationForm from './component/form/BasicInformationForm.svelte';
	import QuestionForm from './component/form/QuestionForm.svelte';
	import FormNav from '../update/components/FormNav.svelte';
	import QuizSettingsForm from './component/form/QuizSettingsForm.svelte';
	import { onMount } from 'svelte';
	import { addFormQuiz, contextMenuOptions } from '$lib/dashboard/stores/store';
	import { toast } from '$lib/dashboard/stores/toast';
	import { api } from '$lib/api/api';
	import { quizzes } from '$lib/dashboard/stores/data.store';
	import axios from 'axios';
	import { goto } from '$app/navigation';

	let imageFormOpen: boolean = $state(false);
	let basicInfoFormOpen: boolean = $state(false);
	let questionFormOpen: boolean = $state(false);
	let settingsFormOpen: boolean = $state(false);

	function resetForm() {
		if (confirm('Czy na pewno chcesz cofnąć wszystkie zmiany?')) {
			addFormQuiz.update(quiz => {
				return {...quiz, title: "", description: "", difficulty: "", author: ""}
			})
			toast.show("Dane quizu zresetowane!")
		}
	}

	onMount(() => {
		$contextMenuOptions = [
			{
				label: 'Prześlij dane z formularza',
				action: async () => {
					toast.show('Przesylanie.....');
					await Save()
					goto("/dashboard/quizzes/update")
				},
				icon: 'mdi:form',
				color: 'text-neutral-500'
			},
			{
				label: 'Anuluj zmiany',
				action: async () => {
					resetForm()
				},
				icon: 'ep:close-bold',
				color: 'text-red-500'
			}
		];
	});

	async function Save() {
		let message = '';
		try {
			const response = await api.quiz.SaveBasicInfo({
				author: $addFormQuiz.author,
				title: $addFormQuiz.title,
				description: $addFormQuiz.description,
				difficulty: $addFormQuiz.difficulty
			})
			$addFormQuiz.id = response.data.id

			if (response.data.message) {
				message = response.data.message;
			}
		} catch (error: unknown) {
			if (axios.isAxiosError(error)) {
				const message = error.response?.data?.message || 'Błąd serwera';
				toast.error(message);
			} else {
				toast.error('Wystąpił nieoczekiwany błąd');
			}
		} finally {
			const response = await api.quiz.GetAll();
			$quizzes = response.data;
			toast.success(message);
		}
	}
</script>

<!-- <QuizForm mode="add" /> -->

<div class="">
	<FormNav renderInfo={false}><Heading>Dodaj nowy quiz</Heading></FormNav>
	<section class="mx-auto mt-5 flex w-6/10 flex-col gap-5">
		<ResiableList bind:isOpen={basicInfoFormOpen} section_name="Podstawowe informacje">
			{#snippet nav_icon()}
				<Icon icon="material-symbols:chat-info" height="30" width="30" />
			{/snippet}
			<BasicInformationForm
				onSave={() => {
					basicInfoFormOpen = false;
				}}
			/>
		</ResiableList>
		<ResiableList disabled={!$addFormQuiz.id} disabledMessage="W pierwszej kolejności uzupełnij podstawowe informacje i zapisz zmiany!" bind:isOpen={imageFormOpen} section_name="Multimedia">
			{#snippet nav_icon()}
				<Icon icon="material-symbols:image" height="30" width="30" />
			{/snippet}
			<ImageForm
				onSave={() => {
					imageFormOpen = !imageFormOpen;
				}}
			/>
		</ResiableList>
		
		<ResiableList disabled={!$addFormQuiz.id} disabledMessage="W pierwszej kolejności uzupełnij podstawowe informacje i zapisz zmiany!" bind:isOpen={questionFormOpen} section_name="Pytania">
			{#snippet nav_icon()}
				<Icon icon="ix:question-filled" height="30" width="30" />
			{/snippet}
			<QuestionForm
				onSave={() => {
					questionFormOpen = !questionFormOpen;
				}}
			/>
		</ResiableList>
		<ResiableList disabled={!$addFormQuiz.id} disabledMessage="W pierwszej kolejności uzupełnij podstawowe informacje i zapisz zmiany!" bind:isOpen={settingsFormOpen} section_name="Statystyki i ustawienia">
			{#snippet nav_icon()}
				<Icon icon="material-symbols:settings-rounded" height="30" width="30" />
			{/snippet}
			<QuizSettingsForm
				onSave={() => {
					settingsFormOpen = !settingsFormOpen;
				}}
			/>
		</ResiableList>
	</section>
</div>

<svelte:window
	on:beforeunload={(e) => {
		e.preventDefault();
		e.returnValue = '';
	}}
/>
