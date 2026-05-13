import type { BaseInfo } from '../quizzes/types';

export type HeroType = {
	quote: string;
	author: string;
	
	image_url: string;
} & BaseInfo;
