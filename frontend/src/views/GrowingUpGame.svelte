<script>
  import { navigate } from 'svelte-routing';
  import { onMount } from 'svelte';
  import { userStore } from '../store/store';
  import { growingUpGameStore } from '../store/store';
  import InputWord from '../components/common/InputWord.svelte';
  import LabelWord from '../components/common/LabelWord.svelte';

  growingUpGameStore.update(state => ({
    ...state,
    hits: $growingUpGameStore.games[$growingUpGameStore.win].word.length,
  }));
  onMount(() => {
    if (!$userStore.userName) {
      navigate('/login', { replace: true });
    }
  });
</script>

<main>
  <div class="information-container">
    <h1>Adivina Adivinador</h1>
    {#each $growingUpGameStore.words as word}
      <LabelWord {word} />
    {/each}
  </div>
  <div>
    <div class="game-container">
      <p>{$growingUpGameStore.games[$growingUpGameStore.win].sentence}</p>
      <div>
        {#each $growingUpGameStore.games[$growingUpGameStore.win].word as letter, index}
          <InputWord {letter} {index} />
        {/each}
      </div>
    </div>
  </div>
</main>

<style>
  main {
    height: 100%;
    width: 100%;
    background-color: var(--pink);
    display: flex;
    flex-direction: row-reverse;
  }

  main > div:nth-child(2) {
    width: 70%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .information-container {
    height: 100%;
    width: 30%;
    text-align: center;
    padding: 2rem 4rem 4rem 0;
    display: flex;
    flex-direction: column;
    justify-content: space-around;
    align-items: center;
  }
  .game-container {
    height: 80%;
    width: 90%;
    background-color: var(--white);
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: space-evenly;
  }
  .information-container > h1 {
    color: var(--purple);
    font-size: 3.5rem;
  }

  .game-container > p {
    font-size: 2rem;
    color: var(--purple);
    text-align: center;
    padding: 3rem;
  }
</style>
