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
	import { blur } from 'svelte/transition';
</script>

{#if $dashboard_config.dock}
	<div
		in:blur={{ duration: 300 }}
		out:blur={{ duration: 300 }}
		class="m-3 h-17 p-4 flex rounded-xl items-center sticky justify-between border-neutral-700/60 bg-neutral-900"
	>
		<div class="flex h-full">
			<div class="flex flex-col justify-end">
				{@render Tabs()}
			</div>
			<div class="flex items-center gap-5">
				<div class="flex gap-2">
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
					<EditingFile />
				</div>
			</div>
		</div>
	</div>
{/if}

{#snippet Tabs()}{/snippet}
