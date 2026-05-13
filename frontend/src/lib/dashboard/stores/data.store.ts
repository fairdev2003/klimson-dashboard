import { writable } from 'svelte/store';
import { type Answer, type Question, type Quiz } from '../../../routes/dashboard/quizzes/types';
import type { RoutesResponse } from '../../../routes/dashboard/routes/types';
import type { BlogType } from '../../../routes/dashboard/blog/types';
import type { HeroType } from '../../../routes/dashboard/hero/types';
import { type Permission, type Contributor } from '../../../routes/dashboard/contributors/types';

export const dashboardLoadState = writable<string>('Ładowanie komponentów');
export const dashboardLoaded = writable<boolean>(false)

export const quizzes = writable<Quiz[]>();
export const questions = writable<Question[] | undefined>();
export const answers = writable<Answer[] | undefined>();
export const routes = writable<RoutesResponse[] | undefined>();
export const blogs = writable<BlogType[] | undefined>();
export const heros = writable<HeroType[] | undefined>();
export const contributors = writable<Contributor[]>();
export const permissionList = writable<Permission[]>()
export const searchOpen = writable<boolean>();

export type TimeResponse = "quizzesResponseTime" |
    "questionsResponseTime" |
    "answersResponseTime" |
    "routesResponseTime" |
    "blogResponseTime" |
    "heroResponseTime" |
    "contributorsResponseTime" |
    "permissionListResponseTime"

export const requestTimes = writable<Record<
    TimeResponse, number>>(
        {
            quizzesResponseTime: 0,
            questionsResponseTime: 0,
            answersResponseTime: 0,
            routesResponseTime: 0,
            blogResponseTime: 0,
            heroResponseTime: 0,
            contributorsResponseTime: 0,
            permissionListResponseTime: 0
        }
    )

export const updateResponseTime = (key: TimeResponse, value: number) => {
    requestTimes.update(s => ({ ...s, [key]: value }));
};
