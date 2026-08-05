import { dev } from '$app/environment';
import { persistedWritable } from '$lib/dashboard/stores/persist';

export const base_url = persistedWritable(
	'baseURL',
	dev ? 'http://localhost:8090' : 'https://api.klimson.dev'
);
