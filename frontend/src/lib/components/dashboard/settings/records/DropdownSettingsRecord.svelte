<script lang="ts">
	import Heading from '../../typography/Heading.svelte';
	import DropdownButton from '../components/DropdownButton.svelte';

	type DropdownOption = {
		key: string;
		value: string;
	};

	type Props = {
		options: DropdownOption[];
		onchoose?: (e: DropdownOption) => void;
		current_value: any;
		label?: string;
		title: string;
		description: string;
	};

	let {
		options,
		onchoose,
		current_value = $bindable(),
		label,
		title,
		description
	}: Props = $props();
</script>

<div class="py-10 flex flex-col lg:flex-row lg:items-center lg:gap-0 gap-2 justify-between">
	<div class="flex flex-col lg:w-150">
		<Heading>{title}</Heading>
		<p class="font-medium text-sm text-neutral-300">{description}</p>
	</div>
	<div class="">
		<DropdownButton bind:current_value {label} {options} {onchoose} />
	</div>
</div>
