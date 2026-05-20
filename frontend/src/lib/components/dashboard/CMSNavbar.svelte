<script lang="ts">
	import SearchNavbarBox from './SearchNavbarBox.svelte';
	import UserSettings from './UserSettings.svelte';
	import ContextMenu from './HashMenu.svelte';
	import SessionExpireCounter from './SessionExpireCounter.svelte';
	import type { Attachment } from 'svelte/attachments';
	import Icon from '@iconify/svelte';
	import QuickActionButton from './QuickActionButton.svelte';
	import { userInfo } from '$lib/dashboard/stores/store';

	const SendHeightInfo: Attachment<HTMLDivElement> = (element) => {
		console.log('Attached element:', element.clientHeight);

		return () => {
			console.log('Detached element:', element);
		};
	};
</script>

<div
	{@attach SendHeightInfo}
	class="fixed top-0 z-50 flex w-full justify-between gap-1 border-b-1 border-neutral-700 bg-neutral-900 px-5 text-[14px] text-white"
>
	<!-- left navbar conten -->
	<div class="flex items-center">
		<div class="flex items-center gap-3">
			<span class="text-blue-500">
				<Icon icon="majesticons:lightning-bolt" width="30" height="30" />
			</span>
			{@render CMSTextLogo($userInfo.contributor)}
		</div>
		<div class="ml-5 flex gap-2">
			<SearchNavbarBox />
			<QuickActionButton />
		</div>
	</div>
	<!-- right navbar content -->
	<div class="flex items-center gap-3">
		<div class="hidden lg:flex">
			<SessionExpireCounter />
		</div>
		<UserSettings />
	</div>
</div>

{#snippet CMSTextLogo(contributor: boolean = false)}
	<a href="/dashboard" class="w-20 cursor-pointer py-3 font-semibold text-blue-500 select-none">
		{#if contributor}
			<p>Harcquiz</p>
			<p>for Contr.</p>
		{:else}
			<p>Klimson</p>
			<p>CMS</p>
		{/if}
	</a>
{/snippet}
