import { persistedWritable } from '$lib/dashboard/stores/persist';
import type { Component } from 'svelte';

export type SettingKey = 'main' | 'server' | 'customization' | 'account';

export type DropdownOption = {
	key: string;
	value: string;
};

export type Setting = {
	component: Component;
	slug: SettingKey;
	name: string;
	description: string;
};

export const settings_page_open = persistedWritable<SettingKey>('settings_page_open', 'main');
export const nickname = persistedWritable<string>('nickname', 'Jakub');
