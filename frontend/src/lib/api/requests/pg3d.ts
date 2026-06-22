import type { AxiosInstance } from 'axios';
import type { ClanInfo, ClanResponse, PlayerData } from '../../../routes/dashboard/pg3d/pg3d.types';
import type { ServerResponse } from '../types';

class PG3D {
	constructor(private api: AxiosInstance) {
		this.api = api;
	}

	public async GetClanInfo(clan_id: string): Promise<ServerResponse<ClanResponse>> {
		const response: ServerResponse<ClanResponse> = await this.api.get(`/pg3d/clan_info/${clan_id}`);

		return response;
	}

	public async HardWarInfo(
		second_clan_id: string
	): Promise<{
		blackout: ServerResponse<ClanResponse>;
		second_clan: ServerResponse<ClanResponse>;
	}> {
		const blackout_id = '31259536';

		const [blackout, second_clan] = await Promise.all([
			this.GetClanInfo(blackout_id),
			this.GetClanInfo(second_clan_id)
		]);

		return { blackout, second_clan };
	}

	public async GetPlayerData(player_id: string): Promise<ServerResponse<PlayerData>> {
		const response: ServerResponse<PlayerData> = await this.api.get(
			`/pg3d/player_data/${player_id}`
		);

		return response;
	}
}

export { PG3D };
