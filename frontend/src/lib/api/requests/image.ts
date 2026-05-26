/**
 * ImageApi - klasa do obsługi uploadu oraz pobierania obrazów
 * z backendu (upload, generowanie URL, listowanie plików).
 */

import type { AxiosInstance, AxiosResponse } from 'axios';
import type { ApiConfig, ImageKey, ImageList } from '../types';
import type { UploadedImageResponse } from '$lib/types/image';
import { Api } from '../api';

class ImageApi {
	private api: AxiosInstance;
	private api_config: ApiConfig;

	/**
	 * Tworzy nową instancję ImageApi.
	 *
	 * @param {AxiosInstance} api - Instancja axios używana do wykonywania requestów.
	 * @param {ApiConfig} api_config - Konfiguracja API (np. host, dev_server, baseURL).
	 */
	constructor(api: AxiosInstance, api_config: ApiConfig) {
		this.api = api;
		this.api_config = api_config;
	}

	/**
	 * Zwraca pełną ścieżkę bazową do katalogu `/uploads/`.
	 *
	 * @returns {string} - URL bazowy do uploadów.
	 *
	 * @example
	 * console.log(api.image.upload_url);
	 */
	public get upload_url() {
		return this.api_config.baseURL + 'uploads/';
	}

	/**
	 * Wysyła obraz do backendu.
	 *
	 * @async
	 * @param {ImageKey} key - Nazwa kategorii obrazów (`quiz`, `question`, `blog`, `hero`).
	 * @param {FormData} data - Dane z plikiem (FormData musi zawierać klucz `image`).
	 * @returns {Promise<UploadedImageResponse>} - Obiekt zwracany przez backend: `{ id, file_name, message }`.
	 *
	 * @example
	 * ```ts
	 * const data = new FormData();
	 * data.append('image', file);
	 * const response = await api.image.SendImage('quiz', data);
	 * console.log(res.file_name);
	 * ```
	 */
	public async SendImage(key: ImageKey, data: FormData): Promise<UploadedImageResponse> {
		const response: AxiosResponse<UploadedImageResponse> = await this.api.post(
			`/admin/upload?key=${key}`,
			data,
			{ headers: { Authorization: `Bearer ${Api.token}` } }
		);

		return response.data;
	}

	/**
	 * Usuwa zdjęcie z serwera.
	 *
	 * @async
	 * @param {ImageKey} key - Nazwa kategorii obrazów (`quiz`, `question`, `blog`, `hero`).
	 * @param {string} file - Nazwa pliku do usunięcia (np `uuid.jpeg`).
	 * @returns {Promise<UploadedImageResponse>} - Obiekt zwracany przez backend: `{ file_name, message }`.
	 *
	 * @example
	 * ```ts
	 * const file_name = '2a8a1b3c-d692-47d2-a10a-a33a057515ac.jpeg';
	 * const response = await api.image.DeleteImage('quiz', file_name);
	 * console.log(res.file_name);
	 * ```
	 */
	public async DeleteImage(key: ImageKey, file: string) {
		const res = await this.api.delete(`/admin/images/delete?key=${key}&file=${file}`, {
			headers: { Authorization: `Bearer ${Api.token}` }
		});

		console.log(res.data);
	}

	/**
	 * Buduje pełny URL do obrazu.
	 *
	 * @param {ImageKey} key - Kategoria obrazu.
	 * @param {string} image_name - Nazwa pliku (np. `uuid.jpg`).
	 * @returns {string} - Pełny URL do obrazka.
	 *
	 * @example
	 * const url = api.image.getImage('quiz', '1234.jpg');
	 * <img src={url} />
	 */
	public getImage(key: ImageKey, image_name: string | undefined) {
		if (image_name && image_name.startsWith('http')) {
			return image_name;
		}
		return `${this.api_config.dev_server}/uploads/${key}/${image_name}`;
	}

	/**
	 * Skrót dla `quiz`.
	 *
	 * @param {string} image_name - Nazwa pliku.
	 * @returns {string}
	 */
	public quiz(image_name: string) {
		return this.getImage('quiz', image_name);
	}

	/**
	 * Skrót dla `question`.
	 */
	public question(image_name: string) {
		return this.getImage('question', image_name);
	}

	/**
	 * Skrót dla `blog`.
	 */
	public blog(image_name: string) {
		return this.getImage('blog', image_name);
	}

	/**
	 * Skrót dla `hero`.
	 */
	public hero(image_name: string) {
		return this.getImage('hero', image_name);
	}

	public static(image_name: string) {
		return this.getImage('static', image_name);
	}

	/**
	 * Pobiera listę obrazów z backendu i zamienia nazwy plików na pełne URL.
	 *
	 * @async
	 * @param {ImageKey} key - Kategoria obrazów do pobrania.
	 * @returns {Promise<string[]>} - Lista pełnych URL obrazów.
	 *
	 * @example
	 */
	public async ListImages(key: ImageKey): Promise<ImageList> {
		const response: AxiosResponse<ImageList> = await this.api.get(`/admin/images/list?key=${key}`, {
			headers: { Authorization: `Bearer ${Api.token}` }
		});

		return response.data;
	}
}

export { ImageApi };
