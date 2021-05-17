import { writable } from 'svelte/store';

const authenticated = writable<boolean>(false);

export default authenticated;
