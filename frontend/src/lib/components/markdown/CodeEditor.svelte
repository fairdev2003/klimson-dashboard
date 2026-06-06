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
			editor = monaco.editor.create(editorContainer, {
				value: value,
				language: language,
				theme: 'vs-dark',
				automaticLayout: true
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
	<div bind:this={editorContainer} class="h-[500px] w-full border border-neutral-700"></div>
{/if}
