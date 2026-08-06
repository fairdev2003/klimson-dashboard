<script lang="ts">
	import SearchNavbarBox from './SearchNavbarBox.svelte';
	import UserSettings from './UserSettings.svelte';
	import ContextMenu from './HashMenu.svelte';
	import SessionExpireCounter from './SessionExpireCounter.svelte';
	import type { Attachment } from 'svelte/attachments';
	import Icon from '@iconify/svelte';
	import QuickActionButton from './QuickActionButton.svelte';
	import { userInfo } from '$lib/dashboard/stores/store';
	import { isMobile, mobile_sidebar_open, sidebar_open } from '$lib/dashboard/stores/persist';
	import PowerOffDashboard from './navbar/PowerOffDashboard.svelte';
	import { Dashboard } from '$lib/dashboard/logic';
	import { blur } from 'svelte/transition';
	import Logo from '$lib/assets/klimson_logo.svg';
	import Notification from './navbar/NotificationPanel.svelte';
	import Pixeletedsraka from '$lib/assets/Pixeletedsraka.svelte';
	import { goto } from '$app/navigation';
	import { page } from '$app/state';

	const SendHeightInfo: Attachment<HTMLDivElement> = (element) => {
		return () => {
			console.log('Detached element:', element.clientHeight);
		};
	};

	function toggleSidebar() {
		if ($isMobile) {
			$mobile_sidebar_open = !$mobile_sidebar_open;
			if (!$mobile_sidebar_open) {
				document.body.style.overflow = 'hidden';
				return;
			}
			document.body.style.overflow = 'auto';
		} else {
			$sidebar_open = !$sidebar_open;
		}
	}
</script>

<div
	class="absolute top-0 left-0 flex h-[80px] z-110 w-full items-center justify-between border-b border-neutral-700 bg-neutral-900 px-5 text-[14px] text-white"
>
	{#if Dashboard.state.dashboard_loader_on}
		<div
			out:blur={{ duration: 300 }}
			class="bg-blue-500 duration-1000 top-0 h-0.5 left-0 absolute transition-all"
			style="width: {Dashboard.state.loaded_percent}%;"
		></div>
	{/if}

	<!-- Lewa strona -->
	<div class="flex flex-1 items-center min-w-0">
		<!-- Logo (tylko desktop) -->
		{@render CMSTextLogo()}

		<!-- Wyszukiwarka (min-w-0 pozwala jej się kurczyć) -->
		<div class="flex-1 max-w-sm min-w-0">
			<SearchNavbarBox />
		</div>
	</div>

	<!-- Prawa strona -->
	<div class="flex items-center gap-6 mr-5 shrink-0">
		<Notification />
		<PowerOffDashboard />
	</div>
</div>

{#snippet CMSTextLogo()}
	<a class="pr-5 text-blue-500 pb-2 ml-5" href="/dashboard">
		<Icon icon="dinkie-icons:mango" width="40" height="40" />
	</a>
{/snippet}
