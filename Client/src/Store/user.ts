import { writable } from 'svelte/store';

const user = writable<string>('');

export default user;
