<script lang="ts">
	import { goto } from '$app/navigation';
	import { api } from '$lib/api/api';
	import { base_url } from '$lib/api/api.store';
	import type { StorageRecord } from '$lib/api/requests/storage';
	import Icon from '@iconify/svelte';
	import { onMount } from 'svelte';

	let files: StorageRecord[] = $state([]);

	onMount(async () => {
		const response = await api.storage.GetLatestStorageRecords();

		files = response.data.files;
	});
</script>

<button
	class="relative overflow-hidden p-3 text-start group rounded-xl flex flex-col col-span-3 border gap-3 border-neutral-700 bg-neutral-800/60"
>
	<div class="flex flex-col">
		<h2 class="text-lg font-bold">Latest Files</h2>
		<p class="text-xs text-neutral-300">See what files you added recently!</p>
	</div>
	<div class="flex flex-col gap-3 relative w-full">
		{#if files.length === 0}
			<div class="absolute right-1/2 bottom-1/2 translate-1/2">
				<p class="text-sm font-md text-neutral-400">No records :c</p>
			</div>
		{/if}

		{#if files.length === 0}
			{#each Array(3) as f}
				<button
					class="text-start flex gap-5 h-15 bg-transparent items-center p-3 transition-colors duration-300 borderrounded-lg"
				>
				</button>
			{/each}
		{:else}
			{#each files as { name, is_dir, file_size }}
				<button
					onclick={() => {
						goto('/dashboard/storage');
					}}
					class="text-start flex gap-5 h-15 w-full items-center p-3 bg-neutral-800/60 hover:bg-blue-600/40 hover:border-blue-500 transition-colors duration-300 cursor-pointer border border-neutral-700 rounded-lg"
				>
					{@render FolderIcon(name, is_dir)}
					<div>
						<p class="text-sm font-semibold">{name}</p>
						<p class="text-xs text-neutral-200">{file_size}B</p>
					</div>
				</button>
			{/each}
		{/if}
		<div class="mx-auto flex">
			{#if files.length > 0}
				<a
					href="/dashboard/storage"
					class="text-blue-500 text-xs h-3 mb-4 hover:underline cursor-pointer">See more.</a
				>
			{/if}
		</div>
	</div>
</button>

{#snippet FolderIcon(name: string, is_dir: boolean)}
	{@const size = '30'}
	{#if !is_dir && (name.endsWith('.png') || name.endsWith('.jpg') || name.endsWith('.jpeg') || name.endsWith('.svg') || name.endsWith('.webp'))}{:else if !is_dir && name.endsWith('.pdf')}
		<Icon icon="mingcute:pdf-fill" width={size} height={size} />
	{:else if !is_dir && name.endsWith('.sfm')}
		<Icon icon="mage:compact-disk-fill" width={size} height={size} />
	{:else if !is_dir && name.endsWith('.gif')}
		<Icon icon="fluent:gif-16-filled" width={size} height={size} />
	{:else if !is_dir && name.endsWith('.mp3')}
		<Icon icon="rivet-icons:audio-solid" width={size} height={size} />
	{:else if is_dir}
		<Icon icon="material-symbols:folder" width={size} height={size} />
	{:else}
		<Icon icon="material-symbols:description" width={size} height={size} />
	{/if}
{/snippet}
