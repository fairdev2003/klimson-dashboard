<script lang="ts">
	import { debugOn } from '$lib/dashboard/stores/persist';
	import { toast, type Toast } from '$lib/dashboard/stores/toast';
	import { X } from '@lucide/svelte';
	import { flip } from 'svelte/animate';
	import { fly } from 'svelte/transition';

	const colors = {
		success: 'bg-green-600',
		error: 'bg-red-600',
		info: 'bg-neutral-800 border-1 border-neutral-700',
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
	class="fixed left-5 z-[150] flex flex-col gap-3 transition-all duration-300"
	class:bottom-20={$debugOn}
	class:bottom-5={!$debugOn}
>
	{#each $toast as t (t.id)}
		<div
			animate:flip={{ duration: 300 }}
			transition:fly={{ x: -50, duration: 300 }}
			class="flex min-h-[50px] max-w-[360px] flex-col justify-between overflow-hidden text-white shadow-xl {colors[
				t.type
			]}"
		>
			<div class="flex items-center justify-between px-4 py-3 text-sm">
				<span class="font-medium">{t.message}</span>
				<button
					class="ml-3 flex size-7 cursor-pointer items-center justify-center rounded-full transition-colors hover:bg-white/20"
					onclick={() => toast.remove(t.id)}
				>
					<X size={16} />
				</button>
			</div>

			<div class="h-1 w-full bg-black/10">
				<div class="h-full bg-white/60" style="width: {getProgress(t)}%"></div>
			</div>
		</div>
	{/each}
</div>
