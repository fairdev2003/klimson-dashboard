export const handle = async ({ event, resolve }) => {
	console.log(`[Request]: ${event.url.pathname}`);
	return resolve(event);
};
