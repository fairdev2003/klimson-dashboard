<script lang="ts">
	import { goto } from '$app/navigation';
	import { ChevronRight } from '@lucide/svelte';
	import { blogForm } from '$lib/dashboard/stores/blog';
	import type { BlogType } from '../types';

	type Props = {
		blog_record: BlogType;
	};

	const { blog_record }: Props = $props();
</script>

<button
	class="bg-primary-add-button hover:bg-secondary-add-button flex size-12 cursor-pointer items-center justify-center rounded-lg text-[10px]"
	onclick={async (e) => {
		blogForm.set(blog_record);
		goto('/dashboard/blog/update');
	}}><ChevronRight /></button
>
