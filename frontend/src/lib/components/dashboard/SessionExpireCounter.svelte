<script lang="ts">
	import { jwtDecode } from 'jwt-decode';
	import { onMount, onDestroy } from 'svelte';
	import { browser } from '$app/environment';
	import { goto } from '$app/navigation';
	import { userInfo } from '$lib/dashboard/stores/store';
	import { accounts } from '$lib/dashboard/stores/persist';
	type Props = {
		exp: number;
	};

	let timeLeft = $state(0);
	let clock = $state('');
	let interval: ReturnType<typeof setInterval>;
	let token: string | null = $state('');

	function updateTime() {
		const now = Math.floor(Date.now() / 1000);
		timeLeft = Math.max($userInfo.exp - now, 0);

		if (timeLeft === 0) {
			goto('/login');
		}

		const hours = Math.floor(timeLeft / 3600)
			.toString()
			.padStart(2, '0');
		const minutes = Math.floor((timeLeft % 3600) / 60)
			.toString()
			.padStart(2, '0');
		const seconds = (timeLeft % 60).toString().padStart(2, '0');

		clock = `${hours}:${minutes}:${seconds}`;
	}

	onMount(() => {
		if (!browser) return;

		const tokenString = localStorage.getItem('token');
		token = tokenString;
		if (token) {
			$userInfo = jwtDecode(token);
			console.log($userInfo);
			const exists = $accounts.some(acc => acc.login === $userInfo.login);
			const newData = {...userInfo, token}

			if (!exists) {
				$accounts = [...$accounts];
				console.log($accounts)
			} else {
				console.log("Konto o tym loginie już istnieje!");
				console.log($accounts)
			}
		}

		updateTime();
		interval = setInterval(updateTime, 1000);
	});

	onDestroy(() => {
		clearInterval(interval);
	});
</script>

<div
	class="flex h-8 w-50 items-center justify-center border-1 border-neutral-700/60 bg-neutral-800/60 text-white"
>
	{#if timeLeft > 0}
		Sesja wygaśnie w {clock}
	{/if}
</div>
