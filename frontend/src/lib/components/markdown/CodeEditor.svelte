<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import * as monaco from 'monaco-editor';
	import { browser } from '$app/environment';
	import { codeEditorTheme } from './markdown';
	import { dashboard_config } from '$lib/dashboard/stores/persist';

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
			monaco.editor.defineTheme('classic', {
				base: 'vs-dark',
				inherit: true,
				rules: [{ token: 'comment', foreground: 'ffa500', fontStyle: 'italic' }],
				colors: {
					'editor.background': '#171717'
				}
			});

			monaco.editor.defineTheme('dracula', {
				base: 'vs-dark',
				inherit: true,
				rules: [
					{ token: 'comment', foreground: '6272a4', fontStyle: 'italic' },
					{ token: 'string', foreground: 'f1fa8c' },
					{ token: 'keyword', foreground: 'ff79c6' },
					{ token: 'number', foreground: 'bd93f9' },
					{ token: 'type', foreground: '8be9fd' },
					{ token: 'function', foreground: '50fa7b' }
				],
				colors: {
					'editor.background': '#282a36',
					'editor.foreground': '#f8f8f2',
					'editor.lineHighlightBackground': '#44475a',
					'editorCursor.foreground': '#f8f8f2',
					'editorWhitespace.foreground': '#3b4048',
					'editorIndentGuide.background': '#3b4048',
					'editor.selectionBackground': '#44475a'
				}
			});

			editor = monaco.editor.create(editorContainer, {
				value: value,
				language: language,
				automaticLayout: true,
				theme: $dashboard_config.code_theme,
				fontSize: 16
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

<div class="border border-neutral-700 rounded-xl overflow-hidden">
	{#if browser}
		<div class="w-full bg-neutral-900/50 p-1">
			<code class="text-sm">{file}</code>
		</div>
		<div bind:this={editorContainer} class="h-[500px] w-full"></div>
	{/if}
</div>
