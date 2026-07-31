<script lang="ts">
	import { Settings, Terminal, Trash2, X } from '@lucide/svelte';
	import { terminal } from './console/terminal.svelte';
	import Icon from '@iconify/svelte';
	import { debug } from '$lib/terminal/logic';
	import TerminalSettings from './TerminalSettings.svelte';
	import TerminalPageDropdown from './TerminalPageDropdown.svelte';
	import { blur } from 'svelte/transition';

	function handleMouseDown(e: MouseEvent) {
		e.preventDefault();
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
</script>

<div
	role="presentation"
	onmousedown={handleMouseDown}
	class:w-300={terminal.fullscreen}
	class:w-150={!terminal.fullscreen}
	class="flex relative w-full items-center justify-between border border-blue-800 bg-blue-800 p-2 text-white active:cursor-grabbing"
>
	<TerminalPageDropdown />

	<div class="flex items-center gap-2">
		<div
			class=" text-xs gap-1 bg-blue-400/20 text-blue-400 mr-5 font-black flex items-center justify-center p-1 px-4"
		>
			<span>{`${$debug.length}`}</span>
			<Icon icon="ic:round-message" />
		</div>
		<button
			class="p-1 text-white transition-colors hover:text-orange-500"
			onclick={() => terminal.settings.toggleSettings()}
			title="Clear console - CTRL + X"
		>
			<Settings class="h-3 w-3" />
		</button>
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
				terminal.fullscreen = !terminal.fullscreen;
			}}
			title="Close - F2"
		>
			<Icon icon="material-symbols:fullscreen" />
		</button>

		<button
			title="Close terminal - F2"
			class="p-1 text-white hover:text-blue-500"
			onclick={() => (terminal.terminalOpened = !terminal.terminalOpened)}
		>
			<X class="h-3 w-3" />
		</button>
	</div>
	<TerminalSettings bind:opened={terminal.settings.settings_opened} />
</div>
