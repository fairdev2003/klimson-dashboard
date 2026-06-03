import { dev } from '$app/environment';
import axios, { type AxiosInstance } from 'axios';
import type { ApiClassParams, ApiConfig } from './types';
import { Misc } from './requests/misc';
import { DatabaseClass } from './requests/database';
import { ContextStorage } from './requests/context_storage';
import { PG3D } from './requests/pg3d';
import { Storage } from './requests/storage';

/**
 * Klasa bazowa definiująca konfigurację API.
 *
 * Zawiera podstawowe informacje o serwerach, wersjach i ścieżkach API,
 * a także automatycznie generuje bazowy adres `_baseURL` w zależności od środowiska (`dev` vs `prod`).
 */
class ApiStatic {
	/** Adres serwera deweloperskiego */
	protected _dev_server = 'http://localhost:8090';

	/** Adres serwera produkcyjnego (backend) */
	protected _prod_server = 'https://api.klimson.dev';

	/** Adres frontendu produkcyjnego */
	protected _prod_front = 'https://dashboard.klimson.dev';

	/** Wersja API (np. `/v1`) */
	protected _version = '';

	/** Prefiks ścieżek API (np. `/api`) */
	protected _prefix = '';

	/**
	 * Zwraca aktualny host w zależności od środowiska.
	 * @returns {string} Adres hosta (dev lub prod)
	 */
	protected get _host(): string {
		return dev ? this._dev_server : this._prod_server;
	}

	/**
	 * Zwraca kompletny bazowy URL dla żądań API.
	 * @returns {string} Bazowy adres API
	 */
	protected get _baseURL(): string {
		return `${this._host}${this._version}${this._prefix}/`;
	}

	/**
	 * Zwraca pełną konfigurację API w postaci obiektu.
	 * Przydatne np. do debugowania lub dynamicznych komponentów.
	 */
	public get api_config(): ApiConfig {
		return {
			dev_server: this._dev_server,
			prod_server: this._prod_server,
			prod_front: this._prod_front,
			version: this._version,
			prefix: this._prefix,
			host: this._host,
			baseURL: this._baseURL
		};
	}

	public static get token(): string {
		const token_string = localStorage.getItem('token');
		let token: string = '';
		if (token_string && token_string.length > 10) {
			token = token_string;
		}
		return token;
	}
}

/**
 * Główna klasa API, rozszerzająca konfigurację bazową (`ApiStatic`)
 * i inicjalizująca instancję klienta `axios`.
 */
export class Api extends ApiStatic {
	/** Instancja Axios skonfigurowana dla aktualnego środowiska */
	public api: AxiosInstance;

	/**
	 * Tworzy nowy obiekt API z opcjonalną konfiguracją.
	 *
	 * @param {ApiClassParams} [params] - Obiekt pozwalający nadpisać domyślne ustawienia (np. inne serwery lub wersje API).
	 *
	 * @example
	 * ```ts
	 * const api = new Api({
	 *   dev_server: 'http://localhost:8090',
	 *   prod_server: 'https://api.klimson.dev',
	 * });
	 *
	 * const data = await api.quiz.GetQuiz({ id: 1 }); // sql where
	 * console.log(data);
	 * ```
	 */
	constructor(params?: ApiClassParams) {
		super();

		if (params?.dev_server) this._dev_server = params.dev_server;
		if (params?.prod_server) this._prod_server = params.prod_server;
		if (params?.prod_version) this._version = params.prod_version;
		if (params?.prod_prefix) this._prefix = params.prod_prefix;

		console.log(`Operating API at: ${this._baseURL}`);

		this.api = axios.create({
			baseURL: this._baseURL,
			headers: {
				Accept: 'application/json'
			}
		});

		this.api.interceptors.request.use((config) => {
			(config as any).metadata = { startTime: new Date() };
			return config;
		});
	}

	/**
	 * Zwraca interfejs do endpointów `quiz`.
	 *
	 * @example
	 * ```ts
	 * const quiz = await api.quiz.getQuiz({id: 1});
	 * ```
	 */

	/**
	 * Zwraca interfejs do endpointów `image`.
	 *
	 * @example
	 * ```ts
	 * const formData = new FormData();
	 * formData.append('image', file);
	 * formData.append('id', id);
	 * const response = await api.image.SendImage("quiz", formData);
	 * ```
	 */

	public get misc() {
		return new Misc(this.api);
	}

	public get database() {
		return new DatabaseClass(this.api);
	}

	public get context_storage() {
		return new ContextStorage(this.api);
	}

	public get pg3d() {
		return new PG3D(this.api);
	}

	public get storage() {
		return new Storage(this.api);
	}
}

/**
 * Domyślna instancja API, gotowa do użycia w całej aplikacji.
 *
 * Używa domyślnych adresów zdefiniowanych w klasie `ApiStatic`:
 * - `dev` → `http://localhost:8080`
 * - `prod` → `https://harcquiz-ztai.onrender.com`
 * @example
 * ```ts
 * // Utworzenie niestandardowej konfiguracji (np. środowisko testowe)
 * import { api } from '$lib/api';
 *
 * const testApi = new Api({
 *   dev_server: 'http://127.0.0.1:5000',
 *   prod_server: 'https://staging.harcquiz.dev',
 *   prod_version: '/v2',
 *   prod_prefix: '/api-test'
 * });
 * ```
 */
// const api = new Api({ dev_server: 'https://api.klimson.dev' });
const api = new Api({ dev_server: 'http://localhost:8090' });

export { api };
