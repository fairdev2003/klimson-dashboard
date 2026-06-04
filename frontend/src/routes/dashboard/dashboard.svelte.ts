import { api } from '$lib/api/api';
import type { ServerResponse } from '$lib/api/types';
import {
	dashboardLoadState,
	routes,
	updateResponseTime,
	dashboardLoaded
} from '$lib/dashboard/stores/data.store';
import { debug } from '$lib/dashboard/stores/debug';
import { toast } from '$lib/dashboard/stores/toast';
import { get } from 'svelte/store';

class DashboardClass {
	constructor() {}

	public art: string = `
 ____ ___ _____ __  __    _       ______        _______ _    _   _ 
/ ___|_ _| ____|  \\/  |  / \\     / ___\\ \\      / / ____| |  | | | |
\\___ \\| ||  _| | |\\/| | / _ \\   | |    \\ \\ /\\ / /|  _| | |  | | | |
 ___) | || |___| |  | |/ ___ \\  | |___  \\ V  V / | |___| |__| |_| |
|____/___|_____|_|  |_/_/   \\_\\  \\____|  \\_/\\_/  |_____|_____\\___/ 
`;

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

		const [routesResponse] = await Promise.all([api.misc.GetRoutes()]);

		// quizzes.set(quizzesResponse.data);
		// questions.set(questionsResponse.data as any);
		routes.set(routesResponse.data);
		// blogs.set(blogsResponse.data);
		// heros.set(herosResponse.data);
		// contributors.set(contributorsResponse.data);
		// permissionList.set(permissionListResponse.data);

		// updateResponseTime("quizzesResponseTime", quizzesResponse.duration);
		// updateResponseTime("questionsResponseTime", questionsResponse.duration);
		updateResponseTime('routesResponseTime', routesResponse.duration);
		// updateResponseTime("blogResponseTime", blogsResponse.duration);
		// updateResponseTime("heroResponseTime", herosResponse.duration);
		// updateResponseTime("contributorsResponseTime", contributorsResponse.duration);
		// updateResponseTime("permissionListResponseTime", permissionListResponse.duration);

		debug.system('Zesralem sie');

		toast.success('Aktualne dane załadowane');
		dashboardLoaded.set(true);

		return true;
	}
}

const Dashboard = new DashboardClass();
export default Dashboard;
