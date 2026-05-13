import type { BaseInfo } from '../quizzes/types';

export type Contributor = {
	name: string;
	description: string;
	login: string;
	password: string;
	permissions: string;
	thumbnail: string;
	profile_image: string;
	last_login: string;
	blocked_till: boolean;
	logs?: string[];

	pending_quizzes?: number[];
} & BaseInfo;

export type RoleOption = {
	id: string;
	name: string;
	description: string;
	icon: string;
	color: string;
};

export type ContributorLog = {
	action: string;
	timestamp: Date;
} & BaseInfo;

export type Role = {
	name: string;
} & BaseInfo;

export type OptionType = {
	label: string;
	description: string;
	icon: string;
	action: () => void;
	color?: string;
	visible: boolean;
};

export type Permission = {
	tag: string;
	name: string;
	color: string;
	description: string;
	icon: string;
};
