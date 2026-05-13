import { writable } from 'svelte/store';

export type LogLevel = 'info' | 'success' | 'warn' | 'error' | 'system';

export type DebugEntry = {
	date: number;
	message: string;
	level: LogLevel;
};

function createDebugStore() {
	const MAX_LOGS = 100; // Nie chcemy zapchać RAM-u
	const { subscribe, update } = writable<DebugEntry[]>([]);

	function addLog(message: any, level: LogLevel = 'info') {
		// Jeśli message to obiekt, parsujemy go na string
		const formattedMessage = typeof message === 'object'
			? JSON.stringify(message, null, 2)
			: String(message);

		update((logs) => {
			const newLogs = [
				...logs,
				{
					date: Date.now(),
					message: formattedMessage,
					level
				}
			];
			return newLogs.slice(-MAX_LOGS);
		});
	}

	return {
		subscribe,
		// Helpery dla różnych poziomów
		log: (msg: any) => addLog(msg, 'info'),
		success: (msg: any) => addLog(msg, 'success'),
		warn: (msg: any) => addLog(msg, 'warn'),
		error: (msg: any) => addLog(msg, 'error'),
		system: (msg: any) => addLog(msg, 'system'),
		clear: () => update(() => [])
	};
}

export const debug = createDebugStore();