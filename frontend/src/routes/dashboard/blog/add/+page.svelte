<script lang="ts">
	import { goto } from '$app/navigation';
	import { api } from '$lib/api/api';
	import Button from '$lib/components/Button.svelte';
	import { blogForm, blogFormState } from '$lib/dashboard/stores/blog';
	import { blogs } from '$lib/dashboard/stores/data.store';
	import { debug } from '$lib/dashboard/stores/debug';
	import { toast } from '$lib/dashboard/stores/toast';
	import { marked } from 'marked';
	import { onMount } from 'svelte';
	import { get } from 'svelte/store';

	let previewHtml = $state();

	function updateField(key: 'title' | 'description' | 'html', value: string) {
		blogForm.update((form) => ({
			...form,
			[key]: value
		}));
	}

	onMount(() => {
		blogForm.set({
			title: '',
			html: '',
			description: '',
			public: false
		});
	});

	$effect(() => {
		previewHtml = marked.parse($blogForm.html || '');
	});
</script>

<div class="mx-auto mt-6 flex max-w-5xl flex-col gap-6">
	<!-- FORM -->
	<div class="flex flex-col gap-4 border-1 border-neutral-800 bg-neutral-900 p-6 shadow">
		<h2 class="text-lg font-semibold text-white">Nowy wpis</h2>

		<input
			class="rounded bg-neutral-800 px-3 py-2 text-white ring-1 ring-neutral-700 outline-none focus:ring-blue-500"
			placeholder="Tytuł"
			value={$blogForm.title}
			oninput={(e) => updateField('title', (e.currentTarget as HTMLInputElement).value)}
		/>

		<input
			class="rounded bg-neutral-800 px-3 py-2 text-white ring-1 ring-neutral-700 outline-none focus:ring-blue-500"
			placeholder="Krótki opis"
			value={$blogForm.description}
			oninput={(e) => updateField('description', (e.currentTarget as HTMLInputElement).value)}
		/>

		<textarea
			class="min-h-[200px] resize-none rounded bg-neutral-800 px-3 py-2 text-white ring-1 ring-neutral-700 outline-none focus:ring-blue-500"
			placeholder="Treść (Markdown)"
			value={$blogForm.html}
			oninput={(e) => updateField('html', (e.currentTarget as HTMLTextAreaElement).value)}
		/>

		<div class="mt-5 mb-4 flex items-center">
			<input
				id="default-checkbox"
				type="checkbox"
				bind:checked={$blogForm.public}
				onchange={(e) => {
					debug.log(get(blogForm).public);
				}}
				class="border-default-medium bg-neutral-secondary-medium focus:ring-brand-soft h-4 w-4 rounded-xs border focus:ring-2"
			/>
			<label for="default-checkbox" class="text-heading ms-2 text-sm font-medium select-none"
				>Publiczny</label
			>
		</div>

		<Button
			theme="correct"
			className="mt-2"
			onclick={async () => {
				await api.blog.CreateHero(get(blogForm));
				toast.show('Created Blog', 'info', 4000);
				const response = await api.blog.FetchBlogs();
				const r = response.data;
				blogs.set(r);
				goto('/dashboard/blog');
			}}>Opublikuj wpis</Button
		>
	</div>

	<!-- PREVIEW -->
	<div class="border-1 border-neutral-800 bg-neutral-900 p-6 shadow">
		<h2 class="mb-4 text-lg font-semibold text-white">Podgląd</h2>

		<div class="prose max-w-none prose-invert">{@html previewHtml}</div>
	</div>
</div>
