<script lang="ts">
	import { api } from '$lib/api/api';
	import type { BackendResponse, ServerResponse } from '$lib/api/types';
	import { onMount } from 'svelte';
	import { dockComponent } from '../dashboard.svelte';
	import RedisWritableDocky from '$lib/components/dashboard/dock/boxes/RedisWritableDocky.svelte';
	import Heading from '$lib/components/dashboard/typography/Heading.svelte';
	import Icon from '@iconify/svelte';
	import Button from '$lib/components/Button.svelte';
	import './redis_hub.css';
	import PageLoading from '$lib/components/PageLoading.svelte';
	import Loader from '$lib/components/dashboard/Loader.svelte';
	import { blur, slide } from 'svelte/transition';
	import { toast } from '$lib/dashboard/stores/toast';
	import FancyLoader from './(components)/FancyLoader.svelte';
	import { goto } from '$app/navigation';
	import RDBModal from '$lib/components/modal/RDBModal.svelte';
	import { debug } from '$lib/dashboard/stores/debug';

	async function get_tables(): Promise<ServerResponse<BackendResponse<{ rdbs: string[] }>>> {
		const response: ServerResponse<BackendResponse<{ rdbs: string[] }>> =
			await api.api.get('/redis/keys');

		return response;
	}

	async function get(key: string) {
		const response: ServerResponse<BackendResponse<{ result: any }>> = await api.api.get(
			`/redis/get?key=${key}`
		);

		return response;
	}

	onMount(() => {
		dockComponent.set(RedisWritableDocky);
	});

	type RedisHubView = 'start' | 'react';
	let current_view: RedisHubView = $state('start');

	function change_view(view: RedisHubView) {
		current_view = view;
	}

	let redisWritableForm: { key: string; value: string } = $state({ key: '', value: '' });
	let savedFirstValue = $state('');

	let testModalOpened: boolean = $state(false);
</script>

<div class="p-7 flex flex-col gap-4">
	<div class="flex justify-between items-center border-b border-neutral-700 pb-4">
		<div class="flex-col flex gap-1">
			<Heading>
				<div
					class="flex gap-2 items-center overflow-hidden"
					in:slide={{ duration: 300, delay: 300 }}
				>
					<Icon icon="devicon:redis" />
					<p class="text-red-500">Redis Storage</p>
				</div>
			</Heading>
			<span class="text-sm font-md text-neutral-400"
				><b>Fast</b> and easy solution to store one-way data. Redis data is reactive!</span
			>
			<div class="flex mt-4 gap-2">
				<button
					onclick={() => change_view('start')}
					class="base-label-pill"
					class:normal-label-pill={current_view !== 'start'}
					class:selected-label-pill={current_view === 'start'}>Current Redis Keys</button
				>
				<button
					class:normal-label-pill={current_view !== 'react'}
					class:selected-label-pill={current_view === 'react'}
					onclick={() => change_view('react')}
					class="base-label-pill">Reactive State</button
				>
			</div>
		</div>

		<div class="flex gap-4">
			<!-- <Button theme="base">Dump data</Button>
			<!-- <Button theme="base">Implement role</Button> -->
			<Button
				theme="secondary"
				onclick={() => {
					testModalOpened = !testModalOpened;
				}}>Open test modal</Button
			>
		</div>
	</div>
	{#if current_view === 'start'}
		{#await get_tables()}
			<div class="flex mx-auto">
				<FancyLoader />
			</div>
		{:then ping_data}
			<div class="flex flex-col gap-2 w-2xl mx-auto" in:blur={{ duration: 300 }}>
				{#each ping_data.data.rdbs as rdb}
					{#await get(rdb)}
						<p>Loading</p>
					{:then data}
						<div class="bg-neutral-800 justify-between flex rounded-lg p-3 px-7 items-center">
							<div class="flex flex-col">
								<span class="flex gap-1 items-center">
									<Icon icon="devicon:redis" />
									<p
										onclick={() => {
											goto(`/dashboard/redis_writable/${rdb}/info`);
										}}
										class="font-black text-red-500 hover:underline cursor-pointer"
									>
										{rdb}
									</p>
								</span>
								<p
									onclick={() => {
										navigator.clipboard.writeText(data.data.result);
										toast.success('Copied to clipboard!');
									}}
									class="hover:underline cursor-pointer"
								>
									{data.data.result}
								</p>
							</div>
							<div class="flex items-center gap-2">
								<button
									onclick={() => {
										testModalOpened = !testModalOpened;
										savedFirstValue = data.data.result;
										redisWritableForm = { key: rdb, value: data.data.result };
									}}
									class="p-2 hover:bg-neutral-700/50 hover:text-blue-400 rounded-xl cursor-pointer"
								>
									<Icon icon="boxicons:edit-filled" width="20" height="20" />
								</button>
								<button
									class="p-2 hover:bg-neutral-700/50 hover:text-red-400 rounded-xl cursor-pointer"
								>
									<Icon icon="boxicons:trash-filled" width="20" height="20" />
								</button>
							</div>
						</div>
					{/await}
				{/each}
			</div>
		{/await}
	{/if}

	{#if current_view === 'react'}
		<FancyLoader color="white" />
	{/if}
</div>
<RDBModal
	form_config={{
		onSubmit: () => {
			debug.log('Submit!');
			testModalOpened = !testModalOpened;
		}
	}}
	title={`Editing "${redisWritableForm.key}"`}
	border="borderless"
	size="form_preset"
	bind:opened={testModalOpened}
>
	<div class="flex gap-4 flex-col">
		<div class="flex flex-col gap-1">
			<p class="text-neutral-400 mb-1 uppercase font-bold text-xs">PREVIEW</p>

			<div class="bg-neutral-800 gap-4 flex rounded-lg justify-between p-3 items-center">
				<div class="flex gap-2 items-center">
					<Icon icon="devicon:redis" />
					<p class="font-black text-red-500">{redisWritableForm.key}</p>
				</div>
				{#key redisWritableForm.value}
					<p
						onclick={() => {
							navigator.clipboard.writeText(redisWritableForm.value);
							toast.success('Copied to clipboard!');
						}}
						in:blur={{ duration: 300, delay: 300 }}
						class="hover:underline cursor-pointer"
					>
						{redisWritableForm.value}
					</p>
				{/key}
			</div>
		</div>
		<div class="flex flex-col gap-1">
			<span class="text-neutral-400 uppercase justify-between items-center flex font-bold text-xs">
				<p>VALUE</p>
				<button
					onclick={() => {
						redisWritableForm = { ...redisWritableForm, value: savedFirstValue };
					}}
					class="hover:text-white cursor-pointer mr-1"
				>
					<Icon icon="material-symbols:undo" width="20" height="20" />
				</button>
			</span>

			<input bind:value={redisWritableForm.value} class="rounded-lg border-0 bg-neutral-800" />
		</div>
	</div>
</RDBModal>
