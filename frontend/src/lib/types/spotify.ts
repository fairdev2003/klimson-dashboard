export type SpotifyImage = {
	height: number;
	url: string;
	width: number;
};

export type SpotifyExternalUrls = {
	spotify: string;
};

export type SpotifyArtist = {
	external_urls: SpotifyExternalUrls;
	href: string;
	id: string;
	name: string;
	type: 'artist';
	uri: string;
};

export type SpotifyAlbum = {
	album_type: 'album' | 'single' | 'compilation';
	artists: SpotifyArtist[];
	external_urls: SpotifyExternalUrls;
	href: string;
	id: string;
	images: SpotifyImage[];
	name: string;
	release_date: string;
	release_date_precision: 'day' | 'month' | 'year';
	total_tracks: number;
	type: 'album';
	uri: string;
};

export type SpotifyTrack = {
	album: SpotifyAlbum;
	artists: SpotifyArtist[];
	disc_number: number;
	duration_ms: number;
	explicit: boolean;
	external_ids: {
		isrc: string;
		[key: string]: string;
	};
	external_urls: SpotifyExternalUrls;
	href: string;
	id: string;
	is_local: boolean;
	name: string;
	popularity: number;
	preview_url: string | null;
	track_number: number;
	type: 'track';
	uri: string;
};

export type SpotifyDevice = {
	id: string | null;
	is_active: boolean;
	is_private_session: boolean;
	is_restricted: boolean;
	name: string;
	supports_volume: boolean;
	type: 'Smartphone' | 'Computer' | 'Speaker' | string;
	volume_percent: number | null;
};

export type SpotifyPlaybackActions = {
	disallows: {
		resuming?: boolean;
		skipping_prev?: boolean;
		interrupting_playback?: boolean;
		pausing?: boolean;
		seeking?: boolean;
		skipping_next?: boolean;
		toggling_repeat_context?: boolean;
		toggling_repeat_track?: boolean;
		toggling_shuffle?: boolean;
		[key: string]: boolean | undefined;
	};
};

export type SpotifyPlaybackState = {
	actions: SpotifyPlaybackActions;
	context: any | null;
	currently_playing_type: 'track' | 'episode' | 'ad' | 'unknown';
	device: SpotifyDevice;
	is_playing: boolean;
	item: SpotifyTrack | null;
	progress_ms: number | null;
	repeat_state: 'off' | 'track' | 'context';
	shuffle_state: boolean;
	smart_shuffle: boolean;
	timestamp: number;
};
