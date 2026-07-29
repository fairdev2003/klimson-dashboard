<script lang="ts">
	import RDBInput from '$lib/components/modal/RDBInput.svelte';
	import RDBModal from '$lib/components/modal/RDBModal.svelte';
	import { onMount } from 'svelte';
	import RoleImplementModal from './RoleImplementModal.svelte';

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

<div></div>

<RoleImplementModal bind:implementRoleModalOpened />

<style>
	@import 'tailwindcss';

	.selector {
		@apply;
	}
</style>
