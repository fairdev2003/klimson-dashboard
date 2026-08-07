<script lang="ts">
	import { blur } from 'svelte/transition';
	import DropdownSettingsRecord from '../records/DropdownSettingsRecord.svelte';
	import { debug } from '$lib/dashboard/stores/debug';
	import { animation_preset, dashboard_config } from '$lib/dashboard/stores/persist';
	import ButtonSettingsRecord from '../records/ButtonSettingsRecord.svelte';
	import { toast } from '$lib/dashboard/stores/toast';
	import Modal from '$lib/components/Modal.svelte';
	import InputSettingsRecord from '../records/InputSettingsRecord.svelte';
	import { nickname } from '../store.svelte';
	import { onMount, tick } from 'svelte';
	import { settings_startup_modal } from '../../stores/main';
	import { codeEditorTheme } from '$lib/components/markdown/markdown';
	import MultipleSettingsRecord from '../records/MultipleSettingsRecord.svelte';
	import SampleDragNDrop from './SampleDragNDrop.svelte';
	import CheckboxSettingsRecord from '../records/CheckboxSettingsRecord.svelte';
	let value = $state('');

	let orderOpened: boolean = $state(false);
	let widgetOpened: boolean = $state(false);

	let list = $state(['apple', 'pineapple']);
</script>

<div
	class="flex lg:gap-0 gap-5 pt-5 flex-col border-t border-foreground h-full lg:px-10"
	in:blur={{ duration: 300 }}
>
	<DropdownSettingsRecord
		title="Dashboard Theme"
		description="Choose Theme of your code editor built in dashboard!"
		options={[
			{ key: 'Dark', value: 'dark' },
			{ key: 'Light', value: 'light' },
			{ key: 'System', value: 'system' },
			{ key: 'Candy', value: 'candy' }
		]}
		bind:current_value={$dashboard_config.theme}
		onchoose={(e) => {
			debug.log(e.key);
		}}
	/>
	<DropdownSettingsRecord
		title="Modal Animation"
		description="Choose modal animation that will show in!"
		options={[
			{ key: 'Blur Animation', value: 'blur' },
			{ key: 'Jason Animation', value: 'jason' },
			{ key: 'Klimson Animation', value: 'klimson' }
		]}
		bind:current_value={$animation_preset}
		onchoose={(e) => {
			debug.log(e.key);
		}}
	/>

	<DropdownSettingsRecord
		title="Code Editor Theme"
		description="Choose Theme of your code editor built in dashboard!"
		options={[
			{ key: 'Classic', value: 'classic' },
			{ key: 'Dracula Theme', value: 'dracula' },
			{ key: 'Light (type any)', value: 'normal' }
		]}
		bind:current_value={$dashboard_config.code_theme}
		onchoose={(e) => {
			debug.log(e.key);
		}}
	/>
	<CheckboxSettingsRecord
		beta
		bind:checked={$dashboard_config.dock_on}
		title="Dashboard Dock"
		description="Enables dashboard dock"
	/>

	<MultipleSettingsRecord
		title="Choose sidebar pills"
		description="Choose what pill will be displayed on the top-level of the sidebar"
		options={[
			{ key: 'User Pill', value: 'profile' },
			{ key: 'Storage Pill', value: 'storage' }
		]}
		bind:current_value={$dashboard_config.client_pills}
		onchoose={(e) => {
			debug.log(e.key);
		}}
	/>

	<ButtonSettingsRecord
		onclick={() => {
			orderOpened = !orderOpened;
		}}
		label="Edit"
		title="Sidebar order"
		description="Edit order of sidebar contents. This is saved on frontend client."
	/>
	<ButtonSettingsRecord
		onclick={() => {
			widgetOpened = !widgetOpened;
		}}
		label="Edit"
		title="System Widgets"
		description="Decide what widgets will show up on dashboard home page."
	/>

	<Modal
		exitMode={false}
		title="Edit system widgets"
		bind:opened={widgetOpened}
		onClose={() => {
			widgetOpened = !widgetOpened;
		}}
		className="w-100 h-130 overflow-hidden"
	>
		<SampleDragNDrop />
	</Modal>

	<Modal
		exitMode={false}
		title="Edit sidebar order"
		bind:opened={orderOpened}
		onClose={() => {
			orderOpened = !orderOpened;
		}}
		className="w-100 overflow-hidden"
	>
		<div class="h-100"></div>
	</Modal>
</div>
