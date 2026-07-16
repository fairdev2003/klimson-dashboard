<script lang="ts">
	import type { TerminalNaming } from '../../console/terminal.svelte';

	type Props = {
		date: number;
		color?: string;
		naming: TerminalNaming | undefined;
	};

	let { date, naming, color = 'bg-blue-700' }: Props = $props();
</script>

{#if naming}
	<div class="flex items-center text-xs shrink-0">
		<div class={`${color} px-3 p-1 text-white rounded-full`}>
			{naming.name}/{naming.path} [{new Date(date).toLocaleTimeString()}]
		</div>
	</div>
{/if}
