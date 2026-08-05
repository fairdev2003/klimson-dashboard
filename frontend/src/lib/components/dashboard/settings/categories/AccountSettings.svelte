<script lang="ts">
	import { blur } from 'svelte/transition';
	import { nickname } from '../store.svelte';
	import InputSettingsRecord from '../records/InputSettingsRecord.svelte';
</script>

<div
	class="flex lg:gap-0 gap-5 pt-5 flex-col border-t border-neutral-700 h-full lg:px-10"
	in:blur={{ duration: 300 }}
>
	<InputSettingsRecord
		bind:value={$nickname}
		title="Edit nickname"
		placeholder="nickname"
		description=""
	/>
</div>
