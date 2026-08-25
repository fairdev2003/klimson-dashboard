<script lang="ts">
	import { onMount } from 'svelte';
	import { browser } from '$app/environment';
	import RDBModal from '$lib/components/modal/RDBModal.svelte';
	import { Line } from 'svelte-chartjs';
	import {
		Chart as ChartJS,
		Title,
		Tooltip,
		Legend,
		LineElement,
		LinearScale,
		PointElement,
		CategoryScale
	} from 'chart.js';

	ChartJS.register(Title, Tooltip, Legend, LineElement, LinearScale, PointElement, CategoryScale);

	let data = {
		labels: ['10:00', '10:01', '10:02', '10:03', '10:04', '10:05', '10:06'],
		datasets: [
			{
				label: 'CPU Usage (%)',
				fill: true,
				backgroundColor: 'rgba(54, 162, 235, 0.2)',
				borderColor: 'rgba(54, 162, 235, 1)',
				data: [25, 42, 38, 85, 60, 45, 30]
			},
			{
				label: 'Memory Usage (%)',
				fill: true,
				backgroundColor: 'rgba(255, 159, 64, 0.2)',
				borderColor: 'rgba(255, 159, 64, 1)',
				data: [65, 66, 68, 72, 70, 68, 67]
			}
		]
	};

	type Props = {
		opened: boolean;
	};

	let { opened = $bindable() }: Props = $props();
</script>

<RDBModal
	onClose={() => {
		opened = !opened;
	}}
	size="window"
	bind:opened
	bg_color="blacker"
	border="borderless"
	title="Server Usage"
>
	<div class="w-full h-full bg-neutral-800/60 rounded-lg p-4">
		<Line {data} options={{ responsive: true }} />
	</div>
</RDBModal>
