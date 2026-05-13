<script lang="ts">
	import Modal from '$lib/components/Modal.svelte';
	import { addFormContributor } from '$lib/dashboard/stores/store';

	import Button from '$lib/components/Button.svelte';
	import CreateFormInput from '$lib/components/dashboard/CreateFormInput.svelte';
	import { api } from '$lib/api/api';
	import { toast } from '$lib/dashboard/stores/toast';
	import type { AxiosError } from 'axios';
	import axios from 'axios';

	type Props = {
		onClose?: () => void;
		opened?: boolean;
	};

	let loading: boolean = $state(false);
	let password: string = $state('');

	let { onClose, opened = $bindable(true) }: Props = $props();

	async function CheckPasssword() {
		loading = true;
		try {
			const response = await api.contributor.CheckContributorPassword({
				login: $addFormContributor.login,
				password
			});

			if (response.data.correct) {
				toast.success('Hasło jest poprawne');
				onClose?.();
			} else {
				toast.info('Hasło nie jest poprawne');
			}
		} catch (error: unknown) {
			if (axios.isAxiosError(error)) {
				const message = error.response?.data?.message || 'Błąd serwera';
				toast.error(message);
				onClose?.();
			} else {
				toast.error('Wystąpił nieoczekiwany błąd');
			}
		} finally {
			loading = false;
		}
	}
</script>

<Modal
	title={`Edytuj uprawnienia dla ${$addFormContributor.name}`}
	className="w-[400px]"
	onClose={() => onClose?.()}
	bind:opened
>
	<div>
		<CreateFormInput label="Hasło" bind:value={password} />
	</div>
	<div class="my-5 flex justify-end">
		<Button {loading} onclick={CheckPasssword} size="small" theme="secondary"
			>Sprawdź poprawność hasła</Button
		>
	</div>
</Modal>
