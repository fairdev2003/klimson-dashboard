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

	const SendHeightInfo: Attachment<HTMLDivElement> = (element) => {
		return () => {
			console.log('Detached element:', element.clientHeight);
		};
	};

	function toggleSidebar() {
		if ($isMobile) {
			$mobile_sidebar_open = !$mobile_sidebar_open;
		} else {
			$sidebar_open = !$sidebar_open;
		}
	}
</script>

<div
	class="fixed top-0 left-0 z-50 flex h-[66px] w-full items-center justify-between border-b border-neutral-700 bg-neutral-900 px-4 text-[14px] text-white"
>
	<!-- Lewa strona -->
	<div class="flex flex-1 items-center gap-4 min-w-0">
		<!-- Logo (tylko desktop) -->
		<div class="items-center shrink-0 md:hidden hidden lg:flex">
			<span class="text-blue-500 mr-2">
				<Icon icon="majesticons:lightning-bolt" width="30" height="30" />
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
		<div class="hidden lg:flex">
			<SessionExpireCounter />
		</div>
		<UserSettings />
	</div>
</div>

{#snippet CMSTextLogo(contributor: boolean = false)}
	<a href="/dashboard" class="cursor-pointer py-3 font-semibold text-blue-500 select-none mr-4">
		{#if contributor}
			<p>Harcquiz</p>
			<p>for Contr.</p>
		{:else}
			<p>Klimson</p>
			<p>CMS</p>
		{/if}
	</a>
{/snippet}
