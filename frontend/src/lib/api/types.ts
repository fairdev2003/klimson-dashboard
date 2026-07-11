import type { AxiosResponse } from 'axios';

/**
 * Typ konfiguracji inicjalizacyjnej klasy `Api`.
 *
 * Pozwala nadpisać domyślne wartości środowiskowe i ścieżki API.
 * Każda z właściwości jest opcjonalna — podanie dowolnej spowoduje
 * zastąpienie domyślnej konfiguracji w konstruktorze klasy `Api`.
 *
 * @property {string} [dev_server]   - Adres serwera deweloperskiego (np. `http://localhost:8080`)
 * @property {string} [prod_server]  - Adres serwera produkcyjnego (np. `https://api.mojaapka.com`)
 * @property {string} [prod_version] - Wersja API (np. `/v1`)
 * @property {string} [prod_prefix]  - Prefiks ścieżki API (np. `/api`)
 */
export type ApiClassParams = {
	dev_server?: string;
	prod_server?: string;
	prod_version?: string;
	prod_prefix?: string;
	prod_switch?: boolean;
};

/**
 * Rozszerzony typ odpowiedzi z serwera.
 *
 * Obejmuje standardową odpowiedź `AxiosResponse<T>`, a dodatkowo
 * zawiera czas trwania żądania (`duration`), mierzony w milisekundach.
 * Wartość ta jest ustawiana automatycznie przez interceptory Axios w klasie `Api`.
 *
 * @template T Typ danych zwracanych w odpowiedzi (np. `User`, `Quiz`, `Question[]`)
 *
 * @example
 * ```ts
 * const response: ServerResponse<Quiz> = await api.quiz.getQuiz({author: "Fredi Kamionka"});
 * console.log(response.data.name); // -> dane u Quizie
 * console.log(response.duration); // -> czas odpowiedzi w ms
 * ```
 */
export type ServerResponse<T = {}> = AxiosResponse<T & { message: string; error: string }> & {
	duration: number;
};

export type BackendResponse<T = {}> = {
	message: string;
	succes: boolean;
} & T;

export type ApiConfig = {
	dev_server: string;
	prod_server: string;
	prod_front: string;
	version: string;
	prefix: string;
	host: string;
	baseURL: string;
};

export type ImageKey = 'quiz' | 'question' | 'blog' | 'hero' | 'static';

export type ImageList = {
	name: string;
	is_dir: boolean;
}[];

export type QuizWhereObject = {
	id: number | undefined;
};

export type CheckCorrectObject = {
	id: number;
	is_correct: boolean;
};

export type UpdateOneFieldType<T> = {
	[K in keyof T]: {
		key: K;
		value: T[K];
	};
}[keyof T];
