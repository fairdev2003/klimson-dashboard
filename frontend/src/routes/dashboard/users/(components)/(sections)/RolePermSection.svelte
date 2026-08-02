<script lang="ts">
	import RDBColorInput from '$lib/components/modal/RDBColorInput.svelte';
	import type { Role } from '$lib/types/user';
	import { onMount } from 'svelte';

	type Props = { role: Role };

	let { role = $bindable() }: Props = $props();

	onMount(() => {});
</script>

<div class="h-500"></div>
