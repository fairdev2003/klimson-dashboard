<script lang="ts">
	import { addFormQuiz } from '$lib/dashboard/stores/store';
	import Button from '$lib/components/Button.svelte';
	import { toast } from '$lib/dashboard/stores/toast';
	import { api } from '$lib/api/api';
	import axios from 'axios';
	import QuestionsImport from '$lib/components/dashboard/add/QuestionsImport.svelte';
	import { quizzes } from '$lib/dashboard/stores/data.store';
	import QuestionRecord from './QuestionRecord.svelte';
	import { quiz_manager } from '../../lib/quiz_manager.svelte';

	type Props = {
		onSave?: () => void;
	};

	let { onSave }: Props = $props();
	let loading: boolean = $state(false);
	let modalOpen: boolean = $state(false);
	let loadedImages: boolean[] = $state([]);
	let loadedImagesCount: number = $derived.by(() => {
		let count = loadedImages.filter((loaded) => loaded === true).length;
		return count;
	});
	let infoModalOpen: boolean = $state(false);
	let updateQuestionModalOpen: boolean = $state(false);
	let deleteConfiationOpen: boolean = $state(false);

	async function handleSave() {
		loading = true;

		const response = await api.question.UpdateMany($addFormQuiz.questions);
		console.log(response);

		loading = false;
		toast.show(response.data.message, 'success');
		onSave?.();
	}

	function DeleteQuestionFront(questionId: number) {
		$addFormQuiz.questions = $addFormQuiz.questions.filter((s) => s.id !== questionId);
	}

	async function DeleteQuestionFromServer(id: number) {
		loading = true;

		let message = '';
		onSave?.();
		try {
			const response = await api.question.DeleteQuestion(id);

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
			loading = false;
			toast.success(message);
		}
	}
</script>

<div class="flex flex-col gap-4">
	<div class="flex gap-2">
		<Button
			size="small"
			theme="secondary"
			onclick={() => {
				infoModalOpen = true;
			}}>Pokaż informacje o pytaniach</Button
		>
		<QuestionsImport />
	</div>

	<div class="flex w-full flex-col gap-3 mt-10">
		{#each $addFormQuiz.questions as question, index}
			<QuestionRecord
				{index}
				question_record={question}
				selected={index === quiz_manager.get_question_manager.selectedQuestionId}
				onclick={() => {
					quiz_manager.get_question_manager.toggleSelectedQuestionIndex(index);
				}}
			/>
		{/each}
		{@render AddQuestion()}
	</div>

	<div class="flex gap-4"></div>

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

{#snippet AddQuestion()}
	<div class="flex gap-4">
		<button
			onclick={() => {
				modalOpen = true;
				const newQuestion = {
					answers: [],
					content: 'Siema',
					image_url: '',
					quiz_id: Number($addFormQuiz.id),
					time_limit: 30,
					type: ''
				};
				addFormQuiz.update((quiz) => {
					quiz.questions.push(newQuestion);
					return quiz;
				});

				$addFormQuiz.questions = [...$addFormQuiz.questions];
			}}
			class={`mt-5 flex h-[50px] w-8/10 cursor-pointer items-center justify-center gap-3 border border-neutral-700 bg-neutral-900  transition-colors hover:bg-neutral-700`}
		>
			Wybierz szablon
		</button>
		<button
			onclick={() => {
				modalOpen = true;
				const newQuestion = {
					answers: [],
					content: 'Siema',
					image_url: '',
					quiz_id: Number($addFormQuiz.id),
					time_limit: 30,
					type: ''
				};
				addFormQuiz.update((quiz) => {
					quiz.questions.push(newQuestion);
					return quiz;
				});

				$addFormQuiz.questions = [...$addFormQuiz.questions];
			}}
			class={`mt-5 flex h-[50px] w-2/10 cursor-pointer items-center justify-center gap-3 border border-neutral-700 bg-neutral-900  transition-colors hover:bg-neutral-700`}
		>
			Dodaj pytanie
		</button>
	</div>
{/snippet}
