<script lang="ts">
	import {
		addChildToQuizzes,
		addFormQuiz,
		sidebar_content,
		sidebar_menu_static
	} from '$lib/dashboard/stores/store';
	import Icon from '@iconify/svelte';
	import MovingTooltip from './MovingTooltip.svelte';
	import { slide } from 'svelte/transition';
	import { goto } from '$app/navigation';
	import SidebarUserLogged from './SidebarUserLogged.svelte';
	import { page } from '$app/state';
	import { route } from '$lib/dashboard/stores/persist';

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
		href: string;
		name: string;
	};

	$effect(() => {
		console.log(page.url.pathname);
	});

	const contents: SContent[] = [{ href: '/dashboard/database/witam', name: 'Database Editor' }];
</script>

<div class="sticky top-[65px] h-[calc(100vh-65px)] overflow-hidden flex flex-col m-5">
	<SidebarUserLogged
		name="cwel"
		role="$root"
		pfp_logo="https://api.klimson.dev/storage/interface/random/banana.webp"
	/>

	{#if content_show}
		<nav class="flex flex-col" in:slide={{ duration: 300 }} out:slide={{ duration: 300 }}>
			{#each contents as content, i}
				{@const statement = $route === content.href}
				<div
					onclick={() => {
						goto(content.href);
					}}
					class:normal={!statement}
					class:selected={statement}
					class="flex items-center px-3 cursor-pointer transition-colors border-x border-b h-10 gap-3"
				>
					<Icon icon="material-symbols:database" />
					<p class="text-neutral-300 text-sm">
						{content.name}
					</p>
				</div>
			{/each}
		</nav>
	{/if}
</div>

<style>
	@import 'tailwindcss';

	.link {
		@apply bg-blue-500;
	}

	.selected {
		@apply bg-blue-700/60 hover:bg-blue-700 border-blue-500 border text-white;
	}

	.normal {
		@apply bg-neutral-800 hover:bg-neutral-700 border-neutral-700 text-neutral-500;
	}
</style>
