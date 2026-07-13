import { api } from '$lib/api/api';
import type { BackendResponse, ServerResponse } from '$lib/api/types';
import { debug } from '$lib/dashboard/stores/debug';
import { writable } from 'svelte/store';
import { CommandBuilder } from './command_builder.svelte';
import { goto } from '$app/navigation';

class ConsoleService {
	private commands: Map<string, CommandBuilder> = new Map();
	private unknown_command_handler: (name: string) => void = $state(() => {});

	public loading = $state(false);

	registerCommand(name: string): CommandBuilder {
		const cmd = new CommandBuilder(name);
		this.commands.set(name, cmd);
		return cmd;
	}

	constructor() {
		debug.system('Console Service is initialized successfully.');
	}

	public onUnknownCommand(handler: (name: string) => void) {
		this.unknown_command_handler = handler;
	}

	public getCommandsRegister(): CommandBuilder[] {
		return Array.from(this.commands.values());
	}

	public run(input: string) {
		const parts = input.trim().split(/\s+/);
		const name = parts[0];
		const args = parts.slice(1);
		const command = this.commands.get(name);

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
			command.execute(args);
		} else {
			if (this.unknown_command_handler != undefined) {
				this.unknown_command_handler(name);
			}

			console.warn(`Unknown command: ${name}`);
		}
	}
}

const console_service = new ConsoleService();
console_service.onUnknownCommand((name) => {
	debug.log(`Command with name '${name}' does not exist!`);
	debug.log(`Type 'cmds' to view available commands.`);
});

console_service
	.registerCommand('clear')
	.setDescription('Clears the terminal')
	.setAction(() => {
		debug.clear();
	});

console_service
	.registerCommand('cmds')
	.setDescription('List of all available commands to use in dashboard terminal.')
	.addArgHandler((arg) => arg, {
		customName: 'isDev',
		auto_complete_args: ['true', 'false'],
		strict: false
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
	.setAction(async (args) => {
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
	});

export { console_service };
