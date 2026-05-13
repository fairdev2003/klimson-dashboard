<script lang="ts">
	import { goto } from '$app/navigation';
	import { ChevronRight } from '@lucide/svelte';
	import { blogForm } from '$lib/dashboard/stores/blog';
	import type { HeroType } from '../types';
	import { heroForm } from '$lib/dashboard/stores/hero';

	type Props = {
		hero_record: HeroType;
	};

	const { hero_record }: Props = $props();
</script>

<button
	class="bg-primary-add-button hover:bg-secondary-add-button flex size-12 cursor-pointer items-center justify-center rounded-lg text-[10px]"
	onclick={async (e) => {
		heroForm.set(hero_record);
		goto('/dashboard/hero/update');
	}}><ChevronRight /></button
>
