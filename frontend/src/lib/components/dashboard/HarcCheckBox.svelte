<script lang="ts">
	import { toast } from '$lib/dashboard/stores/toast';
	import Icon from '@iconify/svelte';
	import { fade } from 'svelte/transition';

	type Props = {
		checked: boolean;
		label?: string;
	};

	let { checked = $bindable(false), label }: Props = $props();
</script>

<div class="flex gap-2 items-center">
	<div class="relative w-5 h-5 border-2 border-neutral-700">
		{#if checked}
			<div
				in:fade={{ duration: 100 }}
				out:fade={{ duration: 100 }}
				class="absolute z-8 w-full h-full bottom-3 right-2"
			>
				<Icon icon="material-symbols:check" class="text-green-500" width="37" height="37" />
			</div>
		{/if}

		<div
			class="absolute z-9 w-full h-full cursor-pointer"
			onclick={() => {
				checked = !checked;
			}}
		></div>
	</div>
	<div>
		<p class="text-sm text-neutral-400">
			{label}
		</p>
	</div>
</div>
