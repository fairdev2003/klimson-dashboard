export type Window<T extends AcceptedWindows> = {
	slug: string;
	display_name: string;
	data: T;
};

const window: Window<DatabaseWindow> = {
	data: {
		databases: 'S'
	},
	display_name: 'SIEMA',
	slug: 's'
};

export type DatabaseWindow = {
	// ...
};

export type StorageWindow = {
	// ...
};

export type AcceptedWindows = DatabaseWindow | StorageWindow;
