<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<script lang="ts">
	import { developerView } from '$lib/dashboard/stores/persist';
	import type { Contributor } from '../../types';

	type Props = {
		onclick?: () => void;
		opened?: boolean;
		person: Contributor;
	};

	let loading: boolean = $state(false);

	let { onclick, opened = $bindable(true), person }: Props = $props();
</script>

<td class="flex max-w-[400px] items-center p-4">
	<div class="relative flex max-w-[400px] flex-wrap gap-1">
		<div
			onclick={() => {
				onclick?.();
			}}
			class={`absolute h-full w-full cursor-pointer p-3 ${$developerView ? 'bg-red-500/50' : 'rounded-xl '}`}
		></div>

		{#if person.permissions && person.permissions.length > 1}
			{#each person.permissions.split(',').slice(0, 7) as permission}
				<span
					class="rounded-full border border-blue-500/30 bg-blue-500/10 px-2 py-0.5 text-[10px] text-blue-400"
				>
					{permission}
				</span>
			{/each}
		{:else}
			<span
				class="rounded-full border border-red-500/30 bg-red-500/10 px-2 py-0.5 text-[10px] text-red-400"
			>
				Brak uprawnień!
			</span>
		{/if}
		{#if person.permissions.split(',').slice(0, 6).length >= 6}
			<span
				class="rounded-full border border-neutral-500/30 bg-neutral-500/10 px-2 py-0.5 text-[10px] text-neutral-400"
			>
				...
			</span>
		{/if}
	</div>
</td>
