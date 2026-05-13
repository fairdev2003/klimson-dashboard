<script lang="ts">
	import { goto } from '$app/navigation';
	import { api } from '$lib/api/api';
	import {
		addFormQuiz,
		addFormSection,
		initialFormQuiz,
		questions
	} from '$lib/dashboard/stores/store';
	import { debug } from '$lib/dashboard/stores/debug';
	import Placeholder from '$lib/assets/placeholder.png';
	import type { Quiz } from '../types';
	import MovingTooltip from '$lib/components/dashboard/MovingTooltip.svelte';
	import Icon from '@iconify/svelte';
	import { toast } from '$lib/dashboard/stores/toast';
	import { recordVisibilityDict } from './quiz_record';
	import { quiz_manager } from '../lib/quiz_manager.svelte';

	type Props = {
		quiz_record: Quiz;
		selected?: boolean;
	};

	let { quiz_record, selected = false }: Props = $props();

	function selectQuiz() {
		$questions = quiz_record.questions;
		$addFormSection = 'general';
		$addFormQuiz = quiz_record;
		debug.log(`Zaznaczono quiz: ${quiz_record.title}`);
	}

	function enterQuiz(e: MouseEvent) {
		$addFormQuiz = quiz_record;
		e.stopPropagation();
		if (selected) {
			goto('/dashboard/quizzes/update');
		}
		$addFormSection = 'general';

		$initialFormQuiz = quiz_record;

		$questions = quiz_record.questions;
	}
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
<!-- svelte-ignore a11y_click_events_have_key_events -->

<div
	ondblclick={enterQuiz}
	class:border-neutral-500={selected}
	class:border-neutral-700={!selected}
	class:bg-neutral-800={!selected}
	class:bg-neutral-700={selected}
	class="relative max-w-6xl h-[100px] flex border mx-auto overflow-hidden"
	onclick={selectQuiz}
>
	<MovingTooltip delay={0}>
		{#snippet tooltipContent()}
			<p class="text-bg-neutral-400 text-sm max-w-40">
				{quiz_record.public ? 'Quiz ustawiony na publike' : 'Niedostepny dla użytkownikow'}
			</p>
		{/snippet}
		<div
			class="absolute w-1 h-full mb-0.5 {quiz_record.public ? 'bg-blue-500' : 'bg-neutral-600'}"
		></div>
	</MovingTooltip>
	<div class="flex gap-3 p-3 w-full">
		<img
			class="h-full ml-3 object-cover rounded bg-neutral-900 aspect-square opacity-80 transition-all duration-700 group-hover:opacity-100"
			onerror={(e) => ((e.currentTarget as HTMLImageElement).src = Placeholder)}
			src={api.image.quiz(quiz_record.image_url)}
			alt={`quiz-image-${quiz_record.id}`}
			loading="lazy"
		/>
		{@render InformationBody()}
	</div>
</div>

{#snippet InformationBody()}
	<div class="flex justify-between w-full items-center">
		<MovingTooltip delay={0}>
			{#snippet tooltipContent()}
				<div class="flex flex-col gap-1 text-xs max-w-80">
					{#each Object.entries(quiz_record) as entry}
						<p class="font-bold text-blue-500">
							{entry[0]}: <span class="font-normal text-white">{entry[1]}</span>
						</p>
					{/each}
				</div>
			{/snippet}
			<div class="flex gap-1 items-center">
				{@render DifficultyColorSquares(quiz_record.difficulty)}
				<div class="flex gap-1 items-center ml-2">
					{#if $recordVisibilityDict.questions}
						<p class="text-xs font-semibold">
							Ilośś pytań: {quiz_record.questions && quiz_record.questions.length}
						</p>
					{/if}

					{#if $recordVisibilityDict.stats}
						<p class="text-xs font-semibold">
							Ilośś statystyk: {quiz_record.stats && quiz_record.stats.length}
						</p>
					{/if}
				</div>
			</div>
			<div class="flex flex-col">
				<p class="text-sm font-semibold">{quiz_record.title}</p>
				{#if $recordVisibilityDict.description}
					<p class="text-xs text-neutral-200">{quiz_record.description}</p>
				{/if}
				<p class="text-xs text-neutral-200">Autor: {quiz_record.author}</p>
			</div>
		</MovingTooltip>
		<!-- container for action buttons -->
		{#if selected}
			<div class="flex gap-2">
				<!-- enter the quiz -->
				<MovingTooltip>
					{#snippet tooltipContent()}
						<p class="text-xs">Opcje</p>
					{/snippet}
					<button
						onclick={() => {
							toast.warning('Brak dostepu wypierdalaj!');
						}}
						class="p-3 bg-neutral-500 rounded-md cursor-pointer"
					>
						<Icon icon="tabler:dots-filled" width="16" height="16" />
					</button>
				</MovingTooltip>
				<MovingTooltip>
					{#snippet tooltipContent()}
						<p class="text-xs">Usuń quiz</p>
					{/snippet}
					<button
						onclick={() => quiz_manager.PromptDeleteAndRefresh(quiz_record)}
						class="p-3 bg-red-500 rounded-md cursor-pointer"
					>
						<Icon icon="bx:trash-alt" width="16" height="16" />
					</button>
				</MovingTooltip>

				<MovingTooltip>
					{#snippet tooltipContent()}
						<p class="text-xs">Przejdź do quizu</p>
					{/snippet}
					<button onclick={enterQuiz} class="p-3 cursor-pointer bg-blue-500 rounded-md">
						<Icon icon="lsicon:triangle-right-filled" width="16" height="16" />
					</button>
				</MovingTooltip>
			</div>
		{/if}
	</div>
{/snippet}

{#snippet DifficultyColorSquares(
	difficulty: 'Łatwy' | 'Średni' | 'Trudny' | 'Bardzo trudny' | string
)}
	{@const squareAmount = { Łatwy: 1, Średni: 2, Trudny: 3, 'Bardzo trudny': 4 }[difficulty] ?? 0}
	{@const squreColor =
		{
			Łatwy: 'bg-green-500',
			Średni: 'bg-orange-500',
			Trudny: 'bg-red-500',
			'Bardzo trudny': 'bg-purple-900'
		}[difficulty] ?? null}

	<div class="flex items-center gap-1">
		<p class="text-xs">Trudność: {''}</p>
		{#if squareAmount !== 0}
			<div class="flex items-center bg-neutral-500 w-12">
				{#each Array(squareAmount) as _, i}
					{#if squreColor}
						<div class="size-3 h-2 {squreColor}"></div>
					{/if}
				{/each}
			</div>
		{:else}
			<div class="text-xs">{difficulty}</div>
		{/if}
	</div>
{/snippet}
