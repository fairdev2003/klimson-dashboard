<script lang="ts">
	import { goto } from '$app/navigation';
	import { quizzes } from '$lib/dashboard/stores/data.store';
	import Icon from '@iconify/svelte';
</script>

<button
	onclick={() => {
		goto('/dashboard/quizzes');
	}}
	class="min-h-2 text-start gap-4 p-3 col-span-1 items-center bg-neutral-900 border-neutral-700 border flex"
>
	<div
		class="bg-neutral-800 rounded-xl size-13 border-neutral-700 border justify-center items-center flex"
	>
		<Icon icon="material-symbols:quiz" height="30" width="30" class="text-neutral-600" />
	</div>
	<div class="flex flex-col">
		<h3 class="truncate text-sm font-bold tracking-tight text-neutral-100">Dodane Quizy</h3>
		<p class="truncate text-sm font-bold tracking-tight text-neutral-500">{$quizzes.length}</p>
	</div>
</button>
