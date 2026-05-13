<script lang="ts">
	import type { Stat } from '$lib/types/stats';
	import Button from '$lib/components/Button.svelte';
	import Modal from '$lib/components/Modal.svelte';
	import { addFormQuiz, selectedQuiz } from '$lib/dashboard/stores/store';
	import Icon from '@iconify/svelte';
	import { toast } from '$lib/dashboard/stores/toast';
	import { api } from '$lib/api/api';
	import axios from 'axios';
	import { quizzes } from '$lib/dashboard/stores/data.store';
	import UpdateTags from './UpdateTags.svelte';
	import HarcCheckBox from '$lib/components/dashboard/HarcCheckBox.svelte';

	let modalOpen: boolean = $state(false);
	let statloading: boolean = $state(false);
	type Props = {
		onSave?: () => void;
	};

	let { onSave }: Props = $props();
	let loading: boolean = $state(false);

	function DeleteStatFront(statId: string) {
		$addFormQuiz.stats = $addFormQuiz.stats.filter((s) => s.id !== statId);
	}

	async function DeleteStatFromServer(id: number) {
		loading = true;

		let message = '';
		onSave?.();
		try {
			const response = await api.stats.DeleteStat({ id });

			if (response.data.message) {
				message = response.data.message;
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

	async function SaveSettings() {
		loading = true;

		let message = '';

		try {
			const response = await api.quiz.UpdateSettings(
				{ id: $addFormQuiz.id },
				{
					public: $addFormQuiz.public
				}
			);

			if (response.data.message) {
				message = response.data.message;
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
			onSave?.();
		}
	}

	let tags: string[] = $state(['#tag1', '#tag2']);
</script>

<div class="flex flex-col justify-center gap-5">
	<Button
		theme="secondary"
		size="small"
		onclick={() => {
			if ($addFormQuiz.stats.length > 0) {
				modalOpen = true;
				return;
			}
			toast.show('Ten quiz nie ma statystyk!', 'warning', 5000);
		}}>Statystyki</Button
	>
	<UpdateTags bind:tags />
	<div class="flex items-center">
		<HarcCheckBox bind:checked={$addFormQuiz['public']} label="Ustaw quiz na publiczny" />
	</div>
	<div class="flex justify-end">
		<Button
			{loading}
			theme="secondary"
			size="small"
			onclick={async () => {
				await SaveSettings();
			}}>Zapisz zmiany</Button
		>
	</div>
</div>

<Modal
	draggable
	className={`${$addFormQuiz.stats.length > 0 ? 'w-[400px] h-[600px] pb-5' : 'w-[400px]'} `}
	bind:opened={modalOpen}
	title="Statystyki quizu"
	onClose={() => {
		modalOpen = false;
	}}
>
	<div class="flex flex-col gap-1">
		{#each $addFormQuiz.stats as stat}
			{@render Statistic(stat)}
		{/each}

		{#if $addFormQuiz.stats.length === 0}
			<p class="mx-auto text-center text-neutral-400">Brak relacji statystyk</p>
		{/if}
	</div>
</Modal>

{#snippet Statistic(stat: Stat)}
	{@const date = new Date(stat.created_at).toLocaleDateString('pl-PL', {
		year: '2-digit',
		month: 'long',
		day: 'numeric',
		hour: '2-digit',
		minute: '2-digit',
		numberingSystem: 'latn'
	})}
	<div class="flex items-center justify-between gap-3 bg-neutral-900 p-3 hover:bg-neutral-700">
		<button class="cursor-pointer rounded-lg bg-neutral-600 p-1 text-xs">{date}</button>
		<button
			class="cursor-pointer"
			onclick={async () => {
				const has = confirm('Czy napewno chcesz usunac ta statystyke?');
				if (has) {
					DeleteStatFront(stat.id);
					await DeleteStatFromServer(Number(stat.id));
				}
			}}
		>
			<Icon icon="mdi:trash" width="20" height="20" class="text-red-500" />
		</button>
	</div>
{/snippet}
