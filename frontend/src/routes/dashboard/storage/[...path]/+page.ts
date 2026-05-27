import type { RouteParams } from '$app/types';
import { api } from '$lib/api/api';
import type { PageLoad } from '../$types';

export const load: PageLoad = async ({ params }) => {
	const response = await api.storage.GetStorageRecords(params.path || '');
	const path_table = (params.path as string).split('/');

	return {
		storage_file_list: response.data,
		path_table
	};
};
