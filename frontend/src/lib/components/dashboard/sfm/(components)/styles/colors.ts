import { Keywords, type KeywordType } from '../compiler/keywords';

export enum VSCThemeColor {
	Background = 'bg-[#1e1e1e]', // #1e1e1e — tło edytora
	Foreground = 'text-[#d4d4d4]', // #d4d4d4 — domyślny tekst
	HighlightBlue = 'text-[#569cd6]', // #569cd6 — słowa kluczowe
	StringGreen = 'text-[#6a9955]', // #6a9955 — stringi
	TypePink = 'text-[#c586c0]', // #c586c0 — typy
	VariableOrange = 'text-[#ce9178]', // #ce9178 — zmienne/stringi
	NumberYellow = 'text-[#dcdcaa]', // #dcdcaa — liczby
	CommentGray = 'text-[#6a737d]', // #6a737d — komentarze
	FunctionCyan = 'text-[#4ec9b0]', // #4ec9b0 — funkcje
	VariableBlue = 'text-[#9CDCFE]', // #9CDCFE - variable blue
	MainPurple = 'text-[#c586c0]', // #c586c0 - pink main
	RedRedstone = 'text-red-400'
}

export type VSCThemeColorKey = keyof typeof VSCThemeColor;

export const KeywordColor = {
	[Keywords.EVERY]: VSCThemeColor.MainPurple,
	[Keywords.NUMBER]: VSCThemeColor.NumberYellow,
	[Keywords.TIME]: VSCThemeColor.VariableOrange,
	[Keywords.COMMENT]: VSCThemeColor.StringGreen,
	[Keywords.VARIABLE]: VSCThemeColor.VariableBlue,
	[Keywords.DIRECTION]: VSCThemeColor.NumberYellow,
	[Keywords.OPERATOR]: VSCThemeColor.NumberYellow,
	[Keywords.STRING]: VSCThemeColor.StringGreen,
	[Keywords.REDSTONE]: VSCThemeColor.RedRedstone,
	[Keywords.TIME_TYPE]: VSCThemeColor.VariableOrange
};
