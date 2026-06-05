import { api } from '$lib/api/api';
import { debug } from '$lib/dashboard/stores/debug';
import type { PageLoad } from '../storage/[...path]/$types';

export const load: PageLoad = async ({ params }) => {
	const response = await api.misc.GetTables();
	console.log(response.data);

	return {
		tables: response.data.tables
	};
};
