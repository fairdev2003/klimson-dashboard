import type { TerminalNaming } from '$lib/components/dashboard/dev/console/terminal.svelte';
import type { SidebarItems } from '$lib/components/dashboard/sidebar/sidebar.types';

class Constants {
	public TerminalNaming: TerminalNaming = { name: 'srakens-pierdakens', path: 'main' };
	public SidebarContents: SidebarItems = contents;

	public retrieveSidebarSlugs(): string[] {
		let array: string[] = [];
		this.SidebarContents.forEach((e) => {
			array.push(e.slug);
		});

		return array;
	}

	public retrieveSidebarRoutes(): string[] {
		let array: string[] = [];
		this.SidebarContents.forEach((e) => {
			array.push(e.route);
		});

		return array;
	}
}

export default Constants;

const contents: SidebarItems = [
	{
		icon: 'material-symbols:home',
		href: '/dashboard',
		route: '',
		name: 'Hub',
		disabled: false,
		slug: 'hub'
	},
	{
		icon: 'material-symbols:database',
		href: '/dashboard/database',
		route: 'database',
		name: 'Database Editor',
		disabled: false,
		slug: 'db-database'
	},
	{
		icon: 'devicon:redis',
		href: '/dashboard/redis',
		route: 'redis',
		name: 'Redis',
		disabled: false,
		slug: 'redis'
	},

	{
		icon: 'mdi:files',
		href: '/dashboard/storage',
		route: 'storage',
		name: 'File Storage',
		disabled: false,
		slug: 'file-storage'
	},
	{
		icon: 'mdi:files',
		href: '/dashboard/v2/storage',
		route: 'v2/storage',
		name: 'V2 Storage',
		disabled: false,
		slug: 'file-storage-2'
	},
	{
		icon: 'mdi:tools',
		href: '/dashboard/tools',
		route: 'tools',
		name: 'Tools',
		disabled: false,
		slug: 'tools'
	},

	{
		icon: 'ri:todo-fill',
		href: '/dashboard/todo',
		route: 'todo',
		name: 'Todo List',
		disabled: false,
		slug: 'todo'
	},
	{
		icon: 'mdi:link',
		href: '/dashboard/routes',
		route: 'routes',
		name: 'API Routes',
		disabled: false,
		slug: 'api-routes'
	},
	{
		icon: 'mdi:user-key',
		href: '/dashboard/users?label=acc',
		route: 'users',
		name: 'CMS Access',
		disabled: false,
		slug: 'users-access'
	},
	{
		icon: 'mdi:controller-outline',
		href: '/dashboard/asteroid',
		route: 'asteroid',
		name: 'Asteroid',
		disabled: false,
		slug: 'asteroid'
	}
];
