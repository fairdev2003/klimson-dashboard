<script lang="ts">
	import Icon from '@iconify/svelte';
	import { fade } from 'svelte/transition';
	import type { RoleOption } from '../types';
	import { onMount } from 'svelte';
	import { colorMap } from '../vars';

	interface Props {
		roles: RoleOption[];
		selectedRolesString: string; // Teraz przyjmujemy stringa, np. "role1,role2"
		onChange: (selected: string) => void; // Zwracamy stringa
	}

	let { roles, selectedRolesString = $bindable(''), onChange }: Props = $props();

	// Reaktywna tablica ID wyciągnięta ze stringa
	// let selectedIds = $derived.by(() => {
	// 	if (!selectedRolesString) return;
	// 	selectedRolesString ? selectedRolesString.split(',').filter((id) => id.length > 0) : [];
	// });

	// function toggleRole(id: string) {
	// 	let newIds: string[];

	// 	if (selectedIds.includes(id)) {
	// 		newIds = selectedIds.filter((i) => i !== id);
	// 	} else {
	// 		newIds = [...selectedIds, id];
	// 	}

	// 	// Łączymy z powrotem w stringa i wysyłamy do nadrzędnego komponentu
	// 	onChange(newIds.join(','));
	// }

	// --- Reszta Twojej logiki paginacji bez zmian ---
	const ROLES_PER_PAGE = 8;
	let currentPage = $state(0);
	let totalPages = $derived(Math.ceil(roles.length / ROLES_PER_PAGE));
	let visibleRoles = $derived(
		roles.slice(currentPage * ROLES_PER_PAGE, (currentPage + 1) * ROLES_PER_PAGE)
	);
</script>

<div class="flex flex-col gap-6">
	{#if totalPages > 1}
		<div class="flex items-center justify-center gap-4 border-b border-neutral-800 pb-4">
			<button
				onclick={() => currentPage--}
				disabled={currentPage === 0}
				class="disabled:opacity-20"
			>
				<Icon icon="mdi:chevron-left" width="24" />
			</button>
			<div class="flex gap-2">
				{#each Array(totalPages) as _, i}
					<button
						onclick={() => (currentPage = i)}
						class="h-1 rounded-full transition-all {currentPage === i
							? 'w-8 bg-white'
							: 'w-2 bg-neutral-700'}"
					></button>
				{/each}
			</div>
			<button
				onclick={() => currentPage++}
				disabled={currentPage === totalPages - 1}
				class="disabled:opacity-20"
			>
				<Icon icon="mdi:chevron-right" width="24" />
			</button>
		</div>
	{/if}

	<div class="grid grid-cols-1 gap-3 md:grid-cols-2">
		{#each visibleRoles as role (role.id)}
			{@const isActive = true}
			<button
				onclick={() => toggleRole(role.id)}
				class="relative flex items-start gap-4 rounded-xl border p-4 text-left transition-all duration-300
                {isActive
					? (colorMap[role.color] ?? 'border-blue-500 text-blue-500') +
						' border-opacity-100 shadow-lg'
					: 'border-neutral-800 bg-neutral-900/50 text-neutral-500'}"
			>
				<div class="flex flex-1 flex-col">
					<span class="font-bold {isActive ? 'text-white' : 'text-neutral-400'}">{role.id}</span>
					<p class="mt-1 text-xs text-neutral-500">{role.description}</p>
				</div>
				{#if isActive}
					<Icon icon="mdi:check-circle" class="text-current" width="20" />
				{/if}
			</button>
		{/each}
	</div>
</div>
