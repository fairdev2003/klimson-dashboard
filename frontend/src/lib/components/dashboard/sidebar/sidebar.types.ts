type SidebarBase = {
	icon?: string;
	href: string;
	name: string;
	description?: string;
};

export type SidebarItemType = {
	child?: SidebarItemChildType[];
} & SidebarBase;

export type SidebarItems = SidebarItemType[];

export type SidebarItemChildType = SidebarBase;
