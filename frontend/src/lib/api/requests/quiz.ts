import type { AxiosInstance, AxiosResponse } from 'axios';
import type { Quiz } from '../../../routes/dashboard/quizzes/types';
import { Api } from '../api';
import type { QuizWhereObject, ServerResponse, UpdateOneFieldType } from '../types';

/**
 * Klasa reprezentująca sekcję `quiz` w API.
 *
 * Zawiera metody do wykonywania zapytań HTTP
 * związanych z quizami (pobieranie, tworzenie, aktualizacja itd.).
 *
 * Każda metoda zwraca dane (`res.data`) bezpośrednio z odpowiedzi serwera.
 */
export class QuizApi {
	/**
	 * Tworzy nową instancję sekcji `QuizApi`.
	 *
	 * @param {AxiosInstance} api - Instancja klienta Axios używana do wykonywania żądań HTTP.
	 *
	 * @example
	 * ```ts
	 * const api = new Api();
	 * const quizzes = await api.quiz.getAll();
	 * console.log(quizzes);
	 * ```
	 */
	constructor(private api: AxiosInstance) {}

	/**
	 * Pobiera listę wszystkich quizów z serwera.
	 *
	 * @returns {Promise<any>} Obietnica zwracająca dane z API (`res.data`).
	 *
	 * @example
	 * ```ts
	 * const allQuizzes = await api.quiz.getAll();
	 * console.log(allQuizzes);
	 * ```
	 */
	public async GetAll(): Promise<ServerResponse<Quiz[]>> {
		const response: ServerResponse<Quiz[]> = await this.api.get('/admin/quizzes/all', {
			headers: { Authorization: `Bearer ${Api.token}` }
		});

		return response;
	}

	public async GetPublicQuizzes(): Promise<ServerResponse<Quiz[]>> {
		const response: ServerResponse<Quiz[]> = await this.api.get('/quizzes');

		return response;
	}

	public async GetPublicQuiz(where: QuizWhereObject): Promise<Quiz> {
		const response: ServerResponse<Quiz> = await this.api.get(`/quiz?id=${where.id}`);

		return response.data;
	}

	public async CreateQuiz(data: Quiz): Promise<ServerResponse<any>> {
		const response: ServerResponse<any> = await this.api.post('/admin/quizzes', data, {
			headers: { Authorization: `Bearer ${Api.token}` }
		});

		console.log('Duration ', response.duration);

		return response;
	}

	public async DeleteQuiz(id: number): Promise<any> {
		const response: ServerResponse<any> = await this.api.delete(`/admin/quizzes/${id}`, {
			headers: { Authorization: `Bearer ${Api.token}` }
		});
		return response;
	}

	public async UpdateQuiz(where: QuizWhereObject, quiz: Quiz): Promise<Quiz> {
		const response: ServerResponse<Quiz> = await this.api.put(`/admin/quizzes/${where.id}`, quiz, {
			headers: { Authorization: `Bearer ${Api.token}` }
		});

		console.log(response.duration);

		return response.data;
	}

	public async UpdateBasicInfo(
		where: QuizWhereObject,
		data: Pick<Quiz, 'author' | 'description' | 'title' | 'difficulty'>
	): Promise<ServerResponse<{ message: string; id: string; error?: string }>> {
		const response: ServerResponse<{ message: string; id: string; error?: string }> =
			await this.api.put(`/admin/quizzes/update/basic/${where.id}`, data, {
				headers: { Authorization: `Bearer ${Api.token}` }
			});

		return response;
	}

	public async SaveBasicInfo(
		data: Pick<Quiz, 'author' | 'description' | 'title' | 'difficulty'>
	): Promise<ServerResponse<{ message: string; id: number; error?: string }>> {
		const response: ServerResponse<{ message: string; id: number; error?: string }> =
			await this.api.put(`/admin/quizzes/save/basic`, data, {
				headers: { Authorization: `Bearer ${Api.token}` }
			});

		return response;
	}

	public async UpdateQuizImage(
		where: QuizWhereObject,
		data: Pick<Quiz, 'image_url'>
	): Promise<ServerResponse<{ message: string; id: number; error?: string }>> {
		const response: ServerResponse<{ message: string; id: number; error?: string }> =
			await this.api.put(`/admin/quizzes/update/image/${where.id}`, data, {
				headers: { Authorization: `Bearer ${Api.token}` }
			});

		return response;
	}

	public async GetAdminQuiz(where: QuizWhereObject): Promise<ServerResponse<Quiz>> {
		const response: ServerResponse<Quiz> = await this.api.get(`/admin/quizzes/quiz/${where.id}`, {
			headers: { Authorization: `Bearer ${Api.token}` }
		});

		console.log(response.duration);

		return response;
	}

	public async UpdateOneField(where: QuizWhereObject, data: UpdateOneFieldType<Quiz>) {
		const response: ServerResponse<Quiz> = await this.api.put(
			`/admin/quizzes/update/field/${where.id}`,
			data,
			{
				headers: { Authorization: `Bearer ${Api.token}` }
			}
		);
	}

	public async UpdateSettings(
		where: QuizWhereObject,
		data: Pick<Quiz, 'public'>
	): Promise<ServerResponse<{ message: string; id: string; error?: string }>> {
		const response: ServerResponse<{ message: string; id: string; error?: string }> =
			await this.api.put(`/admin/quizzes/update/settings/${where.id}`, data, {
				headers: { Authorization: `Bearer ${Api.token}` }
			});

		return response;
	}
}
