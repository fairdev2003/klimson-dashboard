<script lang="ts">
	import Heading from '../../typography/Heading.svelte';
	import Checkbox from '../Checkbox.svelte';

	type Props = {
		checked: boolean;
		description: string;
		title: string;
	};

	let { checked = $bindable(), title, description }: Props = $props();
</script>

<div class="lg:py-10 flex flex-col lg:flex-row lg:items-center lg:gap-0 gap-2 justify-between">
	<div class="flex flex-col lg:w-150">
		<Heading>{title}</Heading>
		<p class="font-medium text-sm text-neutral-300">{description}</p>
	</div>
	<div class="">
		<Checkbox bind:checked />
	</div>
</div>
