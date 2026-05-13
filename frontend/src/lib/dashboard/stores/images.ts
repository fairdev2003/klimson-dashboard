import type { ImageList } from '$lib/api/types';
import { writable } from 'svelte/store';

export const images = writable<ImageList>([]);
