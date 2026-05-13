import type { Stat } from '$lib/types/stats';

export type BaseInfo = {
	id?: number;
	created_at?: string;
	updated_at?: string;
};

export type Quiz = {
	title: string;
	description: string;
	image_url: string;
	public: boolean;
	edit_link: string;
	has_time_limit: boolean;
	time_limit: number;
	difficulty: string;
	expected_time_min: string;
	author: string;
	completed_count: number;
	badges: string;
	questions: Question[];
	stats: Stat[];
} & BaseInfo;

export type Question = {
	id?: number;
	createdAt?: string;
	updatedAt?: string;

	quiz_id: number;
	content: string;
	image_url: string;
	type: string;
	time_limit: number;
	answers: Answer[];
};

export type Answer = {
	id?: number;
	createdAt?: string;
	updatedAt?: string;

	question_id?: number;
	content: string;
	is_correct: boolean;
} & BaseInfo;

// old harcQuiz typing
export type RawQuestion = {
	question_id: string;
	quiz_id: string;
	question: string;
	A: string;
	B: string;
	C: string;
	D: string;
	right_answer: string;
	photo: string;
};
