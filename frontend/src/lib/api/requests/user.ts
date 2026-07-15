import type { AxiosInstance } from 'axios';
import type { BackendResponse, ServerResponse, User } from '../types';

export class UserClass {
	constructor(private api: AxiosInstance) {
		this.api = api;
	}

	public async Create(user: Partial<User>): Promise<ServerResponse<BackendResponse<User>>> {
		return await this.api.post(`/admin/users/new-user`, user);
	}

	public async List(): Promise<ServerResponse<BackendResponse<{ users: User[] }>>> {
		return await this.api.get(`/admin/users/get-users`);
	}

	public async GetOne(id: number | string): Promise<ServerResponse<BackendResponse<User>>> {
		return await this.api.get(`/admin/users/get-user/${id}`);
	}

	public async Update(
		id: number | string,
		data: Partial<User>
	): Promise<ServerResponse<BackendResponse<User>>> {
		return await this.api.put(`/admin/users/update-user/${id}`, data);
	}

	public async Delete(id: number | string): Promise<ServerResponse<BackendResponse<{}>>> {
		return await this.api.delete(`/admin/users/delete-user/${id}`);
	}
}
