import { onMount } from 'svelte';
import { console_service } from './console_service.svelte';
import { debug } from '$lib/dashboard/stores/debug';
import { TerminalKeyboardEvents } from './shortcuts.svelte';
import { TerminalSettings } from './settings.svelte';
import { TerminalIntelisense } from './intelisense.svelte';

type TerminalSavedInput = {
	user_input: string;
	id: number;
};

export type TerminalNaming = {
	name: string;
	path: string;
};

export type TerminalPage = 'user' | 'http' | 'only-logs';

export class Terminal {
	public terminalOpened: boolean = $state(false);
	public terminalPage: TerminalPage = $state('user');
	private latest_input_records: TerminalSavedInput[] = $state([]);
	private t_naming: TerminalNaming | undefined = $state();
	private t_disabled: boolean = $state(false);
	public last_record_user_iterator = $state(-1);
	private readonly MAX_INPUT_MEMORY = 30;
	public debugContainer: HTMLDivElement | undefined = $state();
	public loaded = $state(false);
	public boxRef: HTMLDivElement | null = $state(null);
	public terminalFocused: boolean = $state(false);

	public pos = $state({ x: 20, y: 0 });
	public isDragging = $state(false);
	public commandLineValue: string = $state('');
	public inputRef: HTMLInputElement | undefined = $state();

	public fullscreen = $state(false);

	public focusInput() {
		this.inputRef?.focus();
	}

	public unfocusInput() {
		this.inputRef?.blur();
	}

	public inputFocused = $state(false);

	public get debug() {
		return debug;
	}

	public keyboard_event: TerminalKeyboardEvents;
	public settings: TerminalSettings;
	private _intelisense: TerminalIntelisense;

	constructor() {
		this.keyboard_event = new TerminalKeyboardEvents(this);
		this.settings = new TerminalSettings(this);
		this._intelisense = new TerminalIntelisense(this);
	}

	public get console() {
		return console_service;
	}

	public get intelisense() {
		return this._intelisense;
	}

	public set_input(record: TerminalSavedInput) {
		this.latest_input_records = [...this.latest_input_records, record].slice(
			-this.MAX_INPUT_MEMORY
		);
	}

	public get input_history() {
		return this.latest_input_records;
	}

	public toggle_terminal() {
		this.t_disabled = !this.t_disabled;
	}
	public set_terminal_naming(naming: TerminalNaming) {
		this.t_naming = naming;
	}

	public get terminal_naming() {
		return this.t_naming;
	}

	public get terminal_disabled() {
		return this.t_disabled;
	}
}

export const terminal = new Terminal();
