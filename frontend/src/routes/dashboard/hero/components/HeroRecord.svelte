<script lang="ts">
	import { api } from '$lib/api/api';
	import { blogForm } from '$lib/dashboard/stores/blog';
	import { heroForm } from '$lib/dashboard/stores/hero';
	import { get } from 'svelte/store';
	import type { HeroType } from '../types';
	import HeroDeleteButton from './HeroDeleteButton.svelte';
	import HeroEnterButton from './HeroEnterButton.svelte';
	type Props = {
		hero_record: HeroType;
	};

	const { hero_record }: Props = $props();
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
<!-- svelte-ignore a11y_click_events_have_key_events -->
<div
	onclick={() => {
		heroForm.set(hero_record);
		console.log(get(heroForm));
	}}
	class="
		group relative flex h-[120px] cursor-pointer overflow-hidden
		border border-white/5
		transition-all duration-200
	"
	class:bg-neutral-800={hero_record.id !== $heroForm.id}
	class:bg-neutral-600={hero_record.id === $heroForm.id}
	class:ring-2={hero_record.id === $heroForm.id}
	class:ring-primary={hero_record.id === $heroForm.id}
>
	<!-- IMAGE -->
	<div class="h-full w-[140px] flex-shrink-0 overflow-hidden">
		<img
			src={api.image.hero(hero_record.image_url)}
			alt=""
			class="h-full w-full object-cover transition-transform duration-300 group-hover:scale-105"
		/>
	</div>

	<!-- CONTENT -->
	<div class="flex h-full w-full items-center justify-between px-4">
		<div class="flex flex-col gap-1">
			<h3 class="text-sm font-semibold leading-tight">
				{hero_record.quote}
			</h3>

			<p class="line-clamp-2 text-xs opacity-75">
				{hero_record.author}
			</p>
		</div>

		<!-- ACTIONS -->
		{#if hero_record.id === $heroForm.id}
			<div class="flex gap-2">
				<HeroDeleteButton />
				<HeroEnterButton {hero_record} />
			</div>
		{/if}
	</div>
</div>
