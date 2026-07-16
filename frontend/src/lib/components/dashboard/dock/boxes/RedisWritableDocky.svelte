<script lang="ts">
	import { goto } from '$app/navigation';
	import { api } from '$lib/api/api';
	import type { BackendResponse, ServerResponse } from '$lib/api/types';
	import Icon from '@iconify/svelte';

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
	<span
		onclick={() => {
			goto('/dashboard/redis');
		}}
		class="p-1 hover:bg-neutral-600 cursor-pointer px-3 items-center bg-neutral-700 rounded-full flex gap-1"
	>
		<Icon icon="devicon:redis" width="18" height="18" />
		<p class="text-red-500 font-black text-xs">{ping_data.data.rdbs.length}</p>
	</span>
{/await}
