<script lang="ts">
	import { api } from '$lib/api/api';
	import { onMount } from 'svelte';
	import { type PermissionRegistry } from '$lib/api/requests/misc';
	import { debug } from '$lib/dashboard/stores/debug';
	import Icon from '@iconify/svelte';
	import PermissionCategory from './PermissionCategory.svelte';

	let registry: Partial<Record<string, PermissionRegistry[]>> | undefined = $state();

	onMount(async () => {
		const response = await api.misc.GetPermissionRegistry();

		registry = Object.groupBy(response.data.perms, (perm) => perm.category);
		debug.log(registry);
	});
</script>

<div class="flex flex-col gap-5">
	<div class="flex flex-col mx-auto w-2xl gap-4">
		{#each Object.entries(registry ?? {}) as [category, perms]}
			<PermissionCategory {perms} {category} />
		{/each}
	</div>
</div>
