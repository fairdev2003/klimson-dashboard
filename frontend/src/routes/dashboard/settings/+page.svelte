<script lang="ts">
	import CustomizationSettings from '$lib/components/dashboard/settings/categories/CustomizationSettings.svelte';
	import MainSettings from '$lib/components/dashboard/settings/categories/MainSettings.svelte';
	import ServerSettings from '$lib/components/dashboard/settings/categories/ServerSettings.svelte';
	import { settings_page_open, type SettingKey } from '$lib/components/dashboard/settings/store';
	import Heading from '$lib/components/dashboard/typography/Heading.svelte';
	import { blur } from 'svelte/transition';

	const settings_dictonary: Record<SettingKey, string> = {
		main: 'Main Settings',
		server: 'Server Configuration',
		customization: 'Dashboard Customization'
	};
</script>

<!-- category container -->
<div class="col-span-2 lg:p-10 mb-10">
	<h1 class="font-bold">
		<Heading>Categories</Heading>
	</h1>
	<div class="mt-5 flex flex-col gap-1">
		{#each Object.keys(settings_dictonary) as SettingKey[] as k}
			<a
				onclick={() => {
					$settings_page_open = k;
				}}
				class:text-blue-500={$settings_page_open === k}
				class="hover:text-blue-500 text-sm font-semibold cursor-pointer transition-colors"
				>{settings_dictonary[k]}</a
			>
		{/each}
	</div>
</div>

<!-- settings container -->
<div class="col-span-8 lg:p-10 flex flex-col gap-5">
	<h1 class="text-3xl font-bold">
		{settings_dictonary[$settings_page_open]}
	</h1>
	<!-- record concept -->

	{#if $settings_page_open == 'main'}
		<MainSettings />
	{/if}

	{#if $settings_page_open == 'server'}
		<ServerSettings />
	{/if}

	{#if $settings_page_open == 'customization'}
		<CustomizationSettings />
	{/if}
</div>

<style>
	@import 'tailwindcss';

	.selected {
		@apply text-blue-500;
	}
</style>
