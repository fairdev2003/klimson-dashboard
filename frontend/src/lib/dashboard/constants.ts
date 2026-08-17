import type { TerminalNaming } from '$lib/components/dashboard/dev/console/terminal.svelte';
import type {
	SidebarItems,
	SidebarItemType
} from '$lib/components/dashboard/sidebar/sidebar.types';

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

	public findSidebarItem(route: string): SidebarItemType | undefined {
		return this.SidebarContents.find((e) => e.route === route);
	}
}

export default Constants;

const contents: SidebarItems = [
	{
		icon: 'material-symbols:home',
		href: '/dashboard',
		route: '/',
		name: 'Hub',
		disabled: false,
		slug: 'hub',
		category: 'General'
	},
	{
		icon: 'material-symbols:database',
		href: '/dashboard/database',
		route: 'database',
		name: 'Database Editor',
		disabled: false,
		slug: 'db-database',
		category: 'General'
	},
	{
		icon: 'devicon:redis',
		href: '/dashboard/redis',
		route: 'redis',
		name: 'Redis',
		disabled: false,
		slug: 'redis',
		category: 'General'
	},

	{
		icon: 'thesvg:craft-cms',
		href: '/dashboard/cms',
		route: 'cms',
		name: 'Klimson.dev',
		disabled: false,
		slug: 'users-access',
		category: 'General',
		child: [
			{
				icon: 'thesvg:craft-cms',
				href: 'cms/stories',
				route: 'users',
				name: 'Stories',
				disabled: false,
				slug: 'users-access',
				category: 'General'
			},
			{
				icon: 'thesvg:craft-cms',
				href: 'cms/docs',
				route: 'users',
				name: 'Job Documents',
				disabled: false,
				slug: 'users-access',
				category: 'General'
			},
			{
				icon: 'thesvg:craft-cms',
				href: 'cms/blog',
				route: 'users',
				name: 'Blog',
				disabled: false,
				slug: 'users-access',
				category: 'General'
			},
			{
				icon: 'thesvg:craft-cms',
				href: 'cms/server',
				route: 'users',
				name: 'Server Settings',
				disabled: false,
				slug: 'users-access',
				category: 'General'
			}
		]
	},
	{
		icon: 'mdi:files',
		href: '/dashboard/storage',
		route: 'storage',
		name: 'File Storage',
		disabled: false,
		slug: 'file-storage',
		category: 'General'
	},
	{
		icon: 'mdi:files',
		href: '/dashboard/v2/storage',
		route: 'v2/storage',
		name: 'V2 Storage',
		disabled: false,
		slug: 'file-storage-2',
		category: 'General'
	},
	{
		icon: 'mdi:tools',
		href: '/dashboard/tools',
		route: 'tools',
		name: 'Tools',
		disabled: false,
		slug: 'tools',
		category: 'General'
	},

	{
		icon: 'ri:todo-fill',
		href: '/dashboard/todo',
		route: 'todo',
		name: 'Todo List',
		disabled: false,
		slug: 'todo',
		category: 'General'
	},
	{
		icon: 'mdi:link',
		href: '/dashboard/routes',
		route: 'routes',
		name: 'API Routes',
		disabled: false,
		slug: 'api-routes',
		category: 'General'
	},
	{
		icon: 'mdi:user-key',
		href: '/dashboard/users?label=acc',
		route: 'users',
		name: 'CMS Access',
		disabled: false,
		slug: 'users-access',
		category: 'General'
	}
];
