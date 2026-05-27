<script lang="ts">
	import HarcCheckBox from '$lib/components/dashboard/HarcCheckBox.svelte';
	import { debugOn, developerView, sidebar_open } from '$lib/dashboard/stores/persist';
</script>

<div class="flex flex-col gap-5 m-10">
	<HarcCheckBox bind:checked={$debugOn} label="debugOn" />
	<HarcCheckBox bind:checked={$sidebar_open} label="sidebar_open" />
</div>
