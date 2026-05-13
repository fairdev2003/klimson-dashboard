/**
 * ImageApi - klasa do obsługi uploadu oraz pobierania obrazów
 * z backendu (upload, generowanie URL, listowanie plików).
 */

import type { AxiosInstance, AxiosResponse } from 'axios';
import type { BlogType } from '../../../routes/dashboard/blog/types';
import { Api } from '../api';
import type { QuizWhereObject, ServerResponse } from '../types';

class Blog {
	/**
	 * Tworzy nową instancję ImageApi.
	 *
	 * @param {AxiosInstance} api - Instancja axios używana do wykonywania requestów.
	 * @param {ApiConfig} api_config - Konfiguracja API (np. host, dev_server, baseURL).
	 */
	constructor(private api: AxiosInstance) {
		this.api = api;
	}

	public async CreateHero(blog: BlogType) {
		const response: AxiosResponse<BlogType> = await this.api.post('/admin/blog/create', blog, {
			headers: { Authorization: `Bearer ${Api.token}` }
		});

		console.log(response);
	}

	public async FetchBlogs(): Promise<ServerResponse<BlogType[]>> {
		const response: ServerResponse<BlogType[]> = await this.api.get('/blog/all');

		return response;
	}

	public async DeleteBlog(where: QuizWhereObject) {
		// /admin/blog/delete/:blog_id

		const response: AxiosResponse<BlogType> = await this.api.delete(
			`/admin/blog/delete/${where.id}`,
			{
				headers: { Authorization: `Bearer ${Api.token}` }
			}
		);
	}

	public async UpdateBlog(blog: BlogType) {
		console.log(`ID: ${blog.id}`);
		const response: AxiosResponse = await this.api.put(`/admin/blog/update/${blog.id}`, blog, {
			headers: { Authorization: `Bearer ${Api.token}` }
		});
		console.log(response.data);

		return response;
	}
}

export { Blog };
