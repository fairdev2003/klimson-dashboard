<script lang="ts">
	import { api } from '$lib/api/api';
	import { onMount } from 'svelte';
	import { type PermissionRegistry } from '$lib/api/requests/misc';
	import { debug } from '$lib/dashboard/stores/debug';
	import PermissionToggle from '../PermissionToggle.svelte';

	let registry: Partial<Record<string, PermissionRegistry[]>> | undefined = $state();

	onMount(async () => {
		const response = await api.misc.GetPermissionRegistry();

		registry = Object.groupBy(response.data.perms, (perm) => perm.category);
		debug.log(registry);
	});
</script>

<div class="flex flex-col gap-4">
	{#each Object.entries(registry ?? {}).sort() as [category, perms]}
		<PermissionToggle {perms} {category} />
	{/each}
</div>
