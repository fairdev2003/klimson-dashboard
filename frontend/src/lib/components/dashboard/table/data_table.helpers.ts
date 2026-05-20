import type { TableData } from '$lib/api/requests/misc';

export function sortColumnsOrder(columns: TableData[], priorityOrder: string[]): TableData[] {
	const orderMap = new Map<string, number>();

	priorityOrder.forEach((colName, index) => {
		orderMap.set(colName, index);
	});

	return columns.sort((a, b) => {
		const aIdx = orderMap.get(a.name);
		const bIdx = orderMap.get(b.name);

		const aExists = aIdx !== undefined;
		const bExists = bIdx !== undefined;

		if (aExists && bExists) {
			return aIdx! - bIdx!;
		}
		if (aExists) return -1;
		if (bExists) return 1;

		return a.name.localeCompare(b.name);
	});
}
