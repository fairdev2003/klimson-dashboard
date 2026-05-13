import { toast } from '$lib/dashboard/stores/toast';
import type { Question } from '../types';

export class QuestionManager {
	public questions: Question[] | undefined;

	public selectedQuestion: Question | undefined = $state();

	public selectedQuestionId: number = $state(0);

	public onRecordClick(index: number) {
		this.selectedQuestionId = index;
	}

	public changeSelectedQuestionIndex(index: number) {
		this.selectedQuestionId = index;
	}

	public toggleSelectedQuestionIndex(index: number) {
		if (index === this.selectedQuestionId) {
			this.selectedQuestionId = -1;
			return;
		}
		this.selectedQuestionId = index;
	}

	public addNewDebugQuiz() {}
}

export const question_manager = new QuestionManager();
