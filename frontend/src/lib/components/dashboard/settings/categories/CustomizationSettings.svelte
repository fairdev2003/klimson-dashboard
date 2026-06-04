<script lang="ts">
	import { blur } from 'svelte/transition';
	import DropdownSettingsRecord from '../records/DropdownSettingsRecord.svelte';
	import { debug } from '$lib/dashboard/stores/debug';
	import { animation_preset } from '$lib/dashboard/stores/persist';
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
</div>
