import { get, writable, type Readable, type Writable } from 'svelte/store';

export type NotificationUser = {
	username: string;
	avatarUrl: string;
};

export type NotificationContent = {
	headerHtml: string;
	body: string;
};

export type NotificationRecord = {
	id: string;
	user: NotificationUser;
	timestamp: string;
	actionType: 'comment' | 'like' | 'mention';
	isRead: boolean;
	content: NotificationContent;
};

class NotificationService {
	private readonly MAX_LOGS = 300;

	public notifications: NotificationRecord[] = [
		{
			id: '1',
			user: {
				username: 'Klimson',
				avatarUrl: 'https://api.klimson.dev/interface/bucket/random/klimson-chill.jpeg'
			},
			timestamp: '2h ago',
			actionType: 'comment',
			isRead: false,
			content: {
				headerHtml: '<b>Commented</b> on <b>dash/origin</b>',
				body: `
                    Chłopak z Warszawy wygrał pół miliona złotych w oficjalnej aplikacji mobilnej. On siedział na facebook, zobaczył reklamę aplikacji mobilnej, pobrał aplikacje, zaczął grać i wygrał pięć tysięcy złotych. I jak to sie zdarzyło, teraz sie dowiemy. Uwielbiam grać, to moja pasja. Aby wygrywać lepiej nie słuchać ludzi, którzy nic na tym się nie rozumieją, którzy nic nigdy nie wygrywali, ilu ludzi już zostało milionerami. Część pieniędzy spędziliśmy na charytatywności, część inwestowaliśmy w mieszkanie, które w przyszłości będzie przenosić zyski. Jeszcze część spędziliśmy na podróże. Życie jest piękne, zarówno jak przed wygraniem, tak i po nim. Ale teraz mamy lepsze życie, jakościowe, ponieważ możemy sobie pozwolić na to, o czym wcześniej tylko marzyliśmy.
                `
			}
		}
	];
	private store: Writable<NotificationRecord[]> = writable([]);

	public subscribe: Readable<NotificationRecord[]>['subscribe'] = this.store.subscribe;

	public markAsRead(id: string) {
		this.store.update((notifications) =>
			notifications.map((n) => (n.id === id ? { ...n, isRead: true } : n))
		);
	}

	public getUnreadNotifications() {
		return get(this.store).filter((notification) => !notification.isRead);
	}

	public get get() {
		return get(this.store);
	}

	public add(record: NotificationRecord) {
		this.store.update((notifications) => {
			const newNotifications = [record, ...notifications];
			return newNotifications.reverse();
		});
	}

	public set(record: NotificationRecord[]) {
		this.store.set(record);
	}

	public clear() {
		this.store.set([]);
	}
}

const notifications = new NotificationService();
export { notifications };
