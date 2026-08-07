<script lang="ts">
	import Heading from '../../typography/Heading.svelte';
	import Button from '../components/Button.svelte';

	type Props = {
		label: string;
		onclick?: () => void;
		title: string;
		description: string;
	};

	let { onclick, label, title, description }: Props = $props();
</script>

<div class="lg:py-10 flex flex-col lg:flex-row lg:items-center lg:gap-0 gap-2 justify-between">
	<div class="flex flex-col lg:w-150">
		<Heading>{title}</Heading>
		<p class="font-medium text-sm text-text">{description}</p>
	</div>
	<div class="">
		<Button {onclick} {label} />
	</div>
</div>
