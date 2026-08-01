import type { CommandBuilder } from './command_builder.svelte';
import type { Terminal } from './terminal.svelte';

export class TerminalIntelisense {
	constructor(private terminal: Terminal) {}

	public intelisenseKeyValue: string | undefined = $state('');

	public searchForSyntaxAndReturnCommand(prompt: string) {
		const commandName = prompt.trim().split(/\s+/)[0];

		if (!commandName) {
			this.intelisenseKeyValue = undefined;
			return;
		}

		const found = this.terminal.console
			.getCommandsRegister()
			.find((e) => e.name.toLowerCase().startsWith(commandName.toLowerCase()));

		this.intelisenseKeyValue = found?.name;
	}
}
