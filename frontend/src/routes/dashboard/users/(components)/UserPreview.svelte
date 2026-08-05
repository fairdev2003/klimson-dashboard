<script lang="ts">
	import type { User } from '$lib/types/user';
	import Icon from '@iconify/svelte';
	import { blur, slide } from 'svelte/transition';

	type Props = {
		user: User | undefined;
	};

	let { user = $bindable() }: Props = $props();
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
				{#if user.pfp}
					<img src={user.pfp} class="size-10 rounded-full" />
				{:else}
					<div class="size-10 rounded-full bg-neutral-700 flex justify-center items-center">
						<Icon icon="mage:user-question-mark" />
					</div>
				{/if}
				<div class="flex flex-col">
					<span class="flex gap-0.5 items-center text-white">
						<div class="font-black flex items-center gap-2 hover:underline cursor-pointer">
							<p>{user.nickname}</p>
							<span class="text-xs text-neutral-400">{user.first_name} {user.last_name}</span>
						</div>
					</span>
					<p class="hover:underline cursor-pointer">
						<!-- pill -->
						<span
							style="background-color: {user.role?.color || 'gray'};"
							class="rounded-full px-2 text-xs p-0.5">{user.role?.name}</span
						>
					</p>
				</div>
			</div>
		</div>
	</div>
{/if}
