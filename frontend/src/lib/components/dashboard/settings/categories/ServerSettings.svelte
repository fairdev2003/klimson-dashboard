<script lang="ts">
	import { blur } from 'svelte/transition';
</script>

<div
	class="flex flex-col border-t border-neutral-700 lg:w-3/4 lg:px-10"
	in:blur={{ duration: 300 }}
></div>
