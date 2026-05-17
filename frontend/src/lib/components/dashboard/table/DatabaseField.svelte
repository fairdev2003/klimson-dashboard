<script lang="ts">
	import { onMount } from 'svelte';
	import DatabaseColumn from './DatabaseColumn.svelte';
	import MovingTooltip from '../MovingTooltip.svelte';
	import Icon from '@iconify/svelte';
	import FieldEditModal from './FieldEditModal.svelte';
	import SFMCodeArea from '../sfm/(components)/SFMCodeArea.svelte';

	type Props = {
		row: any;
		column: DatabaseColumn;
	};

	let { row, column }: Props = $props();
	let field_id: `${string}-${string}-${string}-${string}-${string}` | undefined = $state();

	let opened: boolean = $state(false);

	function formatFileFromLink(link: any): string {
		if (!link || typeof link !== 'string') {
			return '';
		}

		const linkSplittedTable = link.split('/');
		const splittedTableLength = linkSplittedTable.length;

		const lastElement = linkSplittedTable[splittedTableLength - 1];

		if (!lastElement) {
			return '';
		}

		return lastElement;
	}

	const imageExtensions = ['.png', '.jpg', '.jpeg', '.svg', '.webp'];

	function isImage(fileName: string): boolean {
		const lowerCaseName = fileName.toLowerCase();

		return imageExtensions.some((ext) => lowerCaseName.endsWith(ext));
	}

	const sfmExtensions = ['.sfm'];

	function isSFMCode(fileName: string) {
		const lowerCaseName = fileName.toLowerCase();

		return sfmExtensions.some((ext) => lowerCaseName.endsWith(ext));
	}

	onMount(() => {
		field_id = crypto.randomUUID();
	});
</script>

<td
	onclick={() => {
		opened = !opened;
	}}
	class="whitespace-nowrap cursor-pointer border text-neutral-300 font-medium border-neutral-700 hover:bg-neutral-700"
>
	<MovingTooltip>
		{#snippet tooltipContent()}
			{@const statement =
				(row &&
					row[column.slug] !== undefined &&
					String(row[column.slug]).includes('data:image')) ||
				String(row[column.slug]).includes('https')}
			<div class="flex flex-col gap-3 p-1">
				{#if statement}
					{#if isImage(row[column.slug])}
						<div class="max-w-100">
							<img src={String(row[column.slug])} />
						</div>
					{/if}

					{#if isSFMCode(row[column.slug])}
						<div class="max-w-150">
							<SFMCodeArea program_link={row[column.slug]} />
						</div>
					{/if}
				{:else}
					<p class="text-xs">
						{row[column.slug]}
					</p>
				{/if}

				<p class="text-xs">
					type: {column.type}
				</p>
			</div>
		{/snippet}

		{#if (row && row[column.slug] !== undefined && String(row[column.slug]).includes('data:image')) || String(row[column.slug]).includes('https')}
			<div class="flex items-center gap-1 px-6 py-4 text-blue-500 font-semibold">
				<div class="text-blue-500/60">
					<Icon icon="glyphs-poly:file" width="25" height="25" />
				</div>
				<span>
					{formatFileFromLink(String(row[column.slug]))}
				</span>
			</div>
		{:else}
			<div class="px-6 py-4">
				{row[column.slug]}
			</div>
		{/if}
	</MovingTooltip>
	<div
		onclick={(e) => {
			e.stopPropagation();
		}}
		class="cursor-auto"
	>
		<FieldEditModal {column} {row} bind:opened />
	</div>
</td>
