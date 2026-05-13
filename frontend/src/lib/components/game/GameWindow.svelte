<script lang="ts">
	import { api } from '$lib/api/api';
	import { onMount } from 'svelte';
	import type { Quiz } from '../../../routes/dashboard/quizzes/types';
	import { fade } from 'svelte/transition';
	import Right from '$lib/assets/right.mp3';
	import Wrong from '$lib/assets/wrong.mp3';
	import { debug } from '$lib/dashboard/stores/debug';
	import { Check, Cross, X } from '@lucide/svelte';
	import PageLoading from '../PageLoading.svelte';

	type GameModalProps = {
		quiz: Quiz | undefined;
		showCorrect?: boolean;
		previewMode?: boolean;
	};

	const { quiz, showCorrect = false, previewMode = false }: GameModalProps = $props();

	let correctSound: HTMLAudioElement;
	let wrongSound: HTMLAudioElement;

	let questionIndex = $state(0);
	let selectedAnswerId = $state<number | undefined>();
	let isCorrect = $state<boolean | null>(null);
	let isLocked = $state(false);
	let isCompleted = $state(false);
	let score = $state(0);
	let time_limit = $state(0);
	let canPlaySound = $state(true);

	let timer: ReturnType<typeof setInterval> | null = null;

	onMount(() => {
		correctSound = new Audio(Right);
		wrongSound = new Audio(Wrong);
		correctSound.load();
		wrongSound.load();
	});

	function playCorrect() {
		if (!canPlaySound) return;
		canPlaySound = false;
		correctSound.currentTime = 0;
		correctSound.play();
	}

	function playWrong() {
		if (!canPlaySound) return;
		canPlaySound = false;
		wrongSound.currentTime = 0;
		wrongSound.play();
	}

	function startTimer(limit: number) {
		if (timer) clearInterval(timer);

		time_limit = limit;
		if (!limit || limit <= 0) return;

		timer = setInterval(() => {
			if (time_limit <= 1) {
				clearInterval(timer!);
				timer = null;
				onTimeExpired();
			} else {
				time_limit -= 1;
			}
		}, 1000);
	}

	function onTimeExpired() {
		if (isLocked) return;

		isLocked = true;
		playWrong();

		setTimeout(async () => {
			const isLast = questionIndex + 1 === quiz!.questions.length;

			if (isLast) {
				isCompleted = true;
				await api.stats.NewStat({ id: quiz!.id });
			} else {
				questionIndex++;
				canPlaySound = true;
			}

			selectedAnswerId = null;
			isCorrect = null;
			isLocked = false;
		}, 1200);
	}

	$effect(() => {
		if (!quiz) return;

		const question = quiz.questions[questionIndex];
		startTimer(question.time_limit);

		return () => {
			if (timer) {
				clearInterval(timer);
				timer = null;
			}
		};
	});

	$effect(() => {
		if (quiz && questionIndex + 1 < quiz.questions.length) {
			const img = new Image();
			img.src = api.image.getImage('question', quiz.questions[questionIndex + 1].image_url);
		}
	});
</script>

{#if quiz && !isCompleted}
	<div
		class={`flex h-full w-full items-center justify-center ${previewMode ? 'bg-transparent' : 'bg-zinc-800'}`}
	>
		<div
			class="w-5xl flex flex-col items-center gap-6"
			in:fade={{ duration: 200 }}
			out:fade={{ duration: 200 }}
		>
			<p class="text-white">
				Pytanie {questionIndex + 1}/{quiz.questions.length}
			</p>

			{#if time_limit > 0}
				<div class="text-xl font-bold text-white">
					⏱️ {time_limit}s
				</div>
			{/if}

			<img
				class="rounded-xl {previewMode ? '' : ''} "
				src={api.image.getImage('question', quiz.questions[questionIndex].image_url)}
			/>

			<h3 class="text-center text-2xl text-white">
				{quiz.questions[questionIndex].content}
			</h3>

			<div class="grid w-full grid-cols-2 gap-6">
				{#each quiz.questions[questionIndex].answers as { content, id, is_correct }}
					<button
						onclick={() => {
							if (isLocked) return;

							if (timer) {
								clearInterval(timer);
								timer = null;
							}

							isLocked = true;
							selectedAnswerId = id;
							isCorrect = is_correct;

							if (is_correct) {
								score++;
								playCorrect();
							} else {
								playWrong();
							}

							setTimeout(async () => {
								const isLast = questionIndex + 1 === quiz.questions.length;

								if (isLast) {
									isCompleted = true;
									await api.stats.NewStat({ id: quiz.id });
								} else {
									questionIndex++;
									canPlaySound = true;
								}
							}, 1500);

							setTimeout(() => {
								selectedAnswerId = null;
								isCorrect = null;
								isLocked = false;
							}, 500);
						}}
						class="
							flex w-full cursor-pointer gap-3 rounded-lg p-4 font-medium transition-all duration-300
							{isLocked ? 'pointer-events-none' : ''}
							{previewMode ? '' : ''}
							{selectedAnswerId === id && isCorrect === true ? 'scale-105 bg-green-500 text-white' : ''}
							{selectedAnswerId === id && isCorrect === false ? 'scale-105 bg-red-500 text-white' : ''}
							{selectedAnswerId !== id ? 'bg-white text-black hover:opacity-80' : ''}
						"
					>
						{#if is_correct && showCorrect}
							<Check class="text-green-700" />
						{/if}

						{#if showCorrect && !is_correct}
							<X class="text-red-500" />
						{/if}
						<p>{content}</p>
					</button>
				{/each}
			</div>
		</div>
	</div>
{:else}
	<div class="flex h-full w-full items-center justify-center bg-zinc-800">
		{#if isCompleted}
			{@render CompletedScreen()}
		{:else}
			<PageLoading loading_title="Ładowanie" />
		{/if}
	</div>
{/if}

{#snippet CompletedScreen()}
	<div
		class="flex h-screen w-screen flex-col items-center justify-center gap-6 bg-zinc-800 text-white"
		in:fade={{ duration: 300 }}
		out:fade={{ duration: 300 }}
	>
		<h2 class="text-4xl font-bold">🎉 Quiz ukończony</h2>

		<p class="text-2xl">
			Wynik: <span class="font-bold">{score}</span> / {quiz?.questions.length}
		</p>

		<p class="opacity-80">
			Skuteczność: {Math.round((score / quiz!.questions.length) * 100)}%
		</p>

		<button
			class="mt-6 rounded-lg bg-green-500 px-6 py-3 text-lg font-semibold text-black hover:bg-green-400"
			onclick={() => {
				questionIndex = 0;
				score = 0;
				isCompleted = false;
				canPlaySound = true;
			}}
		>
			Zagraj ponownie
		</button>
	</div>
{/snippet}
