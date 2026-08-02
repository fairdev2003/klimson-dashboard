import type { CommandBuilder } from './command_builder.svelte';
import { console_service } from './console_service.svelte';
import type { Terminal } from './terminal.svelte';

export class TerminalIntelisense {
	constructor(private terminal: Terminal) {}

	public intelisenseValue: string | undefined = $state('');

	public searchForSyntaxAndReturnCommand(prompt: string) {
		const commandName = prompt.trim().split(/\s+/)[0];

		if (!commandName) {
			this.intelisenseValue = undefined;
			return;
		}

		const found = this.terminal.console
			.getCommandsRegister()
			.find((e) => e.name.toLowerCase().startsWith(commandName.toLowerCase()));

		this.intelisenseValue = found?.name;
	}
}
