<script lang="ts">
	import type { TableDataType } from '$lib/api/requests/database.table';
	import type { TableData } from '$lib/api/requests/misc';
	import Icon from '@iconify/svelte';
	import { onMount } from 'svelte';

	type Props = {
		column: TableData;
	};
	let { column }: Props = $props();
	onMount(() => {
		console.log('span', column);
	});
</script>

{#if column.foreign_key === 'true'}
	<Icon icon="fluent-mdl2:relationship" width="12" height="12" />
{/if}

{#if column.type === 'int8'}
	<Icon icon="ix-id-filled" width="12" height="12" />
{/if}

{#if column.type === 'text' || column.type === 'varchar'}
	<Icon icon="codicon-symbol-string" width="14" height="14" />
{/if}

{#if column.type === 'number'}
	<Icon icon="cuida-number-outline" width="14" height="14" />
{/if}

{#if column.type === 'timestamptz'}
	<Icon icon="mdi:clock" width="14" height="14" />
{/if}

{#if column.type === 'bool'}
	<Icon icon="iconamoon:check-bold" width="14" height="14" />
{/if}
