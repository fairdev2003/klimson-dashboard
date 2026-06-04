import { writable } from 'svelte/store';

export type SettingKey = 'main' | 'server';

export const settings_page_open = writable<SettingKey>('main');
