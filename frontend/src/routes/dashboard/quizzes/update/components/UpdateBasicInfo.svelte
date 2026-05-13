<script lang="ts">
	import { api } from '$lib/api/api';
	import Button from '$lib/components/Button.svelte';
	import CreateFormInput from '$lib/components/dashboard/CreateFormInput.svelte';
	import MovingTooltip from '$lib/components/dashboard/MovingTooltip.svelte';
	import { addFormQuiz } from '$lib/dashboard/stores/store';
	import { toast } from '$lib/dashboard/stores/toast';

	type Props = {
		onSave?: () => void;
	};

	let { onSave }: Props = $props();
	let loading: boolean = $state(false);

	async function handleSave() {
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
</script>

<div class="flex flex-col gap-4">
	<div class="flex flex-col gap-3">
		<CreateFormInput disabled bind:value={$addFormQuiz.id} label="ID" />
		<CreateFormInput bind:value={$addFormQuiz.title} label="Nazwa quizu" />
		<CreateFormInput label="Opis quizu" bind:value={$addFormQuiz.description} />
		<CreateFormInput label="Poziom trudności" bind:value={$addFormQuiz.difficulty} />
		<div class="flex gap-2">
			<MovingTooltip>
				{#snippet tooltipContent()}
					<p class="text-xs">Łatwy</p>
				{/snippet}
				<Button
					size="small"
					theme="secondary"
					onclick={() => {
						$addFormQuiz.difficulty = 'Łatwy';
					}}>Łatwy</Button
				>
			</MovingTooltip>
			<Button
				size="small"
				theme="secondary"
				onclick={() => {
					$addFormQuiz.difficulty = 'Średni';
				}}>Średni</Button
			>
			<Button
				size="small"
				theme="secondary"
				onclick={() => {
					$addFormQuiz.difficulty = 'Trudny';
				}}>Trudny</Button
			>
			<Button
				size="small"
				theme="secondary"
				onclick={() => {
					$addFormQuiz.difficulty = 'Bardzo trudny';
				}}>Bardzo trudny</Button
			>
		</div>
		<CreateFormInput label="Autor" bind:value={$addFormQuiz.author} />
	</div>

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
