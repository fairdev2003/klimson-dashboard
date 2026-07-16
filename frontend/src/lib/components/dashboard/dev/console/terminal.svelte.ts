import { console_service } from './console_service.svelte';

type TerminalSavedInput = {
	user_input: string;
	id: number;
};

export type TerminalNaming = {
	name: string;
	path: string;
};

class Terminal {
	private latest_input_records: TerminalSavedInput[] = $state([]);
	private t_naming: TerminalNaming | undefined = $state();
	public last_record_user_iterator = $state(-1);
	private readonly MAX_INPUT_MEMORY = 30;

	public get console_service() {
		return console_service;
	}

	public set_input(record: TerminalSavedInput) {
		this.latest_input_records = [...this.latest_input_records, record].slice(
			-this.MAX_INPUT_MEMORY
		);
	}

	public get input_history() {
		return this.latest_input_records;
	}

	public set_terminal_naming(naming: TerminalNaming) {
		this.t_naming = naming;
	}

	public get terminal_naming() {
		return this.t_naming;
	}
}

export const terminal = new Terminal();
