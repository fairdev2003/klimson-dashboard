<script lang="ts">
	import Modal from '$lib/components/Modal.svelte';
	import { getContext } from 'svelte';
	import type { Quiz } from '../types';

	type Props = {
		key: string;
		value: string;
	};

	let editModalOpened: boolean = $state(false);

	let { key, value }: Props = $props();

	let quiz = getContext<Quiz>('quiz');
</script>

<div class="flex flex-col">
	<p class="text-neutral-400">{key}</p>
	<button
		onclick={() => {
			editModalOpened = true;
		}}
		class="bg-primary border-1 border-secondary w-100 hover:bg-secondary p-2 text-start text-neutral-400"
	>
		{value}
	</button>
</div>

<Modal
	title={`Edytuj: ${key}`}
	className="w-[400px] "
	bind:opened={editModalOpened}
	onClose={() => {
		editModalOpened = false;
	}}
>
	<header>
		<h3 class="text-lg font-semibold text-gray-800">
			Edytuj: <span class="text-blue-600">{key}</span>
		</h3>
		<p class="text-sm text-gray-500">Wprowadź nową wartość poniżej</p>
	</header>

	<div class="flex flex-col gap-2">
		<label class="relative inline-flex cursor-pointer items-center">
			<input type="checkbox" class="peer sr-only" />
			<div
				class="peer h-6 w-11 rounded-full bg-gray-200 after:absolute after:left-[2px] after:top-[2px] after:h-5 after:w-5 after:rounded-full after:border after:border-gray-300 after:bg-white after:transition-all after:content-[''] peer-checked:bg-blue-600 peer-checked:after:translate-x-full peer-checked:after:border-white peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300"
			></div>
			<span class="ml-3 text-sm font-medium text-gray-700"></span>
		</label>

		<input
			type="number"
			class="w-full rounded-lg border border-gray-300 px-3 py-2 outline-none transition-all focus:border-blue-500 focus:ring-2 focus:ring-blue-500"
		/>

		<input
			type="text"
			placeholder="Wpisz wartość..."
			class="w-full rounded-lg border border-gray-300 px-3 py-2 outline-none transition-all focus:border-blue-500 focus:ring-2 focus:ring-blue-500"
		/>
	</div>
</Modal>
