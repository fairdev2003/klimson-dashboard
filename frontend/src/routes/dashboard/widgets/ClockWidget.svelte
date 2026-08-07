<script lang="ts">
	import Heading from '$lib/components/dashboard/typography/Heading.svelte';
	import { locale } from '$lib/dashboard/dashboard.svelte';
	import { persistedWritable } from '$lib/dashboard/stores/persist';
	import Icon from '@iconify/svelte';
	import { blur } from 'svelte/transition';

	type Locale = 'pl-Pl' | 'us-US';
	let currentTime = $state('');
	let currentDate = $state('');
	let currentWeek = $state('');

	const updateDateTime = () => {
		const now = new Date();
		currentTime = now.toLocaleTimeString($locale, {
			hour: '2-digit',
			minute: '2-digit',
			second: '2-digit'
		});
		currentDate = now.toLocaleDateString($locale, {
			day: 'numeric',
			month: 'long',
			year: 'numeric'
		});
		currentWeek = now.toLocaleDateString($locale, {
			weekday: 'long'
		});
	};

	// 3. Efekt startujący zegar
	$effect(() => {
		updateDateTime(); // Ustaw od razu przy starcie
		const interval = setInterval(updateDateTime, 1000);
		return () => clearInterval(interval);
	});
</script>

<button
	class="relative text-text overflow-hidden text-start group rounded-xl flex flex-col h-45 max-w-100 w-full md:w-70 lg:w-70 border gap-3 border-widget-border bg-widget-background"
>
	<div class="flex flex-col gap-3 p-5">
		<div class="flex gap-4 items-center text-secondary-text">
			<Icon icon="material-symbols:database" width="40" height="40" />
			<Heading>Time</Heading>
		</div>
		<div class="flex flex-col gap-1">
			<div class="flex text-start justify-between gap-1 items-center">
				<h2 in:blur={{ duration: 300 }} class="text-3xl text-primary font-bold">
					{currentTime}
				</h2>
			</div>

			<div>
				<h3 class="text-sm text-secondary-text font-semibold">{currentDate}, {currentWeek}</h3>
			</div>
			<div class="hidden gap-2 group-hover:flex">
				{@render LocalePill('pl-Pl', 'Polish')}
				{@render LocalePill('us-US', 'US')}
			</div>
		</div>
	</div>
</button>

{#snippet LocalePill(change: Locale, name: string)}
	<button
		onclick={() => {
			$locale = change;
		}}
		class:locale-selected={$locale === change}
		class="bg-neutral-800 border border-neutral-700 hover:bg-neutral-700 transition-colors cursor-pointer p-0.5 px-2 mt-1 rounded-full"
	>
		<p class="text-xs text-neutral-200">{name}</p>
	</button>
{/snippet}

<style>
	@import 'tailwindcss';

	.locale-selected {
		@apply bg-green-500 border-green-700;
	}
</style>
