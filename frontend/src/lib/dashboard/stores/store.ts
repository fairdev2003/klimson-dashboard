import { get, writable } from 'svelte/store';
import { emptyQuestion, preloaded_quiz_form } from '$lib/static/quiz';
import { browser } from '$app/environment';
import { toast } from './toast';
import { persistedWritable } from './persist';

export function persistentStore<T>(key: string, initial: T) {
	if (!browser) {
		return writable(initial);
	}

	let parsed: T = initial;

	try {
		const stored = localStorage.getItem(key);
		if (stored && stored !== 'undefined' && stored !== 'null') {
			parsed = JSON.parse(stored);
		}
	} catch (err) {
		console.warn(`Błąd przy parsowaniu localStorage[${key}]`, err);
	}

	const store = writable<T>(parsed);

	store.subscribe((value) => {
		try {
			if (value === undefined) {
				localStorage.removeItem(key);
			} else {
				localStorage.setItem(key, JSON.stringify(value));
			}
		} catch (err) {
			console.warn(`Błąd przy zapisie localStorage[${key}]`, err);
		}
	});

	return store;
}

export const page = writable<number>(0);
export const sidebar_open = persistedWritable<boolean>('sidebar_open', true);
export const summary_open = writable<boolean>(false);
export const selectedSummaryType = writable<
	undefined | 'quiz' | 'question' | 'answer' | 'blog' | 'debug'
>();

export const quizView = writable<'table' | 'cards'>('cards');

// question add or edit
export const editMode = writable<boolean>(false);
export const quizStatus = writable<'saving' | 'synced' | 'not-saved'>('synced');

export const addFormSection = writable<'general' | 'questions' | 'creation_summary'>('general');

export const contextMenuOptions = writable<ContextMenu[]>([
	{
		contextMenuName: '',
		items: [
			{
				label: 'Edytuj',
				action: () => {
					toast.show('Edytowanie... :)');
				},
				icon: '',
				color: ''
			},
			{
				label: 'Skopuj adres strony',
				action: () => {
					navigator.clipboard.writeText('siema :)');
					toast.show('Skopiowano do schowka!');
				},
				icon: '',
				color: ''
			},
			{ label: 'Usuń', action: () => {}, icon: '', color: '' }
		]
	}
]);

export type JWT = {
	exp: number;
	contributor: boolean;
	roles: string;
	name: string;
	login: string;
};
export const userInfo = writable<JWT>({
	exp: 0,
	contributor: false,
	roles: '',
	name: 'Paweł Mścichuj',
	login: 'root'
});
