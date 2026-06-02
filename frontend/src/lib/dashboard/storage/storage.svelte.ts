class Storage {
	constructor() {}

	public print_arg<T>(arg: T) {
		console.log(typeof arg);
	}

	public delete_multiple_enabled: boolean = $state(false);
	public edit_enabled: boolean = $state(false); // for mobile
}

const storage_logic = new Storage();
export { storage_logic };
