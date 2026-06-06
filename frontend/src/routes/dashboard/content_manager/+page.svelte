<script lang="ts">
	import { goto } from '$app/navigation';
	import Heading from '$lib/components/dashboard/typography/Heading.svelte';
	import Icon from '@iconify/svelte';
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
</script>

<div class="mx-auto w-3xl mt-10">
	<div class="flex flex-col">
		<div class="flex items-center justify-between">
			<Heading>Blog Posts</Heading>
			<button
				onclick={() => {
					window.location.href = '/dashboard/markdown/editor/example.md';
				}}
				class="p-2 hover:bg-neutral-700 rounded-xl cursor-pointer"
			>
				<Icon icon="ic:round-plus" width="30" height="30" />
			</button>
		</div>
		<div class="space w-full my-4 flex flex-col gap-3">
			{#each Array(4) as example_blog_card}
				<button
					class="bg-neutral-800/60 flex p-3 hover:border-blue-500 hover:bg-blue-600/30 cursor-pointer border w-full rounded-xl border-neutral-700 h-18"
				>
					<div class="flex gap-5 items-center">
						<div
							class="size-12 bg-neutral-700/60 border border-neutral-700 items-center justify-center text-neutral-300 flex rounded-lg"
						>
							<Icon icon="mdi:blog-outline" width="30" height="30" />
						</div>
					</div>
				</button>
			{/each}
		</div>
	</div>
	<div class="flex flex-col">
		<div class="flex items-center justify-between">
			<Heading>Markdown</Heading>

			<button onclick={() => {}} class="p-2 hover:bg-neutral-700 rounded-xl cursor-pointer">
				<Icon icon="ic:round-plus" width="30" height="30" />
			</button>
		</div>
		<div class="space w-full my-4 flex flex-col gap-3">
			{#each Array(4) as example_blog_card}
				<button
					onclick={() => {
						goto('/dashboard/markdown/editor/example.md');
					}}
					class="bg-neutral-800/60 flex p-3 hover:border-blue-500 hover:bg-blue-600/30 cursor-pointer border w-full rounded-xl border-neutral-700 h-18"
				>
					<div class="flex gap-5 items-center">
						<div
							class="size-12 bg-neutral-700/60 border border-neutral-700 items-center justify-center text-neutral-300 flex rounded-lg"
						>
							<Icon icon="material-symbols:markdown" width="30" height="30" />
						</div>
					</div>
				</button>
			{/each}
		</div>
	</div>
</div>
