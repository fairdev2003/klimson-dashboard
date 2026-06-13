<script>
	import { route } from '$lib/dashboard/stores/persist';
	import Icon from '@iconify/svelte';
</script>

{#if $route && $route.includes('/file')}
	<div
		class="py-2 items-center flex gap-1 px-3 rounded-xl bg-neutral-700 border-neutral-800 border text-sm text-neutral-200"
	>
		<Icon icon="mdi:file" width="20" height="20" />
		<p class="">{$route.split('/')[$route.split('/').length - 1]}</p>
	</div>
{/if}
