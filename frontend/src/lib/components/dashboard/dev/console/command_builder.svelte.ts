type CommandLoader = boolean;
type CommandAction = (args: any[]) => void;
type ArgHandler = { handler: (arg: string) => any; config?: CommandConfig };
type LoadingState = {
	set: (value: boolean) => void;
};
type CommandConfig =
	| {
			customName: string;
			strict?: boolean;
			auto_complete_args: string[];
	  }
	| undefined;

const AutoComplete: Record<string, any[]> = {};

export class CommandBuilder {
	public argHandlers: ArgHandler[] = [];
	private action?: CommandAction;
	public loader: CommandLoader = false;
	public description: string | undefined;
	private loadingState?: LoadingState;

	constructor(public name: string) {}

	addArgHandler(handler: (arg: string) => any, config?: CommandConfig): this {
		this.argHandlers.push({ handler, config });
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
	execute(rawArgs: string[]) {
		const processedArgs = rawArgs.map((arg, index) => {
			const argHandler = this.argHandlers[index];

			if (argHandler && typeof argHandler.handler === 'function') {
				return argHandler.handler(arg);
			}

			return arg;
		});

		if (this.loadingState) this.loadingState.set(true);

		try {
			if (this.action) this.action(processedArgs);
		} finally {
			if (this.loadingState) this.loadingState.set(false);
		}
	}
}
