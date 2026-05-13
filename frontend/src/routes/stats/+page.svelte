<script lang="ts">
	import { api } from '$lib/api/api';
	import { onMount } from 'svelte';
	import type { Quiz } from '../dashboard/quizzes/types';
	import Star from './star.svg';
	import {
		CountAll,
		CountDaily,
		CountMonthly,
		CountStats,
		CountThreeDays,
		CountWeekly,
		GetDailyStatsCount,
		GetMonthlyStatsCount,
		GetThreeDaysStatsCount,
		GetWeeklyStatsCount
	} from './helpers/stats.helper';
	let count: number = $state(0);
	let quizzes: Quiz[] = $state([]);

	onMount(async () => {
		const response = await api.stats.Count();
		count = response;

		quizzes = (await api.quiz.GetPublicQuizzes()).data;
		console.log(quizzes);
	});
</script>

<div class="mx-10 mt-10 text-white">
	<div class="mb-5 flex flex-col justify-center gap-3">
		<h1 class="mx-auto text-4xl font-bold text-white">STATYSTYKI STRONY</h1>
		<h2 class="mx-auto text-2xl font-bold text-white">Ilość rozwiązanych QUIZ</h2>
	</div>
	<div class="mx-auto mb-5 flex items-center justify-center gap-2">
		{@render Pasek()}
		<img alt="star" src={Star} class=" size-8" />
		{@render Pasek()}
	</div>
	<div class="mx-auto mb-5 grid grid-cols-5 gap-4">
		{@render StatCount('24h', CountDaily(quizzes))}
		{@render StatCount('3 dni', CountThreeDays(quizzes))}
		{@render StatCount('Łącznie', CountAll(quizzes))}
		{@render StatCount('7 dni', CountWeekly(quizzes))}
		{@render StatCount('30 dni', CountMonthly(quizzes))}
	</div>
	<table class="min-w-full divide-y divide-neutral-700">
		<thead class="bg-neutral-800 text-white">
			<tr>
				<th class="px-4 py-2 text-left text-sm font-medium">Tytul</th>
				<th class="px-4 py-2 text-left text-sm font-medium">24h</th>
				<th class="px-4 py-2 text-left text-sm font-medium">3 dni</th>
				<th class="px-4 py-2 text-left text-sm font-medium">Łącznie</th>
				<th class="px-4 py-2 text-left text-sm font-medium">7 dni</th>
				<th class="px-4 py-2 text-center text-sm font-medium">30 dni</th>
			</tr>
		</thead>
		<tbody class="divide-y divide-gray-700 bg-neutral-900 text-white">
			{#each quizzes as quiz, i}
				<tr class="">
					<td class="px-4 py-2">{i + 1}. {quiz.title}</td>
					<td class="px-4 py-2">{GetDailyStatsCount(quiz)}</td>
					<td class="px-4 py-2">{GetThreeDaysStatsCount(quiz)}</td>
					<td class="px-4 py-2">{CountStats(quiz)}</td>
					<td class="px-4 py-2">{GetWeeklyStatsCount(quiz)}</td>

					<td class="px-4 py-2 text-center">{GetMonthlyStatsCount(quiz)}</td>
				</tr>
			{/each}
		</tbody>
	</table>
</div>

{#snippet StatCount(content: string, stat: number)}
	<div
		class="bg-primary border-1 border-secondary col-span-1 flex flex-col justify-center gap-1 rounded-2xl py-2"
	>
		<h3 class="text-center text-2xl">{content}</h3>
		<h3 class="text-center text-2xl text-yellow-400">{stat}</h3>
	</div>
{/snippet}

{#snippet Pasek()}
	<div class="h-0.5 w-20 bg-yellow-400"></div>
{/snippet}
