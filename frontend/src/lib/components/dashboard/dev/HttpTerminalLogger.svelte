<script lang="ts">
	import { base_url } from '$lib/api/api.store';
	import Dashboard from '$lib/dashboard/dashboard.svelte';
	import Icon from '@iconify/svelte';
	import MovingTooltip from '../MovingTooltip.svelte';
</script>

<div class="text-white text-sm flex flex-col gap-4">
	{#each Dashboard.http.httpRequests as request}
		<div class:error={request.isError} class="flex justify-between items-center w-full">
			<div class="gap-4 flex">
				{@render Method(request.method, request.isError)}
				<div class="flex items-center">
					<MovingTooltip>
						{#snippet tooltipContent()}
							<div class="text-xs flex gap-2">
								<span class="font-black text-purple-500">{$base_url}</span>
								<span>{request.endpoint}</span>
							</div>
						{/snippet}

						<a
							class="font-mono flex items-center gap-1 text-blue-500 text-xs font-black uppercase hover:underline cursor-pointer"
							href={`${$base_url}${request.endpoint}`}
						>
							<p>{request.endpoint}</p>
							{#if request.endpoint.includes('admin')}
								<div class="text-neutral-400">
									<Icon icon="material-symbols:lock" />
								</div>
							{/if}
						</a>
					</MovingTooltip>
				</div>
			</div>
			<div class="mr-5">
				{#if request.duration}
					<p
						class:text-green-500={request.duration > 0 && request.duration < 100}
						class:text-orange-500={request.duration >= 100 && request.duration < 200}
						class:text-red-500={request.duration > 200}
						class="font-black"
					>
						{request.duration}ms
					</p>
				{/if}
			</div>
		</div>
	{/each}
</div>

{#snippet Method(method: string, hasError: boolean)}
	<div
		class:get={method === 'GET' && !hasError}
		class:post={method === 'POST' && !hasError}
		class:put={method === 'PUT' && !hasError}
		class:error={hasError}
		class=" w-20 px-4 p-2 items-center flex justify-center rounded-lg"
	>
		<p>{method}</p>
	</div>
{/snippet}

<style>
	@import 'tailwindcss';
	.error {
		@apply bg-red-500/50 rounded-lg;
	}
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
