import { get, writable, type Writable } from 'svelte/store';

export type EntryType = 'message' | 'success' | 'warn' | 'error' | 'system' | 'console' | 'silent';

export type MessageDebugLogMetadata = { message?: string };
export type TerminaPrefixlDebugLogMetadata = { command?: string };
export type DebugMetadata = MessageDebugLogMetadata & TerminaPrefixlDebugLogMetadata;

export type TerminalEntry = {
	date: number;
	metadata: DebugMetadata;
	type: EntryType;
	id: string;
};

export class DebugService {
	private readonly MAX_LOGS = 300;
	private store: Writable<TerminalEntry[]> = writable([]);

	public subscribe = this.store.subscribe;

	public log(...msg: any[]) {
		this.logHelper('message', ...msg);
	}
	public success(...msg: any[]) {
		this.logHelper('success', ...msg);
	}
	public warn(...msg: any[]) {
		this.logHelper('warn', ...msg);
	}
	public error(...msg: any[]) {
		this.logHelper('error', ...msg);
	}
	public system(...msg: any[]) {
		this.logHelper('system', ...msg);
	}
	public console(...msg: any[]) {
		this.logHelper('console', ...msg);
	}
	public silent(...msg: any[]) {
		this.logHelper('silent', ...msg);
	}

	public clear() {
		this.store.set([]);
	}

	public logStore() {
		this.addLog({ message: JSON.stringify(get(this.store)) }, 'message');
	}

	private logHelper(level: EntryType, ...msg: any[]) {
		const content = this.formatMessage(...msg);
		const metadata: DebugMetadata =
			level === 'console' ? { command: content } : { message: content };

		this.addLog(metadata, level);
	}

	private addLog(metadata: DebugMetadata, type: EntryType) {
		this.store.update((logs) => {
			const newEntry: TerminalEntry = {
				id: crypto.randomUUID(),
				date: Date.now(),
				metadata,
				type
			};
			return [...logs, newEntry].slice(-this.MAX_LOGS);
		});
	}

	private formatMessage(...msg: any[]): string {
		return msg
			.map((m) => {
				if (m instanceof Error) {
					return `${m.message}\nStack trace: ${m.stack}`;
				}
				if (typeof m === 'object' && m !== null) {
					try {
						return JSON.stringify(m, null, 2);
					} catch {
						return String(m);
					}
				}
				return String(m);
			})
			.join(' ');
	}
}

export const debug = new DebugService();
