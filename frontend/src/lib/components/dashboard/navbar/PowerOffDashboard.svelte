<script lang="ts">
	import { goto } from '$app/navigation';
	import RDBModal from '$lib/components/modal/RDBModal.svelte';
	import Icon from '@iconify/svelte';
	import { console_service } from '../dev/console/console_service.svelte';

	let powerOffModalOpened = $state(false);
</script>

<button
	onclick={() => {
		powerOffModalOpened = true;
	}}
	class="size-10 rounded-full flex justify-center hover:bg-neutral-700 transition-colors cursor-pointer items-center text-red-500 bg-neutral-800"
>
	<Icon icon="tdesign:poweroff" width="20" height="20" />
</button>

<RDBModal
	bind:opened={powerOffModalOpened}
	title="Session going to be terminated"
	titleStyle="danger"
	border="borderless"
	size="accept_preset"
	form_config={{
		onAccept: () => {
			console_service.run('logout');
		},
		onCancel: () => {
			powerOffModalOpened = false;
		}
	}}
>
	<p>Do you want to log out?</p>
</RDBModal>
