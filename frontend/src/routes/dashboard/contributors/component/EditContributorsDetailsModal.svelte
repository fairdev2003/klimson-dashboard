<script lang="ts">
	import { api } from '$lib/api/api';
	import Button from '$lib/components/Button.svelte';
	import CreateFormInput from '$lib/components/dashboard/CreateFormInput.svelte';
	import Modal from '$lib/components/Modal.svelte';
	import { contributors } from '$lib/dashboard/stores/data.store';
	import { developerView } from '$lib/dashboard/stores/persist';
	import { addFormContributor } from '$lib/dashboard/stores/store';
	import { toast } from '$lib/dashboard/stores/toast';
	import { onMount } from 'svelte';
	import { fade } from 'svelte/transition';
	import { contributors_loading } from '../vars';
	import Icon from '@iconify/svelte';

	type Props = {
		onClose?: () => void;
		opened?: boolean;
	};

	let { onClose, opened = $bindable(true) }: Props = $props();
	let pass: string = $state('');
	let loading: boolean = $state(false);

	async function UpdateDetails() {
		loading = true;
		$contributors_loading = true;
		onClose?.();

		let { name, description, login, permissions } = $addFormContributor;

		const response = await api.contributor.UpdateContributorDetails(
			{ id: $addFormContributor.id },
			{
				name,
				description,
				login,
				permissions
			}
		);

		loading = false;

		const response2 = await api.contributor.GetContributors();
		$contributors = response2.data;
		$contributors_loading = false;
		toast.show('Dodano!', 'success');
	}

	onMount(() => {
		$addFormContributor.password = '';
	});
</script>

<svelte:boundary
	onerror={(e) => {
		toast.show(e, 'error');
	}}
>
	{#snippet failed(error, reset)}
		<div in:fade={{ duration: 150 }} class="m-5 w-7xl border-1 border-red-400 bg-red-500/70 p-5">
			{error}
		</div>
	{/snippet}
	<Modal
		title="Edytuj szczegóły współtwórcy"
		className="w-[600px] "
		onClose={() => onClose?.()}
		bind:opened
	>
		{#if $addFormContributor.blocked_till}
			<div
				transition:fade={{ duration: 200 }}
				class="mb-4 flex items-center gap-3 rounded-lg border border-orange-500/50 bg-orange-500/10 p-4 text-orange-500 shadow-lg shadow-red-500/5"
			>
				<div class="flex-shrink-0">
					<Icon icon="material-symbols:warning" width="24" class="animate-pulse" />
				</div>

				<div class="flex flex-col">
					<span class="text-sm font-bold tracking-wide">Konto Zawieszone</span>
					<p class="text-xs opacity-90">
						Ten kontrybutor nie ma obecnie dostępu do panelu. Wszystkie jego uprawnienia są
						zawieszone.
					</p>
				</div>
			</div>
		{/if}
		<div class="flex flex-col gap-4">
			<CreateFormInput disabled label="Identyfikator" bind:value={$addFormContributor.id} />
			<CreateFormInput label="Nazwa kontrybutora" bind:value={$addFormContributor.name} />
			<CreateFormInput label="Opis kontrybutora" bind:value={$addFormContributor.description} />
			<CreateFormInput label="Login" bind:value={$addFormContributor.login} />
			{#if $developerView}
				<CreateFormInput
					label="Permisje"
					bind:value={$addFormContributor.permissions}
					disclaimer="*Tylko dla developera"
				/>
			{/if}
		</div>

		<div class="my-5 flex justify-end">
			<Button
				{loading}
				onclick={async () => {
					await UpdateDetails();
				}}
				size="small"
				theme="secondary">Prześlij dane</Button
			>
		</div>
	</Modal>
</svelte:boundary>
