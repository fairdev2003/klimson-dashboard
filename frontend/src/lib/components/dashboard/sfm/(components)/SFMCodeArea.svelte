<script lang="ts">
	import DiskImage from './assets/disk.png';
	import CodeLine from './CodeLine.svelte';
	import { Compiler } from './compiler/sfm.compiler';
	import { VSCThemeColor } from './styles/colors';

	import { onMount } from 'svelte';

	type SFMCodeAreaPage = 'barebone' | 'compiler' | 'data_test';
	let copiedStatus: boolean = $state(false);
	let sfm_program_lines: string[] = $state([]);
	let sfm_program: string = $state('');
	let code_area_page: SFMCodeAreaPage = $state('compiler');
	const handleCodeAreaPageChange = (page: SFMCodeAreaPage) => (code_area_page = page);

	type Props = {
		program_link?: string;
	};

	const { program_link = '' }: Props = $props();

	const SFMCompiler = new Compiler();

	onMount(async () => {
		sfm_program_lines = await SFMCompiler.GetSuperFactoryManagerLines(program_link);
		if (program_link !== '') {
			sfm_program = await SFMCompiler.FetchSuperFactoryManagerContent(program_link);
		}
	});
</script>

<div
	id={SFMCompiler.ExtractFileNameFromURL(program_link)}
	class="h-min-[300px] w-full border-2 border-neutral-800/60 bg-neutral-900/70"
>
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<div class="flex w-full justify-between bg-neutral-800/60 p-2 px-3">
		<div class="flex items-center gap-2">
			<img src={DiskImage} alt="disk" class="h-5 w-5" />
			<p class="program_name text-sm font-thin text-gray-400">
				<a href={program_link} target="_blank" class=" text-gray-400"
					>{SFMCompiler.ExtractFileNameFromURL(program_link)}</a
				>
				<!-- |
				<span class={VSCThemeColor.StringGreen}>{code_area_page}</span> -->
			</p>
		</div>
		<!-- svelte-ignore a11y_click_events_have_key_events -->
		{#if sfm_program !== ''}
			<div
				onclick={() => {
					// TODO: write actual code
					navigator.clipboard.writeText(sfm_program);
					copiedStatus = true;
					setTimeout(() => {
						copiedStatus = false;
					}, 4000);
				}}
				class="flex cursor-pointer items-center gap-1 text-gray-400 transition-colors select-none hover:text-white"
			>
				{#if copiedStatus}
					<p class="text-[13px] text-green-500">Copied!</p>
				{:else}
					<p class="text-[13px]">Copy</p>
				{/if}
			</div>
		{/if}
	</div>

	{#if program_link.includes('.sfm')}
		{#if sfm_program !== ''}
			<!-- Compiler testing -->
			{#if code_area_page === 'compiler'}
				<div class="p-2">
					{#each sfm_program_lines as line}
						<CodeLine>
							{@html SFMCompiler.renderCompiledCode(line)}
						</CodeLine>
					{/each}
				</div>
			{/if}
			{#if code_area_page === 'data_test'}
				<div class="p-2">
					<CodeLine className={VSCThemeColor.MainPurple}>
						{SFMCompiler.GetProgramName()}
					</CodeLine>
				</div>
			{/if}
		{:else}
			<div class="p-2">
				<CodeLine className={VSCThemeColor.VariableBlue}>NO PROGRAM!</CodeLine>
			</div>
		{/if}
	{:else}
		<div class="p-2">
			<CodeLine className={VSCThemeColor.VariableBlue}>NOT VALID PROGRAM!</CodeLine>
		</div>
	{/if}
	<!-- <div class="flex h-[30px] items-center gap-2 bg-slate-700 px-2"></div> -->
</div>

<style>
	a {
		color: oklch(70.7% 0.022 261.325);
		background-color: transparent;
	}

	@font-face {
		font-family: 'Consolas';
		src: url('./assets/consolas.ttf') format('truetype');
	}

	.program_name {
		font-family: 'Consolas', monospace;
	}
</style>
