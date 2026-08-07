<script lang="ts">
	import AccountSettings from '$lib/components/dashboard/settings/categories/AccountSettings.svelte';
	import CustomizationSettings from '$lib/components/dashboard/settings/categories/CustomizationSettings.svelte';
	import MainSettings from '$lib/components/dashboard/settings/categories/MainSettings.svelte';
	import ServerSettings from '$lib/components/dashboard/settings/categories/ServerSettings.svelte';
	import {
		settings_page_open,
		type Setting,
		type SettingKey
	} from '$lib/components/dashboard/settings/store.svelte';
	import Heading from '$lib/components/dashboard/typography/Heading.svelte';
	import RDBModal from '$lib/components/modal/RDBModal.svelte';
	import type { Component, Snippet } from 'svelte';
	import { blur } from 'svelte/transition';

	const settings: Setting[] = [
		{ component: MainSettings, description: '', name: 'Main Settings', slug: 'main' },
		{ component: ServerSettings, description: '', name: 'Server Settings', slug: 'server' },
		{
			component: CustomizationSettings,
			description: '',
			name: 'Customization',
			slug: 'customization'
		},
		{ component: AccountSettings, description: '', name: 'Account', slug: 'account' }
	];

	const current_page = $derived(settings.find((e) => e.slug === $settings_page_open)?.name);
</script>

<!-- category container -->
<div class="col-span-2 lg:p-10 mb-10 sticky top-0">
	<h1 class="font-bold">
		<Heading>Categories</Heading>
	</h1>
	<div class="mt-5 flex flex-col gap-3">
		{#each settings as { slug, name }}
			<button
				onclick={() => {
					$settings_page_open = slug;
				}}
				class:selected={$settings_page_open === slug}
				class="text-start p-3 px-4 rounded-xl bg-neutral-800 hover:bg-neutral-700 text-sm font-semibold cursor-pointer transition-colors"
				>{name}
			</button>
		{/each}
	</div>
</div>

<!-- settings container -->
<div class="col-span-8 lg:p-10 flex flex-col gap-5">
	{#key current_page}
		<h1 in:blur={{ duration: 300 }} class="text-3xl text-text font-bold">
			{current_page}
		</h1>
	{/key}
	<div class="">
		{#each settings as setting}
			{#if setting.slug === $settings_page_open}
				<setting.component />
			{/if}
		{/each}
	</div>
</div>

<style>
	@import 'tailwindcss';

	.selected {
		@apply text-blue-400 bg-blue-500/20 hover:bg-blue-500/40;
	}
</style>
