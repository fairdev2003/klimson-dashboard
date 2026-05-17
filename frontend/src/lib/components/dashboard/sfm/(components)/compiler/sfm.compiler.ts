import axios, { type AxiosResponse } from 'axios';
import { KeywordColor, VSCThemeColor } from '../styles/colors';

class Compiler {
	public sfm_program_content: string = '';
	public sfm_program_lines: string[] = [];

	public ExtractFileNameFromURL(url: string): string {
		if (url === '') {
			return 'No program loaded';
		}
		const parts = url.split('/');
		return parts[parts.length - 1];
	}

	public GetProgramName(): string {
		const match = this.sfm_program_content.match(/^NAME\s+"(.+?)"/im);
		return match ? match[1] : '';
	}

	public ExtractVariablesFromProgram(): string[] {
		const ignoreList = [
			'BOTTOM',
			'TOP',
			'SIDE',
			'LEFT',
			'RIGHT',
			'FRONT',
			'BACK',
			'NORTH',
			'SOUTH',
			'EAST',
			'WEST'
		];

		const matches = [...this.sfm_program_content.matchAll(/\b(?:FROM|TO)\b\s+([^\n]+)/gi)];

		const allVars = new Set<string>();

		for (const match of matches) {
			if (!match[1]) continue;

			const vars = match[1]
				.split(/[\s,]+/)
				.map((v) => v.trim())
				.filter((v) => v !== '' && !ignoreList.includes(v.toUpperCase()));

			vars.forEach((v) => allVars.add(v));
		}

		return [...allVars];
	}

	public async FetchSuperFactoryManagerContent(path: string): Promise<string> {
		const response: AxiosResponse<string> = await axios.get(path, { responseType: 'text' });
		this.sfm_program_content = response.data;
		return response.data;
	}

	public async GetSuperFactoryManagerLines(path: string): Promise<string[]> {
		const content = await this.FetchSuperFactoryManagerContent(path);
		const lines = content.split('\n');
		this.sfm_program_lines = lines;
		return lines;
	}

	public renderCompiledCode(code: string): string {
		code = code.replace(code, `<span class="${KeywordColor.variable}">${code}</span>`);

		const name_string = this.GetProgramName();
		if (code.match(/\bNAME\b/i) && name_string) {
			code = code.replace(new RegExp(`"${name_string}"`, 'i'), '__NAME_STRING__');
		}

		const comments: string[] = [];
		code = code.replace(/--.*/g, (match) => {
			const index = comments.length;
			comments.push(match);
			return `__COMMENT_PLACEHOLDER_${index}__`;
		});

		code = code
			.replace(
				/\b(NAME|SLOTS|EVERY|SOME|OVERALL|ONE|LONE|INPUT|OUTPUT|FROM|TO|END|DO|FORGET|ROUND ROBIN BY|EACH|SIDE|HAS|IF|THEN|WITH|EXCEPT|RETAIN|PULSE|ELSE)\b/gi,

				(match) => `<span class="relative group ${KeywordColor.every}">${match}</span>`
			)
			.replace(/\b(\d+)\b/g, (match) => `<span class="${KeywordColor.number}">${match}</span>`)
			.replace(/::/g, `<span class="text-cyan-300">::</span>`)
			.replace(/,/g, `<span class="text-slate-400">,</span>`)
			.replace(
				/\b(TICKS|SECONDS|MINUTES|HOURS)\b/gi,
				(match) => `<span class="${KeywordColor.time}">${match}</span>`
			)
			.replace(
				/\b(TOP|BOTTOM|NORTH|SOUTH|EAST|WEST|FRONT|BACK|LEFT|RIGHT)\b/gi,
				(match) => `<span class="${KeywordColor.direction}">${match}</span>`
			)
			.replace(
				/\b(TAG|FALSE|TRUE)\b/gi,
				(match) => `<span class="${KeywordColor.variable}">${match}</span>`
			)
			.replace(
				/\b(gt|lt|ge|le|eq|ne|and|or)\b/gi,
				(match) => `<span class="${KeywordColor.operator}">${match}</span>`
			)
			.replace(
				/\bREDSTONE\b/gi,
				(match) => `<span class="${KeywordColor.redstone}">${match}</span>`
			)
			.replace(/[()]/g, (match) => `<span class="text-white">${match}</span>`);

		code = code.replace(/__COMMENT_PLACEHOLDER_(\d+)__/g, (_, index) => {
			const comment = comments[parseInt(index, 10)];
			return `<span class="${KeywordColor.comment}">${comment}</span>`;
		});

		code = code.replace(
			/__NAME_STRING__/g,
			`<span class="${KeywordColor.time}">"${name_string}"</span>`
		);

		return code;
	}
}

export { Compiler };
