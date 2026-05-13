<script lang="ts">
	import { goto } from '$app/navigation';
	import { api } from '$lib/api/api';
	import { routes } from '$lib/dashboard/stores/data.store';
	import Icon from '@iconify/svelte';
</script>

<div
	class="min-h-2 text-start gap-4 p-3 col-span-2 items-center bg-neutral-900 border-neutral-700 border flex"
>
	<div
		class="bg-neutral-800 rounded-xl size-13 border-neutral-700 border justify-center items-center flex"
	>
		<Icon icon="fluent:production-20-filled" height="30" width="30" class="text-neutral-600" />
	</div>
	<div class="flex flex-col">
		<h3 class="truncate text-sm font-bold tracking-tight text-neutral-100">Obecny serwer</h3>
		<a class="text-blue-500 hover:underline" href={api.api_config.dev_server + '/v1/api'}
			>{api.api_config.dev_server}</a
		>
	</div>
</div>
