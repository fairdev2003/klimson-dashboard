<script lang="ts">
	type Props = {
		value?: string;
		label?: string;
	};

	let { value = $bindable(''), label }: Props = $props();
</script>

<div class="flex flex-col gap-1">
	<span class="text-neutral-400 uppercase justify-between items-center flex font-bold text-xs">
		<p>{label}</p>
	</span>

	<input class="rounded-lg border-0 bg-neutral-800 p-2" />
</div>
