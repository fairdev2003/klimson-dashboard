<script lang="ts">
	import Button from '$lib/components/Button.svelte';
	import HarcCheckBox from '$lib/components/dashboard/HarcCheckBox.svelte';
	import SectionList from '$lib/components/dashboard/SectionList.svelte';
	import {
		animation_preset,
		dashboard_config,
		debugOn,
		developerView
	} from '$lib/dashboard/stores/persist';
	import { fade } from 'svelte/transition';
	let value: number = $state(0);
</script>

<div class="p-5 text-white">
	<div class="mb-5">
		<h1 class="text-2xl font-bold">Ustawienia</h1>
	</div>
	<div in:fade={{ duration: 150 }} out:fade={{ duration: 150 }} class="flex flex-col gap-3">
		<HarcCheckBox bind:checked={$debugOn} label="Włącz dziennik zdarzeń" />
		<HarcCheckBox bind:checked={$developerView} label="Widok developera" />

		<div class="flex flex-col gap-2">
			<p>Wybrana animacja okienek: {$animation_preset}</p>
			<div class="flex gap-2">
				<Button onclick={() => ($animation_preset = 'blur')} theme="secondary" size="small"
					>blur</Button
				>
				<Button onclick={() => ($animation_preset = 'klimson')} theme="secondary" size="small"
					>klimson</Button
				>
				<Button onclick={() => ($animation_preset = 'jason')} theme="secondary" size="small"
					>jason</Button
				>
			</div>
		</div>
	</div>
</div>
