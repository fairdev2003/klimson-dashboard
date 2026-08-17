<script lang="ts">
	type Props = {
		onFileSelected: (e: Event) => void;
		uploading: boolean;
		statusMessage: string;
		progress: number;
		handleFileUpload: () => void;
		file: File | null;
	};

	let {
		onFileSelected,
		uploading = $bindable(),
		statusMessage = $bindable(),
		progress = $bindable(),
		handleFileUpload,
		file = $bindable()
	}: Props = $props();
</script>

<div
	class="flex flex-col gap-4 p-6 bg-neutral-900 text-white rounded-xl shadow-lg max-w-md mx-auto"
>
	<h2 class="text-lg font-bold">Wgrywanie binarki na serwer</h2>

	<input
		type="file"
		onchange={onFileSelected}
		disabled={uploading}
		class="block w-full text-sm text-neutral-400 file:mr-4 file:py-2 file:px-4 file:rounded-lg file:border-0 file:text-sm file:font-semibold file:bg-green-600 file:text-white hover:file:bg-green-500 cursor-pointer"
	/>

	{#if uploading}
		<div class="w-full bg-neutral-800 rounded-full h-2.5 overflow-hidden">
			<div class="bg-green-500 h-2.5 transition-all duration-300" style="width: {progress}%"></div>
		</div>
	{/if}

	<p class="text-xs text-neutral-300">{statusMessage}</p>

	<button
		onclick={handleFileUpload}
		disabled={!file || uploading}
		class="py-2 px-4 bg-green-600 hover:bg-green-500 disabled:bg-neutral-700 disabled:cursor-not-allowed font-semibold rounded-lg transition-colors"
	>
		{uploading ? 'Wysyłanie...' : 'Wyślij i zbuduj'}
	</button>
</div>
