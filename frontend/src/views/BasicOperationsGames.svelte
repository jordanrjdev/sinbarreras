<script>
  import { basicOperationsStore, userStore } from '../store/store';
  import { flip } from 'svelte/animate';
  import { crossfade } from 'svelte/transition';
  import { quintOut, elasticOut } from 'svelte/easing';
  import { draggable } from '../js/dragdrop';
  import { navigate } from 'svelte-routing';
  import { onMount } from 'svelte';

  const putInShelf = (item, index) => {
    let items = $basicOperationsStore.items;
    const dropped = $basicOperationsStore.dropped;
    const oldItem = $basicOperationsStore.dropped[index];
    const oldShelfIndex = $basicOperationsStore.dropped.indexOf(item);

    if (items.indexOf(item) !== -1) {
      items = items.filter(i => i != item);
    }
    if (oldShelfIndex !== -1) {
      dropped[oldShelfIndex] = oldItem;
    } else if (oldItem) {
      items.push(oldItem);
    }
    dropped[index] = item;
    basicOperationsStore.update(state => ({
      ...state,
      items,
      dropped,
    }));

    if (dropped.indexOf(null) === -1) {
      console.log(dropped);
      if (
        dropped[0].id === 3 &&
        dropped[1].id === 1 &&
        dropped[2].id === 4 &&
        dropped[3].id === 2
      ) {
        navigate('/congratulations', { replace: true });
      }
      basicOperationsStore.update(state => ({
        ...state,
        items: dropped,
        dropped: [null, null, null, null],
      }));
    }
  };

  const putInInformation = item => {
    const items = $basicOperationsStore.items;
    const dropped = $basicOperationsStore.dropped;
    if (items.indexOf(item) !== -1) items.splice(items.indexOf(item), 1);
    if (dropped.indexOf(item) !== -1) dropped[dropped.indexOf(item)] = null;
    items.push(item);
    basicOperationsStore.update(state => ({
      ...state,
      items,
      dropped,
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
    <h1>ARRASTRA Y RESUELVE LAS OPERACIONES</h1>
    <div class="cart" on:dropped={e => putInInformation(e.detail)}>
      {#each $basicOperationsStore.items as item, index (item.id)}
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
      {#each $basicOperationsStore.dropped as item, index}
        <div
          class="box-slot slot box-{index}"
          on:dropped={e => putInShelf(e.detail, index)}
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
              on:dropped={e => putInShelf(e.detail)}
              style="background-color: var(--{item.color}); color: var(--white)"
            >
              {item.name}
            </div>
          {/if}
        </div>
      {/each}
      <div class="item column1"><span>2</span></div>
      <div class="item column2"><span>+</span></div>
      <div class="item column3"><span>2</span></div>
      <div class="item column4"><span>=</span></div>
      <div class="item column1"><span>9</span></div>
      <div class="item column2"><span>-</span></div>
      <div class="item column4"><span>=</span></div>
      <div class="item column5"><span>2</span></div>
      <div class="item column1"><span>4</span></div>
      <div class="item column3"><span>2</span></div>
      <div class="item column4"><span>=</span></div>
      <div class="item column5"><span>8</span></div>
      <div class="item column2"><span>+</span></div>
      <div class="item column3"><span>2</span></div>
      <div class="item column4"><span>=</span></div>
      <div class="item column5"><span>12</span></div>
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
    display: flex;
    flex-direction: column;
    justify-content: space-around;
    align-items: center;
  }

  .game-container {
    height: 80%;
    width: 90%;
    background-color: var(--white);
    display: grid;
    grid-template: 1fr 1fr 1fr 1fr / 1fr 1fr 1fr 1fr 1fr;
    align-items: center;
    justify-items: center;
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

  .column1 {
    color: var(--pink);
  }
  .column2 {
    color: var(--purple);
  }
  .column3 {
    color: var(--yellow);
  }
  .column4 {
    color: var(--orange);
  }
  .column5 {
    color: var(--cyan);
  }

  .box-container {
    width: 5rem;
    height: 5rem;
    color: var(--white);
    font-size: 3rem;
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;
  }

  .box-slot {
    width: 5rem;
    height: 5rem;
    border: 4px solid var(--shadow);
  }

  .box-0 {
    grid-row: 1;
    grid-column: 5;
  }

  .box-1 {
    grid-row: 2;
    grid-column: 3;
  }

  .box-2 {
    grid-row: 3;
    grid-column: 2;
  }

  .box-3 {
    grid-row: 4;
    grid-column: 1;
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
