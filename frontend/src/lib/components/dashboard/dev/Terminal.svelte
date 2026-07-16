<script lang="ts">
	import { debug } from '$lib/dashboard/stores/debug';
	import { onMount, tick } from 'svelte';
	import { Code, Trash2, X, Terminal, GripHorizontal } from '@lucide/svelte';
	import { debugOn } from '$lib/dashboard/stores/persist';
	import gsap from 'gsap';
	import Icon from '@iconify/svelte';
	import { routes } from '$lib/dashboard/stores/data.store';
	import { api } from '$lib/api/api';
	import Loader from '../Loader.svelte';
	import { console_loading, console_service } from './console/console_service.svelte';
	import TerminalRecord from './(components)/TerminalRecord.svelte';
	import TerminalPrefix from './(components)/TerminalPrefix.svelte';

	let logOpened = $state(false);
	let debugContainer: HTMLDivElement | undefined = $state();
	let loaded = $state(false);
	let boxRef: HTMLDivElement | null = $state(null);

	let pos = $state({ x: 20, y: 0 });
	let isDragging = $state(false);
	let commandLineValue: string = $state('');
	let inputRef: HTMLInputElement | undefined = $state();
	let debugFetchLoading = $state(false);

	let fullscreen = $state(false);
	let terminalFocused = $state(false);

	let inputFocused = $state(false);

	function centerTerminal() {
		fullscreen = !fullscreen;
	}

	$effect(() => {
		if ($debug && debugContainer) {
			tick().then(() => {
				debugContainer!.scrollTo({ top: debugContainer!.scrollHeight, behavior: 'smooth' });
			});
		}
	});

	$effect(() => {
		if (terminalFocused) {
			inputFocused = false;
		} else {
		}
	});

	onMount(() => {
		debug.system('Debug component is ready');
		loaded = true;
	});

	function handleMouseDown(e: MouseEvent) {
		isDragging = true;
		const onMouseMove = (m: MouseEvent) => {
			pos.x += m.movementX;
			pos.y += m.movementY;
		};
		const onMouseUp = () => {
			isDragging = false;
			window.removeEventListener('mousemove', onMouseMove);
			window.removeEventListener('mouseup', onMouseUp);
		};
		window.addEventListener('mousemove', onMouseMove);
		window.addEventListener('mouseup', onMouseUp);
	}
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
{#if $debugOn && loaded}
	<div
		bind:this={boxRef}
		onclick={() => {
			if (inputFocused === true) {
				return;
			}
			inputRef?.focus();
			inputFocused = true;
		}}
		class="fixed z-2000 flex flex-col shadow-2xl transition-shadow"
		style="left: {pos.x}px; bottom: {20 - pos.y}px; {isDragging ? 'z-index: 1000' : ''}"
	>
		{#if logOpened}
			<div
				role="presentation"
				onmousedown={handleMouseDown}
				class:w-300={fullscreen}
				class:w-150={!fullscreen}
				class="flex w-full items-center justify-between rounded-t-xl border-b border-neutral-800 bg-neutral-800 p-2 text-white active:cursor-grabbing"
			>
				<div class="flex items-center gap-2 px-2">
					<Terminal class="h-3 w-3 text-blue-500" />
					<p class="text-[10px] font-bold tracking-widest text-white uppercase">
						Dashboard Terminal
					</p>
				</div>

				<div class="flex items-center gap-2">
					<div
						class=" text-xs gap-1 bg-blue-400/20 text-blue-400 mr-5 font-black flex items-center justify-center p-1 px-4 rounded-md"
					>
						<span>{`${$debug.length}`}</span>
						<Icon icon="ic:round-message" />
					</div>
					<button
						class="p-1 text-white transition-colors hover:text-red-500"
						onclick={() => debug.clear()}
						title="Clear console - CTRL + X"
					>
						<Trash2 class="h-3 w-3" />
					</button>
					<button
						class="p-1 text-white transition-colors hover:text-blue-500"
						onclick={() => {
							centerTerminal();
						}}
						title="Close - F2"
					>
						<Icon icon="material-symbols:fullscreen" />
					</button>

					<button
						title="Close terminal - F2"
						class="p-1 text-white hover:text-blue-500"
						onclick={() => (logOpened = false)}
					>
						<X class="h-3 w-3" />
					</button>
				</div>
			</div>

			<div
				bind:this={debugContainer}
				class:w-300={fullscreen}
				class:h-150={fullscreen}
				class:w-150={!fullscreen}
				class:h-75={!fullscreen}
				class="flex h-150 flex-col gap-1 overflow-y-auto rounded-b-xl border border-t-0 border-neutral-800 bg-neutral-950/95 p-4 font-mono text-[11px] backdrop-blur-md"
			>
				{#each $debug as entry (entry.id)}
					<TerminalRecord {entry} />
				{/each}
				<div
					class="flex gap-3 relative items-center border-b border-white/5 pb-1 leading-relaxed last:border-0"
				>
					<span class="text-neutral-600 flex whitespace-nowrap w-20 shrink-0 text-sm">
						<TerminalPrefix />
					</span>
					<input
						bind:value={commandLineValue}
						bind:this={inputRef}
						class="bg-transparent h-5 px-2 w-full border-0 text-neutral-400 text-xs placeholder-neutral-400 outline-none focus:outline-none focus:ring-0"
					/>
					{#if $console_loading}
						<Loader theme="regular" />
					{/if}
				</div>

				{#if $debug.length === 0}
					<div class="flex h-full flex-col items-center justify-center text-neutral-700 italic">
						<Code class="mb-2 h-8 w-8 opacity-20" />
						<p>No active logs in the buffor</p>
					</div>
				{/if}
			</div>
		{:else}
			<button
				class="group flex items-center gap-2 rounded-full border border-neutral-800 bg-neutral-950 p-2 pr-4 shadow-xl transition-all hover:border-blue-500/50 active:scale-95"
				onclick={() => (logOpened = true)}
			>
				<div class="rounded-full bg-blue-500/20 p-2 transition-colors group-hover:bg-blue-500">
					<Terminal class="h-4 w-4 text-blue-500 group-hover:text-white" />
				</div>
				<span class="text-xs font-bold tracking-tighter text-neutral-400 uppercase">Terminal</span>
			</button>
		{/if}
	</div>
{/if}

<svelte:document
	onkeydown={async (e) => {
		if (commandLineValue.length === 0 && e.key === '/') {
			e.preventDefault();
			await tick();
			inputRef?.focus();
		}
		if (e.key === 'F2') {
			logOpened = !logOpened;
			inputFocused = false;
			await tick();
			inputRef?.focus();
		}
		if (e.key === 'F3') {
			e.preventDefault();
			fullscreen = !fullscreen;
		}
		if (e.ctrlKey && e.key === 'x') {
			debug.clear();
		}
		if (e.key === 'Enter') {
			if (!commandLineValue) return;

			console_service.run(commandLineValue);

			await tick();
			inputRef?.focus();
			commandLineValue = '';
		}
	}}
	onclick={(e) => {
		if (boxRef && !boxRef.contains(e.target as Node)) {
			terminalFocused = false;
			debug.system(terminalFocused);
		}
	}}
/>

<style>
	@import 'tailwindcss';

	/* Custom Scrollbar dla efektu terminala */
	div::-webkit-scrollbar {
		width: 4px;
	}
	div::-webkit-scrollbar-track {
		background: transparent;
	}
	div::-webkit-scrollbar-thumb {
		background: #262626;
		border-radius: 10px;
	}
	div::-webkit-scrollbar-thumb:hover {
		background: #3b82f6;
	}

	.magenta {
		@apply text-violet-500;
	}
</style>
