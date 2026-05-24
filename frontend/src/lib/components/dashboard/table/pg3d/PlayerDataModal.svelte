<script lang="ts">
	import Modal from '$lib/components/Modal.svelte';
	import type { PlayerData, PlayerInfo } from '../../../../../routes/dashboard/pg3d/pg3d.types';
	import ProdWidget from '../../../../../routes/dashboard/widgets/ProdWidget.svelte';
	import Loader from '../../Loader.svelte';
	import Heading from '../../typography/Heading.svelte';

	type Props = {
		title: string;
		onClose: () => void;
		opened: boolean;
		player?: PlayerInfo;
	};

	let { title, onClose, opened = $bindable(), player }: Props = $props();
</script>

<Modal bind:opened {onClose} {title} className="w-150 h-9/10">
	{#if player}
		<div class="bg-neutral-800 w-full p-4 flex gap-4">
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
		<div class="mt-5 gap-4 grid grid-cols-3">
			{@render Currency(
				'Coins',
				'https://modfs.top/static/img/currency/coin.png',
				player.currency.Coins
			)}
			{@render Currency(
				'Gems',
				'https://modfs.top/static/img/currency/gem.png',
				player.currency.Gems
			)}
			{@render Currency(
				'Coupons',
				'https://modfs.top/static/img/currency/coupon.png',
				player.currency.Coupons
			)}
			{@render Currency(
				'Pixel Pass Tickets',
				'https://modfs.top/static/img/currency/pptickets.png',
				player.currency['Pixel Pass Tickets']
			)}
			{@render Currency(
				'Silver',
				'https://modfs.top/static/img/currency/silver.png',
				player.currency.Silver
			)}
		</div>
	{:else}
		<div class="flex justify-center items-center mt-5">
			<Loader />
		</div>
	{/if}
</Modal>

{#snippet Currency(name: string, icon: string, value?: number)}
	{#if value && value > 0}
		<div class="col-span-1 flex-col p-3 h-30 bg-neutral-800 flex justify-between items-center">
			<img class="[image-rendering:pixelated]" alt={name + '-icon'} src={icon} />
			<div class="flex justify-center flex-col items-center">
				<p class="text-xs text-neutral-400">{name}:</p>
				<p class="text-xs font-semibold">{value}</p>
			</div>
		</div>
	{/if}
{/snippet}
