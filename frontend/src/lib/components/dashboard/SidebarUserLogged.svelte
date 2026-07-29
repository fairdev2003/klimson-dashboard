<script lang="ts">
	import { goto } from '$app/navigation';
	import { nickname } from './settings/store.svelte';

	type Props = {
		name: string;
		role: string;
		pfp_logo: string;
	};

	let { name, role, pfp_logo }: Props = $props();
</script>

<div class="bg-neutral-800 gap-3 w-full flex items-center p-2 rounded-lg">
	<img
		alt="pfp"
		src="https://api.klimson.dev/interface/bucket/klimson.dev/mango.jpg"
		class="size-12 rounded-full"
	/>
	<div class="flex flex-col h-full justify-center">
		<p class="text-neutral-400 text-sm text-start">
			Jakub
			<span class="bg-orange-800 p-0.5 text-orange-300 rounded-full px-2 text-xs">{role}</span>
		</p>
		<a
			onclick={() => {
				goto('/dashboard/settings');
			}}
			class="text-xs link hover:underline cursor-pointer">Opcje</a
		>
	</div>
</div>

<style>
	@import 'tailwindcss';

	.link {
		@apply text-blue-500;
	}
</style>
