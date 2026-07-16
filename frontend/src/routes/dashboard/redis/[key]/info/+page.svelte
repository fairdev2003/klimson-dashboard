<script lang="ts">
	import { api } from '$lib/api/api';
	import type { BackendResponse, ServerResponse } from '$lib/api/types';
	import FancyLoader from '../../(components)/FancyLoader.svelte';

	let { params } = $props();

	type RedisKeyInfo = { memory_usage: number; type: string; ttl: string; idle: string };

	async function fetch_info(): Promise<ServerResponse<BackendResponse<RedisKeyInfo>>> {
		const response: ServerResponse<BackendResponse<RedisKeyInfo>> = await api.api.get(
			`/redis/key/info?key=${params.key}`
		);

		return response;
	}
</script>

{#await fetch_info()}
	<div class="flex mx-auto justify-center mt-5 w-full">
		<FancyLoader />
	</div>
{:then data}
	{data.data.succes}
{/await}
