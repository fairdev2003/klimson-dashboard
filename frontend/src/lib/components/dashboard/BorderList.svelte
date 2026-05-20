<script lang="ts">
	import Icon from '@iconify/svelte';
	import { goto } from '$app/navigation';

	type BorderListItem = {
		name: string;
		onclick?: () => void;
		icon: string;
		href?: string;
		description: string;
	};

	type BorderList = BorderListItem[];

	type Props = {
		border_list: BorderList;
	};

	let { border_list }: Props = $props();
</script>

<div class="flex flex-col">
	{#each border_list as content}
		<!-- svelte-ignore a11y_click_events_have_key_events -->
		<!-- svelte-ignore a11y_no_static_element_interactions -->
		<div
			onclick={() => {
				content.onclick?.();
				if (!content.href) return;
				goto(content.href);
			}}
			class="flex items-center px-3 cursor-pointer transition-colors border-x border-b h-10 gap-3"
		>
			{#if content.icon}
				<Icon icon={content.icon} />
			{/if}
			<p class="text-neutral-300 text-sm">
				{content.name}
			</p>
		</div>
	{/each}
</div>
