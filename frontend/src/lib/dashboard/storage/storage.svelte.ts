import type { StorageRecord } from '$lib/api/requests/storage';

class Storage {
	constructor() {}

	public print_arg<T>(arg: T) {
		console.log(typeof arg);
	}

	public view: 'boxes-view' | 'list_view' = $state('boxes-view');

	public delete_multiple_enabled: boolean = $state(false);
	public edit_enabled: boolean = $state(false); // for mobile
	public storage_records: StorageRecord[] | undefined = $state();

	public selected_path: string = $state('/');

	public formatBytes(bytes: number, decimals: number = 2): string {
		if (bytes === 0) return '0 Bytes';

		const k = 1024;
		const dm = decimals < 0 ? 0 : decimals;
		const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];

		const i = Math.floor(Math.log(bytes) / Math.log(k));

		return parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + ' ' + sizes[i];
	}
}

const storage_logic = new Storage();
export { storage_logic };
