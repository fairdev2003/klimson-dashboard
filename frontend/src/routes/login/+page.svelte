<script lang="ts">
	import { goto } from '$app/navigation';
	import { api } from '$lib/api/api';
	import Button from '$lib/components/Button.svelte';
	import type { AxiosResponse } from 'axios';
	import { onMount } from 'svelte';
	import { blur } from 'svelte/transition';

	let pass: string = $state('');
	let login: string = $state('');
	let error: string = $state('');
	let loading: boolean = $state(false);

	type TokenData = {
		token: string;
	};

	let key = $state(false);

	async function handleLogin() {
		if (!login || !pass || loading) return;

		error = '';
		loading = true;

		try {
			const response: AxiosResponse<TokenData> = await api.api.post('/login', {
				password: pass,
				login: login
			});

			const token = response.data.token;
			localStorage.setItem('token', token);

			goto('/dashboard');
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
		key = true;
	});
</script>

<div
	class="mx-auto mt-10 flex w-full max-w-md flex-col items-center justify-center gap-6 rounded-lg border border-neutral-800 bg-neutral-900 p-8 text-white shadow-2xl"
>
	<div class="flex flex-col items-center gap-2">
		<h1 class="text-2xl font-bold tracking-tight">Witaj ponownie</h1>
		<p class="text-sm text-neutral-400">Zaloguj się do DOJEBANEGO panelu</p>
	</div>

	<div class="mt-5 w-full space-y-4">
		<div class="flex flex-col gap-2">
			<label for="login" class="text-xs tracking-widest text-neutral-500 uppercase">Login</label>
			<input
				id="login"
				type="text"
				bind:value={login}
				placeholder="Twój login..."
				class="w-full rounded border border-neutral-700 bg-neutral-800 p-3 text-white transition-all focus:border-blue-500 focus:ring-1 focus:ring-blue-500 focus:outline-none"
			/>
		</div>

		<div class="flex flex-col gap-2">
			<label for="password" class="text-xs tracking-widest text-neutral-500 uppercase">Hasło</label>
			<input
				id="password"
				type="password"
				bind:value={pass}
				placeholder="••••••••"
				class="w-full rounded border border-neutral-700 bg-neutral-800 p-3 text-white transition-all focus:border-blue-500 focus:ring-1 focus:ring-blue-500 focus:outline-none"
			/>
		</div>
	</div>

	<Button
		theme="secondary"
		{loading}
		className="flex justify-center w-full py-3 font-semibold transition-transform  text-center"
		onclick={handleLogin}
	>
		{loading ? 'Logowanie...' : 'Zaloguj się'}
	</Button>

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
		src="https://api.klimson.dev/interface/bucket/random/nugget_cat.png"
	/>
{/key}

<svelte:window
	onkeydown={(e) => {
		if (e.key === 'Enter') handleLogin();
	}}
/>
