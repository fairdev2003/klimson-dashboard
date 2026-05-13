<script lang="ts">
	import { goto } from '$app/navigation';
	import Loader from '$lib/components/dashboard/Loader.svelte';
	import { Check, ChevronLeft, X } from '@lucide/svelte';

	import type { ComponentType } from 'svelte';
	import {
		addFormQuiz,
		addFormSection,
		initialFormQuiz,
		quizStatus,
		selectedQuiz
	} from '$lib/dashboard/stores/store';
	import GeneralSection from '$lib/components/dashboard/add/GeneralSection.svelte';
	import PageLoading from '$lib/components/PageLoading.svelte';
	import QuestionsSection from '$lib/components/dashboard/add/QuestionsSection.svelte';
	import CreationSummary from '$lib/components/dashboard/add/CreationSummary.svelte';
	import Button from '$lib/components/Button.svelte';
	import type { Attachment } from 'svelte/attachments';
	import gsap from 'gsap';
	import QuizFormNavigation from '$lib/components/dashboard/add/QuizFormNavigation.svelte';
	import { api } from '$lib/api/api';
	import { preloaded_quiz_form } from '$lib/static/quiz';
	import { toast } from '$lib/dashboard/stores/toast';
	import { debug } from '$lib/dashboard/stores/debug';

	let loading: boolean = $state(false);
	let showSaveButton: boolean = $state(false);

	async function FakeLoading(): Promise<boolean> {
		// await new Promise((resolve) => setTimeout(resolve, 3000));

		return true;
	}

	const FadeOut: Attachment<HTMLDivElement> = (element) => {
		// setup (np. animacja wejścia)
		gsap.fromTo(element, { opacity: 0 }, { opacity: 1, duration: 0.3 });
	};

	type Props = {
		mode: 'add' | 'edit';
	};

	function checkForm(): boolean {
		if (!$addFormQuiz.title) {
			toast.show('Nie mozna przeslac quizu bez nazwy!', 'error', 5000);
			loading = false;
			return false;
		}
		if (!$addFormQuiz.description) {
			toast.show('Nie mozna przeslac quizu bez opisu!', 'error', 5000);
			loading = false;
			return false;
		}
		if ($addFormQuiz.questions.length === 0) {
			toast.show('Nie mozna przeslac quizu bez pytan!', 'error', 5000);
			loading = false;
			return false;
		}
		return true;
	}

	const { mode = 'add' }: Props = $props();
</script>

<div {@attach FadeOut} class="mx-5">
	<div class="relative mx-auto mt-5 w-full max-w-[1000px] text-white">
		<div
			class="bg-primary border-1 border-secondary h-15 mb-5 flex w-full items-center justify-between px-3"
		>
			<div class="flex items-center">
				<button
					class="cursor-pointer rounded-full p-2 hover:bg-white/10"
					onclick={() => {
						goto('/dashboard/quizzes');
					}}
				>
					<ChevronLeft />
				</button>
			</div>

			<div class="flex gap-3">
				{#if mode === 'edit'}
					<Button
						onclick={() => {
							if (!$addFormQuiz) return;
							window.open(`/quiz?id=${$addFormQuiz.id}`, '_blank', 'noopener,noreferrer');
						}}
						loading={false}
						loadingText="Ładowanie komponentów"
						theme="correct">Podgląd</Button
					>
				{/if}
				<Button
					size="medium"
					theme="secondary"
					{loading}
					onclick={async () => {
						loading = true;
						if (!checkForm()) return;
						if (mode === 'add') {
							const response = await api.quiz.CreateQuiz($addFormQuiz);
							if (response) {
								loading = false;
								$addFormQuiz = preloaded_quiz_form;
								toast.show('Dodano quiz!', 'success');
								goto('/dashboard/quizzes');
							}
						}
						if (mode === 'edit') {
							loading = true;

							const start = performance.now();

							const response = await api.quiz.UpdateQuiz({ id: $addFormQuiz.id }, $addFormQuiz);

							const end = performance.now();
							const durationMs = Math.round(end - start);

							toast.show('Zaaktualizowano quiz!', 'success');

							loading = false;

							if (response) {
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
		<div class="flex w-full flex-col gap-4 lg:grid lg:grid-cols-3">
			{@render ButtonSwitch('general', 'Podstawowe')}
			{@render ButtonSwitch('questions', 'Pytania')}
			{@render ButtonSwitch('creation_summary', 'Podsumowanie')}
		</div>
		{#await FakeLoading()}
			<PageLoading />
		{:then}
			<!-- {@render AddTabs()} -->
			{#if $addFormSection === 'general'}
				<GeneralSection {mode} />
			{/if}
			{#if $addFormSection === 'questions'}
				<QuestionsSection {mode} />
			{/if}
			{#if $addFormSection === 'creation_summary'}
				<CreationSummary {mode} />
			{/if}
		{/await}
	</div>
</div>

{#snippet ButtonSwitch(
	sectionSwitch: 'general' | 'questions' | 'creation_summary',
	sectionName: string
)}
	<Button
		className="w-full"
		theme={sectionSwitch === $addFormSection ? 'secondary' : 'base'}
		onclick={() => {
			$addFormSection = sectionSwitch;
		}}>{sectionName}</Button
	>
{/snippet}

{#snippet AddTabs()}
	<div class=" mt-2 flex flex-col gap-4 md:grid md:grid-cols-3 lg:grid lg:grid-cols-3">
		<div
			onclick={() => {
				$addFormSection = 'general';
			}}
			class:bg-blue-600={$addFormSection === 'general'}
			class:bg-neutral-800={$addFormSection !== 'general'}
			class="cursor-pointer p-2 hover:bg-blue-600"
		>
			Quiz
		</div>
		<div
			onclick={() => {
				$addFormSection = 'questions';
			}}
			class:bg-blue-600={$addFormSection === 'questions'}
			class:bg-neutral-800={$addFormSection !== 'questions'}
			class="cursor-pointer p-2 hover:bg-blue-600"
		>
			Pytania
		</div>
		<div
			onclick={() => {
				$addFormSection = 'creation_summary';
			}}
			class:bg-blue-600={$addFormSection === 'creation_summary'}
			class:bg-neutral-800={$addFormSection !== 'creation_summary'}
			class="cursor-pointer truncate p-2 hover:bg-blue-600"
		>
			Podsumowanie
		</div>
	</div>
{/snippet}

<style>
	input:invalid {
		outline: none;
		box-shadow: none;
	}
</style>
