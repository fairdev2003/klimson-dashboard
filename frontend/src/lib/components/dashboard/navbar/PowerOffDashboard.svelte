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
	class="size-10 rounded-full text-text flex justify-center transition-colors cursor-pointer items-center text-danger bg-foreground"
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
		onLogout: () => {
			console_service.run('logout');
		},
		onLockScreen: () => {
			console_service.run('lockscreen');
		}
	}}
>
	<p>Do you want to log out?</p>
</RDBModal>
