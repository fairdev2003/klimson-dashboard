<script lang="ts">
	import { goto } from '$app/navigation';
	import { api } from '$lib/api/api';
	import { Api } from '$lib/api/api';
	import Button from '$lib/components/Button.svelte';
	import { blogForm } from '$lib/dashboard/stores/blog';
	import { blogs } from '$lib/dashboard/stores/data.store';
	import { heroForm } from '$lib/dashboard/stores/hero';
	import { toast } from '$lib/dashboard/stores/toast';
	import { Delete, Trash, X } from '@lucide/svelte';
	import gsap from 'gsap';
	import { tick } from 'svelte';

	let deleteInputValue: string = $state('');
	let deleteModalOpened: boolean = $state(false);
	let deleteState: string = $state('none');
	let modalEl: HTMLDivElement | undefined = $state();
	let loading: boolean = $state(false);

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
</script>

<button
	class="flex size-12 cursor-pointer items-center justify-center rounded-lg bg-red-500"
	onclick={async () => {
		await openModal();
	}}><Trash /></button
>

{#if deleteModalOpened}
	{@render DeleteModal()}
{/if}

{#snippet DeleteModal()}
	<!-- svelte-ignore a11y_click_events_have_key_events -->
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<div
		onclick={async () => await closeModal()}
		class="fixed inset-0 z-50 flex items-center justify-center bg-black/50 text-white md:backdrop-blur-lg lg:backdrop-blur-lg"
	>
		<!-- svelte-ignore a11y_no_static_element_interactions -->
		<div
			onclick={(e) => {
				e.stopPropagation();
			}}
			bind:this={modalEl}
			class="relative flex w-9/10 flex-col border border-neutral-800/60 bg-neutral-950 md:w-1/2 md:min-w-[400px] md:backdrop-blur-none lg:w-1/4 lg:bg-neutral-950"
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
				<p>Czy napewno chcesz usunac ten</p>

				<div class="mt-5 flex justify-end">
					<Button
						theme="danger"
						loading={deleteState === 'pending'}
						onclick={async () => {
							deleteState = 'pending';
							await api.blog.DeleteBlog({ id: $blogForm.id });

							deleteModalOpened = false;
							deleteState = 'none';
							toast.show(`Usunięto hero z ID: ${$blogForm.id}`, 'success');
							const response = await api.blog.FetchBlogs();
							blogs.set(response.data);
							goto('/dashboard/blog');
						}}
					>
						<p>Usuń</p>
					</Button>
				</div>
			</div>
		</div>
	</div>
{/snippet}
