<script lang="ts">
	import { selectedQuiz, sidebar_open, summary_open } from '$lib/dashboard/stores/store';
	import { ChevronLeft, ChevronRight, Dot, Grip } from '@lucide/svelte';
	import { page } from '$app/stores';
	import Loader from './Loader.svelte';

	let route: string = $state('Start');
	async function FakeLoading(): Promise<boolean> {
		await new Promise((resolve) => setTimeout(resolve, 3000));

		return true;
	}
</script>

<div
	class="border-b-1 h-13 flex w-full items-center justify-between border-neutral-700/60 bg-neutral-900"
>
	<div class="flex h-full">
		<!-- <button
			onclick={() => {
				$sidebar_open = true;
			}}
			class="flex h-full cursor-pointer items-center justify-center bg-neutral-800/60 px-3 transition-colors hover:bg-neutral-700/60"
		>
			<ChevronLeft />
		</button> -->

		<div class="flex flex-col justify-end">
			{@render Tabs()}
		</div>
	</div>
	<button
		onclick={() => {
			if ($selectedQuiz) {
				$summary_open = !$summary_open;
			}
		}}
		class="flex cursor-pointer bg-neutral-800/60 p-2 px-2 transition-colors hover:bg-neutral-700/60"
	>
		{#if !$selectedQuiz}
			<div class="relative rounded-l-full pr-3">
				<Grip />
				<!-- {#if $selectedQuiz}
				<div class="absolute left-0 top-0 text-blue-500">
					<Dot size={60} />
				</div>
			{/if} -->
			</div>
		{:else}
			<div class="relative rounded-l-full pr-3 text-blue-500">
				<Grip />
			</div>
		{/if}
	</button>
</div>

{#snippet Tabs()}{/snippet}
