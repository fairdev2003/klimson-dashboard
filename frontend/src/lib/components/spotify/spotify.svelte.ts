import axios, { type AxiosResponse } from 'axios';

import type { SpotifyPlaybackState } from '$lib/types/spotify';
import type { ServerResponse } from '$lib/api/types';
import { setCustomIconsLoader } from '@iconify/svelte';

class Spotify {
	private readonly apiRoute = 'https://api.klimson.dev/spotify/currently_playing';
	private pollingInterval = 4000;
	private progressInterval = 1000;
	public gradient: string = $state('');

	// Główny stan korzystający z typu, który przygotowaliśmy wcześniej
	public spotify: SpotifyPlaybackState | null = $state(null);
	public progress: number = $state(0);
	public duration: number = $state(0);
	public isLoading: boolean = true;

	private pollingTimer: ReturnType<typeof setTimeout> | null = null;
	private progressTimer: ReturnType<typeof setInterval> | null = null;

	constructor() {}

	private async init() {
		await this.fetchData();
		this.startPolling();
		this.startProgressUpdater();
	}

	public runCallback(callback: () => void) {
		callback();
	}

	public async zamontujKurwe() {
		await this.init();
	}

	private async fetchData() {
		try {
			const res = await axios.get(this.apiRoute);

			const responseData = res.data;

			if (responseData) {
				this.spotify = responseData;

				if (this.spotify && this.spotify.item) {
					this.progress = this.spotify.progress_ms ?? 0;
					this.duration = this.spotify.item.duration_ms;
				}
			}
		} catch (e) {
			console.error('Failed to fetch Spotify status:', e);
		}
	}

	private startPolling() {
		if (this.pollingTimer) clearTimeout(this.pollingTimer);

		const loop = async () => {
			await this.fetchData();
			this.pollingTimer = setTimeout(loop, this.pollingInterval);
		};

		loop();
	}

	private startProgressUpdater() {
		if (this.progressTimer) clearInterval(this.progressTimer);

		this.progressTimer = setInterval(() => {
			if (!this.spotify || !this.spotify.is_playing || !this.spotify.item) {
				this.progress = 0;
				this.duration = 0;
				return;
			}

			if (this.progress < this.duration) {
				this.progress += this.progressInterval;
			} else {
				this.fetchData();
			}
		}, this.progressInterval);
	}

	public formatMs(ms: number): string {
		const totalSec = Math.floor(ms / 1000);
		const min = Math.floor(totalSec / 60);
		const sec = totalSec % 60;
		return `${min}:${sec.toString().padStart(2, '0')}`;
	}

	public getAlbumCover(): string | null {
		console.log(this.spotify?.item?.album.images[0].url);
		return this.spotify?.item?.album.images[0].url;
	}

	public getSong(): string | null {
		return this.spotify?.item?.name ?? null;
	}

	public getArtist(): string | null {
		return this.spotify?.item?.artists?.[0]?.name ?? null;
	}

	public getTrackId(): string | null {
		return this.spotify?.item?.id ?? null;
	}
}

export const spotifyApp = new Spotify();
