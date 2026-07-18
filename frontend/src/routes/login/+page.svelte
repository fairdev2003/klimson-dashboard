<script lang="ts">
	import { goto } from '$app/navigation';
	import { api } from '$lib/api/api';
	import Button from '$lib/components/Button.svelte';
	import type { AxiosResponse } from 'axios';
	import { onMount } from 'svelte';
	import { blur } from 'svelte/transition';
	import AnimatedPadlock from './(components)/AnimatedPadlock.svelte';
	import type { BackendResponse, ServerResponse } from '$lib/api/types';
	import FancyLoader from '../dashboard/redis/(components)/FancyLoader.svelte';
	import axios from 'axios';

	let pass: string = $state('');
	let login: string = $state('');
	let error: string = $state('');
	let loading: boolean = $state(false);
	let cenzurka: boolean = $state(false);
	let session_authorized: boolean = $state(false);
	let disabled_user_input = $state(false);

	type TokenData = {
		token: string;
	};

	let key = $state(false);

	async function CheckIfAuthorized() {
		try {
			const response = await api.api.get('/admin/verify');

			if (response.status === 200) {
				session_authorized = true;
				padlock_loading = false;
				disabled_user_input = true;
			}
		} catch (error) {
			if (axios.isAxiosError(error)) {
				if (error.status === 401) {
					session_authorized = false;
					padlock_loading = false;
					disabled_user_input = false;
				}
			}
		}
	}

	$effect(() => {
		if (pass.length > 0 && pass.length < 3) {
			error = 'Hasło jest za krótkie!';
		} else if (error === 'Hasło jest za krótkie!') {
			error = '';
		}
	});

	async function handleLogin() {
		if (!login || !pass || loading) return;

		error = '';
		loading = true;

		try {
			const response: AxiosResponse<TokenData> = await api.api.post(
				'/login',
				{ password: pass, login: login },
				{ withCredentials: true }
			);

			if (response.status === 200) {
				session_authorized = true;
				padlock_loading = false;
				disabled_user_input = true;
			}
		} catch (e: any) {
			console.log('Pełny obiekt błędu:', e);

			if (e.response && e.response.data) {
				error = e.response.data.message || 'Wystąpił nieznany błąd';
			} else {
				error = 'Błąd połączenia z serwerem';
			}
		} finally {
			loading = false;
		}
	}

	onMount(async () => {
		padlock_loading = true;
		setTimeout(async () => {
			await CheckIfAuthorized();
		}, 2000);
	});

	let padlock_loading = $state(true);
</script>

<div
	class:bg-neutral-950={padlock_loading}
	class:bg-neutral-900={!padlock_loading}
	class="mx-auto mt-10 flex w-full max-w-md flex-col items-center justify-center gap-6 rounded-lg border border-neutral-800 p-8 text-white shadow-2xl"
>
	<div class="pb-0 h-10">
		{#if padlock_loading}
			<FancyLoader color="blue" size="medium" />
		{:else}
			<AnimatedPadlock bind:padlock_loading bind:opened={session_authorized} />
		{/if}
	</div>
	<div class="flex flex-col items-center gap-2">
		<h1 class="text-2xl font-bold tracking-tight">Welcome back nigga</h1>
		<p class="text-sm text-neutral-400">Log into DOJEBANY dashboard</p>
	</div>

	<div class="mt-5 w-full space-y-4">
		<div class="flex flex-col gap-2">
			<label for="login" class="text-xs tracking-widest text-neutral-500 uppercase">Login</label>
			{#key session_authorized}
				<input
					id="login"
					type="text"
					bind:value={login}
					oninput={() => (error = '')}
					placeholder="Your login..."
					disabled={padlock_loading || disabled_user_input}
					class="w-full rounded border disabled:opacity-25 border-neutral-700 bg-neutral-800 p-3 text-white transition-all focus:border-blue-500 focus:ring-1 focus:ring-blue-500 focus:outline-none"
				/>
			{/key}
		</div>

		<div class="flex flex-col gap-2 relative">
			<label
				for="password"
				class="text-xs relative inline-block z-10 tracking-widest text-neutral-500 uppercase"
				>Password</label
			>
			{#key session_authorized}
				<input
					id="password"
					type="password"
					bind:value={pass}
					oninput={() => (error = '')}
					placeholder="••••••••"
					disabled={padlock_loading || disabled_user_input}
					class="w-full rounded border disabled:opacity-25 border-neutral-700 bg-neutral-800 p-3 text-white transition-all focus:border-blue-500 focus:ring-1 focus:ring-blue-500 focus:outline-none"
				/>

				{#if cenzurka}
					<div
						class="marker-overlay top-0 right-10 absolute inset-0 pointer-events-none rounded"
					></div>
					<div class="marker-overlay top-5 absolute inset-0 pointer-events-none rounded"></div>
				{/if}
			{/key}
		</div>
	</div>

	{#if session_authorized}
		<div in:blur={{ duration: 300 }} class="w-full">
			<Button
				theme="base"
				className="flex justify-center w-full py-3 font-semibold transition-transform  text-center"
				onclick={() => {
					goto('/dashboard');
				}}
			>
				Access Dashboard
			</Button>
		</div>
	{/if}

	{#if !session_authorized}
		<Button
			theme="secondary"
			disabled={padlock_loading}
			{loading}
			className="flex justify-center w-full py-3 font-semibold transition-transform  text-center"
			onclick={handleLogin}
		>
			{loading ? 'Authorizing...' : 'Sign in'}
		</Button>
	{/if}

	{#if error}
		<div class="w-full rounded border border-red-500/20 bg-red-500/10 p-3">
			<p class="text-center text-sm text-red-500">{error}</p>
		</div>
	{/if}
</div>

{#key key}
	<img
		in:blur={{ duration: 2000, delay: 2000 }}
		class="lg:flex absolute hidden top-1/2 left-50"
		src="https://api.klimson.dev/interface/bucket/random/banana.webp"
	/>

	<img
		in:blur={{ duration: 2000, delay: 2000 }}
		class="lg:flex absolute hidden top-1/3 right-50 size-1/4"
		src="https://api.klimson.dev/interface/bucket/random/paulinagt.jpg"
	/>
{/key}

<svelte:window
	onkeydown={(e) => {
		if (e.key === 'Enter') handleLogin();
	}}
/>

<style>
	.marker-overlay {
		position: absolute;
		left: 0;
		width: 100%;
		z-index: 1;
		height: 100%;
		pointer-events: none; /* Kliknięcia przechodzą przez marker */

		/* Tekstura markera */
		background: repeating-linear-gradient(15deg, #000, #000 5px, #333 5px, #333 5px);

		/* Klucz do "markerowego" wyglądu */
		filter: blur(2px) contrast(20);
		opacity: 0.8;
	}
</style>
