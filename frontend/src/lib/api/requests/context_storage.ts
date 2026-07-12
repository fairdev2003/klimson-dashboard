import type { AxiosInstance } from 'axios';
import { DatabaseTable } from './database.table';
import type { ServerResponse } from '../types';
import { Api } from '../api';
import type { BaseInfo } from '../../../routes/dashboard/quizzes/types';

export type ContextStorageType = BaseInfo & {
	id?: number;
	key: string;
	value: string;
	category_name: string;
	type: string;
	icon: string;
	is_public: boolean;
};

class ContextStorage {
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

	public async GetPrivateContextStorages(): Promise<ServerResponse<ContextStorageType[]>> {
		const response: ServerResponse<ContextStorageType[]> = await this.api.get(
			`/admin/context_storage/private`
		);

		return response;
	}

	public async CreateContextStorage(
		context_storage: ContextStorageType
	): Promise<ServerResponse<ContextStorageType[]>> {
		const response: ServerResponse<ContextStorageType[]> = await this.api.post(
			`/admin/context_storage/create`
		);

		return response;
	}
	public async DeleteContextStorage(id: number): Promise<ServerResponse<ContextStorageType[]>> {
		const response: ServerResponse<ContextStorageType[]> = await this.api.delete(
			`/admin/context_storage/delete/${id}`
		);

		return response;
	}

	// /context_storage/private/single

	public async GetSinglePrivateContext(key: string): Promise<ServerResponse<ContextStorageType>> {
		const response: ServerResponse<ContextStorageType> = await this.api.get(
			`/admin/context_storage/private/single/${key}`
		);

		return response;
	}

	public async UpdateContextStorage(
		key: string,
		context_storage: ContextStorageType
	): Promise<ServerResponse<ContextStorageType>> {
		const response: ServerResponse<ContextStorageType> = await this.api.put(
			`/admin/context_storage/update/${key}`,
			context_storage
		);

		return response;
	}

	///context_storage/update/:id
}

export { ContextStorage };
