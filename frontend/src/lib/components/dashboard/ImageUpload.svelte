<script module>
	import { toast } from '$lib/dashboard/stores/toast';
	import Icon from '@iconify/svelte';
	import type { ImageKey } from '$lib/api/types';
	import { api } from '$lib/api/api';
</script>

<script lang="ts">
	let {
		label = 'Zdjęcie',
		src = $bindable(),
		sendKey = 'quiz',
		file = $bindable()
	} = $props<{
		label?: string;
		src: string;
		sendKey?: ImageKey;
		file: File | undefined;
	}>();

	let isOver = $state(false);
	let fileInput: HTMLInputElement | undefined = $state();

	const processFile = (f: File) => {
		const reader = new FileReader();
		reader.onload = (e) => {
			src = e.target?.result as string;
		};
		toast.success(src);
		file = f;
		reader.readAsDataURL(f);
	};

	const handleDrop = (e: DragEvent) => {
		e.preventDefault();
		e.stopPropagation();
		isOver = false;

		const f = e.dataTransfer?.files?.[0];
		if (f?.type.startsWith('image/')) {
			processFile(f);
		}
	};

	const handleSelect = (e: Event) => {
		const target = e.target as HTMLInputElement;
		file = target.files?.[0];
		if (file) processFile(file);
	};
</script>

<div class="flex flex-col gap-2">
	<span class="text-sm font-medium text-neutral-400">{label}</span>

	<div
		role="button"
		tabindex="0"
		class="group relative flex h-64 w-full flex-col items-center justify-center overflow-hidden rounded-xl border-2 border-dashed transition-all duration-300 select-none focus:outline-0
        {isOver
			? 'scale-[1.01] border-blue-500 bg-blue-500/10'
			: 'border-neutral-800 bg-neutral-900/50 hover:border-neutral-700'}"
		ondragover={(e) => {
			e.preventDefault();
			isOver = true;
		}}
		ondragenter={(e) => {
			e.preventDefault();
			isOver = true;
		}}
		ondragleave={() => (isOver = false)}
		ondrop={handleDrop}
		onclick={() => fileInput?.click()}
		onkeydown={(e) => e.key === 'Enter' && fileInput?.click()}
	>
		{#if src}
			<img
				{src}
				alt="Preview"
				class="h-full w-full object-contain transition-transform duration-500"
			/>
			<div
				class="absolute inset-0 flex items-center justify-center bg-black/40 opacity-0 backdrop-blur-sm transition-opacity group-hover:opacity-100"
			>
				<div class="flex flex-col items-center p-4 text-center text-white">
					<Icon icon="ri:refresh-line" class="text-3xl" />
					<span class="mt-1 text-sm font-bold">Kliknij lub upuść, aby zmienić</span>
				</div>
			</div>
		{:else}
			<div class="flex flex-col items-center gap-3 text-neutral-500 group-hover:text-neutral-300">
				<div class="rounded-full bg-neutral-800 p-4 transition-colors group-hover:bg-neutral-700">
					<Icon icon="ri:image-add-line" class="text-4xl" />
				</div>
				<div class="text-center">
					<p class="text-sm font-semibold">Przeciągnij i upuść</p>
					<p class="text-xs text-neutral-600">lub kliknij, aby wybrać</p>
				</div>
			</div>
		{/if}

		<input
			type="file"
			accept="image/*"
			class="hidden"
			bind:this={fileInput}
			onchange={handleSelect}
		/>
	</div>
</div>
