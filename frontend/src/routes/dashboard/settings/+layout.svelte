<script>
	let { children } = $props();
</script>

<div class="flex flex-col justify-center grid-cols-10 lg:grid w-full overflow-auto p-10">
	{@render children()}
</div>
