<script lang="ts">
	import Icon from '@iconify/svelte';
	import type { PageData } from './$types';
	import { page } from '$app/state';
	import { goto } from '$app/navigation';
	import StorageRecordTile from '../components/StorageRecordTile.svelte';
	import { type StorageRecord } from '$lib/api/requests/storage';

	let { data, params } = $props();

	let menuVisible = $state<boolean>(false);
	let menuX = $state<number>(0);
	let menuY = $state<number>(0);
	let selectedItem = $state<StorageRecord | null>(null);

	function handleRightClick(e: MouseEvent, item: StorageRecord): void {
		e.preventDefault();
		menuX = e.clientX;
		menuY = e.clientY;
		selectedItem = item;
		menuVisible = true;
	}

	function navigateToDir(pathParts: string[], clickedIndex: number): string {
		return `/dashboard/storage/${pathParts.slice(0, clickedIndex + 1).join('/')}`;
	}
</script>

<div class="p-8 flex flex-col gap-4">
	<p>Listed: {data.storage_file_list && data.storage_file_list.length} files</p>
	<div class="flex gap-3">
		<a
			href="/dashboard/storage"
			class="p-1 px-3 hover:underline cursor-pointer rounded-full border border-neutral-700"
		>
			<p class="text-xs">/</p>
		</a>
		{#if params.path !== ''}
			{#each data.path_table as path, i}
				<a
					href={navigateToDir(data.path_table, i)}
					class="p-1 px-3 hover:underline cursor-pointer rounded-full border border-neutral-700"
				>
					<p class="text-xs">{path}</p>
				</a>
			{/each}
		{/if}
	</div>
	<div class="grid grid-cols-[repeat(auto-fill,minmax(160px,1fr))] gap-4">
		{#each data.storage_file_list as { file_size, is_dir, modified, name }}
			<StorageRecordTile
				onrightclick={(e) => handleRightClick(e, { is_dir, name, file_size, modified })}
				{is_dir}
				{name}
				slug={params.path}
				onclick={() => {
					if (is_dir) {
						const currentPath = page.url.pathname.replace(/\/$/, '');
						goto(`${currentPath}/${name}`);
					}
				}}
			/>
		{/each}
	</div>
</div>

{#if menuVisible}
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<div
		style="position: fixed; top: {menuY}px; left: {menuX}px;"
		class="bg-neutral-800 border border-neutral-700 shadow-xl rounded-lg p-2 z-50 w-40"
		onmouseleave={() => (menuVisible = false)}
	>
		<p class="text-xs">{selectedItem?.name}</p>
		<button
			class="block w-full text-left p-2 hover:bg-neutral-700"
			onclick={() => {
				if (selectedItem?.is_dir) {
					goto(`/dashboard/storage/${params.path ? params.path + '/' : ''}${selectedItem.name}`);
				}
				menuVisible = false;
			}}
		>
			Otwórz
		</button>
		<button class="block w-full text-left p-2 hover:bg-neutral-700">Zmien nazwe</button>
		<button class="block w-full text-left p-2 hover:bg-neutral-700 text-red-400">Usuń</button>
	</div>
{/if}
