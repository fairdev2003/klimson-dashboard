<script lang="ts">
	type Props = {
		src: string;
	};

	let { src }: Props = $props();
</script>

<div class="h-44 w-full sm:w-72 overflow-hidden bg-neutral-900 shadow-lg shadow-cyan-500/10">
	<img
		{src}
		alt="Zdjęcie"
		class="h-full w-full object-cover duration-300 cursor-pointer hover:blur-xs transition-all"
	/>
</div>
