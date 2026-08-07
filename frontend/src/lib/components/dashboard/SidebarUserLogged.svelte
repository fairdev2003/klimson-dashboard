<script lang="ts">
	import { goto } from '$app/navigation';
	import Icon from '@iconify/svelte';
	import { useLocalStorage } from '@ariefsn/svelte-use';

	type Props = {
		name: string;
		role: string;
		pfp_logo: string;
		onclick?: () => void;
		selected: boolean;
	};

	let { name, role, pfp_logo, onclick, selected }: Props = $props();

	const nickname = useLocalStorage('nickname', 'Klimson');
</script>

<button
	{onclick}
	class:selected
	class="bg-foreground hover:bg-neutral-700 cursor-pointer gap-3 w-full flex items-center p-2 rounded-lg"
>
	<img
		alt="pfp"
		src="https://api.klimson.dev/interface/bucket/klimson.dev/mango.jpg"
		class="size-12 rounded-full"
	/>
	<div class="flex flex-col h-full justify-center gap-0.5 items-center">
		<p class="text-text text-sm text-start">
			{nickname.value}
			<span class="bg-orange-800 p-0.5 text-orange-300 rounded-full px-2 text-xs">{role}</span>
		</p>
		<span class="flex gap-1 items-center text-start text-xs text-secondary-text cursor-pointer">
			<Icon icon="material-symbols:edit" width="16" height="16" class="mb-0.5" />
			<p class="text-xs">Edit yourself</p>
		</span>
	</div>
</button>
