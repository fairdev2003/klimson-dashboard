<script lang="ts">
	import { onMount } from 'svelte';
	import Dashboard from '../dashboard.svelte';
	import { questions, quizzes } from '$lib/dashboard/stores/data.store';
	import { api } from '$lib/api/api';
	import UnderConstrcut from '$lib/components/dashboard/UnderConstrcut.svelte';

	onMount(async () => {
		const q = await api.question.GetAllQuestions();
		questions.set(q);
		console.log($questions);
	});
</script>

<div class="py-5 text-white">
	<UnderConstrcut />
</div>
