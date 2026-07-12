import type { AxiosInstance } from 'axios';
import type { ServerResponse } from '../types';
import { Api } from '../api';
import type { TableData } from './misc';

export type ColumnData = {
	name: string;
	type: string;
	slug: string;
};

export type TableDataType = {
	status: 'success' | 'error';
	table: string;
	columns: TableData[];
	count: number;
	data: any[];
};

class DatabaseTable {
	constructor(private api: AxiosInstance) {
		this.api = api;
	}

	public async GetTableData(table_name: string): Promise<ServerResponse<TableDataType>> {
		const response: ServerResponse<TableDataType> = await this.api.get(
			`/admin/database/table/${table_name}`
		);

		return response;
	}
}

export { DatabaseTable };
