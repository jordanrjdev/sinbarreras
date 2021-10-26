<script>
  import { basicOperationsStore } from '../store/store';
  import { flip } from 'svelte/animate';
  import { dndzone } from 'svelte-dnd-action';
  const flipDurationMs = 300;
  const dropTargetStyle = { outline: 'rgba(0, 0, 0, 0) solid 0px' };

  let items = $basicOperationsStore.items;

  function handleDndConsider(e) {
    basicOperationsStore.update(state => ({
      ...state,
      items: e.detail.items,
    }));
    console.log(e.detail.items);
  }

  function handleDndFinalize(e) {
    basicOperationsStore.update(state => ({
      ...state,
      items: e.detail.items,
    }));
    console.log(e.detail.items);
  }
</script>

<!-- on:consider={handleDndConsider}  -->
<section
  use:dndzone={{ items, flipDurationMs, dropTargetStyle }}
  on:consider={handleDndConsider}
  on:finalize={handleDndFinalize}
>
  {#each items as item (item.id)}
    <div
      style={'background-color:' + item.color}
      animate:flip={{ duration: flipDurationMs }}
    >
      {item.name}
    </div>
  {/each}
</section>

<style>
  section {
    width: 80px;
    height: 270px;
    margin-left: 80px;
    align-content: center;
  }

  div {
    color: white;
    width: 60px;
    height: 60px;
    margin: 10px;
    text-align: center;
    justify-content: center;
    cursor: move;
  }
</style>
