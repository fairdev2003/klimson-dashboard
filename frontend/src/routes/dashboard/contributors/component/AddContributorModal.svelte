<script lang="ts">
	import Modal from '$lib/components/Modal.svelte';
	import { addFormContributor } from '$lib/dashboard/stores/store';
	import CreateFormInput from '$lib/components/dashboard/CreateFormInput.svelte';
	import Button from '$lib/components/Button.svelte';
	import RoleMultipleSelect from './RoleMultipleSelect.svelte';
	import type { RoleOption } from '../types';
	import { toast } from '$lib/dashboard/stores/toast';
	import Heading from '$lib/components/dashboard/typography/Heading.svelte';
	import { onMount, tick } from 'svelte';
	import { size } from '@floating-ui/dom';
	import { availableRoles, contributors_loading, roles } from '../vars';
	import { api } from '$lib/api/api';
	import { contributors } from '$lib/dashboard/stores/data.store';
	import { fade } from 'svelte/transition';
	import axios from 'axios';

	type Props = {
		onClose?: () => void;
		opened?: boolean;
	};

	let loading: boolean = $state(false);

	let { onClose, opened = $bindable(true) }: Props = $props();
	let box: HTMLDivElement | null = null;
	let boxHeight: number | undefined = $state(0);
	let sized: boolean = $state(false);
	let password: string = $state('');

	function handleRoleChange(newRolesString: string) {
		$roles = newRolesString;
	}

	function watchSize(node: HTMLElement) {
		if (!sized) return;
		const ro = new ResizeObserver(() => {
			boxHeight = node.offsetHeight;
			sized = true;
			toast.show(boxHeight);
			console.log('Aktualna wysokość diva:', boxHeight);
		});
		ro.observe(node);
		return {
			destroy: () => ro.disconnect()
		};
	}

	$effect(() => {
		if (opened && box) {
			tick().then(() => {
				boxHeight = box?.offsetHeight;
				toast.show(`Wysokość: ${boxHeight}px`);
			});
		}
	});

	async function CreateContributor() {
		loading = true;
		$contributors_loading = true;
		onClose?.();
		let message = '';
		try {
			const response = await api.contributor.CreateContributor($addFormContributor);

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

	onMount(() => {
		$addFormContributor.password = '';
	});
</script>

<svelte:boundary>
	{#snippet failed(error, reset)}
		<div in:fade={{ duration: 150 }} class="m-5 w-7xl border-1 border-red-400 bg-red-500/70 p-5">
			{error}
		</div>
	{/snippet}
	<Modal
		title="Dodaj nowego kontrybutora"
		className="w-7xl"
		onClose={() => onClose?.()}
		bind:opened
	>
		<div bind:this={box} use:watchSize class={`my-4 grid h-[600px] grid-cols-2 flex-col gap-10`}>
			<div class="col-span-1 flex flex-col gap-4">
				<Heading>Informacje</Heading>
				<CreateFormInput label="Nazwa kontrybutora" bind:value={$addFormContributor.name} />
				<CreateFormInput label="Opis kontrybutora" bind:value={$addFormContributor.description} />
				<CreateFormInput label="Login" bind:value={$addFormContributor.login} />
				<CreateFormInput label="Hasło dla użytkownika" bind:value={$addFormContributor.password} />
				<CreateFormInput label="Potwierdź hasło" bind:value={password} />
			</div>

			<div in:fade={{ duration: 150, delay: 500 }}>
				<img
					src="http://imgs.search.brave.com/6KolxIpzcLodT_Q6QMyP0rqNGWfjOPDQDhAHhKIwaiw/rs:fit:500:0:1:0/g:ce/aHR0cHM6Ly9wcmV2/aWV3LnJlZGQuaXQv/c2hvdy1tZS15b3Vy/LWZhdi1jYXQtbWVt/ZS12MC1hcngza3V2/dm8yeGQxLmpwZWc_/d2lkdGg9MTE3MCZm/b3JtYXQ9cGpwZyZh/dXRvPXdlYnAmcz04/NTFlNzNjZjk1ODlh/OGY0YjJlMTlhMjQ2/ODE0YmI4ZTBhMjQ0/Mjdl"
					alt="cat"
				/>
			</div>
		</div>
		<div class="my-5 flex justify-end">
			<Button {loading} onclick={CreateContributor} size="small" theme="secondary"
				>Dodaj nowego kontrybutora</Button
			>
		</div>
	</Modal>
</svelte:boundary>
