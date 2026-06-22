type SidebarBase = {
	icon?: string;
	href: string;
	name: string;
	tooltip?: string;
	description?: string;
	disabled: boolean;
};

export type SidebarItemType = {
	child?: SidebarItemChildType[];
} & SidebarBase;

export type SidebarItems = SidebarItemType[];

export type SidebarItemChildType = SidebarBase;

export const contents: SidebarItems = [
	{ icon: 'material-symbols:home', href: '/dashboard', name: 'Hub', disabled: false },
	{
		icon: 'material-symbols:database',
		href: '/dashboard/database',
		name: 'Database Editor',
		disabled: false

		// child: [
		// 	{
		// 		name: 'Users Database',
		// 		href: '/dashboard/database/users'
		// 	}
		// ]
	},
	{
		icon: 'mdi:bucket',
		href: '/dashboard/context_storage',
		name: 'Context Storage',
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
		icon: 'mdi:link',
		href: '/dashboard/routes',
		name: 'API Routes',
		disabled: false
	},
	{
		icon: 'mdi:user-key',
		href: '/dashboard/users',
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
