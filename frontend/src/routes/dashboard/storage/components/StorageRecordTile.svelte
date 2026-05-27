<script lang="ts">
	import { goto } from '$app/navigation';
	import { api } from '$lib/api/api';
	import MovingTooltip from '$lib/components/dashboard/MovingTooltip.svelte';
	import Icon from '@iconify/svelte';

	type Props = {
		onclick: () => void;
		name: string;
		is_dir: boolean;
		slug: string;
		onrightclick: (e: MouseEvent) => void; // Nowy prop
	};

	let { onclick, name, is_dir, slug, onrightclick }: Props = $props();
</script>

<button
	onclick={() => {
		onclick();
		if (!is_dir) {
			goto(`/dashboard/file/${slug}/${name}`);
		}
	}}
	oncontextmenu={(e) => {
		e.preventDefault();
		onrightclick(e);
	}}
	class="group flex flex-col items-center gap-3 p-4 rounded-md
                       bg-neutral-950 border border-neutral-700 hover:border-neutral-400
                       hover:bg-neutral-800 shadow-sm hover:shadow-md"
>
	<div class="text-neutral-400 group-hover:text-white w-16 h-16 flex items-center justify-center">
		{#if !is_dir && (name.endsWith('.png') || name.endsWith('.jpg') || name.endsWith('.jpeg') || name.endsWith('.svg') || name.endsWith('.webp'))}
			<img
				src="{api.api_config.baseURL}interface/bucket/{slug}/{name}"
				alt={name}
				class="w-full h-full object-cover rounded-md border border-neutral-700"
			/>
		{:else if !is_dir && name.endsWith('.pdf')}
			<Icon icon="mingcute:pdf-fill" width="64" height="64" />
		{:else if !is_dir && name.endsWith('.sfm')}
			<Icon icon="mage:compact-disk-fill" width="64" height="64" />
		{:else if !is_dir && name.endsWith('.gif')}
			<Icon icon="fluent:gif-16-filled" width="64" height="64" />
		{:else if !is_dir && name.endsWith('.mp3')}
			<Icon icon="rivet-icons:audio-solid" width="64" height="64" />
		{:else if is_dir}
			<Icon icon="material-symbols:folder" width="64" height="64" />
		{:else}
			<Icon icon="material-symbols:description" width="64" height="64" />
		{/if}
	</div>

	<p
		class="text-xs font-medium text-neutral-200 text-center w-full truncate px-1 group-hover:text-white"
	>
		{name}
	</p>
</button>
