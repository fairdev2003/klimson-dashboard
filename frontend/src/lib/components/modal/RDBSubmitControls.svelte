<script lang="ts">
	import Icon from '@iconify/svelte';
	import type { FormControlsProps } from './modal.svelte';

	let rdb_submit_controls_props: FormControlsProps = $props();

	let statement =
		rdb_submit_controls_props.onCancel ||
		rdb_submit_controls_props.onDelete ||
		rdb_submit_controls_props.onLog ||
		rdb_submit_controls_props.onSubmit ||
		rdb_submit_controls_props.onDeny ||
		rdb_submit_controls_props.onAccept ||
		rdb_submit_controls_props.onLogout;
</script>

<div class="flex flex-col gap-2 justify-between">
	{#if statement}
		{@render Divider()}
	{/if}
	<div class="flex gap-2 items-center justify-between">
		<div class="flex gap-1 justify-center items-center">
			{#if rdb_submit_controls_props.onCancel}
				<button
					onclick={rdb_submit_controls_props.onCancel}
					class="bg-neutral-800 px-3 p-1 rounded-md text-neutral-400 hover:bg-neutral-700 cursor-pointer"
					>Cancel</button
				>
			{/if}
			{#if rdb_submit_controls_props.onLog}
				<button
					onclick={rdb_submit_controls_props.onLog}
					class="bg-neutral-800 px-4 p-1 hover:bg-neutral-700 rounded-md text-orange-400 cursor-pointer text-sm h-8"
					><Icon icon="catppuccin:log" /></button
				>
			{/if}
		</div>

		<div class="flex gap-1 justify-center items-center">
			{#if rdb_submit_controls_props.onDelete}
				<button
					onclick={rdb_submit_controls_props.onDelete}
					class="bg-neutral-800 px-3 p-1 rounded-md hover:bg-neutral-700 cursor-pointer text-red-400"
					>Delete</button
				>
			{/if}
			{#if rdb_submit_controls_props.onSubmit}
				<button
					onclick={rdb_submit_controls_props.onSubmit}
					class="bg-blue-700 px-3 p-1 hover:bg-blue-700 cursor-pointer rounded-md">Submit</button
				>
			{/if}
			{#if rdb_submit_controls_props.onDeny}
				<button
					onclick={rdb_submit_controls_props.onDeny}
					class="bg-neutral-800 px-3 p-1 rounded-md hover:bg-neutral-700 cursor-pointer text-red-400"
					>No</button
				>
			{/if}
			{#if rdb_submit_controls_props.onAccept}
				<button
					onclick={rdb_submit_controls_props.onAccept}
					class="bg-neutral-800 px-3 p-1 rounded-md hover:bg-neutral-700 cursor-pointer text-green-400"
					>Yes</button
				>
			{/if}
			{#if rdb_submit_controls_props.onLogout}
				<button
					onclick={rdb_submit_controls_props.onLogout}
					class="bg-neutral-800 px-3 p-1 rounded-md hover:bg-neutral-700 cursor-pointer text-red-400"
					>Logout</button
				>
			{/if}
		</div>
	</div>
</div>

{#snippet Divider()}
	<div class="border-t-2 border-neutral-700 my-2"></div>
{/snippet}
