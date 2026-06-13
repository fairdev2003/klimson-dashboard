<script lang="ts">
	import { dashboard_config } from '$lib/dashboard/stores/persist';
</script>

<div class="text-center flex h-full justify-center items-center">
	<p class="animated-gradient-text font-black text-sm">{$dashboard_config.dock_custom_name}</p>
</div>

<style>
	.animated-gradient-text {
		background: linear-gradient(
			90deg,
			#ff0000,
			#ff7f00,
			#ffff00,
			#00ff00,
			#0000ff,
			#4b0082,
			#9400d3
		);
		background-size: 300% 100%;

		-webkit-background-clip: text;
		background-clip: text;

		color: transparent;

		animation: gradient-move 5s linear infinite;
	}

	@keyframes gradient-move {
		0% {
			background-position: 0% 50%;
		}
		50% {
			background-position: 100% 50%;
		}
		100% {
			background-position: 0% 50%;
		}
	}
</style>
