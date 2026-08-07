<script lang="ts">
	import Icon from '@iconify/svelte';
	import type { TitleStyle } from './modal.svelte';
	import { tv } from 'tailwind-variants';

	type Props = {
		onButtonClick: () => void;
		title?: string;
		titleStyle: TitleStyle;
	};

	let { onButtonClick, title, titleStyle }: Props = $props();

	const titleStyles = tv({
		base: 'uppercase text-xs font-black text-white',
		variants: {
			titleStyle: {
				basic: 'text-text',
				danger: 'text-red-400'
			}
		},
		defaultVariants: {
			titleStyle: 'basic'
		}
	});
</script>

<div class="flex justify-between items-center">
	<p class={titleStyles({ titleStyle })}>{title}</p>
	<button
		onclick={onButtonClick}
		class="p-2 hover:bg-neutral-700/50 hover:text-red-400 text-text rounded-xl cursor-pointer"
	>
		<Icon icon="material-symbols:close" />
	</button>
</div>
