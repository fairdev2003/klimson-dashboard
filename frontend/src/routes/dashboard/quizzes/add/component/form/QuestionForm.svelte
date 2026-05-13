<script lang="ts">
	import { addFormQuiz, addFormSection } from '$lib/dashboard/stores/store';
	import { type Question } from '../../../types';
	import Button from '$lib/components/Button.svelte';
	import CreateFormInput from '$lib/components/dashboard/CreateFormInput.svelte';
	import { toast } from '$lib/dashboard/stores/toast';
	import { questions } from '$lib/dashboard/stores/data.store';
	import Modal from '$lib/components/Modal.svelte';
	import QuestionsImport from '$lib/components/dashboard/add/QuestionsImport.svelte';
	import { api } from '$lib/api/api';
	import { Trash } from '@lucide/svelte';

	type Props = {
		onSave?: () => void;
	};
	let loadedImages: boolean[] = $state([]);


	let { onSave }: Props = $props();
	let loading: boolean = $state(false);
	let modalOpen: boolean = $state(false);

	async function handleSave() {
		loading = true;
		await new Promise((resolve) => setTimeout(resolve, 2000));
		loading = false;
		toast.show('Zmiany zostały zapisane pomyślnie!', 'success');
		onSave?.();
	}
</script>

<div class="flex flex-col gap-4">
	<div class="flex gap-2">
		<QuestionsImport/>
	</div>
	<div>
		{#each $addFormQuiz.questions as question, index}
			{@render Question(question, index, $addFormQuiz.questions.length)}
		{/each}
	</div>
	{@render AddQuestion()}
	<div class="flex justify-end">
		<Button
			{loading}
			theme="secondary"
			size="small"
			onclick={function () {
				handleSave ? handleSave() : null;
			}}>Zapisz zmiany</Button
		>
	</div>
</div>

<Modal
	title="Dodaj nowe pytanie"
	className="w-250 h-150"
	onClose={() => {
		modalOpen = false;
	}}
	bind:opened={modalOpen}
></Modal>

{#snippet Question(question: Question, index: number, length: number)}
	<div
		class="flex h-[80px] w-full cursor-pointer items-center border-neutral-700 bg-neutral-900 transition-colors hover:bg-neutral-800
        {index === 0 ? 'border' : 'border-x border-b'}"
	>
		<button
			
			class="flex h-full w-full items-center gap-4 text-left"
		>
			<div class="h-full w-20 flex-shrink-0 bg-neutral-800">
				{#if question.image_url}
					<img
						src={api.image.question(question.image_url)}
						class="h-full w-full object-cover"
						onload={() => loadedImages.push(true)}
						onerror={(e) => ((e.currentTarget as HTMLImageElement).style.display = 'none')}
						alt="Question"
					/>
				{/if}
			</div>

			<div class="flex flex-col justify-center overflow-hidden pr-4">
				<p class="truncate text-lg font-semibold text-white">
					{question.content}
				</p>
				<p class="text-sm text-neutral-400">
					Ilość odpowiedzi: <span class="text-primary">{question.answers.length}</span>
				</p>
			</div>
		</button>
		<button
			onclick={async () => {
				const r = confirm("Czy napewno chcesz usunac to pytanie. Zmiana jest nieodwracalna")
				if (r) {
					toast.warning("To nie koncert zyczen moj drogi. Nie dziala!")
				}
			}}
			class="border-l-1 flex size-[80px] cursor-pointer items-center justify-center border-l-neutral-700 hover:bg-red-500 hover:text-white"
			><Trash /></button
		>
	</div>
{/snippet}

{#snippet AddQuestion()}
	<button
		onclick={() => {
			modalOpen = true;
		
		}}
		class={`mt-5 flex h-[50px] cursor-pointer items-center justify-center gap-3 border border-neutral-700 bg-neutral-900  transition-colors hover:bg-neutral-700`}
	>
		Dodaj pytanie
	</button>
{/snippet}
