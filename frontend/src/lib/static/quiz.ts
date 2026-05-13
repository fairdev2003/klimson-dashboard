import type { Question, Quiz } from '../../routes/dashboard/quizzes/types';

export const preloaded_quiz_form: Quiz = {
	title: '',
	description: '',
	image_url: '',
	public: false,
	edit_link: '',
	has_time_limit: false,
	time_limit: 0,
	difficulty: '',
	expected_time_min: '',
	author: '',
	completed_count: 0,
	badges: '',
	questions: []
};

export const emptyQuestion: Question = {
	id: 0,
	createdAt: undefined,
	updatedAt: undefined,

	quiz_id: 0,
	content: '',
	image_url: '',
	type: '',
	time_limit: null,
	answers: []

	// jeśli BaseInfo ma jakieś pola, dodaj tu np.:
	// created_by: '',
	// description: '',
};
