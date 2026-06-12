<script lang="ts">
	import { dashboard_config } from '$lib/dashboard/stores/persist';
	import { nickname } from '../settings/store.svelte';
	import SidebarUserLogged from '../SidebarUserLogged.svelte';
	import StoragePill from './StoragePill.svelte';
</script>

<div class="flex flex-col gap-3">
	{#if $dashboard_config.sidebar_preference.includes('profile')}
		<SidebarUserLogged
			name={$nickname}
			role="$root"
			pfp_logo="https://klimson.dev/_app/immutable/assets/klimson.CQA0gh-5.jpeg"
		/>
	{/if}

	{#if $dashboard_config.sidebar_preference.includes('storage')}
		<StoragePill />
	{/if}
</div>
