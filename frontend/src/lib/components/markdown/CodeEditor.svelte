<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import * as monaco from 'monaco-editor';
	import { browser } from '$app/environment';

	// 1. Definicja propsów z użyciem interfejsu lub typu
	interface Props {
		value: string;
		file: string;
		language?: string; // Opcjonalny, domyślnie 'markdown'
		onchange: (value: string) => void;
	}

	let { value = $bindable(), language = 'markdown', onchange, file }: Props = $props();

	let editorContainer: HTMLDivElement | undefined = $state();
	let editor: monaco.editor.IStandaloneCodeEditor | undefined = $state();

	onMount(() => {
		if (!editorContainer) return;
		if (browser) {
			monaco.editor.defineTheme('theme', {
				base: 'vs-dark',
				inherit: true,
				rules: [{ token: 'comment', foreground: 'ffa500', fontStyle: 'italic' }],
				colors: {
					'editor.background': '#171717'
				}
			});

			editor = monaco.editor.create(editorContainer, {
				value: value,
				language: language,
				automaticLayout: true,
				theme: 'theme'
			});

			editor.onDidChangeModelContent(() => {
				if (editor) {
					const newValue = editor.getValue();
					onchange(newValue);
				}
			});
		}
	});

	onDestroy(() => {
		editor?.dispose();
	});

	$effect(() => {
		if (editor && value !== editor.getValue()) {
			editor.setValue(value);
		}
	});
</script>

<div class="border border-neutral-700">
	{#if browser}
		<div class="w-full bg-neutral-900/50 p-1">
			<code class="text-sm">{file}</code>
		</div>
		<div bind:this={editorContainer} class="h-[500px] w-full"></div>
	{/if}
</div>
