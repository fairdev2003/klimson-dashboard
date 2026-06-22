<script lang="ts">
	import { clan_info } from '$lib/components/dashboard/table/pg3d/clan.store';
	import BlackoutWidget from '../BlackoutWidget.svelte';
	import LatestFilesWidget from '../LatestFilesWidget.svelte';
</script>

<div class="lg:grid flex flex-col lg:grid-cols-10 w-full gap-5">
	<LatestFilesWidget />
	<div class="flex flex-col h-80 gap-5 col-span-3">
		<BlackoutWidget />
		<div class="border border-neutral-700 bg-neutral-800 rounded-lg h-1/2">s</div>
	</div>
</div>
