<script lang="ts">
	import { X } from '@lucide/svelte';
	import { toast } from '$lib/dashboard/stores/toast';
	import { developerView } from '$lib/dashboard/stores/persist';
	import { onMount } from 'svelte';

	let v: string = $state('');
	let el: HTMLInputElement;

	type Props = {
		label?: string;
		value?: string | number;
		disclaimer?: string;
		disabled?: boolean;
		focus?: boolean;
	};

	let {
		focus = false,
		label = 'LABEL',
		value = $bindable(v),
		disclaimer = '',
		disabled = false
	}: Props = $props();

	$effect(() => {
		if (focus) {
			el.focus();
		}
	});
</script>

<div class="">
	<div class="">
		<p class="mb-2 text-sm text-neutral-400">
			{label} <span class="text-xs text-neutral-300">{disclaimer}</span>
		</p>

		<div class="relative flex">
			{#if disabled}
				<div
					onclick={() => {
						if (disabled && value) {
							navigator.clipboard.writeText(value as string);
							toast.success('Skopiowano do schowka');
						}
					}}
					class={`absolute h-full w-full cursor-pointer rounded-lg  p-3 ${$developerView ? 'bg-red-500/50' : 'rounded-xl '}`}
				></div>
			{/if}
			<input
				bind:this={el}
				{disabled}
				bind:value
				class="h-10 w-full rounded-lg {disabled
					? ' cursor-pointer rounded-lg border-neutral-800 bg-black/10 outline-red-500'
					: 'border-neutral-700/60 bg-transparent placeholder-neutral-700'} {value
					? ''
					: 'rounded-lg'} border-1 outline-none focus:border-white"
				placeholder="..."
			/>
		</div>
	</div>
</div>
