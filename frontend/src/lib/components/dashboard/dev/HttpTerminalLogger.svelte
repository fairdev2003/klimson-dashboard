<script lang="ts">
	import Dashboard from '$lib/dashboard/dashboard.svelte';
	import type { Component } from '@lucide/svelte';
	import GetMethod from './methods/GetMethod.svelte';

	const recordMap: Record<string, any> = {
		GET: GetMethod
	};
</script>

<div class="text-white text-sm flex flex-col gap-4">
	{#each Dashboard.http.httpRequests as request}
		<div class="flex justify-between items-center w-full">
			<div class="gap-4 flex">
				{@render Method(request.method)}
				<div class="flex items-center">
					<p class="font-mono text-blue-500 text-xs font-black uppercase">{request.endpoint}</p>
				</div>
			</div>
			<div class="mr-5">
				<p class="text-neutral-400">{request.duration}ms</p>
			</div>
		</div>
	{/each}
</div>

{#snippet Method(method: string)}
	<div
		class:get={method === 'GET'}
		class:post={method === 'POST'}
		class:put={method === 'PUT'}
		class=" w-20 px-4 p-2 items-center flex justify-center rounded-lg"
	>
		<p>{method}</p>
	</div>
{/snippet}

<style>
	@import 'tailwindcss';
	.get {
		@apply text-green-500 bg-green-500/50 font-black;
	}
	.post {
		@apply text-orange-500 bg-orange-500/50 font-black;
	}
	.put {
		@apply text-yellow-500 bg-yellow-500/50 font-black;
	}
</style>
