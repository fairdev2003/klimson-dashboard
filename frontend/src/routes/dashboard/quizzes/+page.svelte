<script lang="ts">
	import { onMount } from 'svelte';
	import { Plus, Search } from '@lucide/svelte';
	import {
		addFormQuiz,
		addFormSection,
		questions,
		contextMenuOptions
	} from '$lib/dashboard/stores/store';
	import { goto } from '$app/navigation';
	import Button from '$lib/components/Button.svelte';
	import { api } from '$lib/api/api';
	import { preloaded_quiz_form } from '$lib/static/quiz';
	import QuizRecord from './components/QuizRecord.svelte';
	import { quizzes, searchOpen } from '$lib/dashboard/stores/data.store';
	import { toast } from '$lib/dashboard/stores/toast';
	import { debug } from '$lib/dashboard/stores/debug';
	import { sidebar_open } from '$lib/dashboard/stores/persist';
	import HarcCheckBox from '$lib/components/dashboard/HarcCheckBox.svelte';
	import { recordVisibilityDict } from './components/quiz_record';

	let addButtonLoading: boolean = $state(false);
	let refreshloading: boolean = $state(false);
	let counter: number = $state(0);
	let checked: boolean = $state(false);

	onMount(async () => {
		onMount(() => {
			$contextMenuOptions = [
				{
					contextMenuName: 'Ustawienia quizów',
					items: [
						{
							label: 'Odśwież dane o quizach',
							action: async () => {
								const response = await api.quiz.GetAll();
								$quizzes = response.data;
								if (response.status === 200) {
									toast.show('Odświeżono rekordy!', 'success');
								}
							},
							icon: 'lucide:refresh-cw',
							color: 'text-green-500'
						},
						{
							label: 'Podsumowanie (Summary)',
							action: () => sidebar_open.set(!$sidebar_open),
							icon: 'ooui:text-summary-rtl',
							color: 'text-blue-500'
						}
					]
				},
				{
					contextMenuName: 'Inne',
					items: [
						{
							label: 'Kopiuj ścieżkę API',
							action: () => {
								navigator.clipboard.writeText(window.location.pathname);
								toast.show('Skopiowano!', 'success');
							},
							icon: 'lucide:copy',
							color: ''
						},
						{
							label: 'Wyczyść cache panelu',
							action: () => {
								localStorage.clear();
								location.reload();
							},
							icon: 'material-symbols:cached',
							color: 'text-red-500'
						},
						{
							label: 'Wyloguj sesję',
							action: () => {
								localStorage.setItem('token', '');
								goto('/login');
							},
							icon: 'lucide:log-out',
							color: 'text-red-500'
						}
					]
				}
			];
		});
	});
</script>

<div class="mt-5">
	{@render QuizSectionHeading()}
</div>

<div class="w-full">
	<div class="flex w-full flex-col items-center justify-center mx-auto gap-4 overflow-x-auto px-4">
		{#each $quizzes as quiz_record}
			<div class="flex-shrink-0 lg:w-full">
				<QuizRecord selected={quiz_record.id === $addFormQuiz.id} {quiz_record} />
			</div>
		{/each}
	</div>
</div>

{#snippet QuizSectionHeading()}
	<div class="mb-5 flex items-center justify-between gap-5 px-5">
		<div class="flex items-center gap-3">
			<h1 class="text-2xl font-bold">Quizy</h1>
			<Button
				loading={refreshloading}
				onclick={async () => {
					refreshloading = true;
					const response = await api.quiz.GetAll();
					$quizzes = response.data;
					toast.show(`Odświeżono rekordy w ${response.duration}ms`, 'success', 2000);
					refreshloading = false;
				}}>Odswież</Button
			>
			<Button
				className="flex gap-2"
				onclick={() => {
					$searchOpen = true;
				}}
				><Search />
				<p>Szukaj</p></Button
			>
		</div>
		<div class="flex items-center justify-center gap-2">
			<Button
				loading={addButtonLoading}
				onclick={() => {
					addButtonLoading = true;

					addButtonLoading = false;
					$addFormQuiz = preloaded_quiz_form;
					$addFormQuiz.id = undefined;
					$questions = [];
					$addFormSection = 'general';
					toast.show('Utworzono nowy formularz', 'info', 2000);
					debug.log('Nowy formularz');
					goto('/dashboard/quizzes/add');
					console.log('siema');
				}}
				className=""
				size="medium"
				theme="secondary"
			>
				{#if !addButtonLoading}
					<Plus size={20} />
				{/if}
				<p>Dodaj</p></Button
			>
		</div>
	</div>
	<div class="mt-20"></div>
{/snippet}

<style>
	@import 'tailwindcss';
	@plugin '@tailwindcss/forms';
	@plugin '@tailwindcss/typography';

	label {
		color: white;
	}

	input,
	textarea {
		@apply mt-2 border-1 border-neutral-800/60 bg-neutral-900 text-white;
	}
</style>
