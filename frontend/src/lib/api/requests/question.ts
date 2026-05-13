/**
 * ImageApi - klasa do obsługi uploadu oraz pobierania obrazów
 * z backendu (upload, generowanie URL, listowanie plików).
 */

import type { AxiosInstance, AxiosResponse } from 'axios';

import type { RoutesResponse } from '../../../routes/dashboard/routes/types';
import { Api } from '../api';

import { type Question as Q } from '../../../routes/dashboard/quizzes/types';
import type { ServerResponse } from '../types';

class Question {
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

	public async GetAllQuestions(): Promise<ServerResponse<Question[] | undefined>> {
		const response: ServerResponse<Question[]> = await this.api.get('/admin/questions/all', {
			headers: { Authorization: `Bearer ${Api.token}` }
		});

		return response;
	}

	public async UpdateQuestion(q: Q): Promise<ServerResponse<Q | undefined>> {
		const response: ServerResponse<Q> = await this.api.put(`/admin/questions/update/${q.id}`, q, {
			headers: { Authorization: `Bearer ${Api.token}` }
		});

		return response;
	}

	public async UpdateMany(q: Q[]): Promise<ServerResponse<{ message: string; error?: string }>> {
		const response: ServerResponse<{ message: string; error?: string }> = await this.api.put(
			`/admin/questions/update/many`,
			q,
			{
				headers: { Authorization: `Bearer ${Api.token}` }
			}
		);

		return response;
	}

	public async CreateQuestion(
		q: Q,
		question_id: number | undefined
	): Promise<ServerResponse<Q | undefined>> {
		const response: ServerResponse<Q | undefined> = await this.api.post(
			`/admin/questions/create/${question_id}`,
			q,
			{
				headers: { Authorization: `Bearer ${Api.token}` }
			}
		);

		return response;
	}

	public async DeleteQuestion(
		question_id: number | undefined
	): Promise<ServerResponse<{ message: string; error?: string }>> {
		const response: ServerResponse<{ message: string; error?: string }> = await this.api.delete(
			`/admin/questions/delete/${question_id}`,
			{
				headers: { Authorization: `Bearer ${Api.token}` }
			}
		);

		return response;
	}
}

export { Question };
