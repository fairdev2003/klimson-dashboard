export type DatabaseTableProps = {
	database: Database;
	columns: DatabaseTableColumn[];
};

export type DatabaseTableColumn = {
	slug: string;
	name?: string;
	type: string;
};

export type Database = {
	name: string;
};
