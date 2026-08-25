<script lang="ts">
	import { uploader } from '$lib/dashboard/file_upload.svelte';
	import { toast } from '$lib/dashboard/stores/toast';
	import Icon from '@iconify/svelte';

	let opened: boolean = $state(true);
</script>

{#if uploader.uploading}
	<div class="fixed flex flex-col gap-4 bottom-5 right-5 z-201 w-80 p-4 bg-neutral-900 rounded-lg">
		<div class="flex justify-between items-center">
			<p class="text-neutral-400 uppercase text-xs font-black">Uploading photo</p>
			<button
				onclick={() => {
					opened = !opened;
				}}
				class="text-white p-0.5 rounded-md cursor-pointer hover:bg-white/10"
			>
				<Icon icon="material-symbols:close" />
			</button>
		</div>
		<div class="w-full h-1 rounded-full bg-green-700/40">
			<div style="width: {uploader.progress}%;" class="h-1 transition-all bg-green-500"></div>
		</div>
		<p class="text-neutral-400 text-xs">Uploading: smegmanpng.png</p>
	</div>
{/if}
