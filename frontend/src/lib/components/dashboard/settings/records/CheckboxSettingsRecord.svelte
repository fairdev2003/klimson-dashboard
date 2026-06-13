<script lang="ts">
	import Heading from '../../typography/Heading.svelte';
	import Checkbox from '../Checkbox.svelte';

	type Props = {
		checked: boolean;
		description: string;
		title: string;
		beta?: boolean;
	};

	let { checked = $bindable(), title, description, beta = false }: Props = $props();
</script>

<div class="lg:py-10 flex flex-col lg:flex-row lg:items-center lg:gap-0 gap-2 justify-between">
	<div class="flex flex-col lg:w-150">
		<div class="flex gap-3">
			<Heading>{title}</Heading>
			{#if beta}
				<div
					class="p-0.5 bg-blue-600 text-xs font-black items-center flex justify-center px-2 rounded-lg"
				>
					BETA
				</div>
			{/if}
		</div>
		<p class="font-medium text-sm text-neutral-300">{description}</p>
	</div>
	<div class="">
		<Checkbox bind:checked />
	</div>
</div>
