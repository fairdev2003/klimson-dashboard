<script lang="ts">
	import { goto } from '$app/navigation';
	import { api } from '$lib/api/api';
	import Button from '$lib/components/Button.svelte';
	import type { AxiosResponse } from 'axios';
	import { onMount } from 'svelte';
	import { blur } from 'svelte/transition';
	import AnimatedPadlock from './(components)/AnimatedPadlock.svelte';
	import type { BackendResponse, ServerResponse } from '$lib/api/types';
	import FancyLoader from '../dashboard/redis_writable/(components)/FancyLoader.svelte';

	let pass: string = $state('');
	let login: string = $state('');
	let error: string = $state('');
	let loading: boolean = $state(false);
	let cenzurka: boolean = $state(false);
	let session_authorized: boolean = $state(false);

	type TokenData = {
		token: string;
	};

	let key = $state(false);

	async function handleLogin() {
		if (!login || !pass || loading) return;

		error = '';
		loading = true;

		try {
			const response: AxiosResponse<TokenData> = await api.api.post(
				'/login',
				{
					password: pass,
					login: login
				},
				{ withCredentials: true }
			);

			try {
				const auth: ServerResponse<{ access: boolean }> = await api.api.get('/admin/verify');

				session_authorized = auth.data.access;
			} catch (error) {}
		} catch (e: any) {
			if (e.response && e.response.data && e.response.data.error) {
				error = e.response.data.error;
			} else if (e.code === 'ECONNABORTED') {
				error = 'Serwer zbyt długo nie odpowiada...';
			} else {
				error = 'Coś poszło nie tak. Spróbuj ponownie.';
			}
		} finally {
			loading = false;
		}
	}

	onMount(() => {
		padlock_loading = true;
		setTimeout(() => {
			padlock_loading = false;
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
					placeholder="Your login..."
					disabled={padlock_loading}
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
					placeholder="••••••••"
					disabled={padlock_loading}
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
