<script>
  import { navigate } from 'svelte-routing';
  import { onMount } from 'svelte';
  import { userStore } from '../store/store';
  import { memoryGameStore } from '../store/store';
  import Cart from '../components/Card.svelte';

  $: found = 0;
  $: missing = 6;
  $: onLoad = false;
  $: temporaryCard = [];
  $: found = $memoryGameStore.numbers.length;
  $: found ? (missing -= 1) : '';
  let cards = [1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5, 6];
  onMount(() => {
    if (!$userStore.userName) {
      navigate('/login', { replace: true });
    }
    temporaryCard = cards.sort(() => Math.random() - 0.5);
    onLoad = true;
  });
</script>

<main>
  <div class="information-container">
    <h1>PON A PRUEBA TU MEMORIA</h1>
    <div>
      <p>
        Encontrados: <span class="number">{found}</span>
      </p>
      <p>
        Faltantes: <span class="number">{missing}</span>
      </p>
    </div>
  </div>
  <div>
    <div class="game-container">
      {#if onLoad}
        {#each temporaryCard as value}
          <Cart {value} />
        {/each}
      {/if}
    </div>
  </div>
</main>

<style>
  main {
    height: 100%;
    width: 100%;
    background-color: var(--cyan);
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
    display: flex;
    flex-direction: column;
    justify-content: space-around;
    align-items: center;
    text-align: center;
  }

  .information-container > h1 {
    color: var(--purple);
    font-size: 3rem;
  }

  .information-container > div {
    width: 100%;
    height: 30%;
    display: flex;
    flex-direction: column;
    justify-content: space-around;
    align-items: center;
    font-size: 3rem;
  }

  .game-container {
    height: 80%;
    width: 90%;
    background-color: var(--white);
    display: grid;
    grid-template: 1fr 1fr 1fr / 1fr 1fr 1fr 1fr;
    align-items: center;
    justify-items: center;
  }

  p {
    color: var(--white);
  }

  .number {
    color: var(--orange);
  }
</style>
