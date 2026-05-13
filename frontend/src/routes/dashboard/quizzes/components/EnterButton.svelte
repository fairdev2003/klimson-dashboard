<script lang="ts">
	import { goto } from '$app/navigation';
	import {
		addFormQuiz,
		addFormSection,
		initialFormQuiz,
		questions
	} from '$lib/dashboard/stores/store';
	import { ChevronRight } from '@lucide/svelte';
	import type { Quiz } from '../types';

	type Props = {
		quiz_record: Quiz;
		selected?: boolean;
	};

	const { quiz_record, selected = false }: Props = $props();
</script>

<button
	class="bg-primary-add-button hover:bg-secondary-add-button flex size-12 cursor-pointer items-center justify-center rounded-lg text-[10px]"
	onclick={async (e) => {
		$addFormQuiz = quiz_record;
		e.stopPropagation();
		if (selected) {
			goto('/dashboard/quizzes/update');
		}
		$addFormSection = 'general';

		$initialFormQuiz = quiz_record;

		$questions = quiz_record.questions;
	}}><ChevronRight /></button
>
