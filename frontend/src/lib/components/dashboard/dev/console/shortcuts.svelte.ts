import { tick } from 'svelte';
import type { Terminal } from './terminal.svelte';
import { command } from '$app/server';

export class TerminalKeyboardEvents {
	constructor(private terminal: Terminal) {}
	public async onAnyKeyClicked(e: KeyboardEvent) {
		if (e.ctrlKey || e.metaKey) {
			return;
		}

		const selection = window.getSelection();
		if (selection && selection.toString().length > 0) {
			return;
		}

		if (this.terminal.terminalFocused && this.terminal.inputRef) {
			if (document.activeElement !== this.terminal.inputRef) {
				this.terminal.inputRef.focus();
			}
		}
	}

	public async handleTerminalInputEnclosure(e: KeyboardEvent) {
		const pairs: Record<string, { left: string; right: string }> = {
			'"': { left: '"', right: '"' },
			"'": { left: "'", right: "'" },
			'`': { left: '`', right: '`' },
			'(': { left: '(', right: ')' },
			'[': { left: '[', right: ']' },
			'{': { left: '{', right: '}' }
		};

		if (e.shiftKey && Object.keys(pairs).includes(e.key)) {
			e.preventDefault();
			const selectedText = window.getSelection()?.toString();
			if (selectedText) {
				this.terminal.debug.log(selectedText);
				this.terminal.commandLineValue = this.terminal.commandLineValue.replace(
					selectedText,
					`${pairs[e.key].left}${selectedText}${pairs[e.key].right}`
				);
			} else {
				if (!this.terminal.commandLineValue.includes(' ')) {
					this.terminal.commandLineValue =
						this.terminal.commandLineValue + ` ${pairs[e.key].left}${pairs[e.key].right}`;
				} else {
					this.terminal.commandLineValue =
						this.terminal.commandLineValue + `${pairs[e.key].left}${pairs[e.key].right}`;
				}
			}
		}
	}

	public async onEscapeKeyClicked(e: KeyboardEvent) {
		if (this.terminal.inputRef) {
			e.preventDefault();
			this.terminal.unfocusInput();
		}
	}

	public async onSlashKeyClick(e: KeyboardEvent) {
		if (this.terminal.terminalOpened) {
			e.preventDefault();
			await tick();
			this.terminal.focusInput();
			this.terminal.terminalFocused = true;
		}
	}

	public async onF2KeyClicked(e: KeyboardEvent) {
		this.terminal.terminalOpened = !this.terminal.terminalOpened;
		this.terminal.fullscreen = false;
		this.terminal.terminalFocused = true;

		await tick();
		this.terminal.focusInput();
	}
	public async onF3KeyClicked(e: KeyboardEvent) {
		e.preventDefault();
		this.terminal.terminalFocused = true;

		if (!this.terminal.fullscreen) {
			this.terminal.fullscreen = true;
			return;
		}
		this.terminal.fullscreen = true;
		this.terminal.terminalOpened = !this.terminal.terminalOpened;
	}

	private async setCommand(value: string) {
		this.terminal.commandLineValue = value;

		await tick();

		requestAnimationFrame(() => {
			if (this.terminal.inputRef) {
				const len = this.terminal.inputRef.value.length;
				this.terminal.inputRef.setSelectionRange(len, len);
				this.terminal.focusInput();
			}
		});
	}

	public async onArrowDownClicked(e: KeyboardEvent) {
		e.preventDefault();
		if (this.terminal.input_history.length === 0) return;

		if (this.terminal.last_record_user_iterator >= this.terminal.input_history.length - 1) {
			this.terminal.last_record_user_iterator = this.terminal.input_history.length;
			this.setCommand('');
		} else {
			this.terminal.last_record_user_iterator++;
			this.setCommand(
				this.terminal.input_history[this.terminal.last_record_user_iterator].user_input
			);
		}
	}

	public async onArrowUpClicked(e: KeyboardEvent) {
		e.preventDefault();
		if (this.terminal.input_history.length === 0) return;

		if (this.terminal.last_record_user_iterator === -1) {
			this.terminal.last_record_user_iterator = this.terminal.input_history.length - 1;
		} else if (this.terminal.last_record_user_iterator > 0) {
			this.terminal.last_record_user_iterator--;
		}

		this.setCommand(
			this.terminal.input_history[this.terminal.last_record_user_iterator].user_input
		);
	}

	public async onArrowRightClicked(e: KeyboardEvent) {
		if (e.key === 'ArrowRight') {
			if (this.terminal.intelisense.intelisenseValue) {
				this.terminal.commandLineValue = this.terminal.intelisense.intelisenseValue;
			}
		}
	}

	public async onConsoleClearClicked() {
		this.terminal.debug.clear();
	}

	public async onEnter() {
		if (!this.terminal.commandLineValue) return;

		this.terminal.console.run(this.terminal.commandLineValue);

		await tick();
		this.terminal.focusInput();
		this.terminal.commandLineValue = '';
	}
}
