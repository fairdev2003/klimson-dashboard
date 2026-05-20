import type { ContextStorageType } from '$lib/api/requests/context_storage';
import { writable } from 'svelte/store';

export const context_storage = writable<ContextStorageType[]>();
