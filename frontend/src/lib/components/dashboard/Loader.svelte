<script lang="ts">
	type Props = {
		theme?: 'dashboard' | 'regular';
	};

	const { theme = 'dashboard' }: Props = $props();
</script>

<div class:loader={theme === 'dashboard'} class:loader-white={theme === 'regular'}></div>

<style>
	.loader {
		width: 25px;
		aspect-ratio: 1;
		border-radius: 50%;
		border: 3px solid oklch(70.7% 0.165 254.624); /* kolor */
		border-top-color: transparent; /* przezroczysta część do efektu obrotu */
		animation: spin 0.8s linear infinite;
	}

	.loader-white {
		width: 15px;
		aspect-ratio: 1;
		border-radius: 50%;
		border: 2px solid rgba(255, 255, 255, 0.5); /* kolor */
		border-top-color: transparent; /* przezroczysta część do efektu obrotu */
		animation: spin 0.8s linear infinite;
	}

	@keyframes spin {
		to {
			transform: rotate(360deg);
		}
	}
</style>
