<script lang="ts">
	import { dashboard_config } from '$lib/dashboard/stores/persist';
	import { nickname } from '../settings/store.svelte';
	import SidebarUserLogged from '../SidebarUserLogged.svelte';
	import StoragePill from './StoragePill.svelte';
</script>

<div class="flex flex-col gap-3 -z-1">
	{#if $dashboard_config.client_pills}
		{#if $dashboard_config.client_pills.includes('profile')}
			<SidebarUserLogged
				selected={false}
				name={$nickname}
				role="$root"
				pfp_logo="https://klimson.dev/_app/immutable/assets/klimson.CQA0gh-5.jpeg"
			/>
		{/if}

		{#if $dashboard_config.client_pills.includes('storage')}
			<StoragePill />
		{/if}
	{/if}
</div>
