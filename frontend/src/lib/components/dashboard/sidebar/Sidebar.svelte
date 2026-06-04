<script lang="ts">
	import Icon from '@iconify/svelte';
	import MovingTooltip from '../MovingTooltip.svelte';
	import { slide } from 'svelte/transition';
	import { goto } from '$app/navigation';
	import SidebarUserLogged from '../SidebarUserLogged.svelte';
	import { page } from '$app/state';
	import { isMobile, route } from '$lib/dashboard/stores/persist';
	import type { SidebarItems } from './sidebar.types';
	import SidebarItem from './SidebarItem.svelte';
	import { sidebar_open } from '$lib/dashboard/stores/store';
	import { get } from 'svelte/store';
	import { onMount } from 'svelte';

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
		{ icon: 'material-symbols:home', href: '/dashboard', name: 'Main Page', disabled: false },
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
			icon: 'mdi:spotify',
			href: '/dashboard/spotify',
			name: 'Spotify',
			disabled: false
		},

		{
			icon: 'mdi:files',
			href: '/dashboard/storage',
			name: 'File Storage',
			disabled: false
		},

		{
			icon: 'mdi:image',
			href: '/dashboard/images',
			name: 'Image Explorer (useless)',
			disabled: true
		},
		{
			icon: 'at-icons:gun',
			href: '/dashboard/pg3d',
			name: 'PG3D [ ARCHIVED ]',
			disabled: true
		},
		{
			icon: 'mdi:link',
			href: '/dashboard/routes',
			name: 'API Routes',
			disabled: true
		}
	];
</script>

<div class="sticky left-0 top-16.25 h-screen z-500 overflow-hidden flex flex-col m-5">
	<SidebarUserLogged
		name="cwel"
		role="$root"
		pfp_logo="https://api.klimson.dev/storage/interface/random/banana.webp"
	/>

	{#if $sidebar_open}
		<nav class="flex flex-col" in:slide={{ duration: 300 }} out:slide={{ duration: 300 }}>
			{#each contents as content, i}
				<SidebarItem {content} />
			{/each}
		</nav>
	{/if}
</div>
