/**
 * ImageApi - klasa do obsługi uploadu oraz pobierania obrazów
 * z backendu (upload, generowanie URL, listowanie plików).
 */

import type { AxiosInstance, AxiosResponse } from 'axios';

import type { RoutesResponse } from '../../../routes/dashboard/routes/types';
import { Api } from '../api';
import type { ServerResponse } from '../types';

export type TableData = {
	table: string;
	type: string;
	name: string;
	icon: string;
	foreign_key: string;
};

export type DiskData = {
	percentage: string;
	used: string;
	total: string;
	label: string;
	arch: string;
	os: string;
};

export type PermissionRegistry = {
	icon: string;
	name: string;
	tag: string;
	color: string;
	description: string;
	category: string;
};

class Misc {
	/**
	 * Tworzy nową instancję ImageApi.
	 *
	 * @param {AxiosInstance} api - Instancja axios używana do wykonywania requestów.
	 * @param {ApiConfig} api_config - Konfiguracja API (np. host, dev_server, baseURL).
	 */
	constructor(private api: AxiosInstance) {
		this.api = api;
	}

	/**
	 * Zwraca pełną ścieżkę bazową do katalogu `/uploads/`.
	 *
	 * @returns {string} - URL bazowy do uploadów.
	 *
	 * @example
	 * console.log(api.image.upload_url);
	 */

	public async GetRoutes(): Promise<ServerResponse<RoutesResponse[]>> {
		const response: ServerResponse<RoutesResponse[]> = await this.api.get('/admin/routes');

		return response;
	}

	public async GetPermissionRegistry(): Promise<ServerResponse<{ perms: PermissionRegistry[] }>> {
		const response: ServerResponse<{ perms: PermissionRegistry[] }> = await this.api.get(
			'/admin/permissions/all',
			{ withCredentials: true }
		);

		return response;
	}

	public async GetTables(): Promise<ServerResponse<{ tables: TableData[] }>> {
		const response: ServerResponse<{ tables: TableData[] }> = await this.api.get(
			'/admin/database/list/tables'
		);

		return response;
	}

	public async GetDisk(): Promise<ServerResponse<DiskData>> {
		const response: ServerResponse<DiskData> = await this.api.get('/admin/disk');

		return response;
	}
}

export { Misc };
