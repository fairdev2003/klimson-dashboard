import type { Terminal } from './terminal.svelte';

export class TerminalSettings {
	constructor(private terminal: Terminal) {}

	public settings_opened: boolean = $state(false);

	public openSettings() {
		this.settings_opened = true;
	}
	public closeSettings() {
		this.settings_opened = false;
	}
	public toggleSettings() {
		this.settings_opened = !this.settings_opened;
	}
}
