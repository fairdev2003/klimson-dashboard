<script lang="ts">
	import RDBInput from '$lib/components/modal/RDBInput.svelte';
	import RDBModal from '$lib/components/modal/RDBModal.svelte';

	type Props = {
		implementRoleModalOpened?: boolean;
	};

	let { implementRoleModalOpened = $bindable(false) }: Props = $props();
</script>

<RDBModal
	bind:opened={implementRoleModalOpened}
	border="borderless"
	title="Adding new role"
	size="form_preset"
	form_config={{
		onLog: () => {}
	}}
>
	<RDBInput label="ROLE NAME" />
</RDBModal>

<style>
	@import 'tailwindcss';

	.selector {
		@apply;
	}
</style>
