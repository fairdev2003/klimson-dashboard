<script lang="ts">
	type Props = {
		label: string;
		onclick?: () => void;
	};

	let { onclick, label }: Props = $props();
</script>

<button
	class="bg-primary dropdown-button flex items-center gap-1 justify-center px-6 p-2 border cursor-pointer hover:bg-foreground border-border rounded-xl"
	onclick={() => {
		onclick?.();
	}}
>
	<p>{label}</p>
</button>
