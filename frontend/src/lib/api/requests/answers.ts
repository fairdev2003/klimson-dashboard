import type { AxiosInstance } from 'axios';
import type { CheckCorrectObject, QuizWhereObject, ServerResponse } from '../types';
import type { Answer } from '../../../routes/dashboard/quizzes/types';

/**
 * API handler odpowiedzialny za operacje na odpowiedziach (Answers).
 */
export class AnswerApi {
	/**
	 * Tworzy nową instancję AnswerApi.
	 *
	 * @param api - Instancja Axios do wykonywania zapytań HTTP.
	 */
	constructor(private api: AxiosInstance) {}

	/**
	 * Sprawdza czy odpowiedź jest poprawna. Jeżeli jest - w `is_correct` daje `true`, jeżeli nie daje `false`
	 *
	 * Wysyła żądanie GET na endpoint `/check/answer?id={id}` i zwraca obiekt odpowiedzi.
	 *
	 * @param where - Obiekt zawierający identyfikator odpowiedzi (`id`).
	 * @returns Promise z obiektem `Answer` zwróconym przez backend.
	 *
	 * @example
	 * ```ts
	 * const answer = await answerApi.CheckAnswer({ id: 123 });
	 * console.log(answer);
	 * ```
	 */
	public async CheckAnswer(where: QuizWhereObject): Promise<CheckCorrectObject> {
		const response: ServerResponse<CheckCorrectObject> = await this.api.get(
			`/check/answer?id=${where.id}`
		);

		return response.data;
	}
}
