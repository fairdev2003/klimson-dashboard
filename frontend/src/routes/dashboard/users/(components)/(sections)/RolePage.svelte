<script lang="ts">
	import RDBInput from '$lib/components/modal/RDBInput.svelte';
	import RDBModal from '$lib/components/modal/RDBModal.svelte';
	import { onMount } from 'svelte';
	import RoleImplementModal from './RoleImplementModal.svelte';
	import RoleRecord from '../RoleRecord.svelte';

	type Props = {
		implementRoleModalOpened?: boolean;
		roleUpdateModalState?: boolean;
	};

	onMount(() => {
		// fetch role data here
	});

	let {
		implementRoleModalOpened = $bindable(false),
		roleUpdateModalState = $bindable(false)
	}: Props = $props();
</script>

<div class="flex flex-col gap-2">
	<RoleRecord role={{ name: 'Admin', permissions: [{ name: 'siema' }] }} />
	<RoleRecord role={{ name: 'Moderator', permissions: [{ name: 'siema' }] }} />
</div>
<RoleImplementModal bind:implementRoleModalOpened />

<style>
	@import 'tailwindcss';

	.selector {
		@apply;
	}
</style>
