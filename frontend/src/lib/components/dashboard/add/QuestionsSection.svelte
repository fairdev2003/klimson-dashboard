<script lang="ts">
	import { Key, Plus, X } from '@lucide/svelte';
	import QuestionWorkspaceCard from '../QuestionWorkspaceCard.svelte';
	import { tick } from 'svelte';
	import gsap from 'gsap';
	import {
		addFormQuiz,
		addFormSection,
		questions,
		selectedQuestionToEdit,
		updateFormQuiz
	} from '$lib/dashboard/stores/store';
	import CreateFormInput from '../CreateFormInput.svelte';
	import type { Answer, Question } from '../../../../routes/dashboard/quizzes/types';
	import { emptyQuestion } from '$lib/static/quiz';
	import QuizFormNavigation from './QuizFormNavigation.svelte';
	import ImageContainer from './ImageContainer.svelte';
	import Button from '$lib/components/Button.svelte';
	import { Api, api } from '$lib/api/api';
	import { get } from 'svelte/store';
	import { debug } from '$lib/dashboard/stores/debug';
	import { toast } from '$lib/dashboard/stores/toast';
	import QuestionsImport from './QuestionsImport.svelte';
	import Tooltip from '../Tooltip.svelte';
	import Modal from '$lib/components/Modal.svelte';

	type AddPageMode = 'edit' | 'add' | 'none';

	let modal_opened: boolean = $state(false);
	let delete_modal_opened: boolean = $state(false);
	let modalCreator: boolean = $state(false);
	let modalOpenMode: AddPageMode = $state('none');
	let addModalEl: HTMLDivElement | undefined = $state();
	let deleteModalEl: HTMLDivElement | undefined = $state();

	async function OpenDeleteModal() {
		delete_modal_opened = true;
		document.body.style.overflow = 'hidden';
		await tick();
		if (!deleteModalEl) return;

		gsap.fromTo(
			deleteModalEl,
			{ scaleY: 0.2, scaleX: 0.15, transformOrigin: 'bottom', opacity: 0 },
			{ scaleY: 1, scaleX: 1, duration: 0.4, opacity: 1, ease: 'power2.out' }
		);
	}

	async function CloseDeleteModal() {
		await tick();
		if (!deleteModalEl) return;

		gsap.to(deleteModalEl, {
			scaleY: 0.2,
			scaleX: 0.15,
			opacity: 0,
			transformOrigin: 'bottom',
			duration: 0.3,
			ease: 'power2.in',
			onComplete: () => {
				document.body.style.overflow = 'auto';
				delete_modal_opened = false;
			}
		});
	}

	let selected_question_id = $state<number>(0);
	let loading: boolean = $state(false);

	$effect(() => {
		if (modalCreator) {
			OpenModal('add');
		}
	});

	async function OpenModal(mode: AddPageMode) {
		modal_opened = true;
		document.body.style.overflow = 'hidden';
		modalOpenMode = mode;
		await tick();
		if (!addModalEl) return;

		gsap.fromTo(
			addModalEl,
			{ scaleY: 0.2, scaleX: 0.15, transformOrigin: 'bottom', opacity: 0 },
			{ scaleY: 1, scaleX: 1, duration: 0.4, opacity: 1, ease: 'power2.out' }
		);
	}

	async function CloseModal() {
		await tick();
		if (!addModalEl) return;

		gsap.to(addModalEl, {
			scaleY: 0.2,
			scaleX: 0.15,
			opacity: 0,
			transformOrigin: 'bottom',
			duration: 0.3,
			ease: 'power2.in',
			onComplete: () => {
				document.body.style.overflow = 'auto';
				modal_opened = false;
				modalOpenMode = 'none';
			}
		});
	}

	type Props = {
		mode?: 'add' | 'edit';
	};

	const { mode = 'add' }: Props = $props();
</script>

<div class="mb-10 mt-5 flex flex-col gap-y-5">
	<div class="flex gap-4">
		<Tooltip content="Dodaj nowe pytanie">
			<Button
				onclick={async () => {
					$selectedQuestionToEdit = structuredClone(emptyQuestion);
					modalOpenMode = 'add';
					modal_opened = true;
				}}>Dodaj pytanie</Button
			>
		</Tooltip>
		<QuestionsImport />
	</div>
	<div class="grid grid-cols-3 gap-5">
		{#each $addFormQuiz.questions as question, i}
			<div class="col-span-3 lg:col-span-1">
				<QuestionWorkspaceCard
					{question}
					onclick={async (q) => {
						selected_question_id = i;
						$selectedQuestionToEdit = q;
						modalOpenMode = 'edit';
						modal_opened = true;
					}}
				/>
			</div>
		{/each}
	</div>
</div>

{#snippet AddNewQuestionButton()}
	<Tooltip content="Dodaj nowe pytanie">
		<button
			onclick={async () => {
				$selectedQuestionToEdit = structuredClone(emptyQuestion);
				await OpenModal('add');
			}}
			class="button col-span-3 flex min-h-[250px] w-full cursor-pointer items-center justify-center transition-colors lg:col-span-1"
		>
			<div><Plus /></div>
		</button>
	</Tooltip>
{/snippet}

<Modal
	bind:opened={modal_opened}
	onClose={() => {
		modal_opened = false;
	}}
	title="Dodaj nowy quiz"
	><div class="flex flex-1 flex-col justify-between overflow-y-auto">
		<div class="overflow-auto px-6 py-6">
			<p class="mb-6 text-center font-semibold">Pytanie</p>
			<div class="mt-5 grid grid-cols-10 gap-4">
				<div class="col-span-10 lg:col-span-4">
					<ImageContainer type="question" src={$selectedQuestionToEdit.image_url} />
				</div>
				<div class="col-span-10 lg:col-span-6">
					<CreateFormInput bind:value={$selectedQuestionToEdit.content} label="Treść pytania" />
					<CreateFormInput
						bind:value={$selectedQuestionToEdit.time_limit}
						label="Limit czasu (sek)"
					/>
					<CreateFormInput bind:value={$selectedQuestionToEdit.type} label="Typ pytania" />
				</div>
			</div>

			<!-- ANSWERS -->
			<div class="mt-6">
				<p class="mb-6 text-center font-semibold">Odpowiedzi</p>

				{#each $selectedQuestionToEdit.answers as ans, i}
					<div class="mb-3 border border-neutral-700/50 p-3">
						<CreateFormInput bind:value={ans.content} label={'Odpowiedź ' + (i + 1)} />

						<div class="mt-2 flex items-center gap-2">
							<input type="checkbox" bind:checked={ans.is_correct} class="h-4 w-4" />
							<p class="text-sm">Poprawna odpowiedź</p>
						</div>

						<button
							onclick={() => {
								$selectedQuestionToEdit.answers = $selectedQuestionToEdit.answers.filter(
									(_, idx) => idx !== i
								);
							}}
							class="mt-3 text-xs text-red-400 hover:text-red-200"
						>
							Usuń odpowiedź
						</button>
					</div>
				{/each}

				<button
					onclick={() => {
						if ($selectedQuestionToEdit.answers.length >= 5) {
							alert('Możesz dodać maksymalnie 5 odpowiedzi.');
							return;
						}

						$selectedQuestionToEdit.answers = [
							...$selectedQuestionToEdit.answers,
							{ content: '', is_correct: false }
						];
					}}
					class="button mt-3 w-full cursor-pointer"
				>
					+ Dodaj odpowiedź
				</button>
				<p class="my-4 text-center font-semibold">Szybki wybór</p>
				<button
					onclick={() => {
						if ($selectedQuestionToEdit.answers.length >= 5) {
							alert('Możesz dodać maksymalnie 5 odpowiedzi.');
							return;
						}

						$selectedQuestionToEdit.answers = [
							{ content: 'Prawda', is_correct: true },
							{ content: 'Fałsz', is_correct: false }
						];
					}}
					class="button w-full cursor-pointer"
				>
					Prawda
				</button>
				<button
					onclick={() => {
						if ($selectedQuestionToEdit.answers.length >= 5) {
							alert('Możesz dodać maksymalnie 5 odpowiedzi.');
							return;
						}

						$selectedQuestionToEdit.answers = [
							{ content: 'Prawda', is_correct: false },
							{ content: 'Fałsz', is_correct: true }
						];
					}}
					class="button mt-3 w-full cursor-pointer"
				>
					Fałsz
				</button>
			</div>
		</div>

		<div
			class:justify-between={modalOpenMode === 'edit'}
			class:justify-end={modalOpenMode === 'add'}
			class="border-t-1 border-secondary flex p-6 text-sm"
		>
			{#if modalOpenMode === 'edit'}
				<Button
					theme="danger"
					onclick={() => {
						OpenDeleteModal();
					}}
				>
					Usuń
				</Button>
			{/if}

			<Button
				{loading}
				onclick={async () => {
					loading = true;
					await setTimeout(async () => {
						if (modalOpenMode === 'add') {
							$selectedQuestionToEdit.time_limit = Number($selectedQuestionToEdit.time_limit);
							console.log($selectedQuestionToEdit.time_limit);
							$questions = [...$questions, $selectedQuestionToEdit];

							if ($addFormQuiz.id != undefined) {
								const response = await api.question.CreateQuestion(
									$selectedQuestionToEdit,
									$addFormQuiz.id
								);

								console.log(response);
								toast.show(`Dodano nowe pytanie w ${response.duration}ms`, 'success');
							}

							$addFormQuiz = { ...$addFormQuiz, questions: $questions };
							loading = false;
							modal_opened = false;
						}

						if (modalOpenMode === 'edit') {
							$selectedQuestionToEdit.time_limit = Number($selectedQuestionToEdit.time_limit);
							console.log($selectedQuestionToEdit.time_limit);
							$questions = $questions.map((q, index) =>
								index === selected_question_id ? $selectedQuestionToEdit : q
							);
							$addFormQuiz = { ...$addFormQuiz, questions: $questions };
							console.log($selectedQuestionToEdit.image_url);

							const response = await api.question.UpdateQuestion($selectedQuestionToEdit);
							if (response) {
								debug.log(`Zaaktualizowano pytanie w ${response.duration}ms`);
								toast.show(`Zaaktualizowano pytanie w ${response.duration}ms`, 'success');
								loading = false;
								modal_opened = false;
							}
						}
					}, 500);
				}}
				theme="secondary"
			>
				{modalOpenMode === 'add' ? 'Dodaj' : 'Zapisz'}
			</Button>
		</div>
	</div></Modal
>
{#if delete_modal_opened}
	{@render DeleteModal()}
{/if}

{#snippet DeleteModal()}
	<!-- svelte-ignore a11y_click_events_have_key_events -->
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<div
		onclick={async () => await CloseDeleteModal()}
		class="fixed inset-0 z-50 flex items-center justify-center bg-black/50 text-white md:backdrop-blur-lg lg:backdrop-blur-lg"
	>
		<!-- svelte-ignore a11y_no_static_element_interactions -->
		<div
			onclick={(e) => {
				e.stopPropagation();
			}}
			bind:this={deleteModalEl}
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
						await CloseDeleteModal();
					}}
				>
					<X />
				</button>
			</div>
			<!-- content -->
			<div class="flex-1 overflow-y-auto p-6 pt-5">
				<p>Czy napewno chcesz usunąć?</p>
				<div class="flex w-full justify-end">
					<Button
						theme="danger"
						size="small"
						className="flex justify-end mt-3"
						onclick={async () => {
							const response = await api.question.DeleteQuestion($selectedQuestionToEdit.id);
							if (response.data.message) {
								toast.show(response.data.message);
							}

							if (response.status === 200) {
								$questions = $questions.filter((_, index) => index !== selected_question_id);

								$addFormQuiz = { ...$addFormQuiz, questions: $questions };

								$selectedQuestionToEdit = structuredClone(emptyQuestion);

								await CloseDeleteModal();

								modal_opened = false;
							} else {
								toast.show(response.data.error);
							}
						}}
					>
						<p>Usuń</p>
					</Button>
				</div>
			</div>
		</div>
	</div>
{/snippet}
