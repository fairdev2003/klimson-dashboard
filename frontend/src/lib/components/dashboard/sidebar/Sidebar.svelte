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
			icon: 'mdi:link',
			href: '/dashboard/routes',
			name: 'API Routes',
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

<div class="sticky left-0 top-16.25 h-dvh z-500 overflow-hidden flex flex-col justify-between m-5">
	<div>
		<SidebarUserLogged
			name={$nickname}
			role="$root"
			pfp_logo="https://klimson.dev/_app/immutable/assets/klimson.CQA0gh-5.jpeg"
		/>

		{#if $sidebar_open}
			<nav class="flex flex-col" in:slide={{ duration: 300 }} out:slide={{ duration: 300 }}>
				{#each contents as content, i}
					<SidebarItem {content} />
				{/each}
			</nav>
		{/if}
	</div>
	<!-- svelte-ignore a11y_consider_explicit_label -->
	<button
		onclick={() => {
			goto('/dashboard/settings');
		}}
		class="mb-25 bg-neutral-800/60 hover:bg-neutral-800 border border-neutral-700 h-12 w-full active:bg-neutral-800 cursor-pointer transition-colors duration-150"
	></button>
</div>
