<script lang="ts">
	import { quizzes } from '$lib/dashboard/stores/data.store';
	import Icon from '@iconify/svelte';
	import { slide } from 'svelte/transition';
	import QuizRecord from '../../quizzes/components/QuizRecord.svelte';

	type Props = {
		label: string;
	};

	let { label }: Props = $props();
	let isOpen: boolean = $state(true);
	let selectedId: number = $state(0);
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
<div>
	<!-- svelte-ignore a11y_click_events_have_key_events -->
	<div
		onclick={() => {
			isOpen = !isOpen;
		}}
		class="w-3xl transition-all hover:bg-neutral-400/30 rounded-lg p-5 cursor-pointer mx-auto text-white"
	>
		<div class="text-blue-500 flex justify-between items-center">
			<h2 class="text-xl font-bold">{label}</h2>
			<div class="transition-all duration-300" class:rotate-180={isOpen}>
				<Icon icon="ep:arrow-down" width="40" height="40" />
			</div>
		</div>
	</div>
	{#if isOpen}
		<div
			in:slide={{ duration: 300 }}
			out:slide={{ duration: 300 }}
			class="flex flex-col gap-4 w-3xl mx-auto mt-5"
		>
			{#each $quizzes as quiz}
				<div class="w-full">
					<QuizRecord quiz_record={quiz} />
				</div>
			{/each}
		</div>
	{/if}
</div>
