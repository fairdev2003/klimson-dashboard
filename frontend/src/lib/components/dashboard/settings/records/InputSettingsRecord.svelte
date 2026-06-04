<script lang="ts">
	import Heading from '../../typography/Heading.svelte';
	import Checkbox from '../Checkbox.svelte';
	import Button from '../components/Button.svelte';
	import Input from '../components/Input.svelte';

	type Props = {
		value: string;
		title: string;
		description: string;
		placeholder: string;
	};

	let { value = $bindable(), title, description, placeholder }: Props = $props();
</script>

<div class="lg:py-10 flex flex-col lg:flex-row lg:items-center lg:gap-0 gap-2 justify-between">
	<div class="flex flex-col lg:w-150">
		<Heading>{title}</Heading>
		<p class="font-medium text-sm text-neutral-300">{description}</p>
	</div>
	<div class="flex gap-4">
		<Input bind:value {placeholder} />
	</div>
</div>
