<script lang="ts">
	import Modal from '$lib/components/Modal.svelte';
	import { onMount } from 'svelte';
	import CreateFormInput from '../CreateFormInput.svelte';
	import DatabaseColumn from './DatabaseColumn.svelte';
	import Button from '$lib/components/Button.svelte';
	import DatabaseModalInput from './DatabaseModalInput.svelte';
	import type { ColumnData } from '$lib/api/requests/database.table';
	import { highlightedFields } from './data_table.store';

	type Props = {
		opened: boolean;
		row: any;
		onSave: () => void;
		column: ColumnData;
		id: `${string}-${string}-${string}-${string}-${string}` | undefined;
	};

	let { opened = $bindable(), column, row, id, onSave }: Props = $props();

	let key: string = $state('');
	let value: string = $state('');

	onMount(() => {
		if (!row) return;

		key = column.slug;
		value = String(row[column.slug]);
	});
</script>

<Modal
	title="Editing '{key}' - {row[key]}"
	className="w-100 h-auto"
	bind:opened
	onClose={() => (opened = !opened)}
>
	<div class="flex flex-col gap-3">
		<DatabaseModalInput disabled label="key" bind:value={key} />

		<DatabaseModalInput label="value" bind:value />
	</div>

	<div class="flex justify-end mt-3">
		<Button
			size="small"
			theme="correct"
			onclick={() => {
				onSave();
				opened = false;
			}}>Add Change</Button
		>
	</div>
</Modal>
