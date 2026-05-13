<script lang="ts">
	import {
		addFormQuiz,
		selectedQuiz,
		selectedSummaryType,
		summary_open
	} from '$lib/dashboard/stores/store';
	import { X } from '@lucide/svelte';
	import QuizSummary from './QuizSummary.svelte';
</script>

<div class="flex w-full flex-col gap-4 overflow-hidden">
	<!-- kontrolowanie przeplywu sekcji -->
	{#if $selectedSummaryType === 'quiz'}
		<QuizSummary />
	{/if}
	{#if $selectedSummaryType === 'debug'}
		<button
			onclick={() => {
				$summary_open = false;
			}}
			class="m-4 flex gap-2 bg-red-500"
		>
			<p>Zamknij</p>
			<X />
		</button>
		<div class="p-4">
			<p>
				Title: <span class="text-green-500"
					>{($addFormQuiz && $addFormQuiz.title) || '<ni mo>'}</span
				>
			</p>
			<p>
				Desc: <span class="text-green-500"
					>{($addFormQuiz && $addFormQuiz.description) || '<ni mo>'}</span
				>
			</p>
			<p>
				Autor: <span class="text-green-500"
					>{($addFormQuiz && $addFormQuiz.author) || '<ni mo>'}</span
				>
			</p>
			<p>
				Poziom: <span class="text-green-500"
					>{($addFormQuiz && $addFormQuiz.difficulty) || '<ni mo>'}</span
				>
			</p>
			<p>
				URL Obrazka: <span class="text-green-500">{$addFormQuiz && $addFormQuiz.image_url}</span>
			</p>
			<p>
				Publiczny: <span class="text-green-500"
					>{$addFormQuiz && ($addFormQuiz.public ? 'TAK' : 'NIE')}</span
				>
			</p>
			<p>
				Limit czasu: <span class="text-green-500"
					>{$addFormQuiz && ($addFormQuiz.has_time_limit ? 'TAK' : 'NIE')}</span
				>
			</p>
		</div>
	{/if}
</div>
