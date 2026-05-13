<script lang="ts">
	import type { SidebarItemType } from '$lib/types/sidebar';
	import Icon from '@iconify/svelte';

	let { onclick, children, id, icon, label, link }: SidebarItemType = $props();
</script>

<div class="flex gap-3">
	<Icon {icon} />
	<p>{label}</p>
</div>
