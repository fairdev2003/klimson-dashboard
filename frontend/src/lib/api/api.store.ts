import { persistedWritable } from '$lib/dashboard/stores/persist';

export const base_url = persistedWritable('baseURL', 'https://api.klimson.dev');
