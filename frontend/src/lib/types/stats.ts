import type { BaseInfo } from '../../routes/dashboard/quizzes/types';

export type Stat = {
	quiz_id: string;
	created_at: string;
	updated_at: string;
	id: string;
};
export interface KlimsonFetchType {
	system_os: string;
	go_version: string;
	arch: string;
	num_cpu: number;
	goroutines: number;
	thread_count: number;
	memory_alloc: string;
	memory_total: string;
	memory_sys: string;
	heap_objects: number;
	num_gc: number;
	uptime: string;
	server_time: string;
	timestamp: string;
}
