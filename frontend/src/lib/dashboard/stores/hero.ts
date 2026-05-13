import { writable } from 'svelte/store';
import type { HeroType } from '../../../routes/dashboard/hero/types';

export const heroForm = writable<HeroType>({
	author: '',
	image_url: '',
	quote: ''
});
