<script lang="ts">
	import Modal from '$lib/components/Modal.svelte';

	let opened: boolean = $state(false);
</script>

<button
	onclick={() => {
		opened = !opened;
	}}
	class="w-full h-15 cursor-pointer hover:bg-neutral-700 bg-neutral-800 rounded-xl px-2"
>
	<div class="size-10 rounded-lg bg-green-500/50"></div>
</button>

<Modal
	bind:opened
	onClose={() => {
		opened = !opened;
	}}
	className="w-100 "
></Modal>
