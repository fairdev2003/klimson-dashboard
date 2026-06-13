<script lang="ts">
	import Modal from '$lib/components/Modal.svelte';
	import { toast } from '$lib/dashboard/stores/toast';
	import type { PlayerInfo, Weapon } from '../../../../../routes/dashboard/pg3d/pg3d.types';
	import Loader from '../../Loader.svelte';
	import Heading from '../../typography/Heading.svelte';

	type Props = {
		title: string;
		onClose: () => void;
		opened: boolean;
		player?: PlayerInfo;
	};

	function getRarity(name: string): string {
		const found_record = player?.weapons.find((e) => String(e.name) === name);

		return String(found_record?.rarity);
	}

	let { title, onClose, opened = $bindable(), player }: Props = $props();
</script>

<Modal bind:opened {onClose} {title} className="w-150 h-9/10">
	{#if player}
		<div class="bg-neutral-800 rounded-xl w-full p-4 flex gap-4">
			<img src="data:image/png;base64,{player.avatar}" />
			<div class="">
				<Heading>{player.username}</Heading>
				<div class="flex gap-1">
					<p class="text-xs">Level {player.level}</p>
				</div>
				{#if player.clan}
					<span class="flex items-center gap-1">
						<img
							src="data:image/png;base64,{player.clan.clanlogo}"
							alt="Clan Logo"
							class="size-5 [image-rendering:pixelated]"
						/>

						<p>{player.clan.clanname}</p>
					</span>
				{/if}
			</div>
		</div>
		<div class="flex gap-3 flex-wrap mt-5">
			{@render Currency(
				'Coins',
				'https://modfs.top/static/img/currency/coin.png',
				player.currency.Coins,
				'bg-yellow-500/50'
			)}
			{@render Currency(
				'Gems',
				'https://modfs.top/static/img/currency/gem.png',
				player.currency.Gems,
				'bg-blue-500/50'
			)}
			{@render Currency(
				'Coupons',
				'https://modfs.top/static/img/currency/coupon.png',
				player.currency.Coupons
			)}
			{@render Currency(
				'Pixel Pass Tickets',
				'https://modfs.top/static/img/currency/pptickets.png',
				player.currency['Pixel Pass Tickets'],
				'bg-orange-500/50'
			)}
			{@render Currency(
				'Silver',
				'https://modfs.top/static/img/currency/silver.png',
				player.currency.Silver,
				'bg-slate-500/50'
			)}
		</div>
		<div class="grid grid-cols-3 gap-5 mt-5">
			{@render Weapon(
				player.loadouts[player.current_loadout].primary.id,
				player.loadouts[player.current_loadout].primary.name,
				getRarity(player.loadouts[player.current_loadout].primary.name)
			)}
			{@render Weapon(
				player.loadouts[player.current_loadout].backup.id,
				player.loadouts[player.current_loadout].backup.name,
				getRarity(player.loadouts[player.current_loadout].primary.name)
			)}
			{@render Weapon(
				player.loadouts[player.current_loadout].melee.id,
				player.loadouts[player.current_loadout].melee.name,
				getRarity(player.loadouts[player.current_loadout].primary.name)
			)}
			{@render Weapon(
				player.loadouts[player.current_loadout].special.id,
				player.loadouts[player.current_loadout].special.name,
				getRarity(player.loadouts[player.current_loadout].primary.name)
			)}
			{@render Weapon(
				player.loadouts[player.current_loadout].sniper.id,
				player.loadouts[player.current_loadout].sniper.name,
				getRarity(player.loadouts[player.current_loadout].primary.name)
			)}
			{@render Weapon(
				player.loadouts[player.current_loadout].heavy.id,
				player.loadouts[player.current_loadout].heavy.name,
				getRarity(player.loadouts[player.current_loadout].primary.name)
			)}
		</div>
	{:else}
		<div class="flex justify-center items-center mt-5">
			<Loader />
		</div>
	{/if}
</Modal>

{#snippet Weapon(id: number, name: string, rarity?: string)}
	<button
		class="col-span-1 flex cursor-pointer border-2 border-neutral-800 focus:border-blue-500 hover:border-blue-500 flex-col gap-1 items-center p-4 rounded-xl justify-center bg-neutral-800"
	>
		<img
			class="[image-rendering:pixelated] size-15"
			src="https://asteroidpg3d.xyz/api/weapon_icon/{id}"
		/>
		<p class="font-black text-center text-xs">{name}</p>
		<p class="text-neutral-400 text-center text-xs">ID: {id}</p>
	</button>

	<style>
		.common {
			@apply border-gray-400;
		}

		.uncommon {
			@apply border-green-600;
		}

		.rare {
			@apply border-blue-500;
		}

		.epic {
			@apply border-purple-500;
		}

		.legendary {
			@apply border-orange-500;
		}

		.mythical {
			@apply border border-red-600 bg-red-500;
		}
	</style>
{/snippet}

{#snippet Currency(name: string, icon: string, value?: number, color?: string)}
	{#if value && value > 0}
		<div
			class="col-span-1 p-2 rounded-full px-4 {color
				? color
				: 'bg-neutral-800'} flex justify-between items-center"
		>
			<div class="flex justify-center flex-col items-center">
				<p class="text-xs font-black">
					{value}
					<span class="font-normal">
						{name}
					</span>
				</p>
			</div>
		</div>
	{/if}
{/snippet}
