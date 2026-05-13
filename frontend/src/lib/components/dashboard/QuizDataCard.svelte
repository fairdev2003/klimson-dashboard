<script lang="ts">
	import { selectedQuiz, selectedSummaryType, summary_open } from '$lib/dashboard/stores/store';
	import type { Quiz } from '../../../routes/dashboard/quizzes/types';

	type Props = {
		quiz: Quiz | undefined;
	};

	const { quiz }: Props = $props();
</script>

<div
	class="border-1 flex h-[225px] w-[300px] cursor-pointer flex-col border-neutral-700/60 bg-neutral-800/60 transition-colors hover:bg-neutral-700/60"
></div>

<style>
	.selected {
		background-color: rgba(0, 128, 255, 0.1);
		/* background-color: rgba(7, 195, 45, 0.1); */
	}
</style>
