import { writable } from 'svelte/store';

export const loginStore = writable({
  isFirstLogin: false,
  users: [],
});

export const subjectStore = writable({
  subjects: [],
  subject: {},
  games: [],
});

export const themeStore = writable({
  scoreImage: null,
});

export const userStore = writable({
  userName: null,
  avatar: {},
  score: 0,
});

export const avatarStore = writable({
  avatars: [],
  avatar: null,
});

export const memoryGameStore = writable({
  number: 0,
  numbers: [],
  flip: false,
});

export const growingUpGameStore = writable({
  games: [
    {
      sentence:
        'Blanca por dentro, verde por fuera. Si quieres te lo digo, espera.',
      word: 'pera',
    },
    {
      sentence:
        'Llevo a la familia y a su equipaje. Paso todas las noches en el garaje. ¿Quién soy?',
      word: 'carro',
    },
    {
      sentence:
        'Con una manguera, casco y escalera, apago el fuego de la hoguera. ¿Quién soy?',
      word: 'bombero',
    },
  ],
  words: [null, null, null],
  letters: [],
  hits: 0,
  win: 0,
});

export const basicOperationsStore = writable({
  items: [
    { id: 1, name: 7, color: 'pink' },
    { id: 2, name: 10, color: 'yellow' },
    { id: 3, name: 4, color: 'orange' },
    { id: 4, name: 'X', color: 'purple' },
  ],
  dropped: [null, null, null, null],
});

export const syllabeStore = writable({
  items: [
    { id: 1, name: 'CA', color: 'cyan' },
    { id: 2, name: 'GO', color: 'orange' },
    { id: 3, name: 'JU', color: 'orange' },
    { id: 4, name: 'JA', color: 'yellow' },
    { id: 5, name: 'SA', color: 'cyan' },
    { id: 6, name: 'CA', color: 'yellow' },
  ],
  syllabes: [null, null, null, null, null, null],
});

export const livingBeignsStore = writable({
  items: [
    { id: 1, name: 'koala' },
    { id: 2, name: 'butterfly' },
    { id: 3, name: 'chainsaw' },
    { id: 4, name: 'fence' },
    { id: 5, name: 'bee' },
    { id: 6, name: 'axe' },
  ],
  beigns: [],
  inerts: [],
});
