<script lang="ts">
	let checked: boolean = $state(false);

	type Props = {
		for_what?: string;
		label: string;
		onchange?: (val: boolean) => void;
		value: boolean;
	};

	const { for_what, onchange, label, value = $bindable(checked) }: Props = $props();

	$effect(() => {
		checked = value;
	});
</script>

<div class="mt-5">
	<div class="flex flex-col">
		<p class="mb-2 text-sm text-neutral-400">{label}</p>
		<div class="flex">
			<button
				onclick={() => {
					checked = false;
					onchange?.(false);				
				}}
				class:bg-red-500={!checked}
				class="border-l-1 border-y-1 flex w-20 cursor-pointer items-center justify-center border-neutral-700 px-5 py-2 text-sm transition-all hover:bg-red-500"
				>NIE</button
			>
			<button
				onclick={() => {
					checked = true;
					onchange?.(true);
				}}
				class:bg-green-500={checked}
				class="border-r-1 border-y-1 flex w-20 cursor-pointer items-center justify-center border-neutral-700 px-5 py-2 text-sm transition-all hover:bg-green-500"
				>TAK</button
			>
		</div>
	</div>
</div>
