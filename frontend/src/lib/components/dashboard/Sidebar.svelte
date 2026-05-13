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
</script>

<div class="sticky top-[65px] h-[calc(100vh-65px)] overflow-hidden flex flex-col gap-2 m-5">
	<div class="bg-neutral-800 border-neutral-700 gap-3 border-1 w-full flex items-center p-2">
		<img
			alt="pfp"
			src="https://api.klimson.dev/storage/interface/random/banana.webp"
			class="size-12 rounded-full"
		/>
		<div class="flex flex-col h-full justify-center">
			<p class="text-neutral-400 text-sm text-start">
				Paweł Cyngot <span class="bg-orange-800 p-0.5 text-orange-300 rounded-full px-2 text-xs"
					>$root</span
				>
			</p>
			<a class="text-xs link hover:underline cursor-pointer" href="/dashboard/settings">Opcje</a>
		</div>
	</div>
	<button
		class="flex gap-2 items-center focus:outline-none"
		onclick={() => {
			content_show = !content_show;
		}}
	>
		<div class="transition-all duration-300" class:rotate-180={content_show}>
			<Icon
				class="text-blue-500 rotate-180"
				icon="fluent:triangle-down-32-filled"
				width="13"
				height="13"
			/>
		</div>
		<p class="text-blue-500 font-semibold">Spis tresci</p>
	</button>
	{#if content_show}
		<nav class="flex flex-col gap-2" in:slide={{ duration: 300 }} out:slide={{ duration: 300 }}>
			{#each $sidebar_content as item (item.id)}
				<div class="flex flex-col ml-5">
					<div class="flex gap-3 items-center">
						<!-- <div class="absolute -left-[1px] w-4 border-1 border-white"></div> -->
						<Icon icon={item.icon} class="text-neutral-500" />

						<MovingTooltip>
							{#snippet tooltipContent()}
								<p class="text-bg-neutral-400 text-sm max-w-40">
									{item.desc}
								</p>
							{/snippet}
							<a href={item.link} class="text-neutral-500 font-medium">{item.label}</a>
						</MovingTooltip>
					</div>

					{#if item.children.length > 0}
						<div class="ml-7 border-l border-neutral-500 mt-2 flex flex-col">
							{#each item.children as child (child.id)}
								<div class="relative flex items-center h-8">
									<div class="absolute -left-[1px] w-4 border-t border-neutral-500"></div>
									<MovingTooltip>
										{#snippet tooltipContent()}
											<p class="text-bg-neutral-400 text-sm max-w-40">
												{child.desc}
											</p>
										{/snippet}
										<a
											href={child.link}
											class="ml-6 text-sm opacity-70 hover:opacity-100 transition-opacity"
										>
											{child.label}
										</a>
									</MovingTooltip>
								</div>
							{/each}
						</div>
					{/if}
				</div>
			{/each}
		</nav>
	{/if}
	<p>Siema</p>
</div>

<style>
	@import 'tailwindcss';

	a {
		@apply text-neutral-400;
	}

	.link {
		@apply text-blue-500;
	}
</style>
