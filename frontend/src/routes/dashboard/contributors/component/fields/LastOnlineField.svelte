<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<script lang="ts">
	import { developerView } from '$lib/dashboard/stores/persist';
	import type { Contributor } from '../../types';

	type Props = {
		// onclick?: () => void;
		opened?: boolean;
		// person: Contributor;
	};

	let loading: boolean = $state(false);

	let { opened = $bindable(true) }: Props = $props();
</script>

<td class="p-4 text-sm">
	<span
		class="rounded-md border border-green-700 bg-green-600/50 px-2 py-1 text-[10px] font-bold uppercase"
	>
		12 stycznia 2024, 14:23
	</span>
</td>
