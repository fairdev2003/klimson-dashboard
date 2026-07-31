<script lang="ts">
	import { blur } from 'svelte/transition';
	import DropdownSettingsRecord from '../records/DropdownSettingsRecord.svelte';
	import { dev } from '$app/environment';
	import { debug } from '$lib/dashboard/stores/debug';
	import { base_url } from '$lib/api/api.store';
	import Heading from '../../typography/Heading.svelte';

	let selectedServer: 'dev' | 'prod' = $state('prod');
</script>

<div
	class="flex lg:gap-0 gap-5 pt-5 flex-col border-t border-neutral-700 lg:w-3/4 lg:px-10"
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
</div>
