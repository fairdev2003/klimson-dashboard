import { persistedWritable } from '$lib/dashboard/stores/persist';
import { writable } from 'svelte/store';

export type SettingKey = 'main' | 'server' | 'customization';

export const settings_page_open = persistedWritable<SettingKey>('settings_page_open', 'main');
export const nickname = persistedWritable<string>('nickname', 'cwel');
