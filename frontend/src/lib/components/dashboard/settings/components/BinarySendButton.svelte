<script lang="ts">
	import type { Uploader } from '$lib/dashboard/file_upload.svelte';

	type Props = {
		uploader: Uploader;
	};

	let { uploader }: Props = $props();
</script>

<div class="flex flex-col gap-4 p-6 text-white rounded-xl max-w-md mx-auto">
	{#if !uploader.file}
		<input
			type="file"
			onchange={uploader.OnFileSelected}
			disabled={uploader.uploading}
			class="block w-full text-sm text-neutral-400 file:mr-4 file:py-2 file:px-4 file:rounded-lg file:border-0 file:text-sm file:font-semibold file:bg-green-600 file:text-white hover:file:bg-green-500 cursor-pointer"
		/>
	{/if}

	<!-- {#if uploading}
		<div class="w-full bg-neutral-800 rounded-full h-2.5 overflow-hidden">
			<div class="bg-green-500 h-2.5 transition-all duration-300" style="width: {progress}%"></div>
		</div>
	{/if} -->

	<p class="text-xs text-neutral-300">{uploader.statusMessage}</p>
	{#if uploader.file}
		<button
			onclick={uploader.HandleFileUpload}
			disabled={!uploader.file || uploader.uploading}
			class="py-2 px-4 bg-green-600 hover:bg-green-500 disabled:bg-neutral-700 disabled:cursor-not-allowed font-semibold rounded-lg transition-colors"
		>
			{uploader.uploading ? 'Sending...' : 'Send'}
		</button>
	{/if}
</div>
