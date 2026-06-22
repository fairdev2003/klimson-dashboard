<script lang="ts">
	import Heading from '$lib/components/dashboard/typography/Heading.svelte';
	import Icon from '@iconify/svelte';
	import { fade } from 'svelte/transition';

	let selectedLabel: 'acc' | 'roles' = $state('acc');
</script>

<div in:fade={{ duration: 150 }} class="flex flex-col m-8 my-4">
	<div class="flex-col flex gap-1">
		<Heading>
			<div class="flex gap-2 items-center">
				<Icon icon="mdi:user-key" />
				<p>CMS Access</p>
			</div>
		</Heading>
		<span class="text-sm font-md text-neutral-400"
			>Control what users has permission to specific part of the dashboard.</span
		>
	</div>
	<div class="flex mt-4 gap-4">
		<button
			onclick={() => {
				selectedLabel = 'acc';
			}}
			class:selected-label-pill={selectedLabel === 'acc'}
			class:normal-label-pill={selectedLabel !== 'acc'}
			class="px-4 p-2 rounded-full text-sm transition-all cursor-pointer"
		>
			Registered accounts
		</button>
		<button
			onclick={() => {
				selectedLabel = 'roles';
			}}
			class:selected-label-pill={selectedLabel === 'roles'}
			class:normal-label-pill={selectedLabel !== 'roles'}
			class="px-4 p-2 bg-neutral-800 transition-all rounded-full cursor-pointer text-sm"
			>Roles</button
		>
	</div>
</div>

<style>
	@import 'tailwindcss';

	.selected-label-pill {
		@apply text-blue-400 bg-blue-500/20 hover:bg-blue-500/40;
	}

	.normal-label-pill {
		@apply bg-neutral-800 text-neutral-400 hover:bg-neutral-700;
	}
</style>
