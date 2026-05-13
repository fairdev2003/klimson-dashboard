<script lang="ts">
	import {
		addFormQuiz,
		addFormSection,
		questions,
		selectedQuiz,
		selectedSummaryType,
		summary_open
	} from '$lib/dashboard/stores/store';
	import { Ampersands, Delete, Library, X } from '@lucide/svelte';
	import Example from '$lib/assets/example.jpeg';
	import { Api, api } from '$lib/api/api';
	import Dashboard from '../../../../routes/dashboard/dashboard.svelte';
	import gsap from 'gsap';
	import { tick } from 'svelte';
	import Loader from '../Loader.svelte';
	import type { Question, Quiz } from '../../../../routes/dashboard/quizzes/types';
	import type { AxiosResponse } from 'axios';
	import Button from '$lib/components/Button.svelte';
	import { goto } from '$app/navigation';
	import { quizzes } from '$lib/dashboard/stores/data.store';

	let deleteInputValue: string = $state('');
	let deleteModalOpened: boolean = $state(false);
	let deleteState: string = $state('none');
	let modalEl: HTMLDivElement | undefined = $state();
	let loading: boolean = $state(false);

	async function openModal() {
		deleteModalOpened = true;
		await tick();

		if (!modalEl) {
			return;
		}

		gsap.fromTo(
			modalEl,
			{ scaleY: 0.2, scaleX: 0.15, transformOrigin: 'bottom', opacity: 0 },
			{ scaleY: 1, scaleX: 1, duration: 0.4, opacity: 1, ease: 'power2.out' }
		);
	}

	async function closeModal() {
		await tick();

		if (!modalEl) {
			return;
		}

		gsap.to(modalEl, {
			scaleY: 0.2,
			scaleX: 0.15,
			opacity: 0,
			transformOrigin: 'bottom',
			duration: 0.3,
			ease: 'power2.in',
			onComplete: () => {
				deleteModalOpened = false;
			}
		});
	}

	async function DeleteMe() {
		deleteState = 'pending';
		if (deleteInputValue.length === 0) return;

		const response = await api.quiz.DeleteQuiz(deleteInputValue);
		console.log(response);

		deleteState = 'none';
		$selectedQuiz = undefined;
		$summary_open = false;

		$quizzes = (await api.quiz.GetAll()).data
		
	}

	const sumTime = (questions: Question[]): string => {
		let totalTime: number = 0;
		let totalTimeString: string = 'Brak danych';
		questions.map((question) => {
			if (question.time_limit) {
				totalTime = totalTime + question.time_limit;
			}
		});

		if (totalTime) {
			return `${totalTime}`;
		} else {
			return totalTimeString;
		}
	};
</script>

{#if $summary_open && $selectedQuiz}
	<div class="flex w-full flex-col gap-4 overflow-hidden p-4">
		{@render SummaryHeading()}
		{@render Iksior()}
		{@render QuizIcon()}

		<div>
			<Button
				onclick={() => {
					if (!$selectedQuiz) return;
					window.open(`/quiz?id=${$selectedQuiz.id}`, '_blank', 'noopener,noreferrer');
				}}
				className="w-full"
				loading={false}
				loadingText="Ładowanie komponentów"
				theme="correct">Podgląd</Button
			>
		</div>

		<div class="border-t-1 flex flex-col border-neutral-700/60">
			{@render SummaryItemKeyValue('Id quizu', $selectedQuiz?.id)}
			{@render SummaryItemKeyValue('Nazwa quizu', $selectedQuiz?.title)}
			{@render SummaryItemKeyValue('Opis', $selectedQuiz?.description)}
			{@render SummaryItemKeyValue('Poziom trudności', $selectedQuiz?.difficulty)}
			{@render SummaryItemKeyValue('Autor', $selectedQuiz?.author)}
			{@render SummaryItemKeyValue('Publiczny', $selectedQuiz?.public ? 'Tak' : 'Nie')}
			{@render SummaryItemKeyValue(
				'Szacowany czas do rozwiązania',
				sumTime($selectedQuiz?.questions)
			)}
			{@render SummaryItemKeyValue('Ilość pytań', $selectedQuiz?.questions.length)}
		</div>
		<div class="flex justify-between gap-3">
			<Button
				{loading}
				onclick={async () => {
					if (!$selectedQuiz) {
						return;
					}
					// const response = await api.quiz.GetAdminQuiz({ id: $selectedQuiz.id });

					// if (response.id) {
					// 	loading = false;
					// } else {
					// 	return;
					// }

					$addFormQuiz = $selectedQuiz;
					$addFormSection = 'general';
					$questions = $selectedQuiz.questions;
					goto('/dashboard/quizzes/update');
				}}
				theme="secondary">Edytuj</Button
			>
			<button
				onclick={async () => {
					await openModal();
				}}
				class="cursor-pointer bg-red-500 px-5 py-2 text-sm hover:bg-red-500/80">Usuń</button
			>
		</div>
	</div>
{:else}
	<div class="mx-auto mt-8">
		<Loader />
	</div>
{/if}

{#snippet Iksior()}{/snippet}

{#snippet SummaryHeading()}
	<div class="relative flex w-full items-center justify-between">
		<div class="flex items-center gap-2 text-neutral-400">
			<div>
				<Library />
			</div>
			<h3 class="w-full text-[14px] font-medium text-neutral-400">Podsumowanie quizu</h3>
		</div>

		<button
			onclick={() => {
				$selectedQuiz = undefined;
				$summary_open = !$summary_open;
				$selectedSummaryType = undefined;
			}}
			class="cursor-pointer text-neutral-400 hover:text-neutral-200"
		>
			<X />
		</button>
	</div>
{/snippet}

{#snippet QuizIcon()}
	<div class="flex aspect-video h-[150px] w-full items-center justify-center">
		{#if $selectedQuiz}
			<img
				src={api.image.getImage('quiz', $selectedQuiz.image_url)}
				class="mx-auto object-contain lg:size-3/4"
				alt={$selectedQuiz?.image_url}
			/>
		{/if}
	</div>
{/snippet}

{#snippet QuizName()}
	<div class="flex items-center gap-2">
		<h2 class="text-lg font-semibold">
			{$selectedQuiz?.title}
		</h2>
	</div>
{/snippet}

{#snippet SummaryItemKeyValue(key, value)}
	<div class="flex flex-col">
		<div class="border-x-1 border-b-1 w-full border-neutral-700/60 bg-neutral-800/60 px-2">
			<h3 class="font-semibold">{key}</h3>
		</div>
		<div class="border-x-1 border-b-1 w-full border-neutral-700/60 p-2">
			<p>{value}</p>
		</div>
	</div>
{/snippet}

{#if deleteModalOpened}
	{@render DeleteModal()}
{/if}

{#snippet DeleteModal()}
	<!-- svelte-ignore a11y_click_events_have_key_events -->
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<div
		onclick={async () => await closeModal()}
		class="fixed inset-0 z-50 flex items-center justify-center bg-black/50 text-white md:backdrop-blur-lg lg:backdrop-blur-lg"
	>
		<!-- svelte-ignore a11y_no_static_element_interactions -->
		<div
			onclick={(e) => {
				e.stopPropagation();
			}}
			bind:this={modalEl}
			class="w-9/10 relative flex flex-col border border-neutral-800/60 bg-neutral-950 md:w-1/2 md:min-w-[400px] md:backdrop-blur-none lg:w-1/4 lg:bg-neutral-950"
		>
			<!-- title -->
			<div
				class="mb-2 flex h-10 flex-shrink-0 items-center justify-between border-b border-neutral-700/60 bg-neutral-800/60 px-5"
			>
				<p>Usuwanie</p>
				<button
					class="cursor-pointer text-neutral-500 hover:text-neutral-300"
					onclick={async () => {
						await closeModal();
					}}
				>
					<X />
				</button>
			</div>
			<!-- content -->
			<div class="scrollable flex-1 overflow-y-auto p-6 pt-5">
				<p>Z uwagi na to ze usuwasz cały quiz wpisz liczbę {`${$selectedQuiz?.id}`}</p>
				<div class=" mt-5">
					<div class="">
						<input
							bind:value={deleteInputValue}
							class="border-1 w-40 border-neutral-700/60 bg-transparent placeholder-neutral-700"
							placeholder="..."
						/>
					</div>
				</div>
				<div class="mt-5 flex justify-end">
					<Button
						theme="danger"
						loading={deleteState === 'pending'}
						onclick={async () => {
							deleteState = 'pending';
							const promise = await DeleteMe();

							deleteModalOpened = false;
							deleteState = 'none';
							goto('/dashboard/quizzes');
						}}
					>
						<p>Usuń</p>
					</Button>
				</div>
			</div>
		</div>
	</div>
{/snippet}
