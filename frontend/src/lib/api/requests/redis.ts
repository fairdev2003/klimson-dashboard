import type { AxiosInstance } from 'axios';
import { Api } from '../api';
import type { BackendResponse, ServerResponse } from '../types';

export class Redis {
	/**
	 * Tworzy nową instancję ImageApi.
	 *
	 * @param {AxiosInstance} api - Instancja axios używana do wykonywania requestów.
	 * @param {ApiConfig} api_config - Konfiguracja API (np. host, dev_server, baseURL).
	 */
	constructor(private api: AxiosInstance) {
		this.api = api;
	}

	public async GetKeys(): Promise<ServerResponse<BackendResponse<{ rdbs: string[] }>>> {
		const response: ServerResponse<BackendResponse<{ rdbs: string[] }>> =
			await this.api.get(`/admin/redis/keys`);

		return response;
	}
	public async PingRedis(): Promise<ServerResponse<BackendResponse<{ ping: string }>>> {
		const response: ServerResponse<BackendResponse<{ ping: string }>> =
			await this.api.get(`/redis/ping`);

		return response;
	}
	public async Get(key: string): Promise<ServerResponse<BackendResponse<{ result: any }>>> {
		const response: ServerResponse<BackendResponse<{ result: any }>> = await this.api.put(
			`/admin/redis/get?key=${key}`
		);

		return response;
	}
	public async Set(key: string, value: string): Promise<ServerResponse<BackendResponse<{}>>> {
		const response: ServerResponse<BackendResponse<{}>> = await this.api.put(
			`/admin/redis/set?key=${key}&value=${value}`
		);

		return response;
	}
	public async Del(key: string): Promise<ServerResponse<BackendResponse<{}>>> {
		const response: ServerResponse<BackendResponse<{}>> = await this.api.delete(
			`/admin/redis/del?key=${key}`
		);

		return response;
	}

	public async KeyInfo(key: string): Promise<ServerResponse<BackendResponse<{}>>> {
		const response: ServerResponse<BackendResponse<{}>> = await this.api.delete(
			`/admin/redis/key/info?key=${key}`
		);

		return response;
	}
}
