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
	onLockScreen?: () => void;
	initialForm?: any;
	currentForm?: any;
};

export type TitleStyle = 'basic' | 'danger';

export type ModalProps = {
	opened: boolean;
	children?: Snippet;
	onClose?: () => void;
	className?: string;
	size: 'auto' | 'accept_preset' | 'form_preset' | 'window';
	border: 'normal' | 'borderless' | 'form';
	padding_preset: 'zero' | 'normal' | 'small' | 'big';
	title?: string;
	form_config?: FormControlsProps;
	backgroundExitLocked: boolean;
	titleStyle?: TitleStyle;
	stickyBar?: Snippet;
	toolbarHidden?: boolean;
	bg_color?: 'classic' | 'blacker';
};

export class ModalLogic {
	public modalContainer: HTMLDivElement | undefined;
	public backdropContainer: HTMLDivElement | undefined;
	public props = $state<ModalProps>({ opened: true });
	constructor(public initialProps: ModalProps) {
		this.props = initialProps;
	}

	public static animateBackdrop(node: HTMLElement) {
		gsap.fromTo(node, { opacity: 0 }, { opacity: 1, duration: 0.25, ease: 'power2.out' });

		return () => {
			gsap.to(node, { opacity: 0, duration: 0.2, ease: 'power2.in' });
		};
	}

	public static animateModalWindow(node: HTMLElement) {
		gsap.fromTo(
			node,
			{ opacity: 0, scale: 0.95, y: 20 },
			{ opacity: 1, scale: 1, y: 0, duration: 0.3, ease: 'back.out(1.2)' }
		);

		return () => {
			gsap.to(node, { opacity: 0, scale: 0.95, y: 15, duration: 0.2, ease: 'power2.in' });
		};
	}

	public async handleCloseAnimation() {
		if (!this.backdropContainer || !this.modalContainer) {
			this.props.opened = false;
			return;
		}

		await Promise.all([
			gsap.to(this.backdropContainer, { opacity: 0, duration: 0.2, ease: 'power2.in' }),
			gsap.to(this.modalContainer, {
				opacity: 0,
				scale: 0.95,
				y: 15,
				duration: 0.2,
				ease: 'power2.in'
			})
		]);

		this.props.opened = false;
	}

	private ShakeModalContainer() {
		if (!this.modalContainer) return;

		gsap.killTweensOf(this.modalContainer);

		gsap.to(this.modalContainer, {
			x: -10,
			duration: 0.1,
			repeat: 3,
			yoyo: true,
			ease: 'power1.inOut',
			onComplete: () => {
				if (!this.modalContainer) return;
				gsap.set(this.modalContainer, { x: 0 });
			}
		});
	}

	public async on_background_click() {
		if (this.props.backgroundExitLocked) {
			this.ShakeModalContainer();
			return;
		}

		await this.handleCloseAnimation();
	}

	public async on_exit_icon_click() {
		if (this.props.onClose) {
			this.props.onClose();
		}
		await this.handleCloseAnimation();
	}

	public async on_content_click(e: MouseEvent) {
		e.stopPropagation();
	}
}
