<script lang="ts">
	import type { Role } from '$lib/types/user';
	import { onMount } from 'svelte';

	type Props = { role: Role };

	let { role = $bindable() }: Props = $props();

	onMount(() => {});
</script>
