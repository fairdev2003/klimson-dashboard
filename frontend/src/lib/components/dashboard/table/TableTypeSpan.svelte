<script lang="ts">
	import Icon from '@iconify/svelte';

	type Props = {
		type: string;
	};
	let { type }: Props = $props();
</script>

{#if type === 'id'}
	<Icon icon="ix-id-filled" width="12" height="12" />
{/if}

{#if type === 'string'}
	<Icon icon="codicon-symbol-string" width="14" height="14" />
{/if}

{#if type === 'number'}
	<Icon icon="cuida-number-outline" width="14" height="14" />
{/if}
