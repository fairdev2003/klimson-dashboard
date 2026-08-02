<script lang="ts">
	import type { PermissionRegistry } from '$lib/api/requests/misc';
	import { onMount } from 'svelte';
	import account_controller from '../helpers/access.svelte';

	type Props = {
		checked: boolean;
		perm: PermissionRegistry;
		onclick?: (e?: PermissionRegistry) => void;
	};

	let { checked = $bindable(), onclick, perm }: Props = $props();

	onMount(() => {
		checked = account_controller.roleHasPermission(perm);
	});
</script>

<button
	onclick={() => {
		checked = !checked;
		if (checked) {
			account_controller.ImplementPermToRole(perm);
		} else {
			account_controller.RemovePermFromRole(perm);
		}
		if (onclick) onclick(perm);
	}}
	class="transition-colors duration-300 cursor-pointer
           flex items-center rounded-full w-12 h-6 border-1 {checked
		? 'bg-blue-500/60 border-transparent'
		: 'bg-transparent border-neutral-700'} p-0.5 relative"
>
	<div
		class="h-full rounded-full w-4.5 flex items-center justify-center
               transition-all duration-300 ease-in-out shadow-sm
               {checked ? 'translate-x-6 bg-white' : 'translate-x-0 bg-white'}"
	></div>
</button>
