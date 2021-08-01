import { writable } from 'svelte/store';

const state = { userName: null };

export const userStore = writable(state);
