import { api } from '$lib/api/api';
import type { Quiz } from '../types';
import { QuestionManager } from './question_manager.svelte';

class QuizController {
	public deleteButtonLoadingState: boolean = $state(false);

	public get currenlyEditingQuiz(): Quiz | undefined {
		return this.quiz;
	}

	private _question_class = new QuestionManager();

	public get get_question_manager(): QuestionManager {
		return this._question_class;
	}

	public quiz: Quiz | undefined = $state();

	public async PromptDeleteAndRefresh(quiz_record: Quiz) {
		const id = prompt(`Aby usunąć ten quiz wpisz id quizu ${quiz_record.id}`);

		const response = await api.quiz.DeleteQuiz(Number(id));
		console.log(response);
	}

	constructor() {}
}

export const quiz_manager = new QuizController();
