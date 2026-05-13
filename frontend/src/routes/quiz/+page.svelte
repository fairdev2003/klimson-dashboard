<script lang="ts">
	import { page } from '$app/state';
	import { api } from '$lib/api/api';
	import GameWindow from '$lib/components/game/GameWindow.svelte';
	import { onMount } from 'svelte';
	import type { Quiz } from '../dashboard/quizzes/types';

	let quiz: Quiz | undefined = $state();

	onMount(async () => {
		const id = Number(page.url.searchParams.get('id'));

		const response = await api.quiz.GetPublicQuiz({ id });
		console.log(response);
		quiz = response;
	});
</script>

<div class="flex flex-col items-center justify-center gap-5 text-white">
	<GameWindow {quiz} showCorrect={true} />
</div>
