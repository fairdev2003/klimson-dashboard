<script lang="ts">
	import Modal from '$lib/components/Modal.svelte';
	import { dashboard_config, type Bookmark } from '$lib/dashboard/stores/persist';
	import Icon from '@iconify/svelte';
	import Input from '../../settings/components/Input.svelte';
	import Button from '$lib/components/Button.svelte';

	let bookmarkModalOpened: boolean = $state(false);

	let selectedColor: string = $state('');

	let bookmark: Bookmark = $state({
		color: '',
		href: '',
		name: 'Bookmark 1',
		slug: ''
	});

	let bookmark_slug = $derived(bookmark.name.replaceAll(' ', '-').toLowerCase());

	function SubmitNewBookmark(bookmark: Bookmark) {
		const currentBookmarks = $dashboard_config.bookmarks || [];
		$dashboard_config.bookmarks = [...currentBookmarks, bookmark];

		bookmarkModalOpened = false;
	}
	let deleteOn: boolean = $state(false);
	let deleteModalOpened: boolean = $state(false);
	let selectedBookmarkToDelete: Bookmark | undefined = $state();

	const bookmark_color_pallete = [
		'bg-violet-800',
		'bg-red-800',
		'bg-green-800',
		'bg-indigo-800',
		'bg-amber-800',
		'bg-orange-800',
		'bg-yellow-800',
		'bg-neutral-800'
	];

	function DeleteBookmark(slug: string) {
		const currentBookmarks = $dashboard_config.bookmarks || [];

		$dashboard_config.bookmarks = currentBookmarks.filter((bookmark) => bookmark.slug !== slug);

		deleteModalOpened = false;
	}
</script>

<div class="justify-between w-full text-start lg:flex md:flex gap-1 h-full sm:hidden">
	<div class="flex gap-2 items-center justify-between">
		<div class="flex gap-10 items-center">
			<p class="font-black">Dashboard Hub</p>
			<div class="flex gap-2">
				{#each $dashboard_config.bookmarks as bookmark}
					<button
						onclick={() => {
							if (deleteOn) {
								selectedBookmarkToDelete = bookmark;
								deleteModalOpened = !deleteModalOpened;
								return;
							}
							window.open(bookmark.href, '_blank', 'noopener,noreferrer');
						}}
						class="p-2 px-3 hidden lg:flex gap-1 transition-colors {bookmark.color} cursor-pointer rounded-lg items-center"
					>
						<p class="text-xs">{bookmark.name}</p>
					</button>
				{/each}
			</div>
		</div>
		<div class="flex gap-2">
			<button
				onclick={() => {
					bookmarkModalOpened = !bookmarkModalOpened;
				}}
				class="p-2 px-3 hidden lg:flex gap-1 transition-colors bg-neutral-800 cursor-pointer rounded-lg items-center hover:bg-neutral-700"
			>
				<Icon icon="material-symbols:bookmark" width="15" height="15" />
				<p class="text-xs">Add Bookmark</p>
			</button>
			<button
				onclick={() => {
					deleteOn = !deleteOn;
				}}
				class:delete-on={deleteOn}
				class="p-2 px-3 hidden border-2 border-neutral-800 lg:flex gap-1 transition-colors bg-neutral-800 cursor-pointer rounded-lg items-center hover:bg-neutral-700"
			>
				<Icon icon="mdi:trash" width="15" height="15" />
			</button>
		</div>
	</div>
	<!-- <p class="animated-gradient-text font-black text-sm">{$dashboard_config.dock_custom_name}</p> -->
</div>

<Modal
	bind:opened={bookmarkModalOpened}
	onClose={() => {
		bookmarkModalOpened = !bookmarkModalOpened;
	}}
	title="Add bookmark"
	className="w-100 "
>
	<div class="flex flex-col gap-2">
		<div class="flex flex-col gap-2">
			<p class="font-black uppercase text-xs text-neutral-400">bookmark name</p>
			<input
				bind:value={bookmark.name}
				placeholder="Bookmark name..."
				class="bg-neutral-800 input p-2 font-medium dropdown-button flex items-center gap-1 justify-center w-full border hover:bg-neutral-800 border-neutral-700 rounded-xl"
			/>
		</div>
		<div class="flex flex-col gap-2">
			<p class="font-black uppercase text-xs text-neutral-400">link</p>
			<input
				bind:value={bookmark.href}
				placeholder="Bookmark link..."
				class="bg-neutral-800 input p-2 dropdown-button font-medium flex items-center gap-1 justify-center w-full border hover:bg-neutral-800 border-neutral-700 rounded-xl"
			/>
		</div>
		<div class="flex flex-col gap-2 mb-5">
			<p class="font-black uppercase text-xs text-neutral-400">color</p>
			<div class="flex flex-wrap gap-2">
				{#each bookmark_color_pallete as color}
					<button
						onclick={() => {
							if (bookmark.color === color) {
								bookmark.color = '';
								return;
							}

							bookmark.color = color;
						}}
						class:selected-color-box={color === bookmark.color}
						class="relative size-7 rounded-lg cursor-pointer flex justify-center items-center {color}"
					>
						{#if color === bookmark.color}
							<div class="text-green-500">
								<Icon icon="material-symbols:check" width="20" height="20" />
							</div>
						{/if}
					</button>
				{/each}
			</div>
		</div>
	</div>
	<div class="flex justify-end">
		<Button
			onclick={() =>
				SubmitNewBookmark({
					color: bookmark.color,
					href: bookmark.href,
					name: bookmark.name,
					slug: bookmark_slug
				})}
			size="small"
			theme="secondary">Add bookmark</Button
		>
	</div>
</Modal>

<Modal
	bind:opened={deleteModalOpened}
	onClose={() => {
		deleteModalOpened = !deleteModalOpened;
	}}
	title="Are u sure?"
	className="w-100"
>
	<p>Are you sure you want to delete this bookmark?</p>
	<div class="flex justify-end mt-5">
		<Button
			onclick={() => {
				if (selectedBookmarkToDelete) {
					DeleteBookmark(selectedBookmarkToDelete.slug);
				}
			}}
			size="small"
			theme="danger"
		>
			Delete
		</Button>
	</div>
</Modal>

<style>
	@import 'tailwindcss';

	.selected-color-box {
		@apply border-2 border-green-500 bg-green-500/50;
	}
	.delete-on {
		@apply border-2 border-red-500 bg-red-500/50;
	}

	input[type='color'] {
		border: none;
		width: 50px;
		height: 50px;
		cursor: pointer;
		background: none;
	}
	.animated-gradient-text {
		background: linear-gradient(
			90deg,
			#ff0000,
			#ff7f00,
			#ffff00,
			#00ff00,
			#0000ff,
			#4b0082,
			#9400d3
		);
		background-size: 300% 100%;

		-webkit-background-clip: text;
		background-clip: text;

		color: transparent;

		animation: gradient-move 5s linear infinite;
	}

	@keyframes gradient-move {
		0% {
			background-position: 0% 50%;
		}
		50% {
			background-position: 100% 50%;
		}
		100% {
			background-position: 0% 50%;
		}
	}
</style>
