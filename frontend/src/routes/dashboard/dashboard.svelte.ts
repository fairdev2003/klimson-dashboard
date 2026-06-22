import { preloadCode } from '$app/navigation';
import { api } from '$lib/api/api';
import type { ServerResponse } from '$lib/api/types';
import BaseDockComponent from '$lib/components/dashboard/dock/boxes/BaseDockComponent.svelte';
import { dashboard_load_date } from '$lib/components/dashboard/stores/main';
import { clan_info } from '$lib/components/dashboard/table/pg3d/clan.store';
import {
	dashboardLoadState,
	routes,
	updateResponseTime,
	dashboardLoaded
} from '$lib/dashboard/stores/data.store';
import { debug } from '$lib/dashboard/stores/debug';
import { toast } from '$lib/dashboard/stores/toast';
import type { Component } from 'svelte';
import { get, writable } from 'svelte/store';

export const dockComponent = writable<any>(BaseDockComponent);

class DashboardClass {
	constructor() {}

	public art: string = `
 ____ ___ _____ __  __    _       ______        _______ _    _   _ 
/ ___|_ _| ____|  \\/  |  / \\     / ___\\ \\      / / ____| |  | | | |
\\___ \\| ||  _| | |\\/| | / _ \\   | |    \\ \\ /\\ / /|  _| | |  | | | |
 ___) | || |___| |  | |/ ___ \\  | |___  \\ V  V / | |___| |__| |_| |
|____/___|_____|_|  |_/_/   \\_\\  \\____|  \\_/\\_/  |_____|_____\\___/ 
`;

	private GetDateTime(): string {
		const now = new Date();
		const currentTime = now.toLocaleTimeString('us-US', {
			hour: '2-digit',
			minute: '2-digit',
			second: '2-digit'
		});
		const currentDate = now.toLocaleDateString('us-US', {
			day: 'numeric',
			month: 'long',
			year: 'numeric'
		});
		return `${currentTime}, ${currentDate}`;
	}

	public async Load(): Promise<boolean> {
		console.log(this.art);

		dashboardLoadState.set('Autoryzacja');
		dashboardLoaded.set(false);

		const verify: ServerResponse<{ access: boolean }> = await api.api.get('/admin/verify', {
			headers: { Authorization: `Bearer ${localStorage.getItem('token')}` }
		});

		if (!verify.data.access) return false;
		debug.log('Autoryzacja poprawna');

		dashboardLoadState.set('Pobieranie danych panelu...');
		const context_response = await api.context_storage.GetSinglePrivateContext('clan_id');
		const [routesResponse] = await Promise.all([
			api.misc.GetRoutes()
			// api.pg3d.GetClanInfo(context_response.data.value)
		]);

		// clan_info.set(clan_response.data);

		routes.set(routesResponse.data);
		console.log(routesResponse.data);

		updateResponseTime('routesResponseTime', routesResponse.duration);

		debug.system('Zesralem sie');

		toast.success('Aktualne dane załadowane');
		dashboard_load_date.set(this.GetDateTime());

		dashboardLoadState.set('Lazy Loading dashboard components!');
		await preloadCode('/dashboard/database');
		await preloadCode('/dashboard/context_storage');

		dashboardLoaded.set(true);

		return true;
	}
}

const Dashboard = new DashboardClass();
export default Dashboard;
