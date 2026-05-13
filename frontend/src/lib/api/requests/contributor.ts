/**
 * ImageApi - klasa do obsługi uploadu oraz pobierania obrazów
 * z backendu (upload, generowanie URL, listowanie plików).
 */

import type { AxiosInstance, AxiosResponse } from 'axios';

import type { RoutesResponse } from '../../../routes/dashboard/routes/types';
import { Api } from '../api';
import type { Contributor, Permission, Role } from '../../../routes/dashboard/contributors/types';
import type { QuizWhereObject, ServerResponse } from '../types';

class Cntrb {
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


    public async GetContributors(): Promise<ServerResponse<Contributor[]>> {
        const response: ServerResponse<Contributor[]> = await this.api.get('/admin/contributors/all', {
            headers: { Authorization: `Bearer ${Api.token}` }
        });

        return response;
    }

    public async SwitchContributorBan(where: QuizWhereObject): Promise<ServerResponse<{ block: boolean, id?: string, error?: string }>> {
        const response: ServerResponse<{ block: boolean, id?: string, error?: string }> = await this.api.put(`/admin/contributors/switch/block/${where.id}`, {}, {
            headers: {
                'Authorization': `Bearer ${localStorage.getItem('token')}`
            }
        });

        return response;
    }

    public async UpdateContributorDetails(where: QuizWhereObject, contributor: Pick<
        Contributor, "name" | "description" | "login" | "permissions">
    ): Promise<ServerResponse<{ message: string, id?: string, error?: string }>> {
        const response: ServerResponse<{ message: string, id?: string, error?: string }> =
            await this.api.put(`/admin/contributors/update/details/${where.id}`, contributor, {
                headers: { Authorization: `Bearer ${Api.token}` }
            });

        return response;
    }

    public async UpdateContributorPassword(where: QuizWhereObject, form: Pick<
        Contributor, "password" | "login">
    ): Promise<ServerResponse<{ message: string, id?: string, error?: string }>> {
        const response: ServerResponse<{ message: string, id?: string, error?: string }> =
            await this.api.put(`/admin/contributors/update/password/${where.id}`, form, {
                headers: { Authorization: `Bearer ${Api.token}` }
            });

        return response;
    }

    public async GetPermissionList(): Promise<ServerResponse<Permission[]>> {
        const response: ServerResponse<Permission[]> =
            await this.api.get(`/admin/contributors/view/permissions`, {
                headers: { Authorization: `Bearer ${Api.token}` }
            });

        return response;
    }

    public async UpdateContributorPermissions(where: QuizWhereObject, permissions: string): Promise<ServerResponse<{ message: string, error?: string }>> {

        const response: ServerResponse<{ message: string, error?: string }> = await this.api.put(`/admin/contributors/update/permissions/${where.id}`, { permissions }, {
            headers: { Authorization: `Bearer ${Api.token}` }
        });

        return response;
    }

    public async CreateContributor(data: any): Promise<ServerResponse<Contributor>> {
        const response: ServerResponse<Contributor> = await this.api.post(`/admin/contributors/create`, data, {
            headers: { Authorization: `Bearer ${Api.token}` }
        });

        console.log(response.data);

        return response;
    }

    public async DeleteContributor(where: QuizWhereObject): Promise<ServerResponse<{ message: string, error?: string }>> {
        const response: ServerResponse<{ message: string, error?: string }> = await this.api.delete(`/admin/contributors/delete/${where.id}`, {
            headers: { Authorization: `Bearer ${Api.token}` }
        });

        console.log(response.data);

        return response;
    }

    public async CheckContributorPassword(contributor_payload: Pick<Contributor, "login" | "password">): Promise<ServerResponse<{ message: string, error?: string, correct: boolean }>> {
        const response: ServerResponse<{ message: string, error?: string, correct: boolean }> = await this.api.post(`/admin/contributors/check/password`, contributor_payload, {
            headers: { Authorization: `Bearer ${Api.token}` }
        });

        return response;
    }
}


export { Cntrb };