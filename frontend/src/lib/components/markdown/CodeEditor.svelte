<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import * as monaco from 'monaco-editor';
	import { browser } from '$app/environment';

	// 1. Definicja propsów z użyciem interfejsu lub typu
	interface Props {
		value: string;
		language?: string; // Opcjonalny, domyślnie 'markdown'
		onchange: (value: string) => void;
	}

	let { value = $bindable(), language = 'markdown', onchange }: Props = $props();

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

{#if browser}
	<div bind:this={editorContainer} class="h-[500px] w-full border border-neutral-900"></div>
{/if}
