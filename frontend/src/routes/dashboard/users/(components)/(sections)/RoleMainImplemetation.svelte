<script lang="ts">
	import RDBInput from '$lib/components/modal/RDBInput.svelte';
	import type { Role } from '$lib/types/user';
	import { onMount } from 'svelte';
	import account_controller from '../../helpers/access.svelte';
	import { debug } from '$lib/terminal/logic';
	import RDBColorInput from '$lib/components/modal/RDBColorInput.svelte';

	type Props = { role: Role };

	let { role = $bindable() }: Props = $props();

	onMount(() => {});
</script>

<div class="flex flex-col gap-4">
	<RDBInput label="ROLE NAME" bind:value={role.name} />
	<RDBInput label="ROLE ICON (URL)" bind:value={role.icon} />
	<RDBColorInput label="Role color" bind:value={role.color} />
</div>

<style>
	@import 'tailwindcss';

	.selector {
		@apply;
	}
</style>
