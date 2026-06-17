<script lang="ts">
	import Icon from '@iconify/svelte';
	import MovingTooltip from '../MovingTooltip.svelte';
	import { slide } from 'svelte/transition';
	import { goto } from '$app/navigation';
	import SidebarUserLogged from '../SidebarUserLogged.svelte';
	import { page } from '$app/state';
	import { isMobile, mobile_sidebar_open, route } from '$lib/dashboard/stores/persist';
	import type { SidebarItems } from './sidebar.types';
	import SidebarItem from './SidebarItem.svelte';
	import { sidebar_open } from '$lib/dashboard/stores/store';
	import { get } from 'svelte/store';
	import { onMount } from 'svelte';
	import { nickname } from '../settings/store.svelte';
	import SidebarPill from './SidebarPill.svelte';

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

	const contents: SidebarItems = [
		{ icon: 'material-symbols:home', href: '/dashboard', name: 'Hub', disabled: false },
		{
			icon: 'material-symbols:database',
			href: '/dashboard/database',
			name: 'Database Editor',
			disabled: false

			// child: [
			// 	{
			// 		name: 'Users Database',
			// 		href: '/dashboard/database/users'
			// 	}
			// ]
		},
		{
			icon: 'mdi:bucket',
			href: '/dashboard/context_storage',
			name: 'Context Storage',
			disabled: false
		},

		{
			icon: 'mdi:files',
			href: '/dashboard/storage',
			name: 'File Storage',
			disabled: false
		},
		{
			icon: 'mdi:files',
			href: '/dashboard/v2/storage',
			name: 'V2 Storage',
			disabled: false
		},
		{
			icon: 'mdi:tools',
			href: '/dashboard/tools',
			name: 'Tools',
			disabled: false
		},

		{
			icon: 'mdi:link',
			href: '/dashboard/routes',
			name: 'API Routes',
			disabled: false
		},
		{
			icon: 'streamline-block:content-copy',
			href: '/dashboard/content_manager',
			name: 'Content Manager',
			disabled: true
		}
	];

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
		<SidebarPill />

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
