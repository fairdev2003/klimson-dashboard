export type SidebarBaseType = {
	id: string;
	desc: string;
	label: string;
	onclick?: () => void;
	link: string;
};

export type SidebarItemChildrenType = SidebarBaseType & {};

export type SidebarItemType = {
	id: string;
	label: string;
	desc: string;
	icon: string;
	onclick?: () => void;
	link: string;
	children: SidebarItemChildrenType[];
};
