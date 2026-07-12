<script lang="ts">
	import Icon from '@iconify/svelte';
</script>

<div class="flex flex-col gap-2 justify-between">
	<div class="border-t-2 border-neutral-700 my-2"></div>
	<div class="flex gap-2 items-center justify-between">
		<div class="flex gap-1 justify-center items-center">
			<button
				class="bg-neutral-800 px-3 p-1 rounded-md text-neutral-400 hover:bg-neutral-700 cursor-pointer"
				>Cancel</button
			>
			<button
				class="bg-neutral-800 px-4 p-1 hover:bg-neutral-700 rounded-md text-orange-400 cursor-pointer text-sm h-8"
				><Icon icon="catppuccin:log" /></button
			>
		</div>

		<div class="flex gap-1 justify-center items-center">
			<button
				class="bg-neutral-800 px-3 p-1 rounded-md hover:bg-neutral-700 cursor-pointer text-red-400"
				>Delete</button
			>
			<button class="bg-blue-700 px-3 p-1 hover:bg-blue-700 cursor-pointer rounded-md"
				>Submit</button
			>
		</div>
	</div>
</div>
