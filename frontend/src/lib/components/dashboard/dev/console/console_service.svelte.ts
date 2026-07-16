import { api } from '$lib/api/api';
import type { BackendResponse, ServerResponse } from '$lib/api/types';
import { debug } from '$lib/dashboard/stores/debug';
import { writable } from 'svelte/store';
import { AutoComplete, CommandBuilder } from './command_builder.svelte';
import { goto } from '$app/navigation';

class ConsoleService {
	private commands: Map<string, CommandBuilder> = new Map();
	private unknown_command_handler: (user_input: string, name: string) => void = $state(() => {});
	private on_command_handler: (
		command: CommandBuilder | undefined,
		input: string | undefined
	) => void | undefined = () => {};

	public loading = $state(false);

	registerCommand(name: string): CommandBuilder {
		const cmd = new CommandBuilder(name);
		this.commands.set(name, cmd);
		return cmd;
	}

	constructor() {
		debug.system('Console Service is initialized successfully.');
	}

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
}

const console_service = new ConsoleService();
console_service.onUnknownCommand((input, name) => {
	debug.console(input);

	debug.system(`Command with name '${name}' does not exist!`);
	debug.system(`Type 'cmds' to view available commands.`);
});

console_service.onCommand((command, input) => {
	if (!command) {
		return;
	}
	debug.console(input);
});

console_service
	.registerCommand('clear')
	.setDescription('Clears the terminal')
	.setAction(() => {
		debug.clear();
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
	.setAction(() => {
		goto('/login');
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
		const [dev] = args;
		const command_register = console_service.getCommandsRegister();
		debug.log(`\n`);
		debug.log(`(${command_register.length}) Commands: `);
		debug.log(`\n`);
		command_register.forEach((command) => {
			const desc = command.description ? ` - ${command.description}` : '';

			debug.log(`${command.name}${desc}`);
		});

		if (Boolean(dev)) {
			command_register.forEach((command) => {
				debug.log(command.argHandlers);
			});
		}
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
		if (!(args[0] as string).includes('/')) {
			uri = '/' + args[0];
		}
		goto('/dashboard' + uri);
		debug.system(`Moved to '${uri}'`);
	});

export { console_service };
