<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<script lang="ts">
	import { developerView } from '$lib/dashboard/stores/persist';
	import { toast } from '$lib/dashboard/stores/toast';
	import type { Contributor, OptionType } from '../../types';
	import RowSettings from '../RowSettings.svelte';

	type Props = {
		onclick?: () => void;
		opened?: boolean;
		person: Contributor;
		options: OptionType[];
	};

	let loading: boolean = $state(false);

	let { onclick, opened = $bindable(true), person, options }: Props = $props();
</script>

<td class="relative p-4 text-left">
	<RowSettings
		onclick={() => {
			onclick?.();
		}}
		{options}
	/>
</td>
