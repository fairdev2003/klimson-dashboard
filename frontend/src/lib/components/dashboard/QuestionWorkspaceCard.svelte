<script lang="ts">
	import { api } from '$lib/api/api';
	import type { Question } from '../../../routes/dashboard/quizzes/types';

	type Props = {
		onclick: (e: Question) => void;
		question: Question;
	};

	const { question, onclick }: Props = $props();
</script>

<div
	onclick={() => {
		onclick(question);
	}}
	class="bg-primary border-1 border-secondary hover:bg-secondary flex h-auto min-h-[250px] w-full cursor-pointer flex-col gap-5 p-4"
>
	<div class="mx-auto flex aspect-video w-1/2 justify-center overflow-hidden">
		<img
			class="h-full w-full object-contain"
			src={api.image.getImage('question', question.image_url)}
		/>
	</div>
	<div>
		<p class="text-center">{question.content}</p>
	</div>
</div>
