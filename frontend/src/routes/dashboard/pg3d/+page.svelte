<script lang="ts">
	import { onMount } from 'svelte';
	import type { ClanResponse } from './pg3d.types';
	import { api } from '$lib/api/api';
	import Heading from '$lib/components/dashboard/typography/Heading.svelte';
	import Loader from '$lib/components/dashboard/Loader.svelte';
	import PG3DClanMemberTable from '$lib/components/dashboard/table/pg3d/PG3DClanMemberTable.svelte';
	import { clan_info } from '$lib/components/dashboard/table/pg3d/clan.store';
	import PG3DDocky from '$lib/components/dashboard/dock/boxes/PG3DDocky.svelte';
	import { Dashboard } from '$lib/dashboard/logic';

	onMount(async () => {
		Dashboard.state.setDockComponent(PG3DDocky);
		const context_response = await api.context_storage.GetSinglePrivateContext('clan_id');
		const clan_id = context_response.data.value;
		const response = await api.pg3d.GetClanInfo(clan_id);

		$clan_info = response.data;
	});
</script>

<div class="relative w-full h-full overflow-hidden background-container">
	{#if $clan_info}
		<div class="p-4 z-10">
			<PG3DClanMemberTable members={$clan_info.members} />
		</div>
	{:else}
		<div class="flex justify-center items-center mt-5">
			<Loader />
		</div>
	{/if}
</div>

<style>
	.background-container {
		background-image: url('https://cdnb.artstation.com/p/assets/images/images/051/094/853/large/elena-dudnakova-summer-party-remastering-lottery.jpg?1656441766');

		background-size: cover;
		background-position: center;
		background-repeat: no-repeat;

		width: 100%;
		min-height: 100vh;
	}
</style>
