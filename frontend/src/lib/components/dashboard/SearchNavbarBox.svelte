<script lang="ts">
	import { onDestroy, tick } from 'svelte';
	import type { Quiz } from '../../../routes/dashboard/quizzes/types';
	import Placeholder from '$lib/assets/placeholder.png';
	import { api } from '$lib/api/api';

	import { fade } from 'svelte/transition';
	import { goto } from '$app/navigation';
	import { lastSearched, sidebar_open } from '$lib/dashboard/stores/persist';
	import type { Attachment } from 'svelte/attachments';
	import { toast } from '$lib/dashboard/stores/toast';
	import { get } from 'svelte/store';
	import { Trash, Trash2 } from '@lucide/svelte';

	let value: string = $state('');

	let boxRef: HTMLDivElement | null = null;
	let inputRef: HTMLInputElement | null = null;
</script>

<div class="relative hidden h-[50px] w-[500px] border-1 border-neutral-700/60 md:flex lg:flex">
	<input
		class="h-full w-full border-0 bg-transparent placeholder-neutral-500 ring-0 outline-none select-none focus:outline-none"
		bind:value
		placeholder="Kliknij aby wyszukać"
	/>
</div>

<style>
	input:invalid {
		outline: none;
		box-shadow: none;
	}
</style>
