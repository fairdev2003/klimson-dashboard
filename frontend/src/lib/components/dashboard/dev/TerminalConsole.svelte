<script lang="ts">
	import { debug } from '$lib/dashboard/stores/debug';
	import { onMount, tick } from 'svelte';
	import { Trash2, X, Terminal } from '@lucide/svelte';
	import Icon from '@iconify/svelte';
	import TerminalRecord from './(components)/TerminalRecord.svelte';
	import { terminal } from '$lib/terminal/logic';
	import TerminalInput from './(components)/input/TerminalInput.svelte';
	import { console_service, ConsoleService } from './console/console_service.svelte';
	import { Dashboard } from '$lib/dashboard/logic';
	import { AutoComplete } from './console/command_builder.svelte';
	import { bold } from '$lib/terminal/style';

	$effect(() => {
		if ($debug && terminal.debugContainer) {
			tick().then(() => {
				terminal.debugContainer!.scrollTo({ top: terminal.debugContainer!.scrollHeight });
			});
		}
	});

	$effect(() => {
		if (terminal.terminalFocused) {
			terminal.inputFocused = false;
		} else {
		}
	});

	onMount(() => {
		terminal.loaded = true;
		debug.log('Hello, World');
	});

	function handleMouseDown(e: MouseEvent) {
		terminal.isDragging = true;
		const onMouseMove = (m: MouseEvent) => {
			terminal.pos.x += m.movementX;
			terminal.pos.y += m.movementY;
		};
		const onMouseUp = () => {
			terminal.isDragging = false;
			window.removeEventListener('mousemove', onMouseMove);
			window.removeEventListener('mouseup', onMouseUp);
		};
		window.addEventListener('mousemove', onMouseMove);
		window.addEventListener('mouseup', onMouseUp);
	}

	async function setCommand(value: string) {
		terminal.commandLineValue = value;

		await tick();

		requestAnimationFrame(() => {
			if (terminal.inputRef) {
				terminal.inputRef.focus();
				const len = terminal.inputRef.value.length;
				terminal.inputRef.setSelectionRange(len, len);
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
	class:w-300={terminal.fullscreen}
	class:w-150={!terminal.fullscreen}
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
	bind:this={terminal.debugContainer}
	class:w-300={terminal.fullscreen}
	class:h-150={terminal.fullscreen}
	class:w-200={!terminal.fullscreen}
	class:h-100={!terminal.fullscreen}
	class="flex h-150 flex-col gap-1 overflow-y-auto bg-neutral-950/95 p-4 font-mono text-[11px] backdrop-blur-md"
>
	{#each $debug as entry (entry.id)}
		<TerminalRecord naming={{ name: 'kuel', path: 'dashboard' }} {entry} />
	{/each}

	<TerminalInput
		bind:inputRef={terminal.inputRef}
		bind:commandLineValue={terminal.commandLineValue}
	/>
</div>

<svelte:document
	onkeydown={async (e) => {
		if (terminal.commandLineValue.length === 0 && e.key === '/') {
			e.preventDefault();
			await tick();
			terminal.inputRef?.focus();
		}
		if (e.key === 'F2') {
			terminal.terminalOpened = !terminal.terminalOpened;
			terminal.inputFocused = false;
			await tick();
			terminal.inputRef?.focus();
		}
		if (e.key === 'F3') {
			e.preventDefault();
			if (!terminal.terminalOpened) {
				terminal.terminalOpened = true;
				terminal.fullscreen = true;
				return;
			}
			terminal.fullscreen = !terminal.fullscreen;
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
			if (!terminal.commandLineValue) return;

			const public_console_service = new ConsoleService();
			public_console_service
				.registerCommand('hello')
				.setDescription('Program will wave to you!')
				.setAction(() => {
					debug.log('Hello, World');
				});

			public_console_service.onUnknownCommand((input, name) => {
				debug.console(input);

				debug.system(`Command with name '${name}' does not exist!`);
				debug.system(`Type 'cmds' to view available commands.`);
			});

			public_console_service.onCommand((command, input) => {
				if (!command) {
					return;
				}
				if (!input) {
					return;
				}

				terminal.set_input({ user_input: input, id: terminal.input_history.length + 1 });
				debug.console(input);
				terminal.last_record_user_iterator = -1;
			});

			public_console_service
				.registerCommand('cmds')
				.setDescription('List of all available commands to use in dashboard terminal.')
				.addArgHandler((arg) => arg, {
					customName: 'isDev',
					auto_complete_args: AutoComplete.bool,
					required: false
				})
				.setAction((args) => {
					const [dev] = args;
					const command_register = public_console_service.getCommandsRegister();
					debug.log(`\n`);
					debug.log(`(${command_register.length}) Commands: `);
					debug.log(`\n`);

					let cmds_string: string = '';

					command_register.forEach((command) => {
						const desc = command.description ? ` - ${command.description}` : '';

						cmds_string = cmds_string + `${command.name}${desc}\n\n`;
						debug.raw(`${command.name}${desc}`);
					});

					if (Boolean(dev)) {
						debug.format(bold(terminal.console.dumpAvailableCommands()));
					}
				});

			public_console_service.run(terminal.commandLineValue);
			await tick();
			terminal.inputRef?.focus();
			terminal.commandLineValue = '';
		}
	}}
	onclick={(e) => {
		if (terminal.boxRef && !terminal.boxRef.contains(e.target as Node)) {
			terminal.terminalFocused = false;
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
