import { writable } from 'svelte/store';

export const imageFile = writable<File | undefined>();
export const imageSrc = writable<string>();
