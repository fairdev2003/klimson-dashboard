<script lang="ts">
	import Loader from './dashboard/Loader.svelte';

	type Props = {
		loading_title?: string;
	};

	const { loading_title = 'Ładowanie komponentów' }: Props = $props();
</script>

<div class="mx-auto flex justify-center gap-5">
	<Loader />
	<p>{loading_title}</p>
</div>
