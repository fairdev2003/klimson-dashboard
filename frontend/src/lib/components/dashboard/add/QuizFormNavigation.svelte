<script lang="ts">
	type Props = {
		onleftsideclick?: () => void;
		onrightsideclick?: () => void;
		left_name?: string;
		right_name?: string;
		alt_left?: string;
		alt_right?: string;
		href_left?: string;
		href_right?: string;
	};

	const {
		onleftsideclick,
		href_right,
		href_left,
		onrightsideclick,
		left_name,
		right_name,
		alt_left,
		alt_right
	}: Props = $props();
</script>

<div class="flex justify-between gap-1">
	<div>
		<a
			onclick={() => {
				if (onleftsideclick) {
					onleftsideclick();
				}
			}}
			href={href_left && href_left}
			class="flex cursor-pointer gap-1 text-[14px] text-neutral-500"
		>
			{'<'}
			<p class="hover:underline">{alt_left ? alt_left : 'Wróć'}</p>
		</a>
		<h1 class="mb-5 text-xl font-medium">{left_name || 'LEFT_NAME'}</h1>
	</div>
	<div>
		<a
			href={href_right && href_right}
			onclick={() => {
				if (onrightsideclick) {
					onrightsideclick();
				}
			}}
			class="flex cursor-pointer justify-end gap-1 text-[14px] text-neutral-500"
		>
			<p class="hover:underline">{alt_right ? alt_right : 'Następny'}</p>
			{'>'}
		</a>
		<h1 class="mb-5 text-xl font-medium">{right_name || 'RIGHT_NAME'}</h1>
	</div>
</div>
