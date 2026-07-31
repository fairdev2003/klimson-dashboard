<script lang="ts">
	import type { User } from '$lib/api/types';
	import Icon from '@iconify/svelte';
	import { blur } from 'svelte/transition';

	type Props = {
		users: User[];
		usersLoading: boolean;
		onEditButtonClick: (user: User) => void;
	};

	let { users, usersLoading = $bindable(true), onEditButtonClick }: Props = $props();
</script>

<div in:blur={{ duration: 150 }} class="flex flex-col gap-2 lg:w-2xl max-w-2xl mx-auto">
	{#each users as user}
		<div class="bg-neutral-800 justify-between flex rounded-lg p-3 px-6 items-center">
			<div class="flex items-center gap-3">
				<img src={user.pfp} class="size-10 rounded-full" alt="pfp-{user.id}" />
				<div class="flex flex-col">
					<span class="flex gap-0.5 items-center text-white">
						<p class="font-black hover:underline cursor-pointer">
							{user.nickname}
						</p>
					</span>
					<p class="hover:underline cursor-pointer">
						<!-- pill -->
						<span class="rounded-full px-2 text-xs p-0.5 bg-black">{user.role?.name}</span>
					</p>
				</div>
			</div>
			<div class="flex items-center gap-2">
				<button
					onclick={() => {
						onEditButtonClick(user);
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
	{/each}
</div>
