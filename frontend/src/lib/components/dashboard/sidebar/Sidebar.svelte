<script lang="ts">
	import {
		addChildToQuizzes,
		addFormQuiz,
		sidebar_content,
		sidebar_menu_static
	} from '$lib/dashboard/stores/store';
	import Icon from '@iconify/svelte';
	import MovingTooltip from '../MovingTooltip.svelte';
	import { slide } from 'svelte/transition';
	import { goto } from '$app/navigation';
	import SidebarUserLogged from '../SidebarUserLogged.svelte';
	import { page } from '$app/state';
	import { route } from '$lib/dashboard/stores/persist';
	import type { SidebarItems } from './sidebar.types';
	import SidebarItem from './SidebarItem.svelte';

	let content_show: boolean = $state(true);

	$effect(() => {
		addChildToQuizzes(
			`Aktywny formularz`,
			'/dashboard/quizzes/update',
			'Formularz twojego quizu',
			() => {
				// siurek
			}
		);
	});

	type SContent = {
		icon: string;
		href: string;
		name: string;
	};

	$effect(() => {
		console.log(page.url.pathname);
	});

	const contents: SidebarItems = [
		{ icon: 'material-symbols:home', href: '/dashboard', name: 'Main Page' },
		{
			icon: 'material-symbols:database',
			href: '/dashboard/database',
			name: 'Database Editor'

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
			name: 'Context Storage'
		},
		{
			icon: 'pinhead:pixel-circle',
			href: '/dashboard/pg3d',
			name: 'Pixel Gun 3D'
		},
		{
			icon: 'mdi:image',
			href: '/dashboard/images',
			name: 'Image Explorer'
		},
		{
			icon: 'mdi:link',
			href: '/dashboard/routes',
			name: 'API Routes'
		}
	];
</script>

<div class="sticky top-16.25 h-[calc(100vh-65px)] overflow-hidden flex flex-col m-5">
	<SidebarUserLogged
		name="cwel"
		role="$root"
		pfp_logo="https://api.klimson.dev/storage/interface/random/banana.webp"
	/>

	{#if content_show}
		<nav class="flex flex-col" in:slide={{ duration: 300 }} out:slide={{ duration: 300 }}>
			{#each contents as content, i}
				<SidebarItem {content} />
			{/each}
		</nav>
	{/if}
</div>
