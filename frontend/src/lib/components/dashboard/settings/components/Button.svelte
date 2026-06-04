<script lang="ts">
	type Props = {
		label: string;
		onclick?: () => void;
	};

	let { onclick, label }: Props = $props();
</script>

<button
	class="bg-neutral-800 dropdown-button flex items-center gap-1 justify-center px-6 p-2 border cursor-pointer hover:bg-neutral-700 border-neutral-700 rounded-xl"
	onclick={() => {
		onclick?.();
	}}
>
	<p>{label}</p>
</button>
