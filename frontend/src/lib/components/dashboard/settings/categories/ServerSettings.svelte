<script lang="ts">
	import { blur } from 'svelte/transition';
	import DropdownSettingsRecord from '../records/DropdownSettingsRecord.svelte';
	import { dev } from '$app/environment';
	import { debug } from '$lib/dashboard/stores/debug';
	import { base_url } from '$lib/api/api.store';
	import Heading from '../../typography/Heading.svelte';
	import ButtonSettingsRecord from '../records/ButtonSettingsRecord.svelte';
	import BinarySendButton from '../components/BinarySendButton.svelte';
	import { api } from '$lib/api/api';

	let selectedServer: 'dev' | 'prod' = $state('prod');

	let file = $state<File | null>(null);
	let uploading = $state(false);
	let progress = $state(0);
	let statusMessage = $state('');

	async function handleFileUpload() {
		if (!file) {
			statusMessage = 'Wybierz plik binarny!';
			return;
		}

		uploading = true;
		progress = 0;
		statusMessage = 'Rozpoczynanie wysyłania...';

		const CHUNK_SIZE = 1024 * 1024;
		const totalChunks = Math.ceil(file.size / CHUNK_SIZE);
		const fileName = file.name;

		try {
			for (let i = 0; i < totalChunks; i++) {
				const chunkIndex = i + 1;
				const start = i * CHUNK_SIZE;
				const end = Math.min(start + CHUNK_SIZE, file.size);
				const chunk = file.slice(start, end);

				await api.api.post('/admin/dev/send', chunk, {
					params: {
						filename: fileName
					},
					headers: {
						'Content-Type': 'application/octet-stream',
						'X-Chunk-Index': chunkIndex.toString(),
						'X-Total-Chunks': totalChunks.toString()
					}
				});

				progress = Math.round((chunkIndex / totalChunks) * 100);
				statusMessage = `Wysłano chunk ${chunkIndex} z ${totalChunks} (${progress}%)`;
			}

			statusMessage = 'Sukces! Binarka została wysłana, zbudowana i zrestartowana.';
		} catch (error: any) {
			const errorMessage = error.response?.data?.error || error.message;
			statusMessage = `Błąd: ${errorMessage}`;
		} finally {
			uploading = false;
		}
	}

	function onFileSelected(event: Event) {
		const target = event.target as HTMLInputElement;
		if (target.files && target.files[0]) {
			file = target.files[0];
		}
	}
</script>

<div
	class="flex lg:gap-0 gap-5 pt-5 flex-col border-t border-border lg:px-10"
	in:blur={{ duration: 300 }}
>
	<DropdownSettingsRecord
		error_text="Disabled due production frontend"
		title="Connecting environment"
		disabled
		description="Choose the server you want to connect in! It will refresh your page!"
		options={[
			{ key: 'Production', value: 'https://api.klimson.dev' },
			{ key: 'Development', value: 'http://localhost:8090' }
		]}
		bind:current_value={$base_url}
		onchoose={(e) => {
			debug.log(e.key);
			window.location.reload();
		}}
	/>

	<BinarySendButton
		bind:file
		{handleFileUpload}
		{onFileSelected}
		bind:progress
		bind:statusMessage
		bind:uploading
	/>
</div>
