<script lang="ts">
	import { answers, questions } from '$lib/dashboard/stores/data.store';
	import type { Question } from '../../types';
	import Placeholder from '$lib/assets/placeholder.png';
	import { api } from '$lib/api/api';
	import Icon from '@iconify/svelte';
	import MovingTooltip from '$lib/components/dashboard/MovingTooltip.svelte';
	import { slide } from 'svelte/transition';
	import Loader from '$lib/components/dashboard/Loader.svelte';
	import UpdateSingleQuestion from './UpdateSingleQuestion.svelte';
	import { quiz_manager } from '../../lib/quiz_manager.svelte';
	import { stopPropagation } from 'svelte/legacy';

	type Props = {
		onclick?: () => void;
		selected?: boolean;
		question_record: Question;
		index: number;
	};

	let editing: boolean = $state(false);
	let { onclick, selected = false, question_record, index }: Props = $props();

	function toggleEditMode() {
		quiz_manager.get_question_manager.toggleSelectedQuestionIndex(index);
	}
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
<!-- svelte-ignore a11y_click_events_have_key_events -->
<div
	class:bg-neutral-700={selected}
	class:border-neutral-600={selected}
	class:bg-neutral-800={!selected}
	class:border-neutral-700={!selected}
	class="border"
>
	<div class="h-20 w-full mx-auto p-3 flex gap-3 items-center" {onclick}>
		<img
			class="h-full object-cover rounded bg-neutral-900 aspect-square opacity-80 transition-all duration-700 group-hover:opacity-100"
			onerror={(e) => ((e.currentTarget as HTMLImageElement).src = Placeholder)}
			src={api.image.question(question_record.image_url)}
			alt={`question-image-${question_record.id}`}
			loading="lazy"
		/>

		<div class="flex w-full justify-between">
			<MovingTooltip>
				{#snippet tooltipContent()}
					<div class="flex flex-col gap-1 text-xs">
						{#each Object.entries(question_record) as entry}
							{#if entry[0] === 'answers'}
								<p class="font-bold text-blue-500">
									{entry[0]}: <span class="font-normal text-white">{entry[1]}</span>
								</p>
							{:else}
								<p class="font-bold text-blue-500">
									{entry[0]}: <span class="font-normal text-white">{entry[1]}</span>
								</p>
							{/if}
						{/each}
					</div>
				{/snippet}
				<div class="flex flex-col">
					<p class="text-md font-semibold">{question_record.content}</p>
					<p class="text-neutral-200 text-xs">
						Ilosc odpowiedzi: {question_record.answers.length}
					</p>
				</div>
			</MovingTooltip>

			{#if selected}
				<div class="flex gap-3 text-white items-center">
					<MovingTooltip>
						{#snippet tooltipContent()}
							<p class="text-xs">Usuń pytanie</p>
						{/snippet}
						<button
							onclick={(e) => {
								e.stopPropagation();
							}}
							class="p-3 bg-red-500 rounded-md cursor-pointer"
						>
							<Icon icon="bx:trash-alt" width="16" height="16" />
						</button>
					</MovingTooltip>

					<MovingTooltip>
						{#snippet tooltipContent()}
							<p class="text-xs">Edytuj pytanie</p>
						{/snippet}
						<button class="p-3 bg-blue-500 rounded-md cursor-pointer" onclick={toggleEditMode}>
							<div class:rotate-180={editing} class="transition-all duration-300">
								<Icon icon="lsicon:triangle-down-filled" width="16" height="16" />
							</div>
						</button>
					</MovingTooltip>
				</div>
			{/if}
		</div>
	</div>

	{#if quiz_manager.get_question_manager.selectedQuestionId === index}
		<UpdateSingleQuestion {index} {question_record} />
	{/if}
</div>
