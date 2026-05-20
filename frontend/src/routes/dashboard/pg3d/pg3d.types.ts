export type ClanInfo = {
	clan_announcement: string;
	clan_id: string;
	clan_logo: string;
	clan_name: string;
	continent: string;
	division: string;
	fort_energy: number;
	fort_power: number;
	leader_avatar: string;
	leader_coins: number;
	leader_gems: number;
	leader_id: string;
	leader_join_date: string;
	leader_name: string;
	members_count: number;
	members_valor: number;
	motto: string;
	tank_energy: number;
	tank_power: number;
	type: 'Private' | 'Public';
	valor_points: number;
};

export type ClanMember = {
	id: string;
	level: string;
	name: string;
	rank:
		| 'Newbie'
		| 'Lieutenant'
		| 'Captain'
		| 'Major'
		| 'Colonel'
		| 'Lt. Colonel'
		| 'General'
		| 'Emperor'
		| string;
	valor: number;
};

export type ClanSnapshot = {
	clan_info: ClanInfo;
	members: ClanMember[];
	timestamp: number;
};

export type ClanResponse = {
	clan_info: ClanInfo;
	members: ClanMember[];
	search_count: number;
	snapshots: ClanSnapshot[];
	success: boolean;
};

export interface PlayerData {
	info: PlayerInfo;
}

export interface PlayerInfo {
	avatar: string;
	username: string;
	avatar_count: number;
	banned: 'Yes' | 'No' | string;
	cape: string;
	clan: PlayerClan | null;
	creation_date: string;
	creation_date_unix: number;
	currency: PlayerCurrency;
	gadgets_count: number;
	id: string;
	last_online: string;
	level: number;
	loadouts: Record<string, PlayerLoadout>;
	lottery: LotteryItem[];
	match_history: MatchHistoryItem[];
	modules: PlayerModules;
	pets_count: number;
	searched_player_id: string;
	sets_count: number;
	skins: PlayerSkin[];
	weapons_count: number;
	wear: WearItem[];
}

export interface PlayerClan {
	clanid: number;
	clanlogo: string;
	clanname: string;
	clanrank: string;
}

export interface PlayerCurrency {
	'Battle Pass Tickets'?: number;
	Coins?: number;
	Coupons?: number;
	Gems?: number;
	'Pixel Pass Tickets'?: number;
	Silver?: number;
	[key: string]: number | undefined;
}

export interface EquipmentItem {
	id: number;
	name: string | null;
	code_name?: string | null;
}

export interface WeaponItem {
	id: number;
	name: string;
}

export interface PlayerLoadout {
	index: number;
	name: string;
	armor: EquipmentItem;
	backup: WeaponItem;
	boots: EquipmentItem;
	cape: EquipmentItem;
	hat: EquipmentItem;
	heavy: WeaponItem;
	mask: EquipmentItem;
	melee: WeaponItem;
	primary: WeaponItem;
	sniper: WeaponItem;
	special: WeaponItem;
	throwable: EquipmentItem;
	tools: EquipmentItem;
}

export interface LotteryItem {
	amount: number;
	index: number;
	name: string;
}

export interface MatchPlayer {
	id: string;
	kills: number;
	name: string;
	score: number;
	team: number;
}

export interface MatchHistoryItem {
	game_mode: number;
	players: MatchPlayer[];
	timestamp: number;
}

export interface PlayerModules {
	[moduleName: string]: number;
}

export interface PlayerSkin {
	frontview: string;
	skin: string;
}

export interface WearItem {
	code_name: string;
	index: string;
	name: string;
	rarity: number;
}
