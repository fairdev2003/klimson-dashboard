<script lang="ts">
	import { goto } from '$app/navigation';
	import { api } from '$lib/api/api';
	import MovingTooltip from '$lib/components/dashboard/MovingTooltip.svelte';
	import Icon from '@iconify/svelte';
	import { storage_logic } from '$lib/dashboard/storage/storage.svelte';
	import HarcCheckBox from '$lib/components/dashboard/HarcCheckBox.svelte';
	import { base_url } from '$lib/api/api.store';

	type Props = {
		onclick: (e: MouseEvent) => void;
		name: string;
		is_dir: boolean;
		slug: string;
		onrightclick: (e: MouseEvent) => void;
	};

	type Stringu = string;

	let deleteOn: boolean = $state(false);

	let { onclick, name, is_dir, slug, onrightclick }: Props = $props();
</script>

<button
	onclick={(e) => {
		if (name.endsWith('.md') || name.endsWith('.txt') || name.endsWith('.json')) {
			goto(`/dashboard/content_manager/markdown/editor/${slug}/${name}`);
			return;
		}

		if (storage_logic.delete_multiple_enabled) {
			deleteOn = !deleteOn;
			return;
		}

		onclick(e);

		if (!is_dir && !storage_logic.delete_multiple_enabled) {
			goto(`/dashboard/file/${slug}/${name}`);
		}
	}}
	oncontextmenu={(e) => {
		e.preventDefault();
		onrightclick(e);
	}}
	class:delete_enabled={storage_logic.delete_multiple_enabled}
	class:edit_mode_mobile={storage_logic.edit_enabled}
	class:delete_checked={deleteOn}
	class:tile_normal={!deleteOn}
	class="group relative flex flex-col items-center gap-3 p-4 rounded-md
                        shadow-sm hover:shadow-md"
>
	<div class="text-neutral-400 group-hover:text-white w-16 h-16 flex items-center justify-center">
		{@render FolderIcon(name)}
	</div>

	<div
		class="absolute hidden delete_checkbox z-7"
		class:delete_checkbox_visible={storage_logic.delete_multiple_enabled}
	>
		<HarcCheckBox bind:checked={deleteOn} />
	</div>

	<p
		class="text-xs font-medium text-neutral-200 text-center w-full truncate px-1 group-hover:text-white"
	>
		{name}
	</p>
</button>

{#snippet FolderIcon(name: string)}
	{#if !is_dir && (name.endsWith('.png') || name.endsWith('.jpg') || name.endsWith('.jpeg') || name.endsWith('.svg') || name.endsWith('.webp'))}
		<img
			src="{$base_url}/interface/bucket/{slug}/{name}"
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
{/snippet}

<style>
	@import 'tailwindcss';

	.delete_enabled {
		@apply hover:bg-neutral-950 hover:border-neutral-700 cursor-pointer;
	}

	.delete_checked {
		@apply bg-red-500/50 border-red-500 text-neutral-200 hover:text-neutral-200 hover:border-red-500 hover:bg-red-500/50 border;
	}

	.delete_checkbox_visible {
		@apply flex;
	}

	.tile_normal {
		@apply bg-neutral-900 transition-colors hover:bg-neutral-800;
	}

	.delete_checkbox {
		@apply top-2 right-1;
	}

	.edit_mode_mobile {
		@apply hover:bg-neutral-950;
	}
</style>
