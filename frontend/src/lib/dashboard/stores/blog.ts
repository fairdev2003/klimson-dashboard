import { writable } from 'svelte/store';
import type { BlogType } from '../../../routes/dashboard/blog/types';

export const blogForm = writable<BlogType>({
	title: '',
	description: '',
	html: '',
	public: false
});

export const blogFormState = writable<'add' | 'update'>();
