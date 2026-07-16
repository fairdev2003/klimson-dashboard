import { debug } from '$lib/dashboard/stores/debug';

type CommandLoader = boolean;
type CommandAction = (args: any[], flags: CommandFlag<any>[]) => void;
type ArgHandler = { handler: (arg: any) => any; config?: CommandConfig };
type LoadingState = {
	set: (value: boolean) => void;
};
type CommandConfig<T = any> =
	| {
			customName: string;
			required?: boolean;
			type?: T;
			auto_complete_args?: string[];
	  }
	| undefined;

type CommandFlagConfig<T> = {};

type CommandFlag<T> = {
	flag_name: string;
	flag_value?: any;
	with_value: boolean;
	handler: (flag: T) => void;
};

export type AutoCompleteKey = 'bool';

export const AutoComplete: Record<AutoCompleteKey, any[]> = {
	bool: ['true', 'false']
};

export class CommandBuilder {
	public argHandlers: ArgHandler[] = [];
	public flags: CommandFlag<any>[] = [];
	private action?: CommandAction;
	public loader: CommandLoader = false;
	public description: string | undefined;
	private loadingState?: LoadingState;

	constructor(public name: string) {}

	addArgHandler<T>(handler: (arg: T) => any, config?: CommandConfig<T>): this {
		this.argHandlers.push({ handler, config });
		return this;
	}

	addFlagHandler<T>(
		flag: string,
		handler: (flag: T) => any,
		config?: { with_value: boolean }
	): this {
		if (!config) {
			this.flags.push({ flag_name: flag, handler, with_value: false });
			return this;
		}

		this.flags.push({ flag_name: flag, handler, with_value: config.with_value });
		return this;
	}

	setAction(action: CommandAction): this {
		this.action = action;

		return this;
	}

	setDescription(desc: string): this {
		this.description = desc;

		return this;
	}

	bindLoading(store: { set: (v: boolean) => void }): this {
		this.loadingState = store;
		return this;
	}

	// Wykonanie komendy
	execute(rawArgs: string[], rawInput: string) {
		let args = rawArgs;
		let flags: { flag_key: string; flag_value: any }[] = [];

		// flag handling
		rawArgs.forEach((arg, i) => {
			this.flags.forEach((flag) => {
				if (flag.flag_name === arg) {
					const flag_key = flag.flag_name;
					if (flag.with_value) {
						const key_index = i;
						const value_index = i + 1;
						const flag_value = rawArgs[value_index];

						args = rawArgs.filter((raw_arg, raw_arg_index) => {
							const condition = raw_arg_index === key_index || raw_arg_index === value_index;
							if (!condition) {
								return raw_arg;
							}
						});
						flags.push({ flag_key, flag_value });
						this.flags = [
							...this.flags,
							{
								flag_name: flag_key,
								flag_value: flag_value,
								with_value: flag.with_value,
								handler: flag.handler
							}
						];
						return;
					}
					flags.push({ flag_key, flag_value: undefined });
				}
			});
		});

		const processedArgs = args.map((arg, index) => {
			const argHandler = this.argHandlers[index];

			if (argHandler && typeof argHandler.handler === 'function') {
				return argHandler.handler(arg);
			}

			if (arg) return arg;
		});

		if (this.loadingState) this.loadingState.set(true);

		try {
			if (this.action) this.action(processedArgs, this.flags);
		} finally {
			if (this.loadingState) this.loadingState.set(false);
		}
	}
}
