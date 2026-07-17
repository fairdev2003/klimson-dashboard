class Formatter {
	public red(text: string): string {
		return `<span style="color: red;">${text}</span>`;
	}
	public span(text: string): string {
		return `<span>${text}</span>`;
	}
	public orange(text: string): string {
		return `<span style="color: orange;">${text}</span>`;
	}
	public bold(text: string): string {
		return `<span> <b>${text}</b></span>`;
	}
	public italic(text: string): string {
		return `<span> <i>${text}</i></span>`;
	}

	/**
	 * Formats text with tailwind css class tree
	 * @example
	 * ```ts
	 * debug.format(
	 *      tail("Hello, World!", "bg-red-400")
	 * )
	 * ```
	 */
	public tail(text: string, color: string): string {
		return `<span class="${color}">${text}</span>`;
	}
}

const formatter = new Formatter();
export { formatter };
