type SidebarBase = {
	icon?: string;
	href: string;
	name: string;
	tooltip?: string;
	description?: string;
	disabled: boolean;
	slug: string;
	route: string;
	category: string;
};

export type SidebarItemType = {
	child?: SidebarItemChildType[];
} & SidebarBase;

export type SidebarItems = SidebarItemType[];

export type SidebarItemChildType = SidebarBase;
