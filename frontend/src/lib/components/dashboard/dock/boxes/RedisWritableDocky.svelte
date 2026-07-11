<script lang="ts">
	import { api } from '$lib/api/api';
	import type { BackendResponse, ServerResponse } from '$lib/api/types';

	async function ping_redis(): Promise<ServerResponse<BackendResponse<{ rdbs: string[] }>>> {
		const response: ServerResponse<BackendResponse<{ rdbs: string[] }>> =
			await api.api.get('/redis/keys');

		return response;
	}

	async function get(key: string) {
		const response: ServerResponse<BackendResponse<{ result: any }>> = await api.api.get(
			`/redis/get?key=${key}`
		);

		return response;
	}
</script>

{#await ping_redis()}
	<p>Loading</p>
{:then ping_data}
	{ping_data.data.rdbs.length} redis storage exist
{/await}
