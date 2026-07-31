import type { TerminalNaming } from '$lib/components/dashboard/dev/console/terminal.svelte';
import type { SidebarItems } from '$lib/components/dashboard/sidebar/sidebar.types';
import { sidebar_open } from './stores/store';

class Constants {
	public TerminalNaming: TerminalNaming = { name: 'srakens-pierdakens', path: 'main' };
	public SidebarContents: SidebarItems = contents;
}

export default Constants;

const contents: SidebarItems = [
	{ icon: 'material-symbols:home', href: '/dashboard', name: 'Hub', disabled: false },
	{
		icon: 'material-symbols:database',
		href: '/dashboard/database',
		name: 'Database Editor',
		disabled: false
	},
	{
		icon: 'devicon:redis',
		href: '/dashboard/redis',
		name: 'Redis',
		disabled: false
	},

	{
		icon: 'mdi:files',
		href: '/dashboard/storage',
		name: 'File Storage',
		disabled: false
	},
	{
		icon: 'mdi:files',
		href: '/dashboard/v2/storage',
		name: 'V2 Storage',
		disabled: false
	},
	{
		icon: 'mdi:tools',
		href: '/dashboard/tools',
		name: 'Tools',
		disabled: false
	},

	{
		icon: 'ri:todo-fill',
		href: '/dashboard/todo',
		name: 'Todo List',
		disabled: false
	},
	{
		icon: 'mdi:link',
		href: '/dashboard/routes',
		name: 'API Routes',
		disabled: false
	},
	{
		icon: 'mdi:user-key',
		href: '/dashboard/users?label=acc',
		name: 'CMS Access',
		disabled: false
	},
	{
		icon: 'mdi:controller-outline',
		href: '/dashboard/asteroid',
		name: 'Asteroid',
		disabled: false
	}
];
