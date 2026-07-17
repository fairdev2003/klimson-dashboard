class Misc {
	public art: string = `
 ____ ___ _____ __  __    _       ______        _______ _    _   _ 
/ ___|_ _| ____|  \\/  |  / \\     / ___\\ \\      / / ____| |  | | | |
\\___ \\| ||  _| | |\\/| | / _ \\   | |    \\ \\ /\\ / /|  _| | |  | | | |
 ___) | || |___| |  | |/ ___ \\  | |___  \\ V  V / | |___| |__| |_| |
|____/___|_____|_|  |_/_/   \\_\\  \\____|  \\_/\\_/  |_____|_____\\___/ 
`;

	public GetDateTime(): string {
		const now = new Date();
		const currentTime = now.toLocaleTimeString('us-US', {
			hour: '2-digit',
			minute: '2-digit',
			second: '2-digit'
		});
		const currentDate = now.toLocaleDateString('us-US', {
			day: 'numeric',
			month: 'long',
			year: 'numeric'
		});
		return `${currentTime}, ${currentDate}`;
	}
}

const DashboardMisc = new Misc();
export { DashboardMisc };
