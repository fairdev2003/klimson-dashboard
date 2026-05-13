<script lang="ts">
	import Button from '$lib/components/Button.svelte';
	import Heading from '$lib/components/dashboard/typography/Heading.svelte';
	import Modal from '$lib/components/Modal.svelte';
	import { addFormQuiz } from '$lib/dashboard/stores/store';
	import Icon from '@iconify/svelte';
	import { onMount, type Snippet } from 'svelte';
	import type { Question, Quiz } from '../../types';
	import ButtonWithOptions from '$lib/components/ButtonWithOptions.svelte';
	import { goto } from '$app/navigation';
	import GameWindow from '$lib/components/game/GameWindow.svelte';
	import { developerView } from '$lib/dashboard/stores/persist';
	import MovingTooltip from '$lib/components/dashboard/MovingTooltip.svelte';
	import { quiz_manager } from '../../lib/quiz_manager.svelte';

	type Props = {
		children: Snippet;
		renderInfo?: boolean;
	};

	let { children, renderInfo = true }: Props = $props();
	let modalOpen: boolean = $state(false);
	let previewModalOpen: boolean = $state(false);
	let quiz: Quiz = $state($addFormQuiz);

	function formatTime(questions: Question[]): string {
		const totalSeconds = questions.reduce((acc, q) => acc + (q.time_limit || 0), 0);

		const minutes = Math.floor(totalSeconds / 60);

		const seconds = totalSeconds % 60;

		if (minutes === 0) return `${seconds} sek`;
		if (seconds === 0) return `${minutes} min`;

		return `${minutes} min ${seconds} sek`;
	}

	function formatDate(dateString: string | undefined): string {
		if (!dateString) return 'Brak danych';
		const date = new Date(dateString);
		return date.toLocaleDateString('pl-PL', {
			year: 'numeric',
			month: 'long',
			day: 'numeric',
			hour: '2-digit',
			minute: '2-digit',
			numberingSystem: 'latn'
		});
	}

	function newWindow() {
		const href = `/quiz?id=${$addFormQuiz.id}`;
		window.open(href, `Quiz nr ${$addFormQuiz.id}`, 'width=1500,height=1000,noopener,noreferrer');
	}
</script>

<div class="flex items-center justify-between border-b border-secondary p-3">
	<div class="flex items-center gap-3">
		<Button className="size-12" onclick={() => goto('/dashboard/quizzes')}>
			<Icon icon="weui:back-filled" width="12" height="24" />
		</Button>
		{@render children()}
		{#if renderInfo}
			<div>
				<Button size="small" theme="base" onclick={() => (modalOpen = true)}
					><Icon icon="fluent:info-48-regular" width="24" height="24" /></Button
				>
			</div>
		{/if}
	</div>
	<div class="flex items-center gap-2">
		<Button
			onclick={() => quiz_manager.PromptDeleteAndRefresh($addFormQuiz)}
			theme="danger"
			size="small"
		>
			<div class="flex items-center justify-center">Usuń quiz</div>
		</Button>
		<MovingTooltip>
			{#snippet tooltipContent()}
				<p class="text-xs">Funkcja ta zapisuje formularz quizu na pozniej</p>
			{/snippet}
			<Button size="small" theme="secondary" disabled>
				<div class="flex items-center gap-1">
					<Icon icon="ri:save-3-line" width="13" height="13" />
					<span>Zapisz na później</span>
				</div>
			</Button>
		</MovingTooltip>
	</div>
</div>

<Modal
	title={`Podgląd quizu: ${quiz?.title}`}
	className="w-9/10 h-9/10 p-0"
	bind:opened={previewModalOpen}
	onClose={() => {
		previewModalOpen = false;
	}}
>
	<GameWindow quiz={$addFormQuiz} previewMode showCorrect />
</Modal>

<Modal
	className="w-[500px]"
	bind:opened={modalOpen}
	title="Szczegóły quizu"
	onClose={() => {
		modalOpen = false;
	}}
>
	<div class="flex flex-col gap-1 p-3">
		{#if $addFormQuiz.public}
			<a
				onclick={newWindow}
				target="_blank"
				rel="noopener noreferrer"
				class="flex gap-1 text-sm items-center mb-4"
			>
				<Icon icon="akar-icons:link-out" width="15" height="15" />
				<span>Przejdz do quizu</span>
			</a>
		{/if}
		<p>
			Identyfikator: <span class="text-neutral-400">{quiz?.id}</span>
		</p>
		<p>
			Nazwa: <span class="text-neutral-400">{quiz?.title}</span>
		</p>
		<p>
			Autor: <span class="text-neutral-400">{quiz?.author}</span>
		</p>
		<p>
			Poziom trudności: <span class="text-neutral-400">{quiz?.difficulty}</span>
		</p>
		<p>
			Data utworzenia: <span class="text-neutral-400">{formatDate(quiz?.created_at)}</span>
		</p>
		<p>
			Data modyfikacji: <span class="text-neutral-400">{formatDate(quiz?.updated_at)}</span>
		</p>
		<p>
			Przewidywany czas na rozwiazanie quizu: <span class="text-neutral-400"
				>{formatTime(quiz.questions)}</span
			>
		</p>
		{#if $developerView}
			<p>
				System Time: <span class="text-neutral-400"
					>{new Date().toLocaleString('pl-PL', {
						year: 'numeric',
						month: 'long',
						day: 'numeric',
						hour: '2-digit',
						minute: '2-digit',
						numberingSystem: 'latn'
					})}</span
				>
			</p>
		{/if}
	</div>
</Modal>

<style>
	@import 'tailwindcss';

	a {
		@apply text-blue-500 cursor-pointer;

		:hover {
			text-decoration: underline;
		}
	}
</style>
