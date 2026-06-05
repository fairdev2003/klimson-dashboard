import { writable } from 'svelte/store';
import type { SettingsStartupSlug } from '../types/main.types';

export const settings_startup_modal = writable<SettingsStartupSlug>('none');
export const dashboard_load_date = writable<string>('');
