<script>
  import { syllabeStore, userStore } from '../store/store';
  import { flip } from 'svelte/animate';
  import { crossfade } from 'svelte/transition';
  import { quintOut, elasticOut } from 'svelte/easing';
  import { draggable } from '../js/dragdrop';
  import { navigate } from 'svelte-routing';
  import { onMount } from 'svelte';

  const putInSyllabe = (item, index) => {
    console.log(index);
    console.log(item);
    let items = $syllabeStore.items;
    const syllabes = $syllabeStore.syllabes;
    const oldItem = $syllabeStore.syllabes[index];
    const oldShelfIndex = $syllabeStore.syllabes.indexOf(item);

    if (items.indexOf(item) !== -1) {
      items = items.filter(i => i != item);
    }
    if (oldShelfIndex !== -1) {
      syllabes[oldShelfIndex] = oldItem;
    } else if (oldItem) {
      items.push(oldItem);
    }
    syllabes[index] = item;
    syllabeStore.update(state => ({
      ...state,
      items,
      syllabes,
    }));
    if (syllabes.indexOf(null) === -1) {
      console.log(syllabes);
      if (
        (syllabes[0].id === 1 || syllabes[0].id === 6) &&
        syllabes[1].id === 5 &&
        (syllabes[2].id === 1 || syllabes[2].id === 6) &&
        syllabes[3].id === 4 &&
        syllabes[4].id === 3 &&
        syllabes[5].id === 2
      ) {
        navigate('/congratulations', { replace: true });
      }
      syllabeStore.update(state => ({
        ...state,
        items: syllabes,
        syllabes: [null, null, null, null, null, null],
      }));
    }
  };

  const putInItems = item => {
    const items = $syllabeStore.items;
    const syllabes = $syllabeStore.syllabes;
    if (items.indexOf(item) !== -1) items.splice(items.indexOf(item), 1);
    if (syllabes.indexOf(item) !== -1) syllabes[syllabes.indexOf(item)] = null;
    items.push(item);
    syllabeStore.update(state => ({
      ...state,
      items,
      syllabes,
    }));
  };

  const [send, receive] = crossfade({
    duration: d => 600,
    easing: elasticOut,
    fallback(node, params) {
      const style = getComputedStyle(node);
      const transform = style.transform === 'none' ? '' : style.transform;

      return {
        duration: 600,
        easing: quintOut,
        css: t => `
        transform: ${transform} scale(${t});
        opacity: ${t}
      `,
      };
    },
  });
  onMount(() => {
    if (!$userStore.userName) {
      navigate('/login', { replace: true });
    }
  });
</script>

<main>
  <div class="information-container">
    <h1>ARRASTRA Y ARMA LA PALABRA</h1>
    <div class="cart" on:dropped={e => putInItems(e.detail)}>
      {#each $syllabeStore.items as item, index (item.id)}
        <div
          class="box-container item"
          style="background-color: var(--{item.color});"
          animate:flip
          use:draggable={{ data: item, targets: ['.slot', '.slot .item'] }}
          in:receive={item.id}
          out:send={item.id}
        >
          {item.name}
        </div>
      {/each}
    </div>
  </div>
  <div>
    <div class="game-container ">
      <div class="game-images">
        <img src="images/casa.svg" alt="casa" />
        <img src="images/caja.svg" alt="caja" />
        <img src="images/jugo.svg" alt="jugo" />
      </div>
      <div class="game-syllabe">
        {#each $syllabeStore.syllabes as item, index}
          <div
            class="box-slot slot box-{index}"
            on:dropped={e => putInSyllabe(e.detail, index)}
          >
            {#if item}
              <div
                class="item"
                use:draggable={{
                  data: item,
                  targets: ['.cart', '.slot', '.slot .item'],
                }}
                in:receive={item.id}
                out:send={item.id}
                on:dropped={e => putInSyllabe(e.detail)}
                style="background-color: var(--{item.color}); color: var(--white)"
              >
                {item.name}
              </div>
            {/if}
          </div>
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
    display: flex;
    padding: 5rem 0;
    flex-direction: column;
    justify-content: space-between;
    align-items: center;
    text-align: center;
  }

  .information-container > h1 {
    color: var(--purple);
    font-size: 3rem;
  }

  .information-container > div {
    width: 100%;
    height: 80%;
    display: grid;
    grid-template: 1fr 1fr 1fr / 1fr 1fr;
    justify-items: center;
    align-items: center;
    color: var(--white);
  }

  .game-container {
    height: 80%;
    width: 90%;
    background-color: var(--white);
    display: flex;
    justify-content: center;
    align-items: center;
  }
  .game-images {
    display: flex;
    flex-direction: column;
    align-items: center;
    height: 60%;
    justify-content: space-around;
  }
  .game-syllabe {
    display: grid;
    grid-template: 1fr 1fr 1fr/ 1fr 1fr;
    height: 60%;
    width: 30%;
    justify-items: center;
    align-items: center;
    margin-left: 3rem;
  }
  .box-slot {
    width: 5rem;
    height: 5rem;
    border: 4px solid var(--shadow);
  }
  .item {
    width: 5rem;
    height: 5rem;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 3rem;
    text-shadow: 0 0 0.1rem var(--shadow);
  }
  .slot {
    position: relative;
  }
  .cart {
    position: relative;
  }
  .item {
    position: relative;
  }
  .slot .item {
    position: absolute;
  }
  .dragged {
    pointer-events: none;
    z-index: 100;
  }
</style>
