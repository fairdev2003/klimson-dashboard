<script lang="ts">
	import { onDestroy, tick } from 'svelte';
	import type { Quiz } from '../../../routes/dashboard/quizzes/types';
	import { quizzes, searchOpen } from '$lib/dashboard/stores/data.store';
	import Placeholder from '$lib/assets/placeholder.png';
	import { api } from '$lib/api/api';
	import {
		addFormQuiz,
		addFormSection,
		initialFormQuiz,
		questions
	} from '$lib/dashboard/stores/store';
	import { fade } from 'svelte/transition';
	import { goto } from '$app/navigation';
	import { lastSearched, sidebar_open } from '$lib/dashboard/stores/persist';
	import type { Attachment } from 'svelte/attachments';
	import { toast } from '$lib/dashboard/stores/toast';
	import { get } from 'svelte/store';
	import { Trash, Trash2 } from '@lucide/svelte';

	let value: string = $state('');

	let boxRef: HTMLDivElement | null = null;
	let inputRef: HTMLInputElement | null = null;

	function handleClickOutside(e: MouseEvent) {
		if (!boxRef || !boxRef.contains(e.target as Node)) {
			$searchOpen = false;
		}
	}

	const Mount: Attachment<HTMLDivElement> = () => {
		inputRef?.focus();
	};

	let searchRecords: Quiz[] = $state($quizzes);

	async function openBox(e: MouseEvent) {
		e.stopPropagation();

		$searchOpen = true;

		await tick();

		document.addEventListener('mousedown', handleClickOutside);
	}

	$effect(() => {
		searchRecords = $quizzes;
	});

	$effect(() => {
		const searchTerm = value.trim().toLowerCase();

		if (searchTerm === '') {
			searchRecords = $quizzes;
		} else {
			searchRecords = $quizzes.filter(
				(quiz) =>
					quiz.title.toLowerCase().includes(searchTerm) ||
					quiz.author.toLowerCase().includes(searchTerm) ||
					String(quiz.id).includes(searchTerm)
			);
		}
	});
</script>

<div class="relative hidden h-[50px] w-[500px] border-1 border-neutral-700/60 md:flex lg:flex">
	<input
		class="h-full w-full border-0 bg-transparent placeholder-neutral-500 ring-0 outline-none select-none focus:outline-none"
		onclick={openBox}
		bind:value
		placeholder="Kliknij aby wyszukać"
	/>

	{#if $searchOpen}
		<div
			{@attach function (e) {
				inputRef?.focus();
			}}
			bind:this={boxRef}
			in:fade={{ duration: 150 }}
			out:fade={{ duration: 150 }}
			class="absolute top-0 z-10 w-[500px] border-1 border-neutral-700/60 bg-neutral-900 shadow-[0_10px_30px_rgba(0,0,0,0.65)]"
		>
			<div class="font-sm h-[50px] w-full border-b-1 border-neutral-700/60">
				<input
					bind:this={inputRef}
					bind:value
					placeholder="Szukaj"
					class="h-[50px] w-full border-0 bg-transparent placeholder-neutral-500 ring-0 outline-none focus:outline-none"
				/>
			</div>
			{#if searchRecords.length > 0}
				<p class="my-1 mt-2 px-2 text-sm font-bold">WYNIKI WYSZUKIWANIA</p>
			{:else}
				<p class="p-2">Brak rekordow</p>
			{/if}

			<div class="flex w-full flex-col gap-2 p-2">
				{#each searchRecords.slice(0, 3) as record}
					{@render Record(record)}
				{/each}
				{#if searchRecords.length > 0}
					<a
						onclick={() => {
							$searchOpen = false;
						}}
						href="/dashboard/quizzes"
						class="cursor-pointer text-center text-blue-700 hover:underline"
						>Zobacz wszystkie quizy</a
					>
				{/if}
			</div>
			{#if $lastSearched.length > 0}
				<div class="flex items-center justify-between">
					<p class="my-1 px-2 text-sm font-bold">OSTATNIE</p>
					<button
						onclick={() => {
							$lastSearched = [];
						}}
						class="mr-4 cursor-pointer"
					>
						<Trash2 color="red" size={20} />
					</button>
				</div>
				<div class="flex w-full flex-col gap-2 p-2">
					{#each $lastSearched.slice(0, 3) as record}
						{@render Record(record)}
					{/each}
				</div>
				<p class="mx-2 mb-2 text-center text-[12px] text-neutral-400">HARCQUIZ CMS</p>
			{/if}
		</div>
	{/if}
</div>

{#snippet Record(quiz: Quiz)}
	<button
		onclick={(e) => {
			$addFormQuiz = quiz;

			goto('/dashboard/quizzes/update');

			$addFormSection = 'general';

			$initialFormQuiz = quiz;

			if (!$lastSearched.slice(0, 3).includes(quiz)) {
				lastSearched.update((e) => {
					return [quiz, ...e];
				});
			}

			$questions = quiz.questions;
			$searchOpen = false;
		}}
		class="flex h-15 w-full cursor-pointer gap-2 bg-primary p-1 hover:bg-secondary"
	>
		<img
			class="size-14 object-contain p-1"
			src={api.image.quiz(quiz.image_url)}
			alt={`image-${quiz.id}`}
			onerror={(e) => ((e.currentTarget as HTMLImageElement).src = Placeholder)}
		/>
		<div class="flex flex-col justify-center gap-1 text-start">
			<p class="text-sm">{quiz.title}</p>
			<p class="text-sm text-neutral-400">{quiz.author}</p>
		</div>
	</button>
{/snippet}

<svelte:document
	onkeydown={(e) => {
		if (e.key === 'Escape') {
			$searchOpen = false;
		}
	}}
/>

<style>
	input:invalid {
		outline: none;
		box-shadow: none;
	}
</style>
