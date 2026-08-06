<script lang="ts">
	import { sidebar_open, summary_open } from '$lib/dashboard/stores/store';
	import { ChevronLeft, ChevronRight, Dot, Grip } from '@lucide/svelte';
	import { page } from '$app/stores';
	import Loader from '../Loader.svelte';
	import { dashboard_config, debugOn, route } from '$lib/dashboard/stores/persist';
	import Icon from '@iconify/svelte';
	import { base_url } from '$lib/api/api.store';
	import { toast } from '$lib/dashboard/stores/toast';
	import EditingFile from './boxes/EditingFile.svelte';
	import { goto } from '$app/navigation';
	import { blur, fly, slide } from 'svelte/transition';
	import Modal from '$lib/components/Modal.svelte';
	import { Dashboard } from '$lib/dashboard/logic';
	import { terminal } from '$lib/terminal/logic';
	import DockMenu from './DockMenu.svelte';
	import DockBackground from './DockBackground.svelte';
	import MobileDock from './MobileDock.svelte';
	import SettingsModal from '../settings/(modal)/SettingsModal.svelte';

	let mobileDockOpened = $state(false);
	let settingsModalOpened = $state(false);

	function toggleMobileDockSection(state?: boolean) {
		if (!state) {
			mobileDockOpened = !mobileDockOpened;
		} else {
			mobileDockOpened = state;
		}

		if (mobileDockOpened) {
			document.documentElement.style.overflow = 'hidden';
			document.body.style.overflow = 'hidden';
		} else {
			document.documentElement.style.overflow = '';
			document.body.style.overflow = '';
		}
	}
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
{#if $dashboard_config.dock}
	<div
		style="position: sticky; top: 10px;"
		class="h-17 p-10 px-5 flex justify-between items-center z-100 inset-0 border-neutral-700/60 bg-neutral-900"
	>
		<div class="flex h-full items-center">
			<div class="flex items-center gap-5">
				<div class="flex justify-center gap-2 items-center">
					{@render Controls()}
					<div class="lg:flex hidden">
						<Dashboard.state.dockComponent />
					</div>
				</div>
			</div>
		</div>
		<div class="flex gap-4">
			<button
				title="Toggle terminal"
				onclick={() => {
					$debugOn = !$debugOn;
				}}
				class:terminal-button-base={!$debugOn}
				class:terminal-button-selected={$debugOn}
				class="size-10 cursor-pointer rounded-lg flex items-center justify-center"
			>
				<Icon icon="ri:terminal-line" width="25" height="25" />
			</button>
			<div
				onclick={() => {
					settingsModalOpened = true;
				}}
				class="size-10 rounded-lg hover:bg-neutral-700 flex items-center justify-center"
			>
				<Icon icon="material-symbols:settings" width="25" height="25" />
			</div>
		</div>
	</div>
	<MobileDock bind:mobileDockOpened />
{/if}

<SettingsModal bind:opened={settingsModalOpened} />

{#snippet Controls()}
	<div>
		<!-- PC & Large Devices View -->
		<div class="lg:flex hidden gap-2 items-center">
			<div
				onclick={() => {
					window.history.back();
				}}
				class="w-10 h-10 min-w-10 min-h-10 rounded-lg hover:bg-neutral-700 flex items-center justify-center cursor-pointer"
			>
				<Icon icon="mingcute:left-line" width="24" height="24" />
			</div>
			<div
				onclick={() => {
					window.history.forward();
				}}
				class="w-10 h-10 min-w-10 min-h-10 rounded-lg hover:bg-neutral-700 flex items-center justify-center cursor-pointer"
			>
				<Icon icon="mingcute:right-line" width="24" height="24" />
			</div>

			<div
				onclick={() => {
					window.location.reload();
				}}
				class="w-10 h-10 min-w-10 min-h-10 rounded-lg hover:bg-neutral-700 flex items-center justify-center cursor-pointer"
			>
				<Icon icon="zondicons:reload" width="25" height="25" />
			</div>
		</div>

		<!-- Mobile View -->
		<button
			onclick={() => {
				toggleMobileDockSection();
			}}
			class="lg:hidden flex active:bg-neutral-400 duration-450 p-2 rounded-full"
		>
			<Icon icon="mdi:hamburger-menu" width="25" height="25" />
		</button>
	</div>
{/snippet}

<style>
	@import 'tailwindcss';

	.terminal-button-base {
		@apply hover:bg-neutral-700 hover:text-white;
	}

	.terminal-button-selected {
		@apply text-white bg-blue-700 hover:bg-blue-600;
	}
</style>
