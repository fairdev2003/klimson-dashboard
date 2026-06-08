import { persistedWritable } from '$lib/dashboard/stores/persist';

export type CodeEditorTheme = 'dracula' | 'classic';

export const codeEditorTheme = persistedWritable<CodeEditorTheme>('code_theme', 'classic');
