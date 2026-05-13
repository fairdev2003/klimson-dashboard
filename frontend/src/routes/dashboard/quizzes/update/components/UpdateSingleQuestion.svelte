<script lang="ts">
	import { blur, slide } from 'svelte/transition';
	import type { Question } from '../../types';
	import Placeholder from '$lib/assets/placeholder.png';
	import { api } from '$lib/api/api';
	import CreateFormInput from '$lib/components/dashboard/CreateFormInput.svelte';
	import { question_manager } from '../../lib/question_manager.svelte';
	import { quiz_manager } from '../../lib/quiz_manager.svelte';
	import { addFormQuiz } from '$lib/dashboard/stores/store';
	import { toast } from '$lib/dashboard/stores/toast';
	import ImageContainer from '$lib/components/dashboard/add/ImageContainer.svelte';

	type Props = {
		question_record: Question;
		index: number;
	};

	let { question_record, index }: Props = $props();

	function addAnswer() {
		const qId = quiz_manager.get_question_manager.selectedQuestionId;

		const newAnswer = {
			content: '',
			is_correct: false
		};

		addFormQuiz.update((quiz) => {
			if (quiz.questions[qId].answers.length === 5) {
				toast.warning('Nie mozesz dodac wiecej niz 5 pytań ty cwelu');
				return quiz;
			}
			if (!quiz.questions[qId].answers) {
				quiz.questions[qId].answers = [];
			}
			quiz.questions[qId].answers.push(newAnswer);
			return quiz;
		});
	}
</script>

<div transition:blur class="mt-4 flex flex-col gap-2 bg-neutral-900/70 p-4">
	{#if $addFormQuiz.questions[quiz_manager.get_question_manager.selectedQuestionId].id}
		<CreateFormInput
			label="ID"
			disabled
			bind:value={$addFormQuiz.questions[quiz_manager.get_question_manager.selectedQuestionId].id}
		/>
	{/if}

	<CreateFormInput
		label="Tresc pytania"
		bind:value={
			$addFormQuiz.questions[quiz_manager.get_question_manager.selectedQuestionId].content
		}
	/>
	<div>
		<ImageContainer
			bind:src={
				$addFormQuiz.questions[quiz_manager.get_question_manager.selectedQuestionId].image_url
			}
			type="question"
		/>
	</div>
	<p class="text-xs font-bold text-neutral-500 uppercase mt-5">Odpowiedzi</p>
	{#each $addFormQuiz.questions[quiz_manager.get_question_manager.selectedQuestionId].answers as _, answerIndex}
		<div class="flex gap-2 items-center" transition:slide>
			<div class="flex-1">
				<CreateFormInput
					label="Odpowiedź {answerIndex + 1}"
					bind:value={
						$addFormQuiz.questions[quiz_manager.get_question_manager.selectedQuestionId].answers[
							answerIndex
						].content
					}
				/>
			</div>

			<input
				type="checkbox"
				class="checkbox checkbox-primary"
				bind:checked={
					$addFormQuiz.questions[quiz_manager.get_question_manager.selectedQuestionId].answers[
						answerIndex
					].is_correct
				}
			/>

			<button
				class="text-red-500"
				onclick={() => {
					$addFormQuiz.questions[
						quiz_manager.get_question_manager.selectedQuestionId
					].answers.splice(answerIndex, 1);
					$addFormQuiz = $addFormQuiz; // Wyzwalacz reaktywności w storze
				}}
			>
				✕
			</button>
		</div>
	{/each}

	{#if $addFormQuiz.questions[quiz_manager.get_question_manager.selectedQuestionId].answers.length < 5}
		<button
			transition:blur
			class="btn btn-outline btn-sm mt-2 p-2 border border-neutral-700 cursor-pointer border-dashed hover:bg-neutral-800"
			onclick={addAnswer}
		>
			+ Dodaj odpowiedź
		</button>
	{/if}
</div>
