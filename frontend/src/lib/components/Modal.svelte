<!-- svelte-ignore a11y_no_static_element_interactions -->
<!-- svelte-ignore a11y_click_events_have_key_events -->
<script lang="ts">
	import { X } from '@lucide/svelte';
	import { tick, type Snippet } from 'svelte';
	import gsap from 'gsap';
	import Button from './Button.svelte';
	import { animation_preset, type AnimationPresetType } from '$lib/dashboard/stores/persist';

	type Props = {
		opened: boolean;
		children?: Snippet;
		onClose?: () => void;
		title?: string;
		className?: string;
		animationPreset?: AnimationPresetType;
		editMode?: boolean;
		exitMode?: boolean;
		draggable?: boolean;
	};

	let pos = $state({ x: 20, y: 0 });
	let isDragging = $state(false);

	$effect(() => {
		if (modalEl) {
			tick().then(() => {
				modalEl!.scrollTo({ top: modalEl!.scrollHeight, behavior: 'smooth' });
			});
		}
	});

	function handleMouseDown(e: MouseEvent) {
		isDragging = true;
		const onMouseMove = (m: MouseEvent) => {
			pos.x += m.movementX;
			pos.y += m.movementY;
		};
		const onMouseUp = () => {
			isDragging = false;
			window.removeEventListener('mousemove', onMouseMove);
			window.removeEventListener('mouseup', onMouseUp);
		};
		window.addEventListener('mousemove', onMouseMove);
		window.addEventListener('mouseup', onMouseUp);
	}

	const {
		opened = $bindable(),
		children,
		title,
		onClose,
		exitMode = true,
		className = 'max-w-[1400px] h-[95%]',
		editMode = false,
		// experimental
		draggable = false
	}: Props = $props();

	let visible = $state(false);
	let modalEl: HTMLDivElement;
	let confirmModalEl: HTMLDivElement;
	let confirm_modal_opened: boolean = $state(false);

	function blurAnimation(side: 'in' | 'out') {
		if (side === 'in') {
			gsap.fromTo(
				modalEl,
				{ scale: 0.8, opacity: 0 },
				{ scale: 1, opacity: 1, duration: 0.35, ease: 'power2.out' }
			);
		}

		if (side === 'out') {
			gsap.to(modalEl, {
				scale: 0.8,
				opacity: 0,
				duration: 0.35,
				ease: 'power2.in',
				onComplete: () => {
					visible = false;
					document.body.style.overflow = 'auto';
				}
			});
		}
	}

	function jasonAnimation(side: 'in' | 'out') {
		if (side === 'in') {
			gsap.fromTo(
				modalEl,
				{ scaleY: 0.2, scaleX: 0.15, transformOrigin: 'bottom', opacity: 0, y: 200 },
				{ scaleY: 1, scaleX: 1, duration: 0.4, opacity: 1, ease: 'power2.out', y: 0 }
			);
		}

		if (side === 'out') {
			gsap.to(modalEl, {
				scaleY: 0.2,
				scaleX: 0.15,
				y: 200,
				opacity: 0,
				transformOrigin: 'bottom',
				duration: 0.3,
				ease: 'power2.in',
				onComplete: () => {
					visible = false;
					document.body.style.overflow = 'auto';
				}
			});
		}
	}

	function klimsonAnimation(side: 'in' | 'out') {
		if (side === 'in') {
			gsap.fromTo(
				modalEl,
				{ scaleY: 0.2, scaleX: 0.15, transformOrigin: 'bottom', opacity: 0 },
				{ scaleY: 1, scaleX: 1, duration: 0.4, opacity: 1, ease: 'power2.out' }
			);
		}

		if (side === 'out') {
			gsap.to(modalEl, {
				scaleY: 0.2,
				scaleX: 0.15,
				opacity: 0,
				transformOrigin: 'bottom',
				duration: 0.3,
				ease: 'power2.in',
				onComplete: () => {
					visible = false;
					document.body.style.overflow = 'auto';
				}
			});
		}
	}

	async function animate(side: 'in' | 'out', prs: AnimationPresetType) {
		if (prs === 'blur') {
			blurAnimation(side);
		}
		if (prs === 'klimson') {
			klimsonAnimation(side);
		}
		if (prs === 'jason') {
			jasonAnimation(side);
		}
	}

	$effect(async () => {
		if (opened) {
			visible = true;
			document.body.style.overflow = 'hidden';
			await tick();

			animate('in', $animation_preset);
		} else if (visible) {
			animate('out', $animation_preset);
		}
	});

	$effect(async () => {
		if (confirm_modal_opened) {
			visible = true;
			document.body.style.overflow = 'hidden';
			await tick();

			animate('out', $animation_preset);
		} else if (!confirm_modal_opened) {
			animate('out', $animation_preset);
		}
	});
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
{#if visible}
	<!-- svelte-ignore a11y_click_events_have_key_events -->
	{#if draggable}
		{@render Modal(true)}
	{:else}
		<div
			class="fixed inset-0 z-50 flex items-center justify-center backdrop-blur-sm"
			onclick={() => onClose?.()}
		>
			{@render Modal(false)}
		</div>
	{/if}
{/if}

{#if confirm_modal_opened}
	{@render ConfirmModal()}
{/if}

{#snippet Modal(d: boolean)}
	<!-- svelte-ignore a11y_click_events_have_key_events -->
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<div
		bind:this={modalEl}
		onclick={(e) => {
			e.stopPropagation();
		}}
		style="left: {pos.x}px; bottom: {20 - pos.y}px; {isDragging ? 'z-index: 1000' : ''}"
		class={`${d ? 'absolute' : 'relative'} ${className}  rounded-lg border border-neutral-800 bg-neutral-950 text-white `}
	>
		<!-- HEADER -->
		<!-- svelte-ignore a11y_no_static_element_interactions -->
		<div
			onmousedown={handleMouseDown}
			class="flex h-10 cursor-grab items-center justify-between border-b border-neutral-800 px-4 active:cursor-grabbing"
		>
			<p class="truncate text-sm font-semibold">{title}</p>
			<button
				onclick={() => {
					if (editMode) {
						confirm_modal_opened = true;
						return;
					}

					onClose?.();
				}}
				class="text-neutral-400 hover:text-white"
			>
				<X />
			</button>
		</div>

		<!-- CONTENT -->
		<div class="h-[95%] overflow-y-scroll p-4">
			{#if children}
				{@render children()}
			{/if}
		</div>
	</div>
{/snippet}

{#snippet ConfirmModal()}
	<!-- svelte-ignore a11y_click_events_have_key_events -->
	<div
		class="fixed inset-0 z-100 flex items-center justify-center bg-black/50 backdrop-blur"
		onclick={() => {
			confirm_modal_opened = false;
		}}
	>
		<div
			bind:this={confirmModalEl}
			onclick={(e) => {
				e.stopPropagation();
			}}
			class={`relative w-[400px] rounded-lg border border-neutral-800 bg-neutral-950 text-white shadow-xl`}
		>
			<!-- HEADER -->
			<div class="flex h-10 items-center justify-between border-b border-neutral-800 px-4">
				<p class="text-sm font-semibold">Czy na pewno chcesz wyjść</p>
				<button
					onclick={() => {
						confirm_modal_opened = false;
					}}
					class="text-neutral-400 hover:text-white"
				>
					<X />
				</button>
			</div>

			<!-- CONTENT -->
			<div class="h-[95%] overflow-y-scroll p-4">
				<p>Twoje zmiany nie zostana zapisane</p>
			</div>
			<div class="flex w-full justify-end gap-2 p-4">
				<Button
					size="small"
					onclick={async () => {
						confirm_modal_opened = false;
					}}>Anuluj</Button
				>
				<Button
					size="small"
					theme="secondary"
					onclick={async () => {
						confirm_modal_opened = false;
						onClose?.();
					}}>Nie zapisuj</Button
				>
			</div>
		</div>
	</div>
{/snippet}

<svelte:document
	onkeydown={(e) => {
		if (e.key === 'Escape') {
			if (editMode) {
				confirm_modal_opened = true;
				return;
			}
			onClose?.();
		}
	}}
/>
