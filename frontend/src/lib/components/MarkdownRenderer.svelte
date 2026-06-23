<script lang="ts">
	import { api } from '$lib/api/api';
	import { Marked } from 'marked';
	import { onMount } from 'svelte';
	import { markedHighlight } from 'marked-highlight';
	import hljs from 'highlight.js';
	import 'highlight.js/styles/atom-one-dark.min.css';

	const marked = new Marked(
		markedHighlight({
			langPrefix: 'hljs language-',
			highlight(code, lang) {
				const language = hljs.getLanguage(lang) ? lang : 'plaintext';
				return hljs.highlight(code, { language }).value;
			}
		})
	);

	type Props = {
		markdown: string;
	};

	let { markdown }: Props = $props();

	let markdownText: string = $state('');
	let markdownHTML: string = $state('');

	onMount(async () => {
		const response = await api.api.get(markdown);

		markdownText = response.data;
		markdownHTML = await marked.parse(markdownText);
	});
</script>

<article
	class="article prose prose-invert prose-p:font-[Georgia] prose-p:text-[17px] prose-headings:font-bold mt-5 max-w-none"
>
	{@html markdownHTML}
</article>

<style>
	@import 'tailwindcss';
	@plugin '@tailwindcss/typography';

	.prose p {
		font-family: 'Georgia', serif !important;
	}

	.article {
		@apply prose-a:text-blue-500 prose-img:border-1 prose-img:border-neutral-700 prose-a:text-[17px] prose-pre:bg-transparent prose-code:bg-neutral-900 prose-code:shadow-2xl;
	}
</style>
