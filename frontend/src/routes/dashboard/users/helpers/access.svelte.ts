import { api } from '$lib/api/api';
import { debug } from '$lib/dashboard/stores/debug';

type Account = {};

type UserRole = {};

type PermissionRecord = {};

class AControllerData {
	public accounts: Account[] = $state([]);
	public roles: UserRole[] = $state([]);
	public permission_registry: PermissionRecord[] = $state([]);
}

class AccountController extends AControllerData {
	super() {}

	public async DumpData() {
		console.log('permission_registry: ', this.permission_registry);
		console.log('accounts: ', this.accounts);
		console.log('roles: ', this.roles);

		debug.log('permission_registry: ', this.permission_registry);
		try {
			await api.api.get('/costam');
			throw new Error('Nie udało się pobrać danych');
		} catch (err) {
			debug.error({ message: err });
		}
	}
}

const account_controller = new AccountController();

export default account_controller;
