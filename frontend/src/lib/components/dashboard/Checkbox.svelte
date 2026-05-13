<script lang="ts">
	import type { Snippet } from 'svelte';

	type Props = {
		checked: boolean;
		label: string;
		delete_fragment: Snippet;
		edit_fragment: Snippet;
	};

	let { checked = $bindable(), label, delete_fragment, edit_fragment }: Props = $props();
</script>

<div class="flex gap-2">
	<input
		id="default-checkbox"
		type="checkbox"
		bind:checked
		class="border-default-medium bg-neutral-secondary-medium focus:ring-brand-soft h-4 w-4 rounded-xs border focus:ring-2"
	/>
	<label for="default-checkbox" class="text-heading ms-2 text-sm font-medium select-none"
		>{label}</label
	>
	{@render edit_fragment()}
	{@render delete_fragment()}
</div>
