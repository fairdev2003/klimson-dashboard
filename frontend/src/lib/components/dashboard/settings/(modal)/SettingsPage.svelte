<script lang="ts">
	import RDBModal from '$lib/components/modal/RDBModal.svelte';
	import Heading from '../../typography/Heading.svelte';

	import AccountSettings from '$lib/components/dashboard/settings/categories/AccountSettings.svelte';
	import CustomizationSettings from '$lib/components/dashboard/settings/categories/CustomizationSettings.svelte';
	import MainSettings from '$lib/components/dashboard/settings/categories/MainSettings.svelte';
	import ServerSettings from '$lib/components/dashboard/settings/categories/ServerSettings.svelte';
	import { type Setting, type SettingKey } from '$lib/components/dashboard/settings/store.svelte';
	import { blur } from 'svelte/transition';
	import SidebarUserLogged from '../../SidebarUserLogged.svelte';
	import { useLocalStorage } from '@ariefsn/svelte-use';

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

	let current_page = $derived(settings.find((e) => e.slug === settings_page_open.value));

	type Props = {
		opened?: boolean;
	};

	let { opened = $bindable(false) }: Props = $props();

	const settings_page_open = useLocalStorage<SettingKey>('settings_page_open', 'main');
</script>

<div class="flex flex-col justify-center grid-cols-10 lg:grid w-full">
	<div class="col-span-2 mb-10">
		<div class="mt-5 flex flex-col gap-3">
			<SidebarUserLogged
				selected={settings_page_open.value === 'account'}
				onclick={() => {
					settings_page_open.set('account');
				}}
				name="Klimson"
				role="$root"
				pfp_logo="https://klimson.dev/_app/immutable/assets/klimson.CQA0gh-5.jpeg"
			/>
			{#each settings as { slug, name }}
				<button
					onclick={() => {
						settings_page_open.set(slug);
					}}
					class:selected={settings_page_open.value === slug}
					class="text-start p-3 px-4 rounded-xl bg-background hover:bg-foreground text-sm text-text font-semibold cursor-pointer transition-colors"
					>{name}
				</button>
			{/each}
		</div>
	</div>

	<!-- settings container -->
	<div class="col-span-8 lg:p-10 flex flex-col gap-5 px-20">
		{#key current_page}
			<h1 in:blur={{ duration: 300 }} class="text-3xl font-bold text-text">
				{current_page?.name}
			</h1>
		{/key}
		<div class="max-h-[60vh]">
			{#each settings as setting}
				{#if setting.slug === settings_page_open.value}
					<setting.component />
				{/if}
			{/each}
		</div>
	</div>
</div>

<style>
	@import 'tailwindcss';

	.selected {
		background-color: var(--color-primary);
	}
</style>
