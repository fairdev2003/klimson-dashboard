<script lang="ts">
	import { goto } from '$app/navigation';
	import { toast } from '$lib/dashboard/stores/toast';
	import { fade } from 'svelte/transition';
	import Icon from '@iconify/svelte';
	let btn: HTMLButtonElement | null = null;
	let boxRef: HTMLDivElement | null = null;
	let open = $state(false);

	type QuickAction = {
		label: string;
		description: string;
		icon: string;
		action: () => void;
		color?: string;
	};

	function handleClickOutside(e: MouseEvent) {
		if (!boxRef || !boxRef.contains(e.target as Node)) {
			if (!btn || !btn.contains(e.target as Node)) {
				open = false;
			}
		}
	}

	export const quickActions: QuickAction[] = [
		{
			label: 'Nowy Quiz',
			description: 'Stwórz od zera nowy zestaw pytań',
			icon: 'mdi:plus-circle-outline',
			action: () => {
				goto('/dashboard/quizzes/add');
			},
			color: 'text-blue-400'
		},
		{
			label: 'Panel twórcy',
			description: 'Zaakceptuj zmiany w quizach',
			icon: 'mdi:check-circle-outline',
			action: () => {
				goto('/dashboard/creator');
			},
			color: 'text-green-400'
		},
		{
			label: 'Lista Kontrybutorów',
			description: 'Zarządzaj uprawnieniami zespołu',
			icon: 'mdi:account-group',
			action: () => {
				goto('/dashboard/contributors');
			},
			color: 'text-purple-400'
		},
		{
			label: 'Ustawienia Panelu',
			description: 'Konfiguracja ogólna systemu harc',
			icon: 'mdi:cog',
			action: () => {
				goto('/dashboard/settings');
			},
			color: 'text-neutral-400'
		},
		{
			label: 'Wyloguj',
			description: 'Zakończ bezpiecznie sesję',
			icon: 'mdi:logout',
			action: () => {
				localStorage.setItem('token', '');
				goto('/login');
			},
			color: 'text-red-400'
		}
	];
</script>

<div class="relative">
	<button
		bind:this={btn}
		onclick={() => {
			open = !open;
		}}
		class="border-1 border-secondary flex size-[50px] cursor-pointer items-center justify-center transition-colors hover:bg-neutral-700"
	>
		<Icon icon="icon-park-outline:hamburger-button" width="35" height="35" />
	</button>

	{#if open}
		<div
			in:fade={{ duration: 150 }}
			out:fade={{ duration: 150 }}
			bind:this={boxRef}
			class="border-1 absolute top-12 z-10 w-[300px] border-neutral-700/60 bg-neutral-800"
		>
			{#each quickActions as action}
				{@render Option(action)}
			{/each}
		</div>
	{/if}
</div>

{#snippet Option(action: QuickAction)}
	<button
		class="w-full text-start"
		onclick={() => {
			action.action();
			open = false;
		}}
	>
		<div class="flex cursor-pointer items-center gap-3 p-4 hover:bg-neutral-700">
			<Icon icon={action.icon} width="24" height="24" class={action.color} />
			<div>
				<h3 class="font-semibold">{action.label}</h3>
				<p class="text-sm text-neutral-400">{action.description}</p>
			</div>
		</div>
	</button>
{/snippet}

<svelte:document
	onmousedown={(e) => {
		handleClickOutside(e);
	}}
/>
