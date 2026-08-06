import { writable } from 'svelte/store';
import { browser } from '$app/environment';
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

export type Bookmark = {
	name: string;
	slug: string;
	href: string;
	color: string;
};

export type DashboardSettings = {
	code_theme: CodeEditorTheme;
	client_pills: SideBarPillPreferences[];
	dock_on: boolean;
	bookmarks: Bookmark[];
};
export type AnimationPresetType = 'blur' | 'klimson' | 'jason';

export const debugOn = persistedWritable<boolean>('debug_on', true);
export const dashboard_config = persistedWritable<DashboardSettings>('dashboard_settings', {
	code_theme: 'classic',
	client_pills: ['profile'],
	dock_on: false,
	bookmarks: []
});
export const developerView = persistedWritable<boolean>('dev_view', false);

export const sidebar_open = persistedWritable<boolean>('sidebar_open', true);
export const mobile_sidebar_open = writable<boolean>(false);
export const isMobile = writable(false);
export const animation_preset = persistedWritable<AnimationPresetType>('animationPreset', 'blur');

export const accounts = persistedWritable<JWT[]>('accounts', []);
export const route = persistedWritable<string>('route', '/dashboard');
export const redirectTo = persistedWritable<string>('redirectTo', '/dashboard');
export const userLogin = persistedWritable<string>('userLogin', '');
