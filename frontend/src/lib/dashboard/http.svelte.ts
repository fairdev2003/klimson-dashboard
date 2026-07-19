export type HTTPRequest = {
	id: string;
	method: string;
	endpoint: string;
	startTime: Date;
	duration?: number;
	isError: boolean;
};
export class DashboardHttpLogger {
	public httpRequests: HTTPRequest[] = $state([]);

	public pushRequest(request: HTTPRequest) {
		this.httpRequests.push(request);
	}
}
