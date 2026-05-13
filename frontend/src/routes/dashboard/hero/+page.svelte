<script>
	import { goto } from '$app/navigation';
	import Button from '$lib/components/Button.svelte';
	import Check from '$lib/components/dashboard/Check.svelte';
	import { heros } from '$lib/dashboard/stores/data.store';
	import HeroRecord from './components/HeroRecord.svelte';
</script>

<div class="py-5 text-white">
	{@render HeroSectionHeading()}

	<div>
		{#each $heros as hero_record}
			<HeroRecord {hero_record} />
		{/each}
	</div>
</div>

{#snippet HeroSectionHeading()}
	<div class="mb-5 flex items-center justify-between gap-5">
		<div class="flex items-center gap-3">
			<h1 class="text-2xl font-bold">Hero</h1>
		</div>
		<div class="flex items-center justify-center gap-2">
			<Button
				onclick={() => {
					goto('/dashboard/hero/add');
				}}
				size="medium"
				theme="secondary"
			>
				<p>Dodaj</p></Button
			>
		</div>
	</div>
{/snippet}
