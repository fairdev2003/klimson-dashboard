<script lang="ts">
	import { nickname } from './settings/store';

	type Props = {
		name: string;
		role: string;
		pfp_logo: string;
	};

	let { name, role, pfp_logo }: Props = $props();
</script>

<div class="bg-neutral-800 border-neutral-700 gap-3 border-1 w-full flex items-center p-2">
	<img alt="pfp" src={pfp_logo} class="size-12 rounded-full" />
	<div class="flex flex-col h-full justify-center">
		<p class="text-neutral-400 text-sm text-start">
			Jakub
			<span class="bg-orange-800 p-0.5 text-orange-300 rounded-full px-2 text-xs">{role}</span>
		</p>
		<a class="text-xs link hover:underline cursor-pointer" href="/dashboard/settings">Opcje</a>
	</div>
</div>

<style>
	@import 'tailwindcss';

	.link {
		@apply text-blue-500;
	}
</style>
