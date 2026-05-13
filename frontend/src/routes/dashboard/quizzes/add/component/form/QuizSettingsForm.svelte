<script lang="ts">
	import Button from '$lib/components/Button.svelte';
	import { addFormQuiz } from '$lib/dashboard/stores/store';
	import { toast } from '$lib/dashboard/stores/toast';

	let modalOpen: boolean = $state(false);

	type Props = {
		onSave?: () => void;
	};

	let { onSave }: Props = $props();
	let loading: boolean = $state(false);

	async function handleSave() {
		loading = true;
		await new Promise((resolve) => setTimeout(resolve, 2000));
		loading = false;
		toast.show('Zmiany zostały zapisane pomyślnie!', 'success');
		onSave?.();
	}
</script>

<div class="flex flex-col justify-center gap-5">
	<div class="flex items-center">
		<input
			id="default-checkbox"
			type="checkbox"
			bind:checked={$addFormQuiz['public']}
			class="border-default-medium rounded-xs bg-neutral-secondary-medium focus:ring-brand-soft h-4 w-4 border focus:ring-2"
		/>
		<label for="default-checkbox" class="text-heading ms-2 select-none text-sm font-medium"
			>Ustaw quiz na publiczny</label
		>
	</div>

	<div class="flex justify-end">
		<Button
			{loading}
			theme="secondary"
			size="small"
			onclick={function () {
				handleSave ? handleSave() : null;
			}}>Zapisz zmiany</Button
		>
	</div>
</div>
