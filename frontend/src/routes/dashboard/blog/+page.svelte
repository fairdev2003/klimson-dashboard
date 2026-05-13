<script lang="ts">
	import { goto } from '$app/navigation';
	import Button from '$lib/components/Button.svelte';
	import { blogs } from '$lib/dashboard/stores/data.store';
	import { onMount } from 'svelte';

	import { blogFormState } from '$lib/dashboard/stores/blog';
	import BlogRecord from './components/BlogRecord.svelte';

	onMount(() => {
		console.log($blogs);
	});
</script>

<div class="py-5 text-white">
	{@render QuizSectionHeading()}

	<div>
		{#each $blogs as blog_record}
			<BlogRecord {blog_record} />
		{/each}
	</div>
</div>

{#snippet QuizSectionHeading()}
	<div class="mb-5 flex items-center justify-between gap-5">
		<div class="flex items-center gap-3">
			<h1 class="text-2xl font-bold">Blog</h1>
		</div>
		<div class="flex items-center justify-center gap-2">
			<Button
				onclick={() => {
					blogFormState.set('add');

					goto('/dashboard/blog/add');
				}}
				size="medium"
				theme="secondary"
			>
				<p>Dodaj</p></Button
			>
		</div>
	</div>
{/snippet}
