import type { AxiosInstance } from 'axios';
import type { ServerResponse } from '../types';
import { Api } from '../api';

export type StorageRecord = {
	name: string;
	is_dir: boolean;
	file_size: number;
	modified: string;
};

export type V2StorageRecord = {
	name: string;
	file_type: string;
	file_name: string;
	path: string;
	description: string;
};

type FolderCreationBody = {
	folder_name: string;
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

	public async GetLatestStorageRecords(): Promise<ServerResponse<{ files: StorageRecord[] }>> {
		const response: ServerResponse<{ files: StorageRecord[] }> = await this.api.get(
			`/admin/storage/latest`,
			{
				headers: { Authorization: `Bearer ${Api.token}` }
			}
		);

		return response;
	}

	public async PushChangedTextFile(
		content: { content: string },
		path: string
	): Promise<ServerResponse> {
		const response: ServerResponse = await this.api.post(
			`/admin/storage/interface/edit-file/${path}`,
			content,
			{
				headers: { Authorization: `Bearer ${Api.token}` }
			}
		);

		return response;
	}

	public async AddFile(path: string): Promise<ServerResponse<{}>> {
		const response: ServerResponse<{}> = await this.api.post(
			`/admin/storage/interface/new-file/${path}`,
			{ content: '' },
			{
				headers: { Authorization: `Bearer ${Api.token}` }
			}
		);

		return response;
	}

	public async CreateFolder(
		path: string,
		folder: FolderCreationBody
	): Promise<ServerResponse<{ message: string; success: boolean }>> {
		const pathname = path ? path : '';

		const response: ServerResponse<{ message: string; success: boolean }> = await this.api.post(
			`/admin/storage/interface/create-folder/${pathname}`,
			folder,
			{
				headers: { Authorization: `Bearer ${Api.token}` }
			}
		);

		return response;
	}

	// W pliku, gdzie masz klasę Storage
	public async UploadFile(
		path: string,
		file: File
	): Promise<ServerResponse<{ message: string; success: boolean }>> {
		const pathname = path ? path : '';

		const formData = new FormData();
		formData.append('file', file);

		const response: ServerResponse<{ message: string; success: boolean }> = await this.api.post(
			`/admin/storage/interface/upload-file/${pathname}`,
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

	public async DeleteItem(
		path: string
	): Promise<ServerResponse<{ success: boolean; message: string }>> {
		const response: ServerResponse<{ success: boolean; message: string }> = await this.api.delete(
			`/admin/storage/interface/delete/${path}`,
			{
				headers: { Authorization: `Bearer ${Api.token}` }
			}
		);
		return response;
	}

	public async RenameItem(
		path: string,
		newName: string
	): Promise<ServerResponse<{ success: boolean }>> {
		const response: ServerResponse<{ success: boolean }> = await this.api.post(
			`/admin/storage/interface/rename/${path}`,
			{ newName },
			{ headers: { Authorization: `Bearer ${Api.token}` } }
		);
		return response;
	}
}

export { Storage };
