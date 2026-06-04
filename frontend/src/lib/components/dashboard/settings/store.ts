import { writable } from 'svelte/store';

export type SettingKey = 'main' | 'server' | 'customization';

export const settings_page_open = writable<SettingKey>('main');
