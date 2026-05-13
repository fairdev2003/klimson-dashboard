<script lang="ts">
	import Button from '$lib/components/Button.svelte';
	import CreateFormInput from '$lib/components/dashboard/CreateFormInput.svelte';
	import { addFormQuiz } from '$lib/dashboard/stores/store';
	import { toast } from '$lib/dashboard/stores/toast';
	import { onMount } from 'svelte';
	import { fade } from 'svelte/transition';
	import { api } from '$lib/api/api';
	import axios from 'axios';
	import { quizzes } from '$lib/dashboard/stores/data.store';
	import { goto } from '$app/navigation';

	type Props = {
		onSave?: () => void;
	};

	let { onSave }: Props = $props();
	let loading: boolean = $state(false);

	let initialDataString: string = $state('');

	onMount(() => {
		initialDataString = JSON.stringify($addFormQuiz);
	});

	let hasChanges = $derived(initialDataString !== JSON.stringify($addFormQuiz));

	async function HandleUpdate() {
		loading = true;
		const response = await api.quiz.UpdateBasicInfo(
			{ id: $addFormQuiz.id },
			{
				title: $addFormQuiz.title,
				author: $addFormQuiz.author,
				description: $addFormQuiz.description,
				difficulty: $addFormQuiz.difficulty
			}
		);
		loading = false;
		toast.show(`${response.data.message} [ ${response.duration}ms ]`, 'success');
		onSave?.();
	}

	function resetForm() {
		if (confirm('Czy na pewno chcesz cofnąć wszystkie zmiany?')) {
			addFormQuiz.update((quiz) => {
				return { ...quiz, title: '' };
			});
			toast.show('Dane quizu zresetowane!');
		}
	}

	async function Save() {
		loading = true;

		let message = '';
		onSave?.();
		try {
			const response = await api.quiz.SaveBasicInfo({
				author: $addFormQuiz.author,
				title: $addFormQuiz.title,
				description: $addFormQuiz.description,
				difficulty: $addFormQuiz.difficulty
			});
			$addFormQuiz.id = response.data.id;

			if (response.data.message) {
				message = response.data.message;
				onSave?.();
			}
		} catch (error: unknown) {
			if (axios.isAxiosError(error)) {
				const message = error.response?.data?.message || 'Błąd serwera';
				toast.error(message);
			} else {
				toast.error('Wystąpił nieoczekiwany błąd');
			}
		} finally {
			const response = await api.quiz.GetAll();
			$quizzes = response.data;
			loading = false;
			toast.success(message);
		}
	}
</script>

<div class="flex flex-col gap-4">
	<div class="flex flex-col gap-3">
		<CreateFormInput
			disabled
			value={$addFormQuiz.id}
			label="ID"
			disclaimer={!$addFormQuiz.id ? '*(Id generowane jest losowo)' : ''}
		/>
		<CreateFormInput bind:value={$addFormQuiz.title} label="Nazwa quizu" />
		<CreateFormInput label="Opis quizu" bind:value={$addFormQuiz.description} />

		<div class="flex flex-col gap-1">
			<span class="text-sm text-neutral-400">Poziom trudności: {$addFormQuiz.difficulty}</span>
			<div class="flex gap-2">
				{#each ['Łatwy', 'Średni', 'Trudny', 'Bardzo trudny'] as diff}
					<Button
						size="small"
						theme={$addFormQuiz.difficulty === diff ? 'correct' : 'secondary'}
						onclick={() => {
							$addFormQuiz.difficulty = diff;
						}}
					>
						{diff}
					</Button>
				{/each}
			</div>
		</div>

		<CreateFormInput label="Autor" bind:value={$addFormQuiz.author} />
	</div>

	<div class="flex h-12 items-center justify-end gap-3 border-t border-neutral-800 pt-4">
		{#if hasChanges}
			<div transition:fade={{ duration: 200 }} class="flex gap-2">
				<button
					onclick={resetForm}
					class="px-3 text-sm text-neutral-400 transition-colors hover:text-white"
				>
					Anuluj
				</button>

				<Button
					{loading}
					theme="secondary"
					size="small"
					onclick={async () => {
						await Save();
						goto('/dashboard/quizzes/update');
					}}>Prześlij zmiany</Button
				>
			</div>
		{/if}
	</div>
</div>
