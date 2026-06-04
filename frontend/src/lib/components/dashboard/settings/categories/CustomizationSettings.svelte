<script lang="ts">
	import { blur } from 'svelte/transition';
	import DropdownSettingsRecord from '../records/DropdownSettingsRecord.svelte';
	import { debug } from '$lib/dashboard/stores/debug';
	import { animation_preset } from '$lib/dashboard/stores/persist';
	import ButtonSettingsRecord from '../records/ButtonSettingsRecord.svelte';
	import { toast } from '$lib/dashboard/stores/toast';
	import Modal from '$lib/components/Modal.svelte';
	import InputSettingsRecord from '../records/InputSettingsRecord.svelte';
	import { nickname } from '../store';
	let value = $state('');

	let orderOpened: boolean = $state(false);
</script>

<div
	class="flex flex-col border-t border-neutral-700 h-full lg:w-3/4 lg:px-10"
	in:blur={{ duration: 300 }}
>
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
	<ButtonSettingsRecord
		onclick={() => {
			orderOpened = !orderOpened;
		}}
		label="Edit"
		title="Sidebar order"
		description="Edit order of sidebar contents. This is saved on frontend client."
	/>
	<InputSettingsRecord
		bind:value={$nickname}
		title="Edit nickname"
		placeholder="nickname"
		description=""
	/>
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
