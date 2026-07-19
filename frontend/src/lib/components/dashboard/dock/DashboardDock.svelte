<script lang="ts">
	import { sidebar_open, summary_open } from '$lib/dashboard/stores/store';
	import { ChevronLeft, ChevronRight, Dot, Grip } from '@lucide/svelte';
	import { page } from '$app/stores';
	import Loader from '../Loader.svelte';
	import { dashboard_config, route } from '$lib/dashboard/stores/persist';
	import Icon from '@iconify/svelte';
	import { base_url } from '$lib/api/api.store';
	import { toast } from '$lib/dashboard/stores/toast';
	import EditingFile from './boxes/EditingFile.svelte';
	import { goto } from '$app/navigation';
	import { blur } from 'svelte/transition';
	import Modal from '$lib/components/Modal.svelte';
	import { Dashboard } from '$lib/dashboard/logic';

	let bookmarkModalOpened: boolean = $state(false);
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
{#if $dashboard_config.dock}
	<div
		style="position: sticky; top: 10px;"
		class="m-3 h-17 p-4 flex justify-between items-center rounded-xl z-100 inset-0 border-neutral-700/60 bg-neutral-900"
	>
		<div class="flex h-full items-center">
			<div class="flex items-center gap-5">
				<div class="flex justify-center gap-2 items-center">
					<div
						onclick={() => {
							window.history.back();
						}}
						class="size-10 rounded-lg hover:bg-neutral-700 flex items-center justify-center"
					>
						<Icon icon="mingcute:left-line" width="30" height="30" />
					</div>
					<div
						onclick={() => {
							window.history.back();
						}}
						class="size-10 rounded-lg hover:bg-neutral-700 flex items-center justify-center"
					>
						<Icon icon="mingcute:right-line" width="30" height="30" />
					</div>
					<div
						onclick={() => {
							window.location.reload();
						}}
						class="size-10 rounded-lg hover:bg-neutral-700 flex items-center justify-center"
					>
						<Icon icon="zondicons:reload" width="25" height="25" />
					</div>
					<Dashboard.state.dockComponent />
				</div>
			</div>
		</div>
		<div
			onclick={() => {
				goto('/dashboard/settings');
			}}
			class="size-10 rounded-lg hover:bg-neutral-700 flex items-center justify-center"
		>
			<Icon icon="material-symbols:settings" width="25" height="25" />
		</div>
	</div>
{/if}

{#snippet Tabs()}{/snippet}
