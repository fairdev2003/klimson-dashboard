<script lang="ts">
	import { ArrowRight, HashIcon } from '@lucide/svelte';
	import { onDestroy, tick } from 'svelte';
	import Tooltip from './Tooltip.svelte';
	import { fade } from 'svelte/transition';

	let open = $state(false);
	let temp_options: { name: string; path: string }[] = $state([
		{ name: 'start', path: '/dashboard' },
		{ name: 'quiz', path: '/dashboard/quizzes' },
		{ name: 'questions', path: '/dashboard/questions' },
		{ name: 'answers', path: '/dashboard/answers' },
		{ name: 'api', path: '/dashboard/routes' },
		{ name: 'blog', path: '/dashboard/blog' }
	]);

	let hashIcon: HTMLDivElement | null = null;
	let boxRef: HTMLDivElement | null = null;
	let inputRef: HTMLInputElement | null = null;

	function handleClickOutside(e: MouseEvent) {
		if (!boxRef || !boxRef.contains(e.target as Node)) {
			if (!hashIcon || !hashIcon.contains(e.target as Node)) {
				open = false;
			}
		}
	}

	async function openBox(e: MouseEvent) {
		e.stopPropagation();

		open = !open;

		await tick();
		inputRef?.focus();

		document.addEventListener('mousedown', handleClickOutside);
	}

	function closeBox() {
		open = false;
		document.removeEventListener('mousedown', handleClickOutside);
	}
</script>

<Tooltip content="Context" disabled={open} position="bottom">
	<div class="relative">
		<div class="cursor-pointer text-neutral-500 hover:text-neutral-300">
			<div bind:this={hashIcon}>
				<HashIcon onclick={openBox} />
			</div>
		</div>
		{#if open}
			{@render ContextContent()}
		{/if}
	</div>
</Tooltip>

{#snippet ContextContent()}
	<div
		bind:this={boxRef}
		in:fade={{ duration: 150 }}
		out:fade={{ duration: 150 }}
		class="no-hover absolute top-9 -left-2 z-10 w-[180px] border-1 border-neutral-700/60 bg-neutral-800"
	>
		<div class="flex flex-col gap-2 py-2">
			{#each temp_options as { name, path }, num}
				{@render ContextOption(name, path, num)}
			{/each}
		</div>
	</div>
{/snippet}

{#snippet ContextOption(name: string, path: string, index: number)}
	<div>
		{#if index !== temp_options.length - 1}
			<span class="flex items-center gap-1 border-b-1 border-neutral-700/60">
				<span class="px-1 text-green-500">
					{'R/'}
				</span>
				<a href={path}>{name}</a>
			</span>{:else}
			<span class="flex items-center gap-1">
				<span class="px-1 text-green-500">
					{'R/'}
				</span>
				<a href={path}>{name}</a>
			</span>
		{/if}
	</div>
{/snippet}

<style>
	.no-hover:hover {
		cursor: default !important;
	}

	a:hover {
		text-decoration: underline;
	}
</style>
