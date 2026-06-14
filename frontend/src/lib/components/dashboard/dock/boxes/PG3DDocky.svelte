<script lang="ts">
	import { clan_info } from '../../table/pg3d/clan.store';
	import Heading from '../../typography/Heading.svelte';
</script>

{#if $clan_info}
	<div class="flex gap-2 items-center">
		<img
			src="data:image/png;base64,{$clan_info.clan_info.clan_logo}"
			alt="Clan Logo"
			class="size-7 [image-rendering:pixelated]"
		/>
		<Heading>{$clan_info?.clan_info.clan_name}</Heading>
	</div>
{/if}
