import type { AxiosInstance } from 'axios';
import type { ServerResponse } from '../types';
import type { StorageRecord, V2StorageRecord } from './storage';
import { Api } from '../api';

export class StorageV2 {
	/**
	 * Tworzy nową instancję ImageApi.
	 *
	 * @param {AxiosInstance} api - Instancja axios używana do wykonywania requestów.
	 * @param {ApiConfig} api_config - Konfiguracja API (np. host, dev_server, baseURL).
	 */
	constructor(private api: AxiosInstance) {
		this.api = api;
	}

	public async GetStorageRecords(path: string): Promise<ServerResponse<V2StorageRecord[]>> {
		const pathname = path ? path : '';
		console.log(path);

		const response: ServerResponse<V2StorageRecord[]> = await this.api.get(
			`/admin/v2/storage/get/${pathname}`,
			{
				headers: {
					Authorization: `Bearer ${Api.token}`
				}
			}
		);

		return response;
	}

	public async SendImageData(file: File, path: string) {
		const formData = new FormData();
		formData.append('file', file);
		const response: ServerResponse<{ success: boolean; message: string }> = await this.api.post(
			`/admin/v2/storage/new-file/${path}`,
			formData,
			{
				headers: {
					'Content-Type': 'multipart/form-data',
					Authorization: `Bearer ${Api.token}`
				}
			}
		);

		return response;
	}

	public async CreateFolder(name: string, path: string) {
		const pathName: string = '/ ' + path;

		const response: ServerResponse<{ success: boolean; message: string }> = await this.api.post(
			`/admin/v2/storage/create-folder`,
			{ name: name, path: pathName },
			{
				headers: {
					Authorization: `Bearer ${Api.token}`
				}
			}
		);

		return response;
	}
}
