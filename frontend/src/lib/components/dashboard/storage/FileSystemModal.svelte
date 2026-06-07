<script lang="ts">
	import { onMount } from 'svelte';
	import { explorer } from './file.svelte';
	import gsap from 'gsap';
	import { api } from '$lib/api/api';
	import { toast } from '$lib/dashboard/stores/toast';
	import type { StorageRecord } from '$lib/api/requests/storage';

	$effect(() => {
		if ($explorer.open) {
		}
	});

	let storage_records = $state<StorageRecord[]>([]);

	$effect(async () => {
		if (!$explorer.startingPath) {
			toast.show('No starting path defined!');
			return;
		}
		console.log($explorer.startingPath);
		const response = await api.storage.GetStorageRecords($explorer.startingPath);

		storage_records = response.data;
	});
</script>

{#if $explorer.open}
	<div
		onclick={() => {
			explorer.exitExplorer();
		}}
		class="fixed flex flex-col gap-10 z-100 inset-0 bg-black/50 w-full h-full justify-center items-center"
	>
		<div class="">
			<p class="text-md font-bold text-red-500 underline">DEV CONTAINER</p>
		</div>
		<div
			onclick={(e) => {
				e.stopPropagation();
			}}
			class="modal bg-neutral-950 overflow-hidden w-150 border flex flex-col mx-auto border-neutral-700"
		>
			<div class="bg-neutral-800/60 h-10 w-full p-2">
				<p class="text-sm text-neutral-400">Move item to other directory...</p>
			</div>
			<div class="grid grid-cols-6 h-full p-3">
				<div class="col-span-1 flex gap-4 items-center flex-col">
					<button class="size-15 border border-red-500"></button>
					<button class="size-15 border border-red-500"></button>
					<button class="size-15 border border-red-500"></button>
					<button class="size-15 border border-red-500"></button>
					<button class="size-15 border border-red-500"></button>
				</div>
				<div class="col-span-1"></div>
				<div class="flex flex-col gap-2 col-span-4 overflow-y-auto h-150">
					<div class="flex gap-1">
						<button
							onclick={() => {
								if (!$explorer.previousPath) return;
								console.log($explorer.previousPath);
								explorer.gotoPath($explorer.previousPath);
							}}
							class="size-10">{'<'}</button
						>
						<input
							bind:value={$explorer.startingPath}
							class="w-full focus:ring-0 bg-transparent p-2"
						/>
					</div>
					{#each storage_records as record}
						<button
							onclick={() => {
								if ($explorer.startingPath === '/') {
									explorer.setPrevPath('/');
									$explorer.startingPath = $explorer.startingPath + `${record.name}`;
									return;
								}

								explorer.setPrevPath($explorer.startingPath);
								$explorer.startingPath = $explorer.startingPath + `/${record.name}`;
							}}
							class="p-2 w-full bg-neutral-800/60 cursor-pointer flex-gap-2 text-start transition-colors px-4 rounded-lg hover:bg-neutral-700"
						>
							<div class="flex flex-col gap-1">
								<p>
									{record.name}
								</p>
								<p class="text-xs text-neutral-400">
									Last Modified: {new Date(record.modified).toDateString()}
								</p>
							</div>
						</button>
					{/each}
				</div>
			</div>
		</div>
	</div>
{/if}
