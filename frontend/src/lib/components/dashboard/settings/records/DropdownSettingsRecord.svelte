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
		disabled?: boolean;
		error_text?: string;
	};

	let {
		options,
		onchoose,
		current_value = $bindable(),
		label,
		title,
		description,
		error_text = '',
		disabled = false
	}: Props = $props();
</script>

<div
	class:disabled
	class="py-10 relative flex flex-col lg:flex-row lg:items-center lg:gap-0 gap-2 justify-between"
>
	{#if disabled}
		<div class="absolute w-full h-full z-100"></div>
	{/if}
	<div class="flex flex-col lg:w-150">
		<Heading>{title}</Heading>
		<p class="font-medium text-sm text-neutral-300">{description}</p>
		<p class="text-red-500 text-xs">{error_text && disabled ? error_text : ''}</p>
	</div>
	<div class="">
		<DropdownButton bind:current_value {label} {options} {onchoose} {error_text} />
	</div>
</div>

<style>
	@import 'tailwindcss';

	.disabled {
		@apply opacity-40 select-none;
	}
</style>
