<script lang="ts">
	import Icon from '@iconify/svelte';
	import ResiableList from '../quizzes/add/component/ResiableList.svelte';
	import TodoList from './components/TodoList.svelte';

	let saveLaterOpen: boolean = $state(false);
	let contributorsFormOpen: boolean = $state(false);
</script>

<div class="mt-10 max-w-3xl w-full mx-auto flex flex-col gap-5">
	<ResiableList bind:isOpen={saveLaterOpen} section_name="Zapisane na później">
		{#snippet nav_icon()}
			<Icon icon="ri:save-3-line" width="30" height="30" />
		{/snippet}
		<div class="h-100"></div>
	</ResiableList>
	<ResiableList bind:isOpen={contributorsFormOpen} section_name="Formularze kontrybutorów">
		{#snippet nav_icon()}
			<Icon icon="ph:note-pencil-bold" width="30" height="30" />
		{/snippet}
		<div class="h-100"></div>
	</ResiableList>
</div>
