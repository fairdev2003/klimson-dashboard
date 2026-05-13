import type { Quiz } from '../../dashboard/quizzes/types';

export function GetWeeklyStatsCount(quiz: Quiz): number {
	const now = new Date();
	const weekAgo = new Date();
	weekAgo.setDate(now.getDate() - 7);

	return quiz.stats.filter((stat) => {
		if (!stat.created_at) return false;

		const d = new Date(stat.created_at);
		console.log(stat);
		return d > weekAgo;
	}).length;
}

export function GetDailyStatsCount(quiz: Quiz): number {
	const now = new Date();
	const weekAgo = new Date();
	weekAgo.setDate(now.getDate() - 1);

	return quiz.stats.filter((stat) => {
		if (!stat.created_at) return false;

		const d = new Date(stat.created_at);
		console.log(stat);
		return d > weekAgo;
	}).length;
}

export function GetThreeDaysStatsCount(quiz: Quiz): number {
	const now = new Date();
	const weekAgo = new Date();
	weekAgo.setDate(now.getDate() - 3);

	return quiz.stats.filter((stat) => {
		if (!stat.created_at) return false;

		const d = new Date(stat.created_at);
		console.log(stat);
		return d > weekAgo;
	}).length;
}

export function GetMonthlyStatsCount(quiz: Quiz): number {
	const now = new Date();
	const weekAgo = new Date();
	weekAgo.setDate(now.getDate() - 30);

	return quiz.stats.filter((stat) => {
		if (!stat.created_at) return false;

		const d = new Date(stat.created_at);
		console.log(stat);
		return d > weekAgo;
	}).length;
}

export function CountStats(quiz: Quiz): number {
	return quiz.stats.length;
}

export function CountDaily(quizzes: Quiz[]): number {
	let count: number = 0;
	quizzes.map((quiz) => {
		count = count + GetDailyStatsCount(quiz);
	});

	return count;
}

export function CountThreeDays(quizzes: Quiz[]): number {
	let count: number = 0;
	quizzes.map((quiz) => {
		count = count + GetThreeDaysStatsCount(quiz);
	});

	return count;
}

export function CountWeekly(quizzes: Quiz[]): number {
	let count: number = 0;
	quizzes.map((quiz) => {
		count = count + GetWeeklyStatsCount(quiz);
	});

	return count;
}

export function CountMonthly(quizzes: Quiz[]): number {
	let count: number = 0;
	quizzes.map((quiz) => {
		count = count + GetMonthlyStatsCount(quiz);
	});

	return count;
}

export function CountAll(quizzes: Quiz[]): number {
	let count: number = 0;
	quizzes.map((quiz) => {
		count = count + quiz.stats.length;
	});

	return count;
}
