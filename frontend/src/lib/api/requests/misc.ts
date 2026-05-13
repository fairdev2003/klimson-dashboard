/**
 * ImageApi - klasa do obsługi uploadu oraz pobierania obrazów
 * z backendu (upload, generowanie URL, listowanie plików).
 */

import type { AxiosInstance, AxiosResponse } from 'axios';

import type { RoutesResponse } from '../../../routes/dashboard/routes/types';
import { Api } from '../api';
import type { ServerResponse } from '../types';

class Misc {
	/**
	 * Tworzy nową instancję ImageApi.
	 *
	 * @param {AxiosInstance} api - Instancja axios używana do wykonywania requestów.
	 * @param {ApiConfig} api_config - Konfiguracja API (np. host, dev_server, baseURL).
	 */
	constructor(private api: AxiosInstance) {
		this.api = api;
	}

	/**
	 * Zwraca pełną ścieżkę bazową do katalogu `/uploads/`.
	 *
	 * @returns {string} - URL bazowy do uploadów.
	 *
	 * @example
	 * console.log(api.image.upload_url);
	 */

	public async GetRoutes(): Promise<ServerResponse<RoutesResponse[]>> {
		const response: ServerResponse<RoutesResponse[]> = await this.api.get('/admin/routes', {
			headers: { Authorization: `Bearer ${Api.token}` }
		});

		console.log(response.data);

		return response;
	}
}

export { Misc };
