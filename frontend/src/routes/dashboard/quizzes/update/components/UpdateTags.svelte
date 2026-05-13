<script lang="ts">
	import Button from '$lib/components/Button.svelte';

	type Props = {
		tags?: string[];
	};
	let openedtag: string = $state('');
	let { tags = $bindable(['siema']) }: Props = $props();
</script>

<div>
	<p class="mb-1">Tagi</p>

	<div class="flex gap-1">
		{#each tags as tag, i}
			<div class="relative">
				<div
					class:flex={openedtag === tag}
					class:hidden={openedtag !== tag}
					class="absolute -top-10 bg-primary text-white px-2 py-1 rounded p-2"
				>
					<input
						type="text"
						class="px-2 py-1 rounded text-black"
						value={tag}
						on:input={(e) => {
							tags[i] = (e.target as HTMLInputElement).value;
						}}
					/>
					<Button
						size="small"
						theme="danger"
						onclick={() => {
							tags.splice(i, 1);
							if (openedtag === tag) openedtag = '';
						}}
					>
						Delete
					</Button>
				</div>
				<Button
					onclick={() => {
						openedtag = openedtag === tag ? '' : tag;
					}}
					size="small">{tag}</Button
				>
			</div>
		{/each}
		<Button size="small" onclick={() => tags.push('New Tag')}>Add Tag</Button>
	</div>
</div>
