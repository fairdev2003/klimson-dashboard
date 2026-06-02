<script lang="ts">
	import { api } from '$lib/api/api';
	import SFMCodeArea from '$lib/components/dashboard/sfm/(components)/SFMCodeArea.svelte';
	import Icon from '@iconify/svelte';

	let { params } = $props();

	let path_table = params.path.split('/').filter(Boolean);
	let file_name = path_table[path_table.length - 1];

	function getFileUrl(path: string) {
		return path.split('/').map(encodeURIComponent).join('/');
	}
</script>

<div class="p-8 flex flex-col gap-6 w-full">
	<div class="flex flex-col gap-1 border-b border-neutral-700 pb-4">
		<button
			onclick={() => {
				history.back();
			}}
			class="flex gap-1 items-center text-blue-500 hover:underline"
		>
			<Icon icon="lets-icons:back" />
			<p>Back</p>
		</button>
		<h1 class="text-xl font-bold text-neutral-100">{file_name}</h1>
		<a
			href="{api.api_config.baseURL}interface/bucket/{getFileUrl(params.path)}"
			class="text-sm text-neutral-400 hover:underline cursor-pointer break-all"
		>
			API Reference URL
		</a>
	</div>

	{#if file_name.endsWith('.png') || file_name.endsWith('.jpg') || file_name.endsWith('.webp') || file_name.endsWith('.jpeg') || file_name.endsWith('.gif') || file_name.endsWith('.svg')}
		<div class="flex flex-col justify-center items-center mx-auto rounded-xl shadow-2xl">
			<img
				src="{api.api_config.baseURL}interface/bucket/{getFileUrl(params.path)}"
				alt={file_name}
				class="max-h-[70vh] w-auto rounded-lg object-contain"
			/>
		</div>
	{/if}
	{#if file_name.endsWith('.sfm')}
		<SFMCodeArea
			program_link="{api.api_config.baseURL}interface/bucket/{getFileUrl(params.path)}"
		/>
	{/if}
</div>

<style>
	@import 'tailwindcss';
	a {
		@apply text-blue-500;
	}
</style>
