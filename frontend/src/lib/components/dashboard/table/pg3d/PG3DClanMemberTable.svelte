<script lang="ts">
	import { onMount } from 'svelte';
	import type { ClanMember, PlayerData } from '../../../../../routes/dashboard/pg3d/pg3d.types';
	import ClanTableColumnField from './ClanTableColumnField.svelte';
	import { clan_info } from './clan.store';
	import Modal from '$lib/components/Modal.svelte';
	import { api } from '$lib/api/api';
	import Loader from '../../Loader.svelte';
	import Heading from '../../typography/Heading.svelte';
	import PlayerDataModa from './PlayerDataModal.svelte';
	import PlayerDataModal from './PlayerDataModal.svelte';

	const columns: (keyof ClanMember | 'asteroid')[] = [
		'id',
		'level',
		'name',
		'rank',
		'valor',
		'asteroid'
	];

	const columnsDict = {
		id: 'ID',
		level: 'Member Level',
		name: 'Nickname',
		rank: 'Clan Rank',
		valor: 'Valor Points',
		asteroid: 'Asteroid Record'
	};

	type Props = {
		members: ClanMember[];
	};

	let { members }: Props = $props();
	let playerInfoOpened: boolean = $state(false);
	let selectedName: string = $state('');
	let player_data: PlayerData | undefined = $state();
	let selected_clan_member: ClanMember | undefined = $state();

	let sortedMembers = $derived(
		[...members].sort((a, b) => {
			const leaderId = $clan_info?.clan_info?.leader_id;

			if (a.id === leaderId) return -1;
			if (b.id === leaderId) return 1;

			return b.valor - a.valor;
		})
	);

	async function FetchPlayerData(id: string) {
		const response = await api.pg3d.GetPlayerData(id);
		player_data = response.data;

		console.log(response);
	}
</script>

<div class="w-full max-w-full overflow-x-auto shadow-sm flex flex-col gap-4 pb-10">
	<div class="w-full overflow-x-auto">
		<table
			class="w-7xl mx-auto min-w-max border-collapse text-left rounded-xl overflow-hidden text-sm border border-neutral-600 text-white"
		>
			<thead class="bg-neutral-800 text-xs tracking-wider">
				<tr>
					{#each columns as column}
						<ClanTableColumnField column={columnsDict[column]} />
					{/each}
				</tr>
			</thead>
			{#if sortedMembers && sortedMembers.length}
				<tbody class="bg-neutral-900 text-white">
					{#each sortedMembers as member}
						<tr class="transition-colors">
							{#each columns as column}
								{@const stmt = member.id === $clan_info.clan_info.leader_id}
								<td
									onclick={async () => {
										if (column === 'name') {
											player_data = undefined;
											playerInfoOpened = !playerInfoOpened;
											selectedName = member.name;
											selected_clan_member = member;
											await FetchPlayerData(member.id);
										}
									}}
									class:clickable={column === 'name'}
									class:golden_mayo={stmt}
									class="whitespace-nowrap p-2 text-neutral-300 font-medium"
								>
									<div class="flex items-center gap-1">
										{#if stmt && column === 'name'}
											👑
										{/if}

										<p>
											{member[column]}
										</p>

										{#if column === 'valor'}
											<img src="https://modfs.top/static/img/icons/ClanValor.png" class="size-4" />
										{/if}
									</div>
								</td>
							{/each}
						</tr>
					{/each}
				</tbody>
			{:else}
				<div class="w-full">
					<p>No records :c</p>
				</div>
			{/if}
		</table>
	</div>
</div>

<PlayerDataModal
	bind:opened={playerInfoOpened}
	onClose={() => {
		playerInfoOpened = false;
	}}
	player={player_data?.info}
	title="Viewing {selectedName}"
/>

<style>
	@import 'tailwindcss';

	.golden_mayo {
		@apply bg-amber-400/20;
	}

	.clickable {
		@apply cursor-pointer;
	}
</style>
