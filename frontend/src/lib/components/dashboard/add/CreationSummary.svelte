<script lang="ts">
	import { goto } from '$app/navigation';
	import { api } from '$lib/api/api';
	import Button from '$lib/components/Button.svelte';
	import { quizzes } from '$lib/dashboard/stores/data.store';
	import { addFormQuiz, addFormSection, questions } from '$lib/dashboard/stores/store';
	import { preloaded_quiz_form } from '$lib/static/quiz';
	import QuizFormNavigation from './QuizFormNavigation.svelte';

	function calcTime(): number {
		let time: number = 0;
		$addFormQuiz.questions.map((question) => {
			if (!question.time_limit) return;
			time = time + question.time_limit;
		});
		return time;
	}

	let loading: boolean = $state(false);

	type Props = {
		mode: 'add' | 'edit';
	};

	const { mode = 'add' }: Props = $props();
</script>

<div class="flex flex-col gap-5 py-4">
	{@render SummaryItemKeyValue('Nazwa quizu', $addFormQuiz.title || 'Brak nazwy quizu')}
	{@render SummaryItemKeyValue('Opis quizu', $addFormQuiz.description || 'Brak opisu')}
	{@render SummaryItemKeyValue('Ilość pytań', $addFormQuiz.questions.length || 'Brak pytań')}
	{@render SummaryItemKeyValue('Poziom', $addFormQuiz.difficulty || 'Nie ustawiono poziomu')}
	{@render SummaryItemKeyValue('Autor', $addFormQuiz.author || 'Anonim')}
	{@render SummaryItemKeyValue('Publiczny', $addFormQuiz.public ? 'Tak' : 'Nie')}
	{@render SummaryItemKeyValue('Ustawiono limit czasu', calcTime() + ' sekund')}
	<div class="flex justify-end">
		<Button
			theme="secondary"
			{loading}
			onclick={async () => {
				loading = true;
				if (mode === 'add') {
					const response = await api.quiz.CreateQuiz($addFormQuiz);
					if (response) {
						loading = false;
						$addFormQuiz = preloaded_quiz_form;
						const res = (await api.quiz.GetAll()).data;
						if (res.length > 0) {
							$quizzes = res;
							goto('/dashboard/quizzes');
						}
					}
				}
				if (mode === 'edit') {
					console.log($addFormQuiz.id);
					const response = await api.quiz.UpdateQuiz({ id: $addFormQuiz.id }, $addFormQuiz);
					if (response) {
						loading = false;
						$addFormQuiz = preloaded_quiz_form;
						goto('/dashboard/quizzes');
					}
				}
			}}
		>
			Prześlij quiz
		</Button>
	</div>
</div>

{#snippet SummaryItemKeyValue(key, value)}
	<div class="flex flex-col">
		<div
			class="border-x-1 border-b-1 flex h-10 w-full items-center justify-start border-neutral-700/60 bg-neutral-800/60 px-2"
		>
			<h3 class="font-semibold">{key}</h3>
		</div>
		<div class="border-x-1 border-b-1 min-h-10 w-full border-neutral-700/60 p-2">
			<p>{value}</p>
		</div>
	</div>
{/snippet}
