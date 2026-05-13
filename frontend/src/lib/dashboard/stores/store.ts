import { get, writable } from 'svelte/store';
import { type Question, type Quiz } from '../../../routes/dashboard/quizzes/types';
import { emptyQuestion, preloaded_quiz_form } from '$lib/static/quiz';
import { browser } from '$app/environment';
import type { Contributor } from '../../../routes/dashboard/contributors/types';
import { roles } from '../../../routes/dashboard/contributors/vars';
import { toast } from './toast';
import type { SidebarItemType } from '$lib/types/sidebar';
import { quizzes } from './data.store';

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
export const sidebar_open = writable<boolean>(false);
export const summary_open = writable<boolean>(false);
export const selectedSummaryType = writable<
	undefined | 'quiz' | 'question' | 'answer' | 'blog' | 'debug'
>();

export const quizView = writable<'table' | 'cards'>('cards');
export const selectedQuiz = writable<Quiz | undefined>();
export const initialFormQuiz = writable<Quiz>();
export const addFormQuiz = writable<Quiz>(preloaded_quiz_form);

export function updateFormQuiz(key: keyof Quiz, value: any) {
	addFormQuiz.update((quiz) => ({ ...quiz, [key]: value }));
}

// question add or edit
export const editMode = writable<boolean>(false);
export const selectedQuestionToEdit = writable<Question>(emptyQuestion);
export const questions = writable<Question[]>([]);
export const addingQuestionForm = writable<Question>();
export const quizStatus = writable<'saving' | 'synced' | 'not-saved'>('synced');

export const addFormSection = writable<'general' | 'questions' | 'creation_summary'>('general');

export const previewQuiz = writable<Quiz | undefined>();

export const addFormContributor = writable<Contributor>({
	name: '',
	description: '',
	blocked_till: false,
	last_login: '',
	login: '',
	logs: [],
	password: '',
	pending_quizzes: [],
	profile_image: '',
	permissions: '',
	thumbnail: ''
});

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

export function resetContributor() {
	addFormContributor.set({
		name: '',
		description: '',
		blocked_till: false,
		last_login: '',
		login: '',
		logs: [],
		password: '',
		pending_quizzes: [],
		profile_image: '',
		permissions: '',
		thumbnail: ''
	});
}
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

export const sidebar_menu_static: SidebarItemType[] = [
	{
		id: 'start',
		label: 'Start',
		desc: 'Informacje o stanie serwera, szybki dostep do poszczegolnych sekcji i podsumowanie',
		icon: 'ph:list-numbers-bold', // Przykład ikony z Phosphor
		link: '/dashboard',
		children: []
	},
	{
		id: 'quizzes',
		label: `Quizy`,
		desc: 'Sekcja z quizami',
		icon: 'tabler:campfire-filled', // Material Design
		link: '/dashboard/quizzes',
		children: [
			{
				id: 'q-list',
				label: `Lista (${get(quizzes) && get(quizzes).length})`,
				link: '/dashboard/quizzes',
				desc: 'Przegladaj liste quizów'
			},
			{ id: 'q-add', label: 'Pytania', link: '/dashboard/questions', desc: 'Przeglądaj pytania' },
			{
				id: 'q-answers',
				label: 'Odpowiedzi',
				link: '/dashboard/answers',
				desc: 'Przeglądaj Odpowiedzi'
			},
			{
				id: 'q-preview',
				label: 'Zobacz na żywo',
				link: '/dashboard/answers',
				desc: 'Zobacz jak uzytkownik widzi quizy'
			}
		]
	},
	{
		id: 'contributors',
		desc: 'Sekcja z kontrybutorami. Dodawaj, usuwaj i zarządzaj kontrybutorami',
		label: 'Kontrybutorzy',
		icon: 'ph:users-three-bold',
		link: '/dashboard/contributors',
		children: []
	},
	{
		id: 'blog',
		label: 'Blog',
		desc: 'Zarządzaj blogową sekcją strony harcquiz',
		icon: 'mdi:newspaper-variant-outline',
		link: '/dashboard/blog',
		children: []
	},
	{
		id: 'hero',
		label: 'Hero',
		desc: 'Dodawaj zdjecie oraz cytat na strone glowna',
		icon: 'mdi:image-filter-hdr',
		link: '/dashboard/hero',
		children: []
	},
	{
		id: 'api',
		label: 'API',
		desc: 'Rozpiska wszystkich endpointów api. Sekcja wyłącznie dla developera',
		icon: 'ph:code-bold',
		link: '/dashboard/routes',
		children: []
	},
	{
		id: 'images',
		label: 'Zdjęcia',
		desc: 'Dodawaj statyczne zdęcia. które sa pomocne w dodawaniu zdjec na bloga. Zarządzaj równiez zdjeciami quizów oraz pytan!',
		icon: 'ph:image-bold',
		link: '/dashboard/images',
		children: []
	},
	{
		id: 'settings',
		label: 'Ustawienia',
		desc: 'Globalne ustawienia panelu harcquiz',
		icon: 'ph:gear-six-bold',
		link: '/dashboard/settings',
		children: []
	},
	{
		id: 'info',
		label: 'Informacje',
		desc: 'Informacje od developera',
		icon: 'ph:info-bold',
		link: '/dashboard/info',
		children: []
	}
];

export const sidebar_content = writable<SidebarItemType[]>(sidebar_menu_static);

export function addChildToQuizzes(label: string, link: string, desc: string, onclick: () => void) {
	sidebar_content.update((menu) => {
		return menu.map((item) => {
			if (item.id === 'quizzes') {
				// SPRAWDZENIE: Czy dziecko o tym labelu już jest w menu?
				const alreadyExists = item.children.some((child) => child.label === label);

				// Jeśli istnieje, zwracamy item bez zmian (żeby nie dublować)
				if (alreadyExists) return item;

				// Jeśli nie istnieje, dodajemy nowy element
				return {
					...item,
					children: [
						...item.children,
						{
							// Używamy stałego ID opartego na labelu zamiast Date.now(),
							// co dodatkowo zabezpiecza przed błędami w kluczach pętli {#each}
							id: `dynamic-${label.toLowerCase().replace(/\s+/g, '-')}`,
							label,
							link,
							desc,
							onclick
						}
					]
				};
			}
			return item;
		});
	});
}
