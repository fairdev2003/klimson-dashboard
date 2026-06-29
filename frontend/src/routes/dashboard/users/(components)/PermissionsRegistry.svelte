<script lang="ts">
	import { api } from '$lib/api/api';
	import { onMount } from 'svelte';
	import { type PermissionRegistry } from '$lib/api/requests/misc';
	import { debug } from '$lib/dashboard/stores/debug';
	import Icon from '@iconify/svelte';

	let registry: PermissionRegistry[] = $state([]);

	onMount(async () => {
		const response = await api.misc.GetPermissionRegistry();

		registry = response.data;
		debug.log(registry);
	});
</script>

<div class="flex flex-col gap-5">
	<div class="flex flex-col mx-auto w-2xl gap-2">
		{#each registry as permission}
			<div class="w-full h-20 flex items-center bg-neutral-800 gap-4 rounded-xl p-4">
				<!-- Icon -->
				<div
					class="text-blue-400 bg-blue-500/30 size-10 flex items-center justify-center rounded-lg"
				>
					<Icon icon={permission.icon} width="20" height="20" />
				</div>

				<!-- Content -->
				<div class="flex flex-col">
					<p class="font-black">{permission.name}</p>
					<p class="font-mono text-xs text-neutral-400">{permission.tag}</p>
				</div>
			</div>
		{/each}
	</div>
</div>
