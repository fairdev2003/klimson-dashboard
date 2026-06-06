<script lang="ts">
	import { api } from '$lib/api/api';
	import Button from '$lib/components/dashboard/settings/components/Button.svelte';
	import CodeEditor from '$lib/components/markdown/CodeEditor.svelte';
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
	}
</script>

<div class="h-dvh p-10">
	<div class="relative w-full h-1/2 z-1">
		{#if loading}
			<div class="w-full h-full absolute z-2 bg-blue-500/50"></div>
		{/if}
		<!-- <textarea bind:value={markdown} class="h-full bg-transparent w-full"></textarea> -->
		<CodeEditor
			bind:value={markdown}
			language="markdown"
			onchange={(e: string) => (markdown = e)}
		/>
	</div>
	<div class="flex justify-end">
		<Button label="Wyslij dane" onclick={SendFile}></Button>
	</div>
	<div class="prose prose-invert">
		{@html markdownHTML}
	</div>
</div>
