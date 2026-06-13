<script lang="ts">
	import { User } from '@lucide/svelte';
	import { onDestroy, tick } from 'svelte';
	import BlankUser from '$lib/assets/no-pfp-provided.webp';
	import { goto } from '$app/navigation';
	import { fade } from 'svelte/transition';
	import { userInfo } from '$lib/dashboard/stores/store';
	import Modal from '../Modal.svelte';
	import { accounts } from '$lib/dashboard/stores/persist';

	let open: boolean = $state(false);

	let boxRef: HTMLDivElement | null = $state(null);
	let inputRef: HTMLInputElement | null = $state(null);
	let btn: HTMLButtonElement | null = $state(null);
	let changeAccountModalOpen: boolean = $state(false);

	function handleClickOutside(e: MouseEvent) {
		if (!boxRef || !boxRef.contains(e.target as Node)) {
			if (!btn || !btn.contains(e.target as Node)) {
				open = false;
			}
		}
	}

	async function openBox(e: MouseEvent) {
		e.stopPropagation();
		open = !open;
		await tick();
	}

	function closeBox() {
		open = false;
		document.removeEventListener('mousedown', handleClickOutside);
	}
</script>

<div class="relative">
	<button
		bind:this={btn}
		onclick={openBox}
		class="relative flex cursor-pointer items-center gap-1 mr-5 bg-neutral-800 p-3 px-7 rounded-xl"
	>
		<User size={15} />
		<p class="font-black truncate max-w-15">{$userInfo.name}</p>
	</button>
	{#if open}
		{@render SearchContent()}
	{/if}
</div>

<Modal
	className="w-100"
	title="Zmień konto"
	bind:opened={changeAccountModalOpen}
	onClose={() => (changeAccountModalOpen = false)}
>
	<div class="flex flex-col gap-4 p-3 pt-6 rounded-b-xl rounded-t-xl">
		<div class="max-h-60 overflow-y-auto flex flex-col gap-2 custom-scrollbar pr-1">
			{#each $accounts as account}
				<button
					class="flex items-center justify-between p-3 rounded-xl border border-neutral-800 bg-neutral-800/40 hover:bg-neutral-800 hover:border-neutral-700 transition-all group shadow-lg"
				>
					<div class="flex items-center gap-3">
						<div
							class="w-10 h-10 rounded-full bg-blue-600 text-white flex items-center justify-center font-bold shadow-md group-hover:scale-105 transition-transform"
						>
							{account.name.charAt(0).toUpperCase()}
						</div>

						<div class="text-left">
							<p class="font-semibold text-neutral-100 leading-tight">{account.name}</p>
							<p class="text-sm text-neutral-400">@{account.login}</p>
						</div>
					</div>

					<div class="flex items-center gap-2">
						{#if $userInfo.login === account.login}
							<div
								class="flex items-center gap-1.5 bg-blue-500/10 text-blue-400 px-3 py-1 rounded-full border border-blue-500/20"
							>
								<span
									class="w-2 h-2 bg-blue-400 rounded-full animate-pulse shadow-[0_0_8px_rgba(96,165,250,0.6)]"
								></span>
								<span class="text-[10px] font-bold uppercase tracking-widest">Active</span>
							</div>
						{:else}
							<span
								class="text-neutral-600 group-hover:text-neutral-300 group-hover:translate-x-1 transition-all"
							>
								<svg
									xmlns="http://www.w3.org/2000/svg"
									width="18"
									height="18"
									viewBox="0 0 24 24"
									fill="none"
									stroke="currentColor"
									stroke-width="2.5"
									stroke-linecap="round"
									stroke-linejoin="round"><path d="M5 12h14" /><path d="m12 5 7 7-7 7" /></svg
								>
							</span>
						{/if}
					</div>
				</button>
			{/each}
		</div>

		<div class="h-[1px] bg-neutral-800 w-full my-1"></div>

		<button
			class="flex items-center justify-center gap-2 p-3 rounded-xl border-2 border-dashed border-neutral-700 text-neutral-400 hover:border-blue-500/50 hover:bg-blue-500/5 hover:text-blue-400 transition-all font-medium group"
		>
			<span class="text-xl group-hover:rotate-90 transition-transform duration-300">+</span>
			Dodaj inne konto
		</button>
	</div>
</Modal>

{#snippet SearchContent()}
	<div
		bind:this={boxRef}
		in:fade={{ duration: 150 }}
		out:fade={{ duration: 150 }}
		class="absolute right-0 -top-1 z-10 w-[400px] border-1 border-neutral-700/60 rounded-xl bg-neutral-800"
	>
		<div class="flex flex-col">
			{@render DisplayBasicUserInfo()}

			<div class="flex cursor-pointer flex-col">
				<button
					onclick={() => {
						goto('/dashboard/settings');
					}}
					class="cursor-pointer border-b-1 border-neutral-700/60 p-3 text-start text-neutral-400 hover:text-white"
					>Ustawienia</button
				>
				<button
					onclick={() => {
						changeAccountModalOpen = true;
					}}
					class="cursor-pointer border-b border-neutral-700/60 p-3 text-start text-neutral-400 hover:text-white"
					>Zmień konto</button
				>
				<button
					onclick={() => {
						localStorage.setItem('token', '');
						goto('/login');
					}}
					class="cursor-pointer border-neutral-700/60 p-3 text-start text-neutral-400 hover:text-white"
					>Wyloguj</button
				>
			</div>
		</div>
	</div>
{/snippet}

{#snippet DisplayBasicUserInfo()}
	<div class="flex gap-2 border-b-1 border-neutral-700/60 p-4">
		<img alt="basic-pfp" class="size-12 rounded-full" src={BlankUser} />
		<div class="flex flex-col gap-0.5 py-1 text-start">
			<h3 class="text-[15px] font-semibold">{$userInfo.name}</h3>
			<p class="text-[12px]">{$userInfo.login}</p>
		</div>
	</div>
{/snippet}

<svelte:document
	onmousedown={(e) => {
		handleClickOutside(e);
	}}
/>
