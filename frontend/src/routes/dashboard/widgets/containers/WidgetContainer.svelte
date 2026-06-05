<script lang="ts">
	import type { Widget } from '$lib/components/dashboard/types/main.types';
	import SpotifyStatus from '$lib/components/spotify/SpotifyStatus.svelte';
	import ClockWidget from '../ClockWidget.svelte';
	import CPUWidget from '../CPUWidget.svelte';
	import DatabaseWidget from '../DatabaseWidget.svelte';
	import DiskSpaceWidget from '../DiskSpaceWidget.svelte';
	import SpotifyWidget from '../SpotifyWidget.svelte';

	const widgets: Widget[] = [
		{ component: DiskSpaceWidget, description: '', name: 'Disk Space Widget', slug: 'disk' },
		{ component: CPUWidget, description: '', name: 'CPU Widget', slug: 'cpu' },
		{
			component: DatabaseWidget,
			description: '',
			name: 'Database Widget',
			slug: 'database'
		},
		{ component: SpotifyWidget, description: '', name: 'Account', slug: 'spotify' },
		{ component: ClockWidget, description: '', name: 'Clock', slug: 'clock' }
	];
</script>

<div class="flex gap-5 lg:flex-row flex-col flex-wrap">
	{#each widgets as widget}
		<widget.component />
	{/each}
</div>
