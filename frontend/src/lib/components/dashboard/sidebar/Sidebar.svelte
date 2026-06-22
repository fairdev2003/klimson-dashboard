<script lang="ts">
	import Icon from '@iconify/svelte';
	import MovingTooltip from '../MovingTooltip.svelte';
	import { slide } from 'svelte/transition';
	import { goto } from '$app/navigation';
	import SidebarUserLogged from '../SidebarUserLogged.svelte';
	import { page } from '$app/state';
	import { isMobile, mobile_sidebar_open, route } from '$lib/dashboard/stores/persist';
	import { contents, type SidebarItems } from './sidebar.types';
	import SidebarItem from './SidebarItem.svelte';
	import { sidebar_open } from '$lib/dashboard/stores/store';
	import { get } from 'svelte/store';
	import { onMount } from 'svelte';
	import { nickname } from '../settings/store.svelte';
	import SidebarPill from './SidebarPill.svelte';
	import PagesSelector from './PagesSelector.svelte';
	import SidebarToggler from './SidebarToggler.svelte';

	let content_show: boolean = $state(true);

	export function setupMedia() {
		if (typeof window === 'undefined') return;

		const mql = window.matchMedia('(max-width: 768px)'); // 768px to domyślny 'md' w Tailwind
		isMobile.set(mql.matches);

		mql.addEventListener('change', (e) => {
			isMobile.set(e.matches);
		});
	}

	onMount(() => {
		setupMedia();
	});

	type SContent = {
		icon: string;
		href: string;
		name: string;
	};

	$effect(() => {
		if ($mobile_sidebar_open) {
			document.body.style.overflow = 'hidden';
		} else {
			document.body.style.overflow = '';
		}
	});
</script>

<div class="sticky left-0 top-16.25 h-dvh z-500 overflow-hidden flex flex-col m-5">
	<div>
		<SidebarToggler />
		<SidebarPill />
		<PagesSelector />

		{#if $sidebar_open}
			<nav
				class="flex flex-col gap-2 mt-3"
				in:slide={{ duration: 300 }}
				out:slide={{ duration: 300 }}
			>
				{#each contents as content, i}
					<SidebarItem {content} />
				{/each}
			</nav>
		{/if}
	</div>
	<!-- svelte-ignore a11y_consider_explicit_label -->
</div>
