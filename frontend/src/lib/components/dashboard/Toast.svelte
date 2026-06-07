<script lang="ts">
	import { debugOn } from '$lib/dashboard/stores/persist';
	import { toast, type Toast } from '$lib/dashboard/stores/toast';
	import { X } from '@lucide/svelte';
	import { flip } from 'svelte/animate';
	import { blur, fly, scale } from 'svelte/transition';

	const colors = {
		success: 'bg-green-600',
		error: 'bg-red-600',
		info: 'bg-neutral-800',
		warning: 'bg-yellow-700 text-black'
	};

	// Reaktywne odliczanie czasu tylko gdy są toasty
	let now = $state(Date.now());
	let interval: any;

	$effect(() => {
		if ($toast.length > 0) {
			interval = setInterval(() => {
				now = Date.now();
				// Sprawdzanie czy któryś toast powinien zniknąć
				$toast.forEach((t) => {
					if (now - t.createdAt >= t.duration) {
						toast.remove(t.id);
					}
				});
			}, 50);
		} else {
			clearInterval(interval);
		}
		return () => clearInterval(interval);
	});

	function getProgress(t: Toast) {
		const elapsed = now - t.createdAt;
		const percentage = 100 - (elapsed / t.duration) * 100;
		return Math.max(0, Math.min(100, percentage));
	}
</script>

<div
	class="fixed top-5 right-1/2 translate-x-1/2 z-[150] flex flex-col gap-3 transition-all duration-300"
	class:bottom-20={$debugOn}
	class:bottom-5={!$debugOn}
>
	{#each $toast as t (t.id)}
		<div
			animate:flip={{ duration: 300 }}
			transition:blur={{ duration: 400 }}
			class="flex min-w-50 flex-col rounded-lg justify-between overflow-hidden text-white shadow-xl {colors[
				t.type
			]}"
		>
			<div class="flex items-center p-3 justify-between text-sm">
				<div class="flex gap-1 items-center">
					<div class="h-full w-2 bg-red-500"></div>
					<span class="font-medium">{t.message}</span>
				</div>
				<button
					class="ml-3 flex size-7 cursor-pointer items-center justify-center rounded-full transition-colors hover:bg-white/20"
					onclick={() => toast.remove(t.id)}
				>
					<X size={16} />
				</button>
			</div>

			<div class="h-1 w-full">
				<div class="h-full bg-white/60" style="width: {getProgress(t)}%"></div>
			</div>
		</div>
	{/each}
</div>
