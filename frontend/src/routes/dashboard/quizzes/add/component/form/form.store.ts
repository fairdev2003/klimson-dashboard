import { writable } from "svelte/store";

// writables
function imageFormFunction() {
    const { subscribe, set, update } = writable<boolean>(false);

    function open() {
        set(true);
    }

    function close() {
        set(false);
    }

    function getState(): boolean {
        let currentState: boolean = false;
        subscribe((value) => {
            currentState = value;
        });
        return currentState;
    }

    return {
        subscribe,
        getState,
        update,
        open,
        close,
    };
}

export const imageForm = imageFormFunction();