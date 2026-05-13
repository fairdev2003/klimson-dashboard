import { writable } from 'svelte/store';

export const recordVisibilityDict = writable<{
	questions: boolean;
	stats: boolean;
	description: boolean;
	tooltip: boolean;
}>({
	questions: false,
	stats: false,
	description: true,
	tooltip: true
});
