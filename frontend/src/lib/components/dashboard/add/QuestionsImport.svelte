<script lang="ts">
	import Button from '$lib/components/Button.svelte';
	import { X } from '@lucide/svelte';
	import type { Question } from '../../../../routes/dashboard/quizzes/types';
	import { convertToQuestionType } from './question.helper';
	import { addFormQuiz } from '$lib/dashboard/stores/store';
	import Tooltip from '../Tooltip.svelte';
	import Modal from '$lib/components/Modal.svelte';
	import MovingTooltip from '../MovingTooltip.svelte';

	let modalOpen: boolean = $state(false);

	// Importuj swoje typy i funkcję konwertującą
	// import { convertToQuestionType } from './converter';

	let jsonInput = $state(''); // To co wklejasz
	let questions = $state<Question[]>([]); // Wynik konwersji
	let error = $state('');

	async function handleConvert() {
		try {
			error = '';
			const parsed = JSON.parse(jsonInput);

			questions = convertToQuestionType(parsed);
			const questionsFromOriginalForm = $addFormQuiz.questions;
			$addFormQuiz.questions = [...questionsFromOriginalForm, ...questions];

			console.log('Skonwertowano pomyślnie!', questions);
		} catch (e) {
			error = 'Błąd: Niepoprawny format JSON lub brak wymaganych pól.';
			setTimeout(() => {
				error = '';
			}, 10000);
			console.error(e);
		} finally {
			modalOpen = false;
			jsonInput = '';
		}
	}
</script>

<MovingTooltip>
	{#snippet tooltipContent()}
		<p class="text-xs text-white">Funckja niedostępna</p>
	{/snippet}
	<Button
		disabled
		className="w-auto"
		theme="secondary"
		size="small"
		onclick={() => {
			modalOpen = !modalOpen;
		}}>Importuj pytania</Button
	>
</MovingTooltip>
<Modal
	className="h-auto w-150"
	onClose={() => {
		modalOpen = false;
	}}
	bind:opened={modalOpen}
	title="Importowanie"
	><div class="flex flex-1 flex-col gap-4">
		<p>Wklej tresc swojego pytania JSON</p>
		<textarea
			bind:value={jsonInput}
			class="h-40 w-full border-neutral-700/60 bg-transparent placeholder-neutral-700"
		></textarea>
		<div class="h-5">
			<p class="text-center text-red-500">
				{error}
			</p>
		</div>
		<div class="flex w-full justify-end">
			<Button
				size="small"
				onclick={async () => {
					await handleConvert();
				}}>Importuj</Button
			>
		</div>
	</div>
</Modal>

{#snippet DeleteModal()}
	<!-- svelte-ignore a11y_click_events_have_key_events -->
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<div
		onclick={(e) => {
			modalOpen = false;
		}}
		class="fixed inset-0 z-50 flex items-center justify-center bg-black/50 text-white md:backdrop-blur-lg lg:backdrop-blur-lg"
	>
		<!-- svelte-ignore a11y_no_static_element_interactions -->
		<div
			onclick={(e) => {
				e.stopPropagation();
			}}
			class="w-9/10 relative flex flex-col border border-neutral-800/60 bg-neutral-950 md:w-1/2 md:min-w-[600px] md:backdrop-blur-none lg:w-1/4 lg:bg-neutral-950"
		>
			<!-- title -->
			<div
				class="mb-2 flex h-10 flex-shrink-0 items-center justify-between border-b border-neutral-700/60 bg-neutral-800/60 px-5"
			>
				<p>Importowanie</p>
				<button
					onclick={() => {
						modalOpen = !modalOpen;
					}}
					class="cursor-pointer text-neutral-500 hover:text-neutral-300"
				>
					<X />
				</button>
			</div>
			<!-- content -->
		</div>
	</div>
{/snippet}
