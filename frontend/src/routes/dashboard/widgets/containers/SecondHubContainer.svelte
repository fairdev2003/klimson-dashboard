<script lang="ts">
	import { clan_info } from '$lib/components/dashboard/table/pg3d/clan.store';
	import BlackoutWidget from '../BlackoutWidget.svelte';
	import LatestFilesWidget from '../LatestFilesWidget.svelte';
</script>

<div class="lg:grid flex flex-col lg:grid-cols-10 w-full gap-5">
	<LatestFilesWidget />
</div>
