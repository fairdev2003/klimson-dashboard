import { writable } from 'svelte/store';

export type LogLevel = 'info' | 'success' | 'warn' | 'error' | 'system';

export type DebugEntry = {
	date: number;
	message: string;
	level: LogLevel;
	id: string;
};

function createDebugStore() {
	const MAX_LOGS = 100;
	const { subscribe, update } = writable<DebugEntry[]>([]);

	async function addLog(message: any, level: LogLevel = 'info') {
		let formattedMessage: string;

		if (message instanceof Error) {
			formattedMessage = `${message.message}\nStack trace: ${message.stack}`;
		} else if (typeof message === 'object') {
			formattedMessage = JSON.stringify(message, null, 2);
		} else {
			formattedMessage = String(message);
		}

		update((logs) => {
			const newLog = {
				id: crypto.randomUUID(),
				date: Date.now(),
				message: formattedMessage,
				level
			};

			const newLogs = [...logs, newLog];
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

	const logHelper = (level: LogLevel, ...msg: any[]) => {
		addLog(formatMessage(...msg), level);
	};

	return {
		subscribe,
		log: (...msg: any[]) => logHelper('info', ...msg),
		success: (...msg: any[]) => logHelper('success', ...msg),
		warn: (...msg: any[]) => logHelper('warn', ...msg),
		error: (...msg: any[]) => logHelper('error', ...msg),
		system: (...msg: any[]) => logHelper('system', ...msg),
		clear: () => update(() => [])
	};
}

export const debug = createDebugStore();
