import type { BaseInfo } from '../quizzes/types';

export type BlogType = {
	html: string;
	title: string;
	description: string;
	public: boolean;
} & BaseInfo;
