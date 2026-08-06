<script lang="ts">
	import { base_url } from '$lib/api/api.store';
	import Icon from '@iconify/svelte';

	type Props = {
		name: string;
		is_dir: boolean;
		slug: string;
	};

	let { name, is_dir, slug }: Props = $props();
</script>

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
