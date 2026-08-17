<script lang="ts">
	import Heading from '../../typography/Heading.svelte';
	import Checkbox from '../Checkbox.svelte';
	import BinarySendButton from '../components/BinarySendButton.svelte';
	import Button from '../components/Button.svelte';
	import Input from '../components/Input.svelte';

	type Props = {
		title: string;
		description: string;
		onFileSelected: (e: Event) => void;
		uploading: boolean;
		statusMessage: string;
		progress: number;
		handleFileUpload: () => void;
		file: File | null;
	};

	let {
		title,
		description,
		onFileSelected,
		uploading = $bindable(),
		statusMessage = $bindable(),
		progress = $bindable(),
		handleFileUpload,
		file = $bindable()
	}: Props = $props();
</script>

<div class="flex flex-col gap-4">
	<div class="flex flex-col lg:flex-row lg:items-center lg:gap-0 gap-2 justify-between">
		<div class="flex flex-col lg:w-150">
			<Heading>{title}</Heading>
			<p class="font-medium text-sm text-neutral-300">{description}</p>
		</div>
		<div class="flex gap-4">
			<BinarySendButton
				bind:file
				{handleFileUpload}
				{onFileSelected}
				bind:progress
				bind:statusMessage
				bind:uploading
			/>
		</div>
	</div>

	{#if uploading}
		<div class="w-full bg-neutral-800 rounded-full h-2.5 overflow-hidden">
			<div class="bg-green-500 h-2.5 transition-all duration-300" style="width: {progress}%"></div>
		</div>
	{/if}
</div>
