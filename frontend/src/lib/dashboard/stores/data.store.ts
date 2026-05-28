import { writable } from 'svelte/store';
import type { RoutesResponse } from '../../../routes/dashboard/routes/types';

export const dashboardLoadState = writable<string>('Ładowanie komponentów');
export const dashboardLoaded = writable<boolean>(false);

export const routes = writable<RoutesResponse[] | undefined>();
export const searchOpen = writable<boolean>();

export type TimeResponse =
	| 'quizzesResponseTime'
	| 'questionsResponseTime'
	| 'answersResponseTime'
	| 'routesResponseTime'
	| 'blogResponseTime'
	| 'heroResponseTime'
	| 'contributorsResponseTime'
	| 'permissionListResponseTime';

export const requestTimes = writable<Record<TimeResponse, number>>({
	quizzesResponseTime: 0,
	questionsResponseTime: 0,
	answersResponseTime: 0,
	routesResponseTime: 0,
	blogResponseTime: 0,
	heroResponseTime: 0,
	contributorsResponseTime: 0,
	permissionListResponseTime: 0
});

export const updateResponseTime = (key: TimeResponse, value: number) => {
	requestTimes.update((s) => ({ ...s, [key]: value }));
};
