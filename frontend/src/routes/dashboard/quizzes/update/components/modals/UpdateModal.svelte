<script lang="ts">
	import Button from '$lib/components/Button.svelte';
	import Checkbox from '$lib/components/dashboard/Checkbox.svelte';
	import CreateFormInput from '$lib/components/dashboard/CreateFormInput.svelte';
	import ImageUpload from '$lib/components/dashboard/ImageUpload.svelte';
	import StickyModal from '$lib/components/dashboard/StickyModal.svelte';
	import Heading from '$lib/components/dashboard/typography/Heading.svelte';
	import Modal from '$lib/components/Modal.svelte';
	import { selectedQuestionToEdit } from '$lib/dashboard/stores/store';
	import Icon from '@iconify/svelte';
	import { tick } from 'svelte';
	import { get } from 'svelte/store';

	type Props = {
		opened: boolean;
		onClose: () => void;
	};

	let file: File | undefined = $state();
	let newAnswerModalOpened: boolean = $state(false);

	let { opened = $bindable(), onClose }: Props = $props();

	let component = Heading;
	let updateAnswerContentModalOpened: boolean = $state(false);
	let addAnswerContentModalOpened: boolean = $state(false);

	let addAnswerContent: string = $state('');
	let answerToEdit: string = $state('');
	let el: HTMLDivElement;
	let editingAnswerIndex: number = $state(0);
</script>

<StickyModal
	bind:opened={updateAnswerContentModalOpened}
	onClose={() => (updateAnswerContentModalOpened = false)}
>
	<CreateFormInput focus={true} label="Tresc pytania" bind:value={answerToEdit} />
	<div class="flex justify-end mt-5">
		<Button
			onclick={() => {
				$selectedQuestionToEdit.answers[editingAnswerIndex].content = answerToEdit;
				updateAnswerContentModalOpened = false;
			}}
			size="small"
			theme="secondary">Zapisz</Button
		>
	</div>
</StickyModal>

<StickyModal
	bind:opened={addAnswerContentModalOpened}
	onClose={() => (addAnswerContentModalOpened = false)}
>
	<CreateFormInput label="Tresc pytania" bind:value={addAnswerContent} />
	<div class="flex justify-end mt-5">
		<Button size="small" theme="secondary">Zapisz</Button>
	</div>
</StickyModal>

<Modal
	title={`Edytujesz: ${get(selectedQuestionToEdit).content}`}
	onClose={() => {
		onClose();
	}}
	className="w-250"
	bind:opened
>
	<div class="flex flex-col gap-4">
		<ImageUpload bind:src={$selectedQuestionToEdit.image_url} bind:file />
		<CreateFormInput label="Treść pytania" bind:value={$selectedQuestionToEdit.content} />
		{#each $selectedQuestionToEdit.answers as answer, index}
			<Checkbox bind:checked={answer.is_correct} label={answer.content}>
				{#snippet edit_fragment()}
					<Icon
						onclick={() => {
							answerToEdit = answer.content;
							editingAnswerIndex = index;
							updateAnswerContentModalOpened = true;
						}}
						icon="material-symbols:edit"
						class="text-blue-500 cursor-pointer"
					/>
				{/snippet}

				{#snippet delete_fragment()}
					<Icon
						onclick={() => {
							$selectedQuestionToEdit.answers = $selectedQuestionToEdit.answers.filter(
								(e) => e.content !== answer.content
							);
						}}
						icon="material-symbols:delete"
						class="text-red-500"
					/>
				{/snippet}
			</Checkbox>
		{/each}
		<Button
			onclick={() => {
				addAnswerContentModalOpened = true;
			}}
			size="small"
			theme="secondary">Dodaj odpowiedz</Button
		>
	</div>
</Modal>

<Modal
	className="w-150"
	bind:opened={newAnswerModalOpened}
	onClose={() => {
		newAnswerModalOpened = false;
	}}
	title="Nowa odpowiedz"
>
	<CreateFormInput label="Tresc pytania" />
</Modal>
