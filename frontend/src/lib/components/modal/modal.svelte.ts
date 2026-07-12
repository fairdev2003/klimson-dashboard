import { debug } from '$lib/dashboard/stores/debug';
import type { Snippet } from 'svelte';

export type ModalProps = {
	opened: boolean;
	children?: Snippet;
	onClose?: () => void;
	className?: string;
	size: 'auto' | 'accept_preset' | 'form_preset';
	border: 'normal' | 'borderless';
	padding_preset: 'zero' | 'normal' | 'small' | 'big';
	title?: string;
};

export class ModalLogic {
	public props = $state<ModalProps>({ opened: true });
	constructor(public initialProps: ModalProps) {
		this.props = initialProps;
		$effect(() => {
			debug.log(this.props.opened);
		});
	}

	public async on_background_click() {
		this.props.opened = false;
	}

	public async on_exit_icon_click() {
		if (this.props.onClose) {
			this.props.onClose();
		}
		this.props.opened = false;
	}

	public async on_content_click(e: MouseEvent) {
		e.stopPropagation();
	}
}
