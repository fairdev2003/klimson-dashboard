import type { Question, RawQuestion } from '../../../../routes/dashboard/quizzes/types';

export const convertToQuestionType = (rawItems: RawQuestion[]): Question[] => {
	return rawItems.map((item) => {
		// Mapowanie liter na treści odpowiedzi
		const options = [
			{ key: 'A', content: item.A },
			{ key: 'B', content: item.B },
			{ key: 'C', content: item.C },
			{ key: 'D', content: item.D }
		];

		return {
			// image_url: znane_postacie_harcerskie_cz__1__(podstawowe)/Robert_Baden_Powell.jpg
			quiz_id: parseInt(item.quiz_id),
			content: item.question.trim(),
			image_url: 'https://harc-quiz.pl/images/' + item.photo,
			type: 'multiple_choice', // Domyślny typ
			time_limit: 30, // Domyślny czas w sekundach
			answers: options.map((opt) => ({
				content: opt.content,
				is_correct: opt.key === item.right_answer,
				// Pola BaseInfo (zakładam, że są opcjonalne lub puste)
				createdAt: new Date().toISOString(),
				updatedAt: new Date().toISOString()
			}))
		};
	});
};
