<script lang="ts">
	import type { Role } from '$lib/types/user';
	import Icon from '@iconify/svelte';
	import { onMount } from 'svelte';

	type Props = { role: Role };

	let { role }: Props = $props();

	function onEditButtonClick(role: Role) {}
</script>

<div class="bg-neutral-800 justify-between flex rounded-lg p-3 px-6 items-center">
	<div class="flex items-center gap-3">
		<div class="flex flex-col">
			<span class="flex gap-0.5 items-center text-white">
				<p class="font-black hover:underline cursor-pointer">
					{role.name}
				</p>
			</span>
			<p class="hover:underline cursor-pointer">
				<!-- pill -->
			</p>
		</div>
	</div>
	<div class="flex items-center gap-2">
		<button
			onclick={() => {
				onEditButtonClick(role);
			}}
			class="p-2 hover:bg-neutral-700/50 hover:text-blue-400 rounded-xl cursor-pointer"
		>
			<Icon icon="eos-icons:role-binding" width="20" height="20" />
		</button>
		<button
			onclick={() => {
				onEditButtonClick(role);
			}}
			class="p-2 hover:bg-neutral-700/50 hover:text-blue-400 rounded-xl cursor-pointer"
		>
			<Icon icon="boxicons:edit-filled" width="20" height="20" />
		</button>
		<button class="p-2 hover:bg-neutral-700/50 hover:text-red-400 rounded-xl cursor-pointer">
			<Icon icon="boxicons:trash-filled" width="20" height="20" />
		</button>
	</div>
</div>

<style>
	@import 'tailwindcss';

	.selector {
		@apply;
	}
</style>
