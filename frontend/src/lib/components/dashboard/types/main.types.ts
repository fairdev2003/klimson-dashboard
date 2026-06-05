import type { Component } from 'svelte';
import type { SettingKey } from '../settings/store.svelte';

export type WidgetSlug = 'database' | 'cpu' | 'spotify' | 'disk' | 'clock';
export type SettingsStartupSlug = 'widgets' | 'contents' | 'none';

export type Widget = {
	component: Component;
	slug: WidgetSlug;
	name: string;
	description: string;
};
