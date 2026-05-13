<script lang="ts">
	import { Trash, X } from '@lucide/svelte';
	import type { Quiz } from '../types';
	import { addFormQuiz, selectedQuiz, summary_open } from '$lib/dashboard/stores/store';
	import gsap from 'gsap';
	import { tick } from 'svelte';
	import { api } from '$lib/api/api';
	import { quizzes } from '$lib/dashboard/stores/data.store';
	import Button from '$lib/components/Button.svelte';
	import { goto } from '$app/navigation';
	import { toast } from '$lib/dashboard/stores/toast';
	import axios from 'axios';
	import StickyModal from '$lib/components/dashboard/StickyModal.svelte';
	import Heading from '$lib/components/dashboard/typography/Heading.svelte';

	let deleteInputValue: string = $state('');
	let deleteModalOpened: boolean = $state(false);
	let deleteState: string = $state('none');
	let modalEl: HTMLDivElement | undefined = $state();
	let loading: boolean = $state(false);
	let confirmDeleteModalOpened: boolean = $state(false);

	async function openModal() {
		deleteModalOpened = true;
		await tick();

		if (!modalEl) {
			return;
		}

		gsap.fromTo(
			modalEl,
			{ scaleY: 0.2, scaleX: 0.15, transformOrigin: 'bottom', opacity: 0 },
			{ scaleY: 1, scaleX: 1, duration: 0.4, opacity: 1, ease: 'power2.out' }
		);
	}

	async function closeModal() {
		await tick();

		if (!modalEl) {
			return;
		}

		gsap.to(modalEl, {
			scaleY: 0.2,
			scaleX: 0.15,
			opacity: 0,
			transformOrigin: 'bottom',
			duration: 0.3,
			ease: 'power2.in',
			onComplete: () => {
				deleteModalOpened = false;
			}
		});
	}

	async function DeleteMe() {
		if (!$addFormQuiz.id) return;
		let message = '';
		try {
			const response = await api.quiz.DeleteQuiz($addFormQuiz.id);
			$addFormQuiz.id = response.data.id;

			if (response.data.message) {
				message = response.data.message;
			}
		} catch (error: unknown) {
			if (axios.isAxiosError(error)) {
				const message = error.response?.data?.message || 'Błąd serwera';
				toast.error(message);
			} else {
				toast.error('Wystąpił nieoczekiwany błąd');
			}
		} finally {
			const response = await api.quiz.GetAll();
			$quizzes = response.data;
			loading = false;
			toast.success(message);
		}
	}
</script>

<StickyModal
	onClose={() => {
		confirmDeleteModalOpened = false;
	}}
	bind:opened={confirmDeleteModalOpened}
>
	<h1>HarcQuiz mówi:</h1>
	<p>Czy napewno chcesz usunac?</p>
	<Button
		theme="danger"
		onclick={async () => {
			confirmDeleteModalOpened = false;
			await openModal();
		}}
	/>
</StickyModal>

<button
	class="flex size-12 items-center justify-center rounded-lg bg-red-500"
	onclick={async (e) => {
		confirmDeleteModalOpened = true;
	}}><Trash /></button
>

{#if deleteModalOpened}
	{@render DeleteModal()}
{/if}

{#snippet DeleteModal()}
	<!-- svelte-ignore a11y_click_events_have_key_events -->
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<div
		onclick={async (e) => {
			e.stopPropagation();
			await closeModal();
		}}
		class="z-100 fixed inset-0 flex cursor-default items-center justify-center bg-black/50 text-white md:backdrop-blur-lg lg:backdrop-blur-lg"
	>
		<!-- svelte-ignore a11y_no_static_element_interactions -->
		<div
			onclick={(e) => {
				e.stopPropagation();
			}}
			bind:this={modalEl}
			class="w-9/10 relative flex flex-col border border-neutral-800/60 bg-neutral-950 md:w-1/2 md:min-w-100 md:backdrop-blur-none lg:w-1/4 lg:bg-neutral-950"
		>
			<!-- title -->
			<div
				class="mb-2 flex h-10 flex-shrink-0 items-center justify-between border-b border-neutral-700/60 bg-neutral-800/60 px-5"
			>
				<p>Usuwanie</p>
				<button
					class="cursor-pointer text-neutral-500 hover:text-neutral-300"
					onclick={async () => {
						await closeModal();
					}}
				>
					<X />
				</button>
			</div>
			<!-- content -->
			<div class="scrollable flex-1 overflow-y-auto p-6 pt-5">
				<p>Z uwagi na to ze usuwasz cały quiz wpisz liczbę {`${$addFormQuiz.id}`}</p>
				<div class=" mt-5">
					<div class="">
						<input
							bind:value={deleteInputValue}
							class="border-1 w-40 border-neutral-700/60 bg-transparent placeholder-neutral-700"
							placeholder="..."
						/>
					</div>
				</div>
				<div class="mt-5 flex justify-end">
					<Button
						theme="danger"
						loading={deleteState === 'pending'}
						onclick={async () => {
							deleteState = 'pending';
							const promise = await DeleteMe();

							deleteModalOpened = false;
							deleteState = 'none';
							toast.show(`Usunięto quiz z ID: ${$addFormQuiz.id}`, 'success');
							goto('/dashboard/quizzes');
						}}
					>
						<p>Usuń</p>
					</Button>
				</div>
			</div>
		</div>
	</div>
{/snippet}
