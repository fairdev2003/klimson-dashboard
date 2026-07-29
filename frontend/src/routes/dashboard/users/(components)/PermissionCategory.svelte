<script lang="ts">
	import { type PermissionRegistry } from '$lib/api/requests/misc';
	import Icon from '@iconify/svelte';
	import { slide } from 'svelte/transition';

	type Props = { perms?: PermissionRegistry[]; category?: string };

	let opened: boolean = $state(true);

	let { perms, category = 'CATEGORY' }: Props = $props();
</script>

<div class="flex flex-col gap-4">
	<!-- svelte-ignore a11y_click_events_have_key_events -->
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<div
		onclick={() => {
			opened = !opened;
		}}
		class="flex gap-2 items-center cursor-pointer select-none"
	>
		<div class="text-neutral-400 transition-transform duration-300 {opened ? 'rotate-180' : ''}">
			<Icon icon="mdi:chevron-down" width="25" height="25" />
		</div>

		<h3 class="text-white font-bold capitalize text-lg">{category}</h3>
	</div>

	{#if opened}
		<div transition:slide={{ duration: 300 }} class="flex flex-col gap-4">
			{#each perms ?? [] as permission}
				<div class="w-full h-20 flex items-center bg-neutral-800 gap-4 rounded-xl p-4">
					<div
						class="text-blue-400 bg-blue-500/30 size-10 flex items-center justify-center rounded-lg"
					>
						<Icon icon={permission.icon} width="20" height="20" />
					</div>

					<div class="flex flex-col">
						<p class="font-black text-white">{permission.name}</p>
						<p class="font-mono text-xs text-neutral-400">{permission.tag}</p>
					</div>
				</div>
			{/each}
		</div>
	{/if}
</div>
