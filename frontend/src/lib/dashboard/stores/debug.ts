import { writable } from 'svelte/store';

export type EntryType = 'message' | 'success' | 'warn' | 'error' | 'system' | 'console';

export type MessageDebugLogMetadata = {
	message?: string;
};

export type TerminaPrefixlDebugLogMetadata = {
	command?: string;
};

export type DebugMetadata = MessageDebugLogMetadata & TerminaPrefixlDebugLogMetadata;

export type TerminalEntry = {
	date: number;
	metadata: DebugMetadata;

	type: EntryType;
	id: string;
};

function createDebugStore() {
	const MAX_LOGS = 300;
	const { subscribe, update } = writable<TerminalEntry[]>([]);

	async function addLog(message: DebugMetadata, type: EntryType = 'message') {
		let formattedMessage: string;

		if (message instanceof Error) {
			formattedMessage = `${message.message}\nStack trace: ${message.stack}`;
		} else if (typeof message === 'object') {
			formattedMessage = JSON.stringify(message, null, 2);
		} else {
			formattedMessage = String(message);
		}

		update((logs) => {
			const newLogs = [
				...logs,
				{
					id: crypto.randomUUID(),
					date: Date.now(),
					metadata: message,
					type
				}
			];
			return newLogs.slice(-MAX_LOGS);
		});
	}

	function formatMessage(...msg: any[]): string {
		return msg
			.map((m) => {
				if (typeof m === 'object' && m !== null) {
					try {
						return JSON.stringify(m, null, 4);
					} catch {
						return String(m);
					}
				}
				if (m instanceof Error) {
					return m.message + (m.stack ? `\n${m.stack}` : '');
				}
				return String(m);
			})
			.join(' ');
	}

	const logHelper = (level: EntryType, ...msg: any[]) => {
		if (level === 'console') {
			addLog({ command: formatMessage(...msg) }, level);
			return;
		}

		addLog({ message: formatMessage(...msg) }, level);
	};

	return {
		subscribe,
		log: (...msg: any[]) => logHelper('message', ...msg),
		success: (...msg: any[]) => logHelper('success', ...msg),
		warn: (...msg: any[]) => logHelper('warn', ...msg),
		error: (...msg: any[]) => logHelper('error', ...msg),
		system: (...msg: any[]) => logHelper('system', ...msg),
		console: (...msg: any[]) => logHelper('console', ...msg),
		clear: () => update(() => [])
	};
}

export const debug = createDebugStore();
