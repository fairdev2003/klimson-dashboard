import { tick } from 'svelte';
import type { Terminal } from './terminal.svelte';

export class TerminalKeyboardEvents {
	constructor(private terminal: Terminal) {}

	public async onSlashKeyClick(e: KeyboardEvent) {
		e.preventDefault();
		await tick();
		this.terminal.inputRef?.focus();
	}

	public async onF2KeyClicked(e: KeyboardEvent) {
		this.terminal.terminalOpened = !this.terminal.terminalOpened;
		this.terminal.fullscreen = false;
		await tick();
		this.terminal.inputRef?.focus();
	}
	public async onF3KeyClicked(e: KeyboardEvent) {
		e.preventDefault();

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
				this.terminal.inputRef.focus();
				const len = this.terminal.inputRef.value.length;
				this.terminal.inputRef.setSelectionRange(len, len);
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

	public async onConsoleClearClicked() {
		this.terminal.debug.clear();
	}

	public async onEnter() {
		if (!this.terminal.commandLineValue) return;

		this.terminal.console.run(this.terminal.commandLineValue);

		await tick();
		this.terminal.inputRef?.focus();
		this.terminal.commandLineValue = '';
	}
}
