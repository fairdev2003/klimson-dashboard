<script lang="ts">
	import { debug } from '$lib/dashboard/stores/debug';
	import { onMount, tick } from 'svelte';
	import { Trash2, X, Terminal } from '@lucide/svelte';
	import Icon from '@iconify/svelte';
	import TerminalRecord from './(components)/TerminalRecord.svelte';
	import { terminal } from '$lib/terminal/logic';
	import TerminalInput from './(components)/input/TerminalInput.svelte';
	import { console_service } from './console/console_service.svelte';
	import { Dashboard } from '$lib/dashboard/logic';

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
				debugContainer!.scrollTo({ top: debugContainer!.scrollHeight });
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
		loaded = true;
		debug.log('Hello, World');
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

	async function setCommand(value: string) {
		commandLineValue = value;

		await tick();

		requestAnimationFrame(() => {
			if (inputRef) {
				inputRef.focus();
				const len = inputRef.value.length;
				inputRef.setSelectionRange(len, len);
			}
		});
	}

	const full_terminal_naming = $derived(
		terminal.terminal_naming + Dashboard.state.current_directory
	);

	type Props = {
		standalone?: boolean;
	};

	let { standalone = false }: Props = $props();
</script>

<div
	role="presentation"
	onmousedown={handleMouseDown}
	class:w-300={fullscreen}
	class:w-150={!fullscreen}
	class="flex w-full items-center justify-between border border-blue-800 bg-blue-800 p-2 text-white"
>
	<div class="flex items-center gap-2 px-2">
		<Terminal class="h-3 w-3 text-blue-500" />
		<p class="text-[10px] font-bold tracking-widest text-white uppercase">
			Terminal (Sandbox only)
		</p>
	</div>

	<div class="flex items-center gap-2">
		<div
			class=" text-xs gap-1 bg-blue-400/20 text-blue-400 mr-5 font-black flex items-center justify-center p-1 px-4"
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
		<button class="p-1 text-white transition-colors opacity-50" title="Not available in this view">
			<Icon icon="material-symbols:fullscreen" />
		</button>

		<button title="Not available in this view" class="p-1 opacity-50 text-white">
			<X class="h-3 w-3" />
		</button>
	</div>
</div>

<div
	bind:this={debugContainer}
	class:w-300={fullscreen}
	class:h-150={fullscreen}
	class:w-200={!fullscreen}
	class:h-100={!fullscreen}
	class="flex h-150 flex-col gap-1 overflow-y-auto border border-neutral-800 bg-neutral-950/95 p-4 font-mono text-[11px] backdrop-blur-md"
>
	{#each $debug as entry (entry.id)}
		<TerminalRecord naming={full_terminal_naming} {entry} />
	{/each}

	<TerminalInput bind:inputRef bind:commandLineValue />
</div>

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
			if (!logOpened) {
				logOpened = true;
				fullscreen = true;
				return;
			}
			fullscreen = !fullscreen;
		}

		if (e.key === 'ArrowDown') {
			e.preventDefault();
			if (terminal.input_history.length === 0) return;

			if (terminal.last_record_user_iterator >= terminal.input_history.length - 1) {
				terminal.last_record_user_iterator = terminal.input_history.length;
				setCommand('');
			} else {
				terminal.last_record_user_iterator++;
				setCommand(terminal.input_history[terminal.last_record_user_iterator].user_input);
			}
		}
		if (e.key === 'ArrowUp') {
			e.preventDefault();
			if (terminal.input_history.length === 0) return;

			if (terminal.last_record_user_iterator === -1) {
				terminal.last_record_user_iterator = terminal.input_history.length - 1;
			} else if (terminal.last_record_user_iterator > 0) {
				terminal.last_record_user_iterator--;
			}

			setCommand(terminal.input_history[terminal.last_record_user_iterator].user_input);
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
