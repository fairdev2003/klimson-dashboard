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
	import SettingsPage from './SettingsPage.svelte';
	import { dashboard_config } from '$lib/dashboard/stores/persist';
	import { api } from '$lib/api/api';
	import { toast } from '$lib/dashboard/stores/toast';
	import axios from 'axios';
	import { debug } from '$lib/terminal/logic';

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
	let toolbarName = $derived('Settings: ' + current_page?.name);

	type Props = {
		opened?: boolean;
	};

	let { opened = $bindable(false) }: Props = $props();

	const settings_page_open = useLocalStorage<SettingKey>('settings_page_open', 'main');
</script>

<RDBModal
	bind:opened
	bg_color="classic"
	size="window"
	padding_preset="normal"
	border="borderless"
	bind:title={toolbarName}
>
	<!-- category container -->
	<SettingsPage />
</RDBModal>

<style>
	@import 'tailwindcss';

	.selected {
		@apply text-blue-400 bg-blue-500/20 hover:bg-blue-500/40;
	}
</style>
