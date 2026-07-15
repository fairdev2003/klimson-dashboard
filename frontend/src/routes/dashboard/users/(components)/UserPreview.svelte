<script lang="ts">
	import { type User } from '$lib/api/types';
	import Icon from '@iconify/svelte';

	type Props = {
		user: User | undefined;
	};

	let { user }: Props = $props();
</script>

{#if user}
	<div>
		<span
			class="text-neutral-400 mb-2 uppercase justify-between items-center flex font-bold text-xs"
		>
			<p>PREVIEW</p>
		</span>
		<div class="bg-transparent justify-between flex rounded-lg p-3 items-center">
			<div class="flex items-center gap-3">
				<img src={user.pfp} class="size-10 rounded-full" />
				<div class="flex flex-col">
					<span class="flex gap-0.5 items-center text-white">
						<p class="font-black hover:underline cursor-pointer">
							{user.nickname}
						</p>
					</span>
					<p class="hover:underline cursor-pointer">
						<!-- pill -->
						<span class="rounded-full px-2 text-xs p-0.5 bg-purple-800 text-purple-300"
							>Nigga fucker</span
						>
					</p>
				</div>
			</div>
		</div>
	</div>
{/if}
