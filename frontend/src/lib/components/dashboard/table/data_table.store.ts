import { writable } from 'svelte/store';

export const highlightedFields = writable<`${string}-${string}-${string}-${string}-${string}`[]>(
	[]
);
