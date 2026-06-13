<script lang="ts">
	import Icon from '@iconify/svelte';

	let items = $state([
		{ id: 1, name: 'Hub', icon: 'mynaui:home' },
		{ id: 2, name: 'File Storage', icon: 'mynaui:database' },
		{ id: 3, name: 'Tools', icon: 'mynaui:cog' },
		{ id: 4, name: 'Settings', icon: 'mynaui:setings' },
		{ id: 5, name: 'Users', icon: 'mynaui:users' },
		{ id: 6, name: 'Analytics', icon: 'mynaui:chart-bar' },
		{ id: 7, name: 'Messages', icon: 'mynaui:envelope' },
		{ id: 8, name: 'Security', icon: 'mynaui:lock' }
	]);

	let draggedId = $state(null);
	let overId = $state(null);

	function handleDragStart(id: any) {
		draggedId = id;
	}

	function handleDragOver(e: DragEvent, id: any) {
		e.preventDefault();
		overId = id;
	}

	function handleDrop(targetId: any) {
		const fromIndex = items.findIndex((i) => i.id === draggedId);
		const toIndex = items.findIndex((i) => i.id === targetId);

		if (fromIndex !== toIndex) {
			const item = items[fromIndex];
			items.splice(fromIndex, 1);
			items.splice(toIndex, 0, item);
		}
		draggedId = null;
		overId = null;
	}
</script>

<div class="flex flex-col gap-3 pb-10">
	{#each items as item (item.id)}
		<div
			draggable="true"
			ondragstart={() => handleDragStart(item.id)}
			ondragover={(e) => handleDragOver(e, item.id)}
			ondrop={() => handleDrop(item.id)}
			class="relative flex items-center gap-3 px-4 py-3 bg-neutral-800 border border-neutral-700 rounded-xl cursor-grab transition-all"
			class:opacity-50={draggedId === item.id}
		>
			{#if overId === item.id && draggedId !== item.id}
				<div class="absolute -top-1.5 left-0 right-0 h-1 bg-blue-500 rounded-full z-10"></div>
			{/if}

			<Icon icon="mynaui:grip-dots-vertical" class="text-neutral-500" />
			<Icon icon={item.icon} width="24" />
			<span class="font-medium">{item.name}</span>
		</div>
	{/each}
</div>
