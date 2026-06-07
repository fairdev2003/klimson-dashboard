import { writable } from 'svelte/store';
import gsap from 'gsap';
import { tick } from 'svelte';

type FileExplorer = {
	open: boolean;
	previousPath?: string;
	startingPath?: string;
	acceptedFileTypes?: string[]; // .md, .txt
};

function createExplorerStore() {
	const store = writable<FileExplorer>({
		open: false,
		acceptedFileTypes: [],
		startingPath: '/'
	});

	async function showExplorer(explorer: FileExplorer) {
		gsap.fromTo(
			'.modal',
			{ opacity: 0, duration: 2, scale: 0.8 },
			{ scale: 1, opacity: 1, ease: 'power2.out' }
		);

		store.update((state) => explorer);
	}

	function exitExplorer() {
		gsap.to('.modal', {
			opacity: 0,
			duration: 0.5,
			scale: 0.8,
			onComplete: () => {
				store.update((state) => ({
					...state,
					open: false
				}));
			}
		});
	}

	function backToPreviousPath() {
		store.update((state) => {
			const prevPath = state.previousPath;
			console.log(prevPath);
			return {
				startingPath: prevPath,
				...state
			};
		});
	}

	function setPrevPath(path: string) {
		store.update((state) => {
			console.log(state);
			return {
				...state,
				previousPath: path
			};
		});
	}

	function gotoPath(path: string) {
		store.update((state) => {
			return {
				...state,
				startingPath: path
			};
		});
	}

	return { setPrevPath, gotoPath, showExplorer, exitExplorer, backToPreviousPath, ...store };
}

const explorer = createExplorerStore();

export { explorer };
