import type { AxiosInstance } from 'axios';
import type { QuizWhereObject, ServerResponse } from '../types';

import type { Stat } from '$lib/types/stats';
import { Api } from '../api';

/**
 * API handler odpowiedzialny za operacje na odpowiedziach (Answers).
 */
export class Stats {
	constructor(private api: AxiosInstance) { }

	public async NewStat({ id }: QuizWhereObject) {
		const response = await this.api.post('/stats/create', { quiz_id: id });

		return response;
	}

	public async AllStats(): Promise<Stat[]> {
		const response: ServerResponse<Stat[]> = await this.api.get('/stats/all');

		return response.data;
	}

	public async Count(): Promise<number> {
		const response: ServerResponse<{ count: number }> = await this.api.get('/stats/count');

		return response.data.count;
	}

	// TODO: change any to actual type
	public async DeleteStat(where: QuizWhereObject): Promise<any> {
		const response: ServerResponse<any> = await this.api.delete(`/admin/stats/delete/${where.id}`, {
			headers: { Authorization: `Bearer ${Api.token}` }
		});

		return response
	}
}
