<script lang="ts">
	import { onMount } from 'svelte';
	import type { ClanResponse } from './pg3d.types';
	import { api } from '$lib/api/api';
	import Heading from '$lib/components/dashboard/typography/Heading.svelte';
	import Loader from '$lib/components/dashboard/Loader.svelte';
	import PG3DClanMemberTable from '$lib/components/dashboard/table/pg3d/PG3DClanMemberTable.svelte';
	import { clan_info } from '$lib/components/dashboard/table/pg3d/clan.store';
	import { dockComponent } from '../dashboard.svelte';
	import PG3DDocky from '$lib/components/dashboard/dock/boxes/PG3DDocky.svelte';

	onMount(async () => {
		dockComponent.set(PG3DDocky);
		const context_response = await api.context_storage.GetSinglePrivateContext('clan_id');
		const clan_id = context_response.data.value;
		const response = await api.pg3d.GetClanInfo(clan_id);

		$clan_info = response.data;
	});
</script>

<div>
	{#if $clan_info}
		<div class="p-4">
			<PG3DClanMemberTable members={$clan_info.members} />
		</div>
	{:else}
		<div class="flex justify-center items-center mt-5">
			<Loader />
		</div>
	{/if}
</div>
