<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<script lang="ts">
	import { developerView } from '$lib/dashboard/stores/persist';
	import Icon from '@iconify/svelte';
	import type { Contributor } from '../../types';
	import Tooltip from '$lib/components/dashboard/Tooltip.svelte';

	type Props = {
		onclick?: () => void;
		opened?: boolean;
		person: Contributor;
	};

	let loading: boolean = $state(false);

	let { onclick, opened = $bindable(true), person }: Props = $props();
</script>

<td class="relative p-4">
	<div class="relative flex items-center gap-4">
		<div
			onclick={() => {
				onclick?.();
			}}
			class={`absolute  h-full w-full cursor-pointer p-3 ${$developerView ? 'bg-red-500/50' : 'rounded-xl '}`}
		></div>
		{#if person.blocked_till}
			<Tooltip content="Ten kontrybutor jest zablokowany" position="top">
				<Icon icon="material-symbols:warning" class="text-orange-500/50" height="24" width="24"
				></Icon>
			</Tooltip>
		{/if}
		<div class="flex flex-col">
			<span class="font-medium text-neutral-200">{person.name}</span>
			<span class="text-xs text-neutral-500">{person.login}</span>
		</div>
	</div>
</td>
