import { debug } from '$lib/dashboard/stores/debug';
import type { Snippet } from 'svelte';
import gsap from 'gsap';

export type FormControlsProps = {
	onCancel?: () => void;
	onLog?: () => void;
	onLogout?: () => void;
	onDelete?: () => void;
	onSubmit?: () => void;
	onAccept?: () => void;
	onDeny?: () => void;
	initialForm?: any;
	currentForm?: any;
};

export type TitleStyle = 'basic' | 'danger';

export type ModalProps = {
	opened: boolean;
	children?: Snippet;
	onClose?: () => void;
	className?: string;
	size: 'auto' | 'accept_preset' | 'form_preset';
	border: 'normal' | 'borderless' | 'form';
	padding_preset: 'zero' | 'normal' | 'small' | 'big';
	title?: string;
	form_config?: FormControlsProps;
	backgroundExitLocked: boolean;
	titleStyle?: TitleStyle;
};

export class ModalLogic {
	public modalContainer: HTMLDivElement | undefined;
	public props = $state<ModalProps>({ opened: true });
	constructor(public initialProps: ModalProps) {
		this.props = initialProps;
	}

	private ShakeModalContainer() {
		if (!this.modalContainer) return;
		gsap.to(this.modalContainer, {
			x: -10,
			duration: 0.1,
			repeat: 3,
			yoyo: true,
			ease: 'power1.inOut'
		});
	}

	public async on_background_click() {
		if (this.props.backgroundExitLocked) {
			this.ShakeModalContainer();
			return;
		}

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
