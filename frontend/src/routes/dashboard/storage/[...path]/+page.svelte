<script lang="ts">
	import Icon from '@iconify/svelte';
	import type { PageData } from './$types';
	import { page } from '$app/state';
	import { goto, invalidateAll } from '$app/navigation';
	import StorageRecordTile from '../components/StorageRecordTile.svelte';
	import { type StorageRecord } from '$lib/api/requests/storage';
	import Button from '$lib/components/Button.svelte';
	import Modal from '$lib/components/Modal.svelte';
	import DatabaseModalInput from '$lib/components/dashboard/table/DatabaseModalInput.svelte';
	import { api } from '$lib/api/api';
	import { toast } from '$lib/dashboard/stores/toast';
	import { onMount } from 'svelte';

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

	let folderModalOpen = $state(false);
	let folderName = $state('');
	let fileInput: HTMLInputElement;
	let files: FileList | null = $state(null);
	let uploading = $state(false);

	function navigateToDir(pathParts: string[], clickedIndex: number): string {
		return `/dashboard/storage/${pathParts.slice(0, clickedIndex + 1).join('/')}`;
	}

	async function UploadFile(event: Event) {
		const target = event.target as HTMLInputElement;
		const file = target.files?.[0];

		if (!file) return;

		uploading = true;
		try {
			const response = await api.storage.UploadFile(params.path || '', file);

			if (response.data.success) {
				toast.success('Success!');
				await invalidateAll();
			}
		} catch (e) {
			console.error('Błąd uploadu:', e);
		} finally {
			uploading = false;
			target.value = '';
		}
	}

	async function UploadPastedFile(file: File) {
		uploading = true;
		try {
			const response = await api.storage.UploadFile(params.path || '', file);

			if (response.data.success) {
				toast.success('Plik wklejony pomyślnie!');
				await invalidateAll();
			} else {
				toast.error('Błąd: ' + response.data.message);
			}
		} catch (e) {
			console.error('Błąd wklejania pliku:', e);
			toast.error('Nie udało się wgrać wklejonego pliku.');
		} finally {
			uploading = false;
		}
	}

	async function CreateFolder() {
		try {
			const response = await api.storage.CreateFolder(params.path || '', {
				folder_name: folderName
			});

			if (response.data.success) {
				toast.success('Sukces: ' + response.data.message);
				folderName = '';
				folderModalOpen = false;
				await invalidateAll();
			} else {
				toast.error('Błąd: ' + response.data.message);
			}
		} catch (error: any) {
			console.error('Wystąpił błąd podczas tworzenia folderu:', error);

			const errorMessage = error.response?.data?.error || 'Nieznany błąd';
			toast.error('Nie udało się stworzyć folderu: ' + errorMessage);
		}
	}
	function HandlePaste(event: ClipboardEvent) {
		const items = event.clipboardData?.items;
		if (!items) return;

		const itemList = Array.from(items);

		for (const item of itemList) {
			if (item.kind === 'file') {
				const file = item.getAsFile();
				if (file) {
					event.preventDefault();
					UploadPastedFile(file);
				}
			}
		}
	}

	async function DeleteFileOrFolder(name: string, isDir: boolean) {
		if (!confirm(`Czy na pewno chcesz usunąć ${name}?`)) return;

		try {
			const pathToDelete = `${params.path || ''}/${name}`.replace(/\/+/g, '/');
			const res = await api.storage.DeleteItem(pathToDelete);

			if (res.data.success) {
				toast.success('Usunięto!');
				await invalidateAll();
			}
		} catch (e) {
			toast.error('Błąd podczas usuwania');
			console.error(e);
		}
	}
</script>

<div class="flex flex-col gap-4">
	{@render StorageHeader()}
	<div class="grid grid-cols-[repeat(auto-fill,minmax(160px,1fr))] gap-4 p-4">
		{#each data.storage_file_list as { file_size, is_dir, modified, name }}
			<StorageRecordTile
				onrightclick={(e) => handleRightClick(e, { is_dir, name, file_size, modified })}
				{is_dir}
				{name}
				slug={params.path}
				onclick={(e) => {
					if (is_dir) {
						const currentPath = page.url.pathname.replace(/\/$/, '');
						goto(`${currentPath}/${name}`);
					}
				}}
			/>
		{/each}
	</div>
</div>

{#if menuVisible && selectedItem}
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
		<button
			onclick={async () => {
				if (!selectedItem) return;

				await DeleteFileOrFolder(selectedItem.name, selectedItem.is_dir);
			}}
			class="block w-full text-left p-2 hover:bg-neutral-700 text-red-400">Usuń</button
		>
	</div>
{/if}

{#snippet StorageHeader()}
	<input type="file" class="hidden" bind:this={fileInput} onchange={UploadFile} />
	<div
		class="p-4 bg-neutral-900 border border-neutral-700 lg:flex-row flex flex-col gap-4 justify-between lg:items-center"
	>
		<div class="flex flex-col gap-2">
			<button
				onclick={() => {
					history.back();
				}}
				class="flex gap-1 items-center text-blue-500 hover:underline"
			>
				<Icon icon="lets-icons:back" />
				<p>Back</p>
			</button>
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
		</div>
		<div class="flex gap-3">
			<Button
				theme="base"
				disabled={uploading}
				onclick={() => {
					fileInput.click();
				}}>Upload a file</Button
			>
			<Button
				theme="secondary"
				onclick={() => {
					folderModalOpen = !folderModalOpen;
				}}>Create New Folder</Button
			>
		</div>
	</div>
{/snippet}

<Modal
	bind:opened={folderModalOpen}
	onClose={() => {
		folderModalOpen = false;
	}}
	title="Creting New Folder"
	className="w-100"
>
	<div class="flex flex-col gap-4">
		<DatabaseModalInput bind:value={folderName} label="Folder Name" />
		<div class="flex justify-end">
			<Button onclick={CreateFolder} size="small" theme="secondary">Create</Button>
		</div>
	</div>
</Modal>

<svelte:window onpaste={HandlePaste} />
