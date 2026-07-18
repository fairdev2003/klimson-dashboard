import { goto, preloadCode } from '$app/navigation';
import { api } from '$lib/api/api';
import type { ServerResponse } from '$lib/api/types';
import { terminal } from '$lib/components/dashboard/dev/console/terminal.svelte';
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
import { DashboardState } from '$lib/dashboard/logic';
import { DashboardMisc } from './dashboard_misc.svelte';
import axios from 'axios';

export const dockComponent = writable<any>(BaseDockComponent);

class DashboardClass {
	constructor() {}

	public get state() {
		return DashboardState;
	}

	public get miscellaneous() {
		return DashboardMisc;
	}

	public async Load(): Promise<boolean> {
		debug.image('https://api.klimson.dev/interface/bucket/random/zbysiu.png');

		dashboardLoadState.set('Autoryzacja');
		dashboardLoaded.set(false);
		try {
			const verify: ServerResponse<{ access: boolean }> = await api.api.get('/admin/verify', {
				headers: { Authorization: `Bearer ${localStorage.getItem('token')}` }
			});

			if (!verify.data.access) return false;
		} catch (error) {
			if (axios.isAxiosError(error)) {
				if (error.code === '401') {
					console.log(error.message);
					goto('/login');
				}
			}
		}

		dashboardLoadState.set('Pobieranie danych panelu...');
		const context_response = await api.context_storage.GetSinglePrivateContext('clan_id');
		const [routesResponse] = await Promise.all([api.misc.GetRoutes()]);

		routes.set(routesResponse.data);
		console.log(routesResponse.data);

		updateResponseTime('routesResponseTime', routesResponse.duration);

		debug.system('Server is up and ready!');

		toast.success('Aktualne dane załadowane');
		dashboard_load_date.set(this.miscellaneous.GetDateTime());

		dashboardLoadState.set('Lazy Loading dashboard components!');
		await preloadCode('/dashboard/database');
		await preloadCode('/dashboard/context_storage');

		dashboardLoaded.set(true);
		debug.system("Terminal initialized. Type 'cmds' to view available commands");
		terminal.set_terminal_naming({ name: 'klimson-dashboard', path: 'main' });
		return true;
	}
}

const Dashboard = new DashboardClass();
export default Dashboard;
