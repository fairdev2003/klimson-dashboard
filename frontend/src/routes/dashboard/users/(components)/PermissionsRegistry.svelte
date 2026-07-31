<script lang="ts">
	import { api } from '$lib/api/api';
	import { onMount } from 'svelte';
	import { type PermissionRegistry } from '$lib/api/requests/misc';
	import { debug } from '$lib/dashboard/stores/debug';
	import Icon from '@iconify/svelte';
	import PermissionCategory from './PermissionCategory.svelte';
	import { blur } from 'svelte/transition';

	let registry: Partial<Record<string, PermissionRegistry[]>> | undefined = $state();

	onMount(async () => {
		const response = await api.misc.GetPermissionRegistry();

		registry = Object.groupBy(response.data.perms, (perm) => perm.category);
		debug.log(registry);
	});
</script>

<div class="flex flex-col gap-5" in:blur={{ duration: 150 }}>
	<div class="flex flex-col mx-auto lg:w-2xl md:w-9/10 w-full gap-4">
		{#each Object.entries(registry ?? {}).sort() as [category, perms]}
			<PermissionCategory {perms} {category} />
		{/each}
		<!-- mdi:attachment-lock -->
	</div>
</div>
