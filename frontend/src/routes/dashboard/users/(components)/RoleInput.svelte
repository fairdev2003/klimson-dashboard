<script lang="ts">
	import { api } from '$lib/api/api';
	import RDBInput from '$lib/components/modal/RDBInput.svelte';
	import { debug } from '$lib/terminal/logic';
	import type { Role } from '$lib/types/user';
	import Icon from '@iconify/svelte';
	import { onMount } from 'svelte';
	import { slide } from 'svelte/transition';

	type Props = {
		role_id?: number;
		user_role: Role | undefined;
		label?: string;
		oninput?: (e: string) => void;
	};

	let { role_id = $bindable(), label, oninput, user_role = $bindable() }: Props = $props();

	let searchOn = $state(false);

	function handleClickOutside(event: MouseEvent) {
		const target = event.target as HTMLElement;

		const clickedDropdown = target.closest('.role-container');

		if (!clickedDropdown) {
			searchOn = false;
		}
	}

	async function fetchRoles() {
		try {
			const response = await api.user.ListRoles();
			if (response.status === 200) {
				roles = response.data.data;
			} else {
				debug.error('Failed to fetch roles:', response.statusText);
			}
		} catch (error) {
			debug.error(error);
		}
	}

	let roles: Role[] = $state([]);

	onMount(async () => {
		await fetchRoles();
	});
</script>

<div class="flex flex-col gap-3">
	<span class="text-neutral-400 uppercase justify-between items-center flex font-bold text-xs">
		<p>ROLE</p>
	</span>

	{#if searchOn}
		<div class="flex flex-col gap-2 fixed inset-0 bg-black/60"></div>
	{/if}
	<div class="bg-neutral-900 relative w-full flex flex-col">
		<div
			class:rounded-b-none={searchOn}
			class="flex justify-between items-center bg-neutral-800 w-full p-4 border border-b-none border-neutral-800 rounded-lg"
		>
			<p class="text-neutral-400 text-sm font-black">{user_role?.name}</p>
			<button
				onclick={() => {
					searchOn = !searchOn;
				}}
				class="p-2 hover:bg-neutral-700/50 hover:text-blue-400 rounded-xl cursor-pointer"
			>
				<Icon icon="uil:exchange" width="20" height="20" />
			</button>
		</div>
		{#if searchOn}
			<div
				in:slide={{ duration: 200 }}
				out:slide={{ duration: 200 }}
				class="rounded-lg border border-t-none border-neutral-800 rounded-t-none w-full pt-4 role-container px-4 top-full left-0 rounded-b-lg flex flex-col gap-2 overflow-y-auto"
			>
				<RDBInput placeholder="Search roles..." label="ROLE" />
				<div class="flex flex-col gap-2 pb-10">
					{#each roles as r}
						<button
							onclick={() => {
								user_role = r;
								searchOn = false;
							}}
							class="flex hover:bg-neutral-700/50 justify-between items-center bg-neutral-800 w-full p-4 border-neutral-800 rounded-lg"
						>
							<p class="text-neutral-400 text-sm font-black">{r.name}</p>
						</button>
					{/each}
				</div>
			</div>
		{/if}
	</div>
</div>

<svelte:document onclick={(e) => handleClickOutside(e)} />
