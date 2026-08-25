<script lang="ts">
	import { blur } from 'svelte/transition';
	import DropdownSettingsRecord from '../records/DropdownSettingsRecord.svelte';
	import { debug } from '$lib/dashboard/stores/debug';
	import { base_url } from '$lib/api/api.store';

	import BinarySendSettingsRecord from '../records/BinarySendSettingsRecord.svelte';
	import { uploader } from '$lib/dashboard/file_upload.svelte';
</script>

<div
	class="flex lg:gap-0 gap-5 pt-5 flex-col border-t border-border lg:px-10"
	in:blur={{ duration: 300 }}
>
	<DropdownSettingsRecord
		error_text="Disabled due production frontend"
		title="Connecting environment"
		disabled
		description="Choose the server you want to connect in! It will refresh your page!"
		options={[
			{ key: 'Production', value: 'https://api.klimson.dev' },
			{ key: 'Development', value: 'http://localhost:8090' }
		]}
		bind:current_value={$base_url}
		onchoose={(e) => {
			debug.log(e.key);
			window.location.reload();
		}}
	/>

	<BinarySendSettingsRecord
		title="Send Binary File"
		description="Send binary File description"
		{uploader}
	/>
</div>
