<script lang="ts">
	import { api } from '$lib/api/api';
	import Loader from '$lib/components/dashboard/Loader.svelte';
	import Button from '$lib/components/dashboard/settings/components/Button.svelte';
	import CodeEditor from '$lib/components/markdown/CodeEditor.svelte';
	import { toast } from '$lib/dashboard/stores/toast.js';
	import { marked, type Tokens } from 'marked';
	import { onMount, tick } from 'svelte';

	let { params, data } = $props();

	let requestedMarkdownFile: string = $state('');

	onMount(async () => {
		const response = await api.api.get(`/interface/bucket/${params.path}`);

		requestedMarkdownFile = response.data;
	});

	let markdown = $derived(requestedMarkdownFile || '');

	let loading = $state(false);

	const renderer = {
		heading(token: Tokens.Heading) {
			return `<h${token.depth} class="font-bold text-2xl my-4">${token.text}</h${token.depth}>`;
		},

		link(token: Tokens.Link) {
			return `<a href="${token.href}" class="text-blue-500" target="_blank">${token.text}</a>`;
		}
	};

	marked.use({ renderer });

	let markdownHTML = $derived.by(() => {
		const newMarked = marked.parse(markdown);
		console.log(newMarked);
		return newMarked;
	});

	async function SendFile() {
		loading = true;
		await api.storage.PushChangedTextFile({ content: markdown }, params.path);

		const response = await api.api.get(`/interface/bucket/${params.path}`);

		requestedMarkdownFile = response.data;
		loading = false;
		toast.success('Zapisano');
	}

	function AddImage() {
		const url = prompt('Image link');

		console.log(markdown.split('\n'));
		const newImageLine = `![image](${url})`;
		markdown = markdown + '\n' + newImageLine;
	}
</script>

<div class="h-dvh p-10">
	<div class="relative z-1">
		{#if loading}
			<div class="w-full h-full absolute z-50 bg-blue-500/50 flex justify-center items-center">
				<Loader />
			</div>
		{/if}

		<CodeEditor
			file={params.path.split('/')[params.path.split('/').length - 1]}
			bind:value={markdown}
			language="markdown"
			onchange={(e: string) => (markdown = e)}
		/>
	</div>
	<div class="flex justify-end mt-10">
		<Button label="Wyslij dane" onclick={SendFile}></Button>
		<Button label="Dodaj zdjecie" onclick={AddImage}></Button>
	</div>
	<div class="prose prose-invert flex max-w-none flex-col w-4xl mx-auto justify-center pb-10">
		{@html markdownHTML}
	</div>
</div>

<svelte:document
	onkeydown={async (e) => {
		if ((e.ctrlKey || e.metaKey) && e.key.toLowerCase() === 's') {
			e.preventDefault();
			e.stopPropagation();
			await SendFile();
		}
	}}
/>
