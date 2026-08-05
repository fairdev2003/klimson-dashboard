import { api } from '$lib/api/api';
import type { PermissionRegistry } from '$lib/api/requests/misc';
import { debug } from '$lib/dashboard/stores/debug';
import type { Role } from '$lib/types/user';

type Account = {};

type PermissionRecord = {};

const initialForm = {
	name: '',
	color: '',
	permissions: [],
	icon: ''
};

class AControllerData {
	public accounts: Account[] = $state([]);
	public roles: Role[] = $state([]);
	public permission_registry: PermissionRecord[] = $state([]);

	public role: Role = $state({ ...initialForm });

	public resetRole() {
		this.role = initialForm;
	}
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

	public roleHasPermission(perm: PermissionRegistry): boolean {
		let temp_table: string[] = [];
		this.role.permissions.forEach((e) => {
			temp_table.push(e.name);
		});

		return temp_table.includes(perm.tag);
	}

	public ImplementPermToRole(perm: PermissionRegistry): void {
		if (this.roleHasPermission(perm)) return;

		this.role = { ...this.role, permissions: [...this.role.permissions, { name: perm.tag }] };
		debug.log(this.role);
	}

	public RemovePermFromRole(perm: PermissionRegistry): void {
		if (!this.roleHasPermission(perm)) return;

		const newPerms = this.role.permissions.filter((e) => e.name !== perm.tag);

		this.role = { ...this.role, permissions: newPerms };
		debug.log(this.role);
	}

	public async CreateNewRole(role?: Role) {
		if (role) {
			this.role = role;
		}

		try {
			const response = await api.user.CreateRole(this.role);

			if (response.status === 201) {
				debug.system('User successfully created!');
				debug.log('Response: ', response.data);
			}
		} catch (error) {
			debug.error(error);
		} finally {
			debug.system("Trycatch statement 'AccountController.CreateNewRole()' has ended!");
		}
	}

	public async FetchRolesAndAssign() {
		try {
			const response = await api.user.ListRoles();

			this.roles = response.data.data;
		} catch (error) {
			debug.error(error);
		} finally {
			debug.system("Trycatch statement 'AccountController.CreateNewRole()' has ended!");
		}
	}

	public async DeleteRoleAndFetchNew(id: number): boolean {
		try {
			const delete_response = await api.user.DeleteRole(id);

			if (delete_response.status === 200) {
				debug.system('Roles successfully deleted!');
				await this.FetchRolesAndAssign();
			}
		} catch (error) {
			debug.error(error);
		} finally {
			debug.system("Trycatch statement 'AccountController.CreateNewRole()' has ended!");
		}
		return true;
	}

	public async UpdateRoleAndFetchNew(role: Role): Promise<boolean> {
		try {
			const update_response = await api.user.UpdateRole(role.id, role);

			if (update_response.status === 200) {
				debug.system('Roles successfully updated!');
				await this.FetchRolesAndAssign();
			}
		} catch (error) {
			debug.error(error);
		} finally {
			debug.system("Trycatch statement 'AccountController.CreateNewRole()' has ended!");
		}
		return true;
	}
}

const account_controller = new AccountController();

export default account_controller;
