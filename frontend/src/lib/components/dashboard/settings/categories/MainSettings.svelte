<script lang="ts">
	import { animation_preset, debugOn, sidebar_open } from '$lib/dashboard/stores/persist';
	import Heading from '../../typography/Heading.svelte';
	import Checkbox from '../Checkbox.svelte';
	import { blur } from 'svelte/transition';
	import CheckboxSettingsRecord from '../records/CheckboxSettingsRecord.svelte';
	import DropdownButton from '../components/DropdownButton.svelte';
	import { debug } from '$lib/dashboard/stores/debug';
	import ButtonSettingsRecord from '../records/ButtonSettingsRecord.svelte';
	import { toast } from '$lib/dashboard/stores/toast';
</script>

<div
	class="flex lg:gap-0 gap-5 pt-5 flex-col border-t border-neutral-700 h-full lg:w-3/4 lg:px-10"
	in:blur={{ duration: 300 }}
>
	<CheckboxSettingsRecord
		bind:checked={$debugOn}
		title="Developer Console"
		description="Developers console will apear and will display important logs such as: data from server and
				client logs. Developer console can be moved in any direction on the dashboard."
	/>

	<ButtonSettingsRecord
		title="Reset Dashboard config"
		description="It will dump all of your dashboard configuration"
		label="Process"
		onclick={() => {
			const p = confirm('Are you sure to process with this one?');
			if (p) {
				localStorage.setItem('dashboard_config', '{}');
				window.location.reload();
			} else {
				toast.info('Process is denied!');
			}
		}}
	/>
</div>
