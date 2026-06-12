import { writable } from 'svelte/store';
import { browser } from '$app/environment';
import type { Quiz } from '../../../routes/dashboard/quizzes/types';
import type { Contributor } from '../../../routes/dashboard/contributors/types';
import type { JWT } from './store';
import type { CodeEditorTheme } from '$lib/components/markdown/markdown';

export function persistedWritable<T>(key: string, initialValue: T) {
	let storedValue: T | null = null;

	if (browser) {
		const json = localStorage.getItem(key);
		2;
		if (json) {
			try {
				storedValue = JSON.parse(json);
			} catch (e) {
				console.error(`Nie udało się odczytać localStorage dla klucza ${key}`, e);
			}
		}
	}

	const store = writable<T>(storedValue ?? initialValue);

	if (browser) {
		store.subscribe((value) => {
			localStorage.setItem(key, JSON.stringify(value));
		});
	}

	return store;
}

export type SideBarPillPreferences = 'profile' | 'storage';

export type DashboardSettings = {
	allowToSaveQuizWithoutQuestions: boolean;
	code_theme: CodeEditorTheme;
	sidebar_preference: SideBarPillPreferences;
};
export type AnimationPresetType = 'blur' | 'klimson' | 'jason';

export const debugOn = persistedWritable<boolean>('debug_on', true);
export const lastSearched = persistedWritable<Quiz[]>('last_searched', []);
export const dashboard_config = persistedWritable<DashboardSettings>('dashboard_settings', {
	allowToSaveQuizWithoutQuestions: false,
	code_theme: 'classic',
	sidebar_preference: 'profile'
});
export const developerView = persistedWritable<boolean>('dev_view', false);

export const sidebar_open = persistedWritable<boolean>('sidebar_open', true);
export const mobile_sidebar_open = writable<boolean>(false);
export const isMobile = writable(false);
export const animation_preset = persistedWritable<AnimationPresetType>('animationPreset', 'blur');

export const accounts = persistedWritable<JWT[]>('accounts', []);
export const route = persistedWritable<string>('route', '/dashboard');
