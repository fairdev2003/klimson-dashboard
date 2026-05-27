import type { AxiosInstance } from 'axios';
import type { ServerResponse } from '../types';

export type StorageRecord = {
	name: string;
	is_dir: boolean;
	file_size: number;
	modified: string;
};

class Storage {
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

	public async GetStorageRecords(path: string): Promise<ServerResponse<StorageRecord[]>> {
		const pathname = path ? path : '';

		const response: ServerResponse<StorageRecord[]> = await this.api.get(
			`/storage/list/${pathname}`
		);

		return response;
	}
}

export { Storage };
