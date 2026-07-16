<script lang="ts">
	type Props = {
		value: string;
	};

	let { value = $bindable() }: Props = $props();
</script>

<div>
	<div
		class="bg-transparent h-5 w-full border-0 text-neutral-400 text-xs placeholder-neutral-400 outline-none focus:outline-none focus:ring-0"
	>
		{value}
	</div>
</div>

<!-- <svelte:document
	onkeydown={(e) => {
		value = value + e.key;
	}}
/> -->
