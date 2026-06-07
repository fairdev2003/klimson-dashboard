<script lang="ts">
	import FileSystemModal from '$lib/components/dashboard/storage/FileSystemModal.svelte';
	import type { LayoutProps } from './$types';

	let { data, children }: LayoutProps = $props();
</script>

<div>
	<FileSystemModal />
	{@render children()}
</div>
