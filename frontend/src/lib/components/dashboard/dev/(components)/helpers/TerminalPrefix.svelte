<script lang="ts">
	import type { TerminalNaming } from '../../console/terminal.svelte';

	type Props = {
		naming: TerminalNaming | undefined;
	};

	let { naming }: Props = $props();
</script>

{#if naming}
	<div class="flex items-center text-xs h-10 w-full">
		<div class="bg-blue-700 px-3 p-1 text-white rounded-full">{naming.name}/{naming.path}</div>
	</div>
{/if}
