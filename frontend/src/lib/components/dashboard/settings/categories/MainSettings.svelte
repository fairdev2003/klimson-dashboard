<script lang="ts">
	import { debugOn, sidebar_open } from '$lib/dashboard/stores/persist';
	import Heading from '../../typography/Heading.svelte';
	import Checkbox from '../Checkbox.svelte';
	import { blur } from 'svelte/transition';
	import CheckboxSettingsRecord from '../records/CheckboxSettingsRecord.svelte';
</script>

<div
	class="flex flex-col border-t border-neutral-700 lg:w-3/4 lg:px-10"
	in:blur={{ duration: 300 }}
>
	<CheckboxSettingsRecord
		bind:checked={$debugOn}
		title="Developer Console"
		description="Developers console will apear and will display important logs such as: data from server and
				client logs. Developer console can be moved in any direction on the dashboard."
	/>
	<CheckboxSettingsRecord
		bind:checked={$debugOn}
		title="Dev Mode"
		description="Dev mode allows user to view not visible viewports. Can break normal user experience"
	/>
	<CheckboxSettingsRecord
		bind:checked={$sidebar_open}
		title="Sidebar Force Open"
		description="Force the sidebar to open on mobile devices. Can be broken due to errors now!"
	/>
</div>
