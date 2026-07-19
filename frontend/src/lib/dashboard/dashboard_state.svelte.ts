import BaseDockComponent from '$lib/components/dashboard/dock/boxes/BaseDockComponent.svelte';
import type { Component } from 'svelte';

class State {
	public loaded_percent: number = $state(0);
	public dashboard_loader_on: boolean = $state(false);
	public current_directory: string = $state('/');
	public dockComponent: Component | undefined = $state(BaseDockComponent);

	public async setLoaderPercent(percent: number, hideLoader?: { delay?: number }) {
		this.loaded_percent = 10;
		this.dashboard_loader_on = true;
		this.loaded_percent = percent;

		if (hideLoader) {
			setTimeout(() => {
				this.dashboard_loader_on = false;
			}, hideLoader.delay ?? 1000);
		}
	}

	public setDockComponent(component: Component) {
		this.dockComponent = component;
	}
}

const DashboardState = new State();
export { DashboardState };
