<script lang="ts">
	type Props = {
		value: string;
		placeholder: string;
	};

	let { value = $bindable(), placeholder }: Props = $props();
</script>

<input
	{placeholder}
	{value}
	oninput={(e: InputEvent) => {
		if (!e || !e.target) return;

		value = e.target.value as string;
	}}
	class="bg-neutral-800 input dropdown-button flex items-center gap-1 justify-center w-60 px-6 p-2 border hover:bg-neutral-700 border-neutral-700 rounded-xl"
/>

<style>
	@import 'tailwindcss';

	.input {
		@apply focus:ring-blue-500 transition-colors focus:ring-2 focus:bg-neutral-700 outline-0;
	}
</style>
