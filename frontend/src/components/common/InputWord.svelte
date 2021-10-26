<script>
  import { onMount } from 'svelte';
  import { navigate } from 'svelte-routing';
  import { growingUpGameStore } from '../../store/store';
  export let letter;
  export let index;
  $: disabled = false;
  $: value = '';
  let input;
  const verify = () => {
    if (value.toUpperCase() === letter.toUpperCase()) {
      let letters = $growingUpGameStore.letters;
      letters.push(letter);
      growingUpGameStore.update(state => ({
        ...state,
        letters,
        hits: state.hits + 1,
      }));
      disabled = true;
    } else {
      setTimeout(() => {
        growingUpGameStore.update(state => ({
          ...state,
          hits: state.hits - 1,
        }));
        value = '';
      }, 1000);
    }
  };
  $: {
    if (
      $growingUpGameStore.letters.join('').length ===
      $growingUpGameStore.games[$growingUpGameStore.win].word.length
    ) {
      if (
        value ===
        $growingUpGameStore.letters[$growingUpGameStore.letters.length - 1]
      ) {
        console.log(letter);
        let words = $growingUpGameStore.words;
        words.push($growingUpGameStore.letters.join(''));
        if ($growingUpGameStore.games.length - 1 !== $growingUpGameStore.win) {
          growingUpGameStore.update(state => ({
            ...state,
            win: $growingUpGameStore.win + 1,
            words,
            letters: [],
          }));
        } else {
          growingUpGameStore.update(state => ({
            ...state,
            win: 0,
            words: [],
            letters: [],
            hits: 0,
          }));
          navigate('/congratulations', { replace: true });
        }
      }
      value = '';
      disabled = false;
    }
  }
  onMount(() => {
    if (index == 0) {
      input.focus();
    }
  });
</script>

<input
  type="text"
  bind:this={input}
  bind:value
  maxlength="1"
  on:input={verify}
  {disabled}
/>

<style>
  input[type='text'] {
    background-color: transparent;
    border: none;
    border-bottom: 0.25rem solid var(--black);
    margin: 0 1rem;
    width: 4rem;
    text-align: center;
    border-radius: 15%;
    font-size: 2rem;
    padding: 0 1rem;
  }
</style>
