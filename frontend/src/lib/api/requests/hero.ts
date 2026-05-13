/**
 * ImageApi - klasa do obsługi uploadu oraz pobierania obrazów
 * z backendu (upload, generowanie URL, listowanie plików).
 */

import type { AxiosInstance, AxiosResponse } from 'axios';
import { Api } from '../api';
import type { QuizWhereObject, ServerResponse } from '../types';
import type { HeroType } from '../../../routes/dashboard/hero/types';

class Hero {
	/**
	 * Tworzy nową instancję ImageApi.
	 *
	 * @param {AxiosInstance} api - Instancja axios używana do wykonywania requestów.
	 * @param {ApiConfig} api_config - Konfiguracja API (np. host, dev_server, baseURL).
	 */
	constructor(private api: AxiosInstance) {
		this.api = api;
	}

	public async CreateHero(blog: HeroType) {
		const response: AxiosResponse<HeroType> = await this.api.post('/admin/hero/create', blog, {
			headers: { Authorization: `Bearer ${Api.token}` }
		});

		console.log(response);
	}

	public async FetchHeros(): Promise<ServerResponse<HeroType[]>> {
		const response: ServerResponse<HeroType[]> = await this.api.get('/hero/all');

		return response;
	}

	public async DeleteHero(where: QuizWhereObject) {
		const response: AxiosResponse<HeroType> = await this.api.delete(
			`/admin/hero/delete/${where.id}`,
			{
				headers: { Authorization: `Bearer ${Api.token}` }
			}
		);
	}

	public async UpdateHero(blog: HeroType) {
		console.log(`ID: ${blog.id}`);
		const response: AxiosResponse = await this.api.put(`/admin/hero/update/${blog.id}`, blog, {
			headers: { Authorization: `Bearer ${Api.token}` }
		});
		console.log(response.data);

		return response;
	}
}

export { Hero };
