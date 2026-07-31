import { api } from '$lib/api/api';
import type { BackendResponse, ServerResponse } from '$lib/api/types';
import { debug } from '$lib/dashboard/stores/debug';
import { writable } from 'svelte/store';
import { AutoComplete, CommandBuilder } from './command_builder.svelte';
import { goto } from '$app/navigation';
import { terminal } from './terminal.svelte';
import axios from 'axios';
import Dashboard from '$lib/dashboard/dashboard.svelte';
import { formatter } from './formatter';
import { bold, tail, italic, red, span } from '$lib/terminal/style';
import { notifications, type NotificationRecord } from '../../navbar/notification_service.svelte';
import type { StorageRecord } from '$lib/api/requests/storage';
import { base_url } from '$lib/api/api.store';

export class ConsoleService {
	public activeRequests = $state(new Map<string, AbortController>());
	private commands: Map<string, CommandBuilder> = new Map();
	private unknown_command_handler: (user_input: string, name: string) => void = $state(() => {});
	private on_command_handler: (
		command: CommandBuilder | undefined,
		input: string | undefined
	) => void | undefined = () => {};

	public loading = $state(false);

	public registerCommand(name: string): CommandBuilder {
		const cmd = new CommandBuilder(name);
		this.commands.set(name, cmd);
		return cmd;
	}

	public get hasActiveRequests() {
		return this.activeRequests.size > 0;
	}

	public dumpAvailableCommands() {
		let commandsList: string[] = [];

		this.commands.forEach((e) => {
			commandsList.push(e.name);
		});
		return commandsList.join(', ');
	}

	constructor() {}

	public onUnknownCommand(handler: (user_input: string, name: string) => void) {
		this.unknown_command_handler = handler;
	}

	public onCommand(
		handler: (command: CommandBuilder | undefined, input?: string | undefined) => void,
		input?: string | undefined
	) {
		this.on_command_handler = handler;

		return;
	}

	public getCommandsRegister(): CommandBuilder[] {
		return Array.from(this.commands.values());
	}

	public run(input: string) {
		const regex = /[^\s"]+|"([^"]*)"/g;
		const parts: string[] = [];
		let match;

		while ((match = regex.exec(input.trim())) !== null) {
			parts.push(match[1] ? match[1] : match[0]);
		}

		const name = parts[0];
		const args = parts.slice(1);
		const command = this.commands.get(name);

		if (this.on_command_handler != undefined) {
			this.on_command_handler(command, input);
		}

		if (parts.includes('help')) {
			if (!command) return;

			let usage_string: string = '';
			usage_string = usage_string + name;

			command.argHandlers.forEach((argHandler, i) => {
				if (!argHandler.config?.auto_complete_args) {
					usage_string = usage_string + ` <$arg${i + 1}>`;
					return;
				}
				usage_string = usage_string + ` <${argHandler.config?.auto_complete_args.join(' | ')}>`;
			});
			debug.log('Description: ', command.description);
			debug.log('Usage: ', usage_string);
			return;
		}

		if (command) {
			command.execute(args, input);
		} else {
			if (this.unknown_command_handler != undefined) {
				this.unknown_command_handler(input, name);
			}

			console.warn(`Unknown command: ${name}`);
		}
	}

	public runAndOpenTerminal(input: string) {
		this.run(input);

		terminal.terminalOpened = true;
	}
}

const console_service = new ConsoleService();
console_service.onUnknownCommand((input, name) => {
	debug.console(input);

	debug.system(`Command with name '${name}' does not exist!`);
	debug.system(`Type 'PrettyFormatRecord' to view available commands.`);
});

console_service.onCommand((command, input) => {
	if (!command) {
		return;
	}
	if (!input) {
		return;
	}

	terminal.set_input({ user_input: input, id: terminal.input_history.length + 1 });
	debug.console(input);
	terminal.last_record_user_iterator = -1;
});

console_service
	.registerCommand('clear')
	.setDescription('Clears the terminal')
	.setAction(() => {
		debug.clear();
	});

console_service
	.registerCommand('test')
	.setDescription('Test command')
	.setAction(async (args) => {
		try {
			const response: ServerResponse<
				BackendResponse<{
					token: string;
					claims: {
						name: string;
						login: string;
						exp: string;
					};
				}>
			> = await api.api.get('/admin/users/me');

			if (response.status === 200) {
				if (!args || args[0] === 'claims') {
					debug.silent(response.data.claims);
					return;
				}
			}
		} catch (error) {
			if (axios.isAxiosError(error)) {
				error.cause;
			}
		}
	});

console_service
	.registerCommand('state')
	.setDescription('Golang state management')
	.setAction(async (args) => {
		let [action, key, value] = args;

		if (action === 'get') {
			if (!key) {
				debug.error('Key arg is required to get state from golang server');
				return;
			}

			try {
				const response = await api.api.get(`/state/get/${key}`);

				debug.log(response.data);
			} catch (error) {
				debug.error(error);
			}
			return;
		}
		if (action === 'set') {
			if (!key) {
				debug.error('Key arg is required to get state from golang server');
				return;
			}
			if (!value) {
				debug.error('Key arg is required to set server golang state');
				return;
			}
			try {
				const response = await api.api.post(`/state/set`, {
					key,
					value
				});

				debug.log(response.data);
			} catch (error) {
				debug.error(error);
			}
			return;
		}

		debug.error('Action arg is required!');
	});

console_service
	.registerCommand('history')
	.setDescription('Prinitng user input history.')
	.setAction(() => {
		debug.silent(terminal.input_history);
	});

console_service
	.registerCommand('logs')
	.setDescription('Prinitng all terminal logs in JSON format')
	.setAction(() => {
		debug.logStore();
	});

console_service
	.registerCommand('reload')
	.setDescription('Reload the page.')
	.setAction(() => {
		debug.system('Reloading the page....');
		window.location.reload();
	});

console_service
	.registerCommand('logout')
	.setDescription('Terminating cms session')
	.setAction(async () => {
		try {
			await api.api.post('/auth/logout');
		} catch (error) {
			debug.error('Erorr during logout');
			debug.error(error);
		} finally {
			goto('/login');
		}
	});

console_service
	.registerCommand('lockscreen')
	.setDescription('Terminating cms session')
	.setAction(async () => {
		try {
			await api.api.post('/auth/logout');
		} catch (error) {
			debug.error('Erorr during logout');
			debug.error(error);
		} finally {
			goto('/login/lockscreen');
		}
	});

console_service
	.registerCommand('api')
	.setDescription('Do actions based on api.ts state')
	.setAction((args) => {
		const [action_arg, server_arg] = args;

		type Action = { handler: () => void; name: string };
		type Api = { href: string; name: string };

		const apis: Api[] = [
			{ href: 'https://api.klimson.dev', name: 'prod' },
			{ href: 'https://localhost:8090', name: 'dev' }
		];
		const actions: Action[] = [
			{
				handler: () => {
					const found_server = apis.find((api) => api.name === server_arg);
					if (found_server) {
						base_url.set(found_server.name);
					} else {
						debug.error('Something went wrong');
					}
					window.location.reload();
				},
				name: 'set'
			}
		];

		if (!actions.includes(action_arg)) {
			debug.error("Invalid server_arg value. Available: 'set'");

			return;
		}

		if (!apis.includes(server_arg)) {
			debug.error("Invalid server_arg value. Available 'dev', 'prod'");
			return;
		}

		const found_action = actions.find((action) => action.name === action_arg);
		const found_api = apis.find((api) => api.name === server_arg);

		if (found_action && found_api) {
			found_action.handler;
		}
	});

console_service
	.registerCommand('ls')
	.setDescription('Storage file listing')
	.setAction(async () => {
		let cleaned_records: string[] = [];
		const connection_string = `/storage/list/${Dashboard.state.current_directory}`;
		try {
			const response: ServerResponse<StorageRecord[]> = await api.api.get(connection_string);

			if (response.data === null || response.data.length === 0) {
				debug.format('No files here');
				return;
			}

			if (response.status === 200) {
				response.data.forEach((record) => {
					if (record.is_dir) {
						cleaned_records.push(bold(tail(record.name, 'text-blue-500 col-span-1 font-bold')));
					} else {
						cleaned_records.push(bold(tail(record.name, 'col-span-1 font-bold')));
					}
				});

				debug.format(tail(cleaned_records.join(''), 'flex flex-wrap gap-x-10'));
			}
		} catch (error) {
			if (axios.isAxiosError(error)) {
				debug.error(error.response?.data);
			}
			debug.silent(connection_string);
			debug.error('Error:');
			debug.error(error);
		}
	});

console_service
	.registerCommand('cd')
	.setDescription('Changes current directory')
	.addArgHandler((arg) => arg)
	.setAction(async (args) => {
		let dir: string = args[0] || '';

		if (dir === '..') {
			let splitted = Dashboard.state.current_directory.split('/');

			splitted.pop();

			Dashboard.state.current_directory = splitted.join('/');

			return;
		}

		if (Dashboard.state.current_directory.endsWith('/')) {
			Dashboard.state.current_directory = Dashboard.state.current_directory + `${dir}`;
			return;
		}

		Dashboard.state.current_directory = Dashboard.state.current_directory + `/${dir}`;
		debug.log('Current Directory: ', Dashboard.state.current_directory);
	});

console_service
	.registerCommand('weather')
	.setDescription('Showing current weather in your location')
	.addArgHandler<string>(
		(arg) => {
			return arg;
		},
		{ customName: 'location', required: false, type: 'string' }
	)
	.addFlagHandler<number>('-f', (flag) => {
		return flag;
	})
	.setAction(async (args, flags) => {
		const [l, f] = args;

		const location = l ? l : 'Skawina';
		const format = f ? f : '3';

		const controller = new AbortController();
		console_service.activeRequests.set('weather', controller);

		try {
			const response = await axios.get(`https://wttr.in/${location}?format=${format}`, {
				headers: { 'User-Agent': 'curl/7.64.1' },
				signal: controller.signal
			});

			const parser = new DOMParser();
			const doc = parser.parseFromString(response.data, 'text/html');
			const cleanText = doc.body.textContent || response.data;

			debug.system(cleanText.trim());
			terminal.toggle_terminal();
		} catch (error) {
			if (axios.isAxiosError(error)) {
				debug.error('Error fetching weather:', error.name);
				console_service.activeRequests.delete('weather');
			}
		}
	});

console_service
	.registerCommand('formatter')
	.setDescription('Formatter test')
	.addArgHandler((arg) => arg, { customName: 'run' })
	.addArgHandler((arg) => arg, { customName: 'formatterText' })
	.setAction((args) => {
		let text = args[1] || 'Formatter';
		let tailwind = args[2] || 'bg-orange-500/50 text-orange-200 border-1 p-1';

		debug.format(bold(italic(tail(text, tailwind))));
	});

console_service
	.registerCommand('warn')
	.setDescription('Warn terminal record test')
	.addArgHandler<string>(
		(arg) => {
			return arg;
		},
		{ customName: 'message', required: true, type: 'string' }
	)
	.setAction((args) => {
		debug.warn(args[0]);
	});

console_service
	.registerCommand('image')
	.setDescription('Image terminal record test')
	.addArgHandler<string>(
		(arg) => {
			return arg;
		},
		{ customName: 'imageSrc', required: true, type: 'string' }
	)
	.setAction((args) => {
		const imageSrc: 'cat' | string = args[0];
		type ImageKey = string | 'cat' | 'chill' | 'jewgun' | 'ryba';
		const images: Record<ImageKey, string> = {
			cat: 'https://api.klimson.dev/interface/bucket/random/nugget_cat.png',
			chill: 'https://api.klimson.dev/interface/bucket/random/klimson-chill.jpeg',
			jewgun: 'https://api.klimson.dev/interface/bucket/random/pixelgunicon.png',
			ryba: 'https://api.klimson.dev/interface/bucket/random/681693EF-D9EF-4007-8ACA-031949000CA3.gif'
		};

		if (!imageSrc) {
			debug.error(`No image provided at <$arg1>. Use 'image help' to view usage`);
			return;
		}
		if (imageSrc in images) {
			debug.image(images[imageSrc]);
			return;
		}

		debug.image(args[0]);
	});

console_service
	.registerCommand('error')
	.setDescription('Error terminal record test')
	.addArgHandler<string>(
		(arg) => {
			return arg;
		},
		{ customName: 'message', required: true, type: 'string' }
	)
	.setAction((args) => {
		debug.error(args[0]);
	});

console_service
	.registerCommand('history')
	.setDescription('Error terminal record test')

	.setAction(() => {
		debug.system(terminal.input_history);
	});

console_service
	.registerCommand('n')
	.setDescription('Notification test')

	.setAction((args) => {
		const [action, type] = args;

		if (action === 'test') {
			const mockRecords: Record<string, NotificationRecord> = {
				chlopak: notifications.notifications[0],
				//
				comment: {
					id: Math.random().toString(36),
					user: {
						username: 'Marek Łuszkiewicz',
						avatarUrl: 'https://api.klimson.dev/interface/bucket/random/robalini.png'
					},
					timestamp: '1m ago',
					actionType: 'comment',
					isRead: false,
					content: {
						headerHtml: `${bold('Marek')} mentioned you on ${bold('origin/drama')}`,
						body: span(
							` Nie dales mi zapisać synka do klubu ${tail('@Klimson', 'text-blue-500 bg-blue-800/50 rounded-lg p-0.5 px-1.5 font-black')}. Znajde cie i zajebie!`
						)
					}
				},
				lancuch: {
					id: Math.random().toString(36),
					user: {
						username: 'Robert Łańcuch',
						avatarUrl: 'https://api.klimson.dev/interface/bucket/pedal.png'
					},
					timestamp: '10s ago',
					actionType: 'comment',
					isRead: false,
					content: {
						headerHtml: `${bold('Robert')} mentioned you on ${bold('origin/fotowoltaika')}`,
						body: span(
							`${tail('@Klimson', 'text-blue-500 bg-blue-800/50 rounded-lg p-0.5 px-1.5 font-black')} Mieszkam pod mostem. Można zainstalować te panele fotowoltaiczne na asfalcie na dachu. Zadne pisma mi nie przychodza ani nic tutaj pod tym mostem`
						)
					}
				},
				ram: {
					id: Math.random().toString(36),
					user: {
						username: 'themanofsmegg99',
						avatarUrl:
							'https://api.klimson.dev/interface/bucket/random/1b5217313bab90409e39c48b9351118c.webp'
					},
					timestamp: '1m ago',
					actionType: 'comment',
					isRead: false,
					content: {
						headerHtml: `${bold('themanofsmegg99')} mentioned you on ${bold('vent')}`,
						body: span(
							`${tail('@everyone', 'text-blue-500 bg-blue-800/50 rounded-lg p-0.5 px-1.5 font-black')} sprzedałem 128GB ramu za 100 zl. Teraz ten ram kosztuje więcej. Jak to odkręcić prosze pomozcie 😭`
						)
					}
				},
				like: {
					id: Math.random().toString(36),
					user: {
						username: 'Waifu Otaku :3',
						avatarUrl: 'https://api.klimson.dev/interface/bucket/random/anime.png'
					},
					timestamp: '5m ago',
					actionType: 'like',
					isRead: false,
					content: { headerHtml: '<b>Ania</b> liked your <b>photo</b>', body: '' }
				},
				mention: {
					id: Math.random().toString(36),
					user: {
						username: 'Admin',
						avatarUrl: 'https://api.klimson.dev/interface/bucket/random/klimson-chill.jpeg'
					},
					timestamp: '1h ago',
					actionType: 'mention',
					isRead: false,
					content: { headerHtml: '<b>Admin</b> mentioned you', body: 'Sprawdź nowy regulamin.' }
				},
				component: {
					id: Math.random().toString(36),
					user: {
						username: 'Cwel',
						avatarUrl: 'https://api.klimson.dev/interface/bucket/random/unknown.png'
					},
					timestamp: '1m ago',
					actionType: 'mention',
					isRead: false,
					content: {
						headerHtml: `${bold('Cwel')} mentioned you on ${bold('origin/main')}`,
						body: span(
							`${tail('@Klimson', 'text-blue-500 bg-blue-800/50 rounded-lg p-0.5 px-1.5 font-black')} Wow, ten komponent NotificationPanel wygląda rewelacyjnie! Bardzo podoba mi się czystość designu i płynne przejścia przy dodawaniu nowych elementów - świetnie wpisuje się w nowoczesny stack`
						)
					}
				}
			};

			const record = mockRecords[type || 'comment'];
			if (record) {
				notifications.add(record);
			} else {
				debug.error('Nieznany typ powiadomienia. Użyj: comment, like, mention');
			}
		}

		if (action === 'clear') {
			notifications.clear();
		}
	});

console_service
	.registerCommand('user')
	.setDescription('Fetches specific user')
	.addArgHandler<string>(
		(arg) => {
			return arg;
		},
		{ customName: 'method', required: true, auto_complete_args: ['get'], type: 'string' }
	)
	.addArgHandler<string>(
		(arg) => {
			return arg;
		},
		{ customName: 'arg1', required: false, auto_complete_args: ['number'], type: 'string' }
	)
	.setAction(() => {
		debug.clear();
	})
	.setAction(async (args) => {
		const [method, arg1] = args;

		if (method === 'get') {
			try {
				const response = await api.user.GetOne(arg1);

				if (response.status === 200) {
					debug.log(response.data);
				}
			} catch (error) {
				debug.error(error);
			}
		}
	});

console_service
	.registerCommand('cmds')
	.setDescription('List of all available commands to use in dashboard terminal.')
	.addArgHandler<string>((arg) => arg, {
		customName: 'isDev',
		auto_complete_args: AutoComplete.bool,
		required: false
	})
	.setAction((args) => {
		const command_register = console_service.getCommandsRegister();
		debug.log(`\n`);
		debug.log(`(${command_register.length}) Commands: `);
		debug.log(`\n`);

		debug.pretty_format(terminal.console.dumpAvailableCommands());
		debug.pretty_format(
			`Type ${bold('command name')} and ${bold('help')} next to it to see command usage and help`
		);
	});

export const console_loading = writable<boolean>(false);

console_service
	.registerCommand('redis')
	.bindLoading(console_loading) // binding
	.setDescription('This command will allow you to operate on redis dashboard keys')
	.addArgHandler((arg) => arg) // get, set
	.addArgHandler((arg) => arg) // key
	.addFlagHandler<string>(
		'-t',
		(flag) => {
			return flag;
		},
		{ with_value: true }
	)
	.addFlagHandler<string>(
		'-f',
		(flag) => {
			return flag;
		},
		{ with_value: true }
	)
	.setAction(async (args, flags) => {
		debug.system('Flags: ', flags);
		const [action, key] = args;
		console_loading.set(true);
		try {
			const response = await api.redis.Get(key);

			debug.log(response.data);
		} catch (error) {
			debug.error(error);
		} finally {
			console_loading.set(false);
		}
	});

console_service
	.registerCommand('goto')
	.addArgHandler((arg) => arg)

	.setDescription('Allow you to quickly go to available dashboard enpoints.')
	.setAction((args) => {
		let uri = args[0];
		if (['/', '.'].some((e) => e === uri)) {
			goto('/dashboard' + uri);
			debug.system(`Moved to '${Dashboard.constants.findSidebarItem(uri)?.name}'`);

			return;
		}
		if (!uri) {
			goto('/dashboard');
			debug.system(`Moved to '${Dashboard.constants.findSidebarItem(uri)?.name}'`);

			return;
		}
		if (!Dashboard.constants.retrieveSidebarRoutes().includes(uri)) {
			debug.error(`Route with name '${uri}' does't exists!`);
			return;
		}

		goto('/dashboard' + '/' + uri);
		debug.system(`Moved to '${Dashboard.constants.findSidebarItem(uri)?.name}'`);
	});

export { console_service };
