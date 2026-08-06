<script lang="ts">
	import { goto } from '$app/navigation';
	import { base_url } from '$lib/api/api.store';
	import type { StorageRecord } from '$lib/api/requests/storage';
	import { storage_logic } from '$lib/dashboard/storage/storage.svelte';
	import { toast } from '$lib/dashboard/stores/toast';
	import Icon from '@iconify/svelte';
	import { on } from 'svelte/events';

	type Props = {
		menuVisible: boolean;
		menuX: number;
		menuY: number;
		selectedItem: StorageRecord | null;
		params: { path: string };
		onItemRename: (oldName: string, newName: string) => Promise<void>;
		onItemDelete: (name: string, isDir: boolean) => Promise<void>;
	};

	let {
		menuVisible = $bindable(false),
		menuX = $bindable(0),
		menuY = $bindable(0),
		selectedItem = $bindable(null),
		params = $bindable({ path: '' }),
		onItemRename = $bindable(() => Promise.resolve()),
		onItemDelete = $bindable(() => Promise.resolve())
	}: Props = $props();
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
<div
	onmouseleave={() => (menuVisible = false)}
	style="position: fixed; top: {menuY}px; left: {menuX}px;"
>
	<div class="size-70 absolute bottom-1/2 right-1/2 translate-[50%] -z-4 mx-auto"></div>

	<div class="bg-neutral-800 border border-neutral-700 shadow-xl rounded-lg p-2 z-50 w-80">
		<div class="flex gap-1 mb-2 items-center">
			<div
				class="bg-neutral-900/60 text-neutral-400 flex justify-center items-center rounded-lg size-10"
			>
				<Icon icon="mdi:file" width="20" height="20" />
			</div>
			{#if selectedItem}
				<div class="flex flex-col p-2">
					<p class="text-xs font-black truncate">{selectedItem?.name}</p>
					<p class="text-xs text-neutral-400 truncate">
						{new Date(selectedItem?.modified).toLocaleString('pl-PL', {
							day: '2-digit',
							month: '2-digit',
							year: '2-digit',
							hour: '2-digit',
							minute: '2-digit'
						})} | {storage_logic.formatBytes(selectedItem?.file_size)}
					</p>
				</div>
			{/if}
		</div>
		<button
			class="px-3 block w-full text-left p-2 hover:bg-neutral-700 cursor-pointer rounded-lg transition-colors"
			onclick={() => {
				if (selectedItem?.is_dir) {
					goto(`/dashboard/storage/${params.path ? params.path + '/' : ''}${selectedItem.name}`);
				} else {
					goto(`/dashboard/file/${params.path ? params.path + '/' : ''}${selectedItem.name}`);
				}
				menuVisible = false;
			}}
		>
			Open
		</button>
		<button
			onclick={async () => {
				if (!selectedItem) return;

				onItemRename(selectedItem.name, prompt('Podaj nową nazwę pliku/folderu:') || '');
			}}
			class="px-3 block w-full text-left p-2 hover:bg-neutral-700 cursor-pointer rounded-lg transition-colors"
			>Rename file</button
		>

		<button
			onclick={() => {
				const copyContent = `${$base_url}/interface/bucket/${params.path}/${selectedItem?.name}`;
				navigator.clipboard.writeText(copyContent);
				toast.info(copyContent);
				menuVisible = false;
			}}
			class="px-3 block w-full text-left p-2 hover:bg-neutral-700 cursor-pointer rounded-lg transition-colors"
			>Kopiuj link</button
		>
		<button
			onclick={async () => {
				if (!selectedItem) return;

				onItemDelete(selectedItem.name, selectedItem.is_dir);
			}}
			class="px-3 block w-full text-left p-2 hover:bg-neutral-700 cursor-pointer rounded-lg transition-colors"
			>Usuń</button
		>
	</div>
</div>
