<script lang="ts">
	import { ModalLogic, type ModalProps } from './modal.svelte';
	import { tv } from 'tailwind-variants';
	import RDBModalBar from './RDBModalBar.svelte';
	import RDBSubmitControls from './RDBSubmitControls.svelte';

	let {
		opened = $bindable(false),
		children,
		className,
		onClose,
		title,
		size,
		border,
		padding_preset,
		form_config,
		backgroundExitLocked = $bindable(),
		titleStyle
	}: ModalProps = $props();

	const modal = new ModalLogic({
		get opened() {
			return opened;
		},
		set opened(v) {
			opened = v;
		},
		children,
		className,
		onClose,
		title,
		size,
		border,
		padding_preset,
		backgroundExitLocked,
		titleStyle
	});

	const modalStyles = tv({
		base: 'z-100 flex flex-col gap-2 ',
		variants: {
			bg_color: {
				classic: 'bg-neutral-900'
			},
			round_size: {
				normal: 'rounded-lg',
				large: 'rounded-xl',
				small: 'rounded-md'
			},
			size: {
				auto: 'w-auto',
				accept_preset: 'lg:w-80 ',
				form_preset: 'lg:w-150 h-7/10'
			},
			screen_size: {},
			border: {
				normal: 'border border-neutral-700',
				form: 'border-2 border-blue-500',
				borderless: 'border-0'
			},
			text_color: {
				normal: 'text-white'
			},
			padding_preset: {
				zero: 'p-0',
				normal: 'p-4',
				small: 'p-2',
				big: 'p-6'
			}
		},
		defaultVariants: {
			bg_color: 'classic',
			round_size: 'normal',
			size: 'auto',
			border: 'normal',
			text_color: 'normal',
			padding_preset: 'normal'
		}
	});
</script>

{#if modal.props.opened}
	<div class="fixed inset-0 z-200">
		<div
			class="absolute flex flex-col inset-0 bg-black/50 backdrop-blur-sm"
			onclick={async () => {
				await modal.on_background_click();
			}}
		></div>

		<div class="flex flex-col justify-center items-center w-full h-full">
			<div
				bind:this={modal.modalContainer}
				onclick={(e) => {
					modal.on_content_click(e);
				}}
				class={`${className} ${modalStyles({ size, border, padding_preset })} flex flex-col max-h-[90vh]`}
			>
				<RDBModalBar {title} {titleStyle} onButtonClick={() => modal.on_exit_icon_click()} />

				<div class="flex flex-col flex-1 min-h-0 justify-between w-full">
					<div class="overflow-auto flex-1 min-h-0 w-full px-1">
						{@render children?.()}
					</div>
					<RDBSubmitControls {...form_config} />
				</div>
			</div>
		</div>
	</div>
{/if}
