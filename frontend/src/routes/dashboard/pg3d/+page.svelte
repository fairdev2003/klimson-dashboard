<script lang="ts">
	import { onMount } from 'svelte';
	import type { ClanResponse } from './pg3d.types';
	import { api } from '$lib/api/api';
	import Heading from '$lib/components/dashboard/typography/Heading.svelte';
	import Loader from '$lib/components/dashboard/Loader.svelte';
	import PG3DClanMemberTable from '$lib/components/dashboard/table/pg3d/PG3DClanMemberTable.svelte';
	import { clan_info } from '$lib/components/dashboard/table/pg3d/clan.store';

	onMount(async () => {
		const context_response = await api.context_storage.GetSinglePrivateContext('clan_id');
		const clan_id = context_response.data.value;
		const response = await api.pg3d.GetClanInfo(clan_id);

		$clan_info = response.data;
	});
</script>

<div>
	{#if $clan_info}
		<div class="flex p-4 justify-between border border-neutral-700 bg-neutral-900 items-center">
			<div class="flex gap-2 items-center">
				<img
					src="data:image/png;base64,{$clan_info.clan_info.clan_logo}"
					alt="Clan Logo"
					class="size-7 [image-rendering:pixelated]"
				/>
				<Heading>{$clan_info?.clan_info.clan_name}</Heading>
			</div>

			<p class="text-sm">{$clan_info.clan_info.members_count} members</p>
		</div>
		<div class="p-4">
			<PG3DClanMemberTable members={$clan_info.members} />
		</div>
	{:else}
		<div class="flex justify-center items-center mt-5">
			<Loader />
		</div>
	{/if}
</div>
