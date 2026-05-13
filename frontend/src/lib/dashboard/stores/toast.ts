// toast.ts
import { writable } from 'svelte/store';

export type ToastType = 'success' | 'error' | 'info' | 'warning';
export type Toast = {
	id: string;
	message: string;
	type: ToastType;
	duration: number;
	createdAt: number;
};

function createToastStore() {
	const { subscribe, update } = writable<Toast[]>([]);

	function show(message: any, type: ToastType = 'info', duration = 3000) {
		const id = 'toast-' + crypto.randomUUID();
		update((toasts) => [...toasts, { id, message, type, duration, createdAt: Date.now() }]);
	}

	function info(message: any, duration = 3000) {
		const id = 'toast-' + crypto.randomUUID();
		update((toasts) => [...toasts, { id, message, type: "info", duration, createdAt: Date.now() }]);
	}

	function success(message: any, duration = 3000) {
		const id = 'toast-' + crypto.randomUUID();
		update((toasts) => [...toasts, { id, message, type: "success", duration, createdAt: Date.now() }]);
	}

	function warning(message: any, duration = 3000) {
		const id = 'toast-' + crypto.randomUUID();
		update((toasts) => [...toasts, { id, message, type: "warning", duration, createdAt: Date.now() }]);
	}

	function error(message: any, duration = 3000) {
		const id = 'toast-' + crypto.randomUUID();
		update((toasts) => [...toasts, { id, message, type: "error", duration, createdAt: Date.now() }]);
	}

	function remove(id: string) {
		update((toasts) => toasts.filter((t) => t.id !== id));
	}

	return { subscribe, show, remove, info, success, warning, error };
}

export const toast = createToastStore();
