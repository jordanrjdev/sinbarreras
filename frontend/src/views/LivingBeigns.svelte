<script>
  import { livingBeignsStore, userStore } from '../store/store';
  import { flip } from 'svelte/animate';
  import { crossfade } from 'svelte/transition';
  import { quintOut, elasticOut } from 'svelte/easing';
  import { draggable } from '../js/dragdrop';
  import { navigate } from 'svelte-routing';
  import { onMount } from 'svelte';

  const putInItems = item => {
    let items = $livingBeignsStore.items;
    let beigns = $livingBeignsStore.beigns;
    let inerts = $livingBeignsStore.inerts;
    if (items.indexOf(item) === -1) {
      items.push(item);
      if (beigns.indexOf(item) !== -1) {
        beigns = beigns.filter(i => i != item);
      }
      if (inerts.indexOf(item) !== -1) {
        inerts = inerts.filter(i => i != item);
      }
    }
  };

  const putInBeign = item => {
    let items = $livingBeignsStore.items;
    let beigns = $livingBeignsStore.beigns;
    let inerts = $livingBeignsStore.inerts;
    if (items.indexOf(item) !== -1) {
      items = items.filter(i => i != item);
    }
    if (inerts.indexOf(item) !== -1) {
      inerts = inerts.filter(i => i != item);
    }
    if (beigns.indexOf(item) === -1) {
      beigns.push(item);
    }

    livingBeignsStore.update(state => ({
      ...state,
      items,
      beigns,
      inerts,
    }));
    if (
      beigns.includes(beigns.find(i => i.id === 1)) &&
      beigns.includes(beigns.find(i => i.id === 2)) &&
      beigns.includes(beigns.find(i => i.id === 5)) &&
      inerts.includes(inerts.find(i => i.id === 3)) &&
      inerts.includes(inerts.find(i => i.id === 4)) &&
      inerts.includes(inerts.find(i => i.id === 6))
    ) {
      navigate('/congratulations', { replace: true });
    }
  };
  const putInInert = item => {
    let items = $livingBeignsStore.items;
    let beigns = $livingBeignsStore.beigns;
    let inerts = $livingBeignsStore.inerts;
    if (items.indexOf(item) !== -1) {
      items = items.filter(i => i != item);
    }
    if (beigns.indexOf(item) !== -1) {
      beigns = beigns.filter(i => i != item);
    }
    if (inerts.indexOf(item) === -1) {
      inerts.push(item);
    }
    livingBeignsStore.update(state => ({
      ...state,
      items,
      inerts,
      beigns,
    }));
    if (
      beigns.includes(beigns.find(i => i.id === 1)) &&
      beigns.includes(beigns.find(i => i.id === 2)) &&
      beigns.includes(beigns.find(i => i.id === 5)) &&
      inerts.includes(inerts.find(i => i.id === 3)) &&
      inerts.includes(inerts.find(i => i.id === 4)) &&
      inerts.includes(inerts.find(i => i.id === 6))
    ) {
      navigate('/congratulations', { replace: true });
    }
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
    <h1>ARRASTRA SERES VIVOS E INERTES</h1>
    <div class="cart" on:dropped={e => putInItems(e.detail)}>
      {#each $livingBeignsStore.items as item (item.id)}
        <div
          class="item"
          style="background-image: url('images/{item.name}.svg');"
          animate:flip
          use:draggable={{ data: item, targets: ['.slot', '.slot .item'] }}
          in:receive={item.id}
          out:send={item.id}
        />
      {/each}
    </div>
    <div />
  </div>
  <div>
    <div class="game-container ">
      <div class="game-beign">
        <span>Seres Vivos</span>
        <div class="box-slot slot" on:dropped={e => putInBeign(e.detail)}>
          {#each $livingBeignsStore.beigns as item, index}
            {#if item}
              <div
                class="item"
                use:draggable={{
                  data: item,
                  targets: ['.cart', '.slot', '.slot .item'],
                }}
                in:receive={item.id}
                out:send={item.id}
                on:dropped={e => putInBeign(e.detail)}
                style="background-image: url('images/{item.name}.svg');"
              />
            {/if}
          {/each}
        </div>
      </div>
      <div class="game-inert">
        <span>Seres Inertes</span>
        <div class="box-slot slot" on:dropped={e => putInInert(e.detail)}>
          {#each $livingBeignsStore.inerts as item, index}
            {#if item}
              <div
                class="item"
                use:draggable={{
                  data: item,
                  targets: ['.cart', '.slot', '.slot .item'],
                }}
                in:receive={item.id}
                out:send={item.id}
                on:dropped={e => putInInert(e.detail, index)}
                style="background-image: url('images/{item.name}.svg');"
              />
            {/if}
          {/each}
        </div>
      </div>
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
    display: grid;
    grid-template: 1fr 1fr 1fr/ 1fr 1fr;
    margin-top: 4rem;
    justify-items: center;
    align-items: center;
  }

  .game-container {
    height: 80%;
    width: 90%;
    background-color: var(--white);
    display: flex;
    justify-content: space-around;
  }
  .game-container > div {
    width: 40%;
    height: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: space-between;
  }
  .game-container > div > span {
    font-size: 3rem;
    color: var(--purple);
  }
  .item {
    width: 8rem;
    height: 8rem;
    background-repeat: no-repeat;
    background-size: contain;
  }
  .box-slot {
    width: 100%;
    height: 100%;
    border: 1px solid black;
    display: grid;
    grid-template: 1fr 1fr 1fr / 1fr 1fr;
    justify-items: center;
    align-items: center;
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
</style>
