<script lang="ts">
	import { api } from '$lib/api/api';
	import Button from '$lib/components/Button.svelte';
	import CreateFormInput from '$lib/components/dashboard/CreateFormInput.svelte';
	import Modal from '$lib/components/Modal.svelte';
	import { addFormContributor } from '$lib/dashboard/stores/store';
	import axios from 'axios';
	import { contributors_loading } from '../vars';
	import { toast } from '$lib/dashboard/stores/toast';
	import { contributors } from '$lib/dashboard/stores/data.store';

	type Props = {
		onClose?: () => void;
		opened?: boolean;
	};

	let { onClose, opened = $bindable(true) }: Props = $props();
	let loading: boolean = $state(false);

	let password: string = $state('');
	let repeatPassword: string = $state('');

	async function UpdatePassword() {
		loading = true;
		$contributors_loading = true;

		let message = '';
		if (password !== repeatPassword || password.length < 3) {
			toast.error('Hasła się nie zgadzają!');
			loading = false;
			$contributors_loading = false;
			return;
		}
		onClose?.();
		try {
			const response = await api.contributor.UpdateContributorPassword(
				{ id: $addFormContributor.id },
				{ password, login: $addFormContributor.login }
			);

			if (response.data.message) {
				message = response.data.message;
			}
		} catch (error: unknown) {
			if (axios.isAxiosError(error)) {
				const message = error.response?.data?.message || 'Błąd serwera';
				toast.error(message);
			} else {
				toast.error('Wystąpił nieoczekiwany błąd');
			}
		} finally {
			const response = await api.contributor.GetContributors();
			$contributors = response.data;
			loading = false;
			$contributors_loading = false;
			toast.success(message);
		}
	}
</script>

<Modal
	title="Edytuj szczegóły współtwórcy"
	className="w-150"
	onClose={() => onClose?.()}
	bind:opened
>
	<div class="flex flex-col gap-4">
		<CreateFormInput disabled bind:value={$addFormContributor.name} label="Edytujesz:" />
		<CreateFormInput bind:value={$addFormContributor.login} label="Nowe hasło" />

		<CreateFormInput bind:value={password} label="Nowe hasło" />
		<CreateFormInput bind:value={repeatPassword} label="Powtórz hasło" />
	</div>
	<div class="my-5 flex justify-end">
		<Button {loading} onclick={UpdatePassword} size="small" theme="secondary"
			>Zaaktualizuj hasło</Button
		>
	</div>
</Modal>
