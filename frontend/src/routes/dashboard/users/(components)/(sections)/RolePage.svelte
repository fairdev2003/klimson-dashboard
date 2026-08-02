<script lang="ts">
	import RDBInput from '$lib/components/modal/RDBInput.svelte';
	import RDBModal from '$lib/components/modal/RDBModal.svelte';
	import { onMount } from 'svelte';
	import RoleRecord from '../RoleRecord.svelte';
	import { blur } from 'svelte/transition';
	import RoleImplementModal from '../(modals)/RoleImplementModal.svelte';
	import account_controller from '../../helpers/access.svelte';

	type Props = {
		implementRoleModalOpened?: boolean;
		roleUpdateModalState?: boolean;
	};

	onMount(async () => {
		await account_controller.FetchRolesAndAssign();
	});

	let {
		implementRoleModalOpened = $bindable(false),
		roleUpdateModalState = $bindable(false)
	}: Props = $props();
</script>

<div class="flex flex-col gap-4 w-2xl mx-auto">
	{#each account_controller.roles as role}
		<RoleRecord {role} />
	{/each}
</div>

<RoleImplementModal bind:implementRoleModalOpened />

<style>
	@import 'tailwindcss';

	.selector {
		@apply;
	}
</style>
