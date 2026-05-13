<script lang="ts">
	import { blogForm } from '$lib/dashboard/stores/blog';
	import DeleteButton from '../../quizzes/components/DeleteButton.svelte';
	import EnterButton from '../../quizzes/components/EnterButton.svelte';
	import type { BlogType } from '../types';
	import BlogDeleteButton from './BlogDeleteButton.svelte';
	import BlogEnterButton from './BlogEnterButton.svelte';

	type Props = {
		blog_record: BlogType;
		selected?: boolean;
	};

	const { blog_record, selected = false }: Props = $props();
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
<!-- svelte-ignore a11y_click_events_have_key_events -->
<div
	onclick={() => {
		blogForm.set(blog_record);
	}}
	class="
		group relative flex h-[60px] cursor-pointer overflow-hidden
		border border-white/5
		transition-all duration-200
	"
	class:bg-neutral-800={blog_record.id !== $blogForm.id}
	class:bg-neutral-600={blog_record.id === $blogForm.id}
	class:ring-2={blog_record.id === $blogForm.id}
	class:ring-primary={blog_record.id === $blogForm.id}
>
	<!-- CONTENT -->
	<div class="flex h-full w-full items-center justify-between px-4">
		<div class="flex flex-col gap-1">
			<h3 class="text-sm font-semibold leading-tight">
				{blog_record.title}
			</h3>

			<p class="line-clamp-2 text-xs opacity-75">
				{blog_record.description} | {blog_record.public == true ? 'Publiczny' : 'Nie publiczny'}
			</p>
		</div>

		<!-- ACTIONS -->
		{#if blog_record.id === $blogForm.id}
			<div
				class="
				flex gap-2
				
			"
			>
				<BlogDeleteButton />
				<BlogEnterButton {selected} {blog_record} />
			</div>
		{/if}
	</div>
</div>
