import { writable } from 'svelte/store';
import type { ClanResponse } from '../../../../../routes/dashboard/pg3d/pg3d.types';

export const clan_info = writable<ClanResponse | undefined>();
