import { goto, preloadCode } from '$app/navigation';
import { api } from '$lib/api/api';
import type { BackendResponse, ServerResponse } from '$lib/api/types';
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
import { writable } from 'svelte/store';
import { DashboardState } from '$lib/dashboard/logic';
import { DashboardMisc } from './dashboard_misc.svelte';
import Constants from './constants';
import type { Component } from 'svelte';
import { DashboardHttpLogger } from './http.svelte';
import { userLogin } from './stores/persist';

class DashboardClass {
	private _constants: Constants;
	private _http: DashboardHttpLogger;

	constructor() {
		this._constants = new Constants();
		this._http = new DashboardHttpLogger();
	}

	public get state() {
		return DashboardState;
	}

	public get miscellaneous() {
		return DashboardMisc;
	}

	public get constants() {
		return this._constants;
	}

	public get http() {
		return this._http;
	}

	public async Load(): Promise<boolean> {
		debug.image('https://api.klimson.dev/interface/bucket/random/nugget_cat.png');

		dashboardLoadState.set('Autoryzacja');
		dashboardLoaded.set(false);
		try {
			const verify: ServerResponse<{ access: boolean }> = await api.api.get('/admin/verify');

			if (!verify.data.access) return false;
		} catch (error) {
			console.log(error);
			goto('/login');
		}
		try {
			const current_user: ServerResponse<
				BackendResponse<{
					token: string;
					claims: {
						name: string;
						login: string;
						exp: string;
					};
				}>
			> = await api.api.get('/admin/users/me');

			userLogin.set(current_user.data.claims.login);
		} catch (error) {
			console.log(error);
			goto('/login');
		}

		dashboardLoadState.set('Pobieranie danych panelu...');
		const [routesResponse] = await Promise.all([api.misc.GetRoutes()]);

		routes.set(routesResponse.data);
		console.log(routesResponse.data);

		updateResponseTime('routesResponseTime', routesResponse.duration);

		debug.system('Server is up and ready!');

		dashboard_load_date.set(this.miscellaneous.GetDateTime());

		dashboardLoaded.set(true);
		debug.system("Terminal initialized. Type 'cmds' to view available commands");
		terminal.set_terminal_naming(this.constants.TerminalNaming);
		return true;
	}
}

const Dashboard = new DashboardClass();
export default Dashboard;
