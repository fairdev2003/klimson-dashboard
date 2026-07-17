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
	class="absolute top-0 left-0 flex h-[80px] z-110 w-full items-center justify-between border-b border-neutral-700 bg-neutral-900 px-4 text-[14px] text-white"
>
	{#if Dashboard.state.dashboard_loader_on}
		<div
			out:blur={{ duration: 300 }}
			class="bg-blue-500 duration-1000 bottom-0 h-0.5 left-0 absolute transition-all"
			style="width: {Dashboard.state.loaded_percent}%;"
		></div>
	{/if}

	<!-- Lewa strona -->
	<div class="flex flex-1 items-center min-w-0">
		<!-- Logo (tylko desktop) -->
		<div class="items-center shrink-0 md:hidden hidden lg:flex w-40 justify-center">
			<span class="text-blue-500 mr-2">
				<Icon icon="fluent:brain-20-filled" width="35" height="35" />
			</span>
			{@render CMSTextLogo()}
		</div>

		<!-- Przycisk menu (mobile/desktop toggle) -->
		<button
			onclick={toggleSidebar}
			class="p-2 hover:bg-neutral-800 rounded-md flex lg:hidden transition-colors"
		>
			<Icon icon="material-symbols:menu" width="30" height="30" />
		</button>

		<!-- Wyszukiwarka (min-w-0 pozwala jej się kurczyć) -->
		<div class="flex-1 max-w-sm min-w-0">
			<SearchNavbarBox />
		</div>
	</div>

	<!-- Prawa strona -->
	<div class="flex items-center gap-3 shrink-0">
		<PowerOffDashboard />
	</div>
</div>

{#snippet CMSTextLogo(contributor: boolean = false)}
	<a href="/dashboard" class="cursor-pointer py-3 font-semibold text-blue-500 select-none mr-4">
		{#if contributor}
			<p>Harcquiz</p>
			<p>for Contr.</p>
		{:else}
			<p>klimson.dev</p>
			<p>station</p>
		{/if}
	</a>
{/snippet}
