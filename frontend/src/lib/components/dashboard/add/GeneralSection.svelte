<script lang="ts">
	import { api } from '$lib/api/api';
	import {
		addFormQuiz,
		addFormSection,
		selectedSummaryType,
		summary_open,
		updateFormQuiz
	} from '$lib/dashboard/stores/store';
	import type { Quiz } from '../../../../routes/dashboard/quizzes/types';
	import Check from '../Check.svelte';
	import CreateFormInput from '../CreateFormInput.svelte';
	import ImageUpload from '../ImageUpload.svelte';
	import ImageContainer from './ImageContainer.svelte';
	import QuizFormNavigation from './QuizFormNavigation.svelte';

	const manualForms: { label: string; for: keyof Quiz }[] = [
		{ label: 'Title', for: 'title' },
		{ label: 'Opis', for: 'description' },
		{ label: 'Autor', for: 'author' },
		{ label: 'Poziom', for: 'difficulty' },
		{ label: 'URL Obrazka', for: 'image_url' }
	];

	let checked: boolean = $state(false);

	type Props = {
		mode?: 'add' | 'edit';
	};

	const { mode = 'add' }: Props = $props();
</script>

<div>
	<div class="mt-5">
		<div class="mx-auto flex lg:col-span-4">
			<ImageUpload src={mode === 'edit' ? $addFormQuiz.image_url : undefined} type="quiz" />
		</div>
		<div class="lg:col-span-6">
			{#if mode === 'edit'}
				<p>ID: {$addFormQuiz.id}</p>
			{/if}
			<CreateFormInput bind:value={$addFormQuiz['title']} label={'Nazwa quizu'} />
			<CreateFormInput bind:value={$addFormQuiz['description']} label={'Krótki opis'} />
			<CreateFormInput bind:value={$addFormQuiz['author']} label={'Autor'} />
			<CreateFormInput bind:value={$addFormQuiz['difficulty']} label={'Poziom'} />
		</div>
	</div>

	<div class="mb-4 mt-5 flex items-center">
		<input
			id="default-checkbox"
			type="checkbox"
			bind:checked={$addFormQuiz['public']}
			class="border-default-medium rounded-xs bg-neutral-secondary-medium focus:ring-brand-soft h-4 w-4 border focus:ring-2"
		/>
		<label for="default-checkbox" class="text-heading ms-2 select-none text-sm font-medium"
			>Publiczny</label
		>
	</div>

	<button
		class="border-1 -left-20 top-0 cursor-pointer border-neutral-700/60 bg-neutral-800/60 p-2"
		onclick={async () => {
			if (!$summary_open && $selectedSummaryType === 'debug') {
				$summary_open = true;
				return;
			}

			if ($selectedSummaryType === 'debug') {
				return;
			}

			$summary_open = false;
			await new Promise((resolve) => setTimeout(resolve, 400));

			if ($summary_open) {
				$summary_open = true;
				$selectedSummaryType = 'debug';
				return;
			}
			$selectedSummaryType = 'debug';
			$summary_open = !$summary_open;
		}}>Debug</button
	>
	<button
		class="border-1 -left-40 top-20 mt-10 cursor-pointer border-neutral-700/60 bg-neutral-800/60 p-2"
		onclick={async () => {
			addFormQuiz.set({
				title: 'Podstawy orientacji w terenie',
				description:
					'Sprawdź, jak dobrze znasz zasady orientowania się w lesie i posługiwania się mapą oraz kompasem.',
				image_url: 'http://localhost:5173/src/lib/assets/example.jpeg',
				public: false,
				edit_link: '',
				has_time_limit: false,
				time_limit: 0,
				difficulty: 'Łatwy',
				expected_time_min: '',
				author: 'Zofia Kowalska',
				completed_count: 0,
				badges: '',
				questions: $addFormQuiz.questions
			});
		}}>Wypełnij</button
	>
	<button
		class="border-1 -left-40 top-20 mt-10 cursor-pointer border-neutral-700/60 bg-neutral-800/60 p-2"
		onclick={() => {
			console.log($addFormQuiz);
		}}>Dump</button
	>
</div>
