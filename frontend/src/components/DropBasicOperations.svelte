<script>
  import { basicOperationsStore } from '../store/store';
  import { flip } from 'svelte/animate';
  import { dndzone } from 'svelte-dnd-action';
  $: items = $basicOperationsStore.dropped;
  export let value;
  console.log(value);
  const flipDurationMs = 300;

  function handleDndConsider(e) {
    basicOperationsStore.update(state => ({
      ...state,
      dropped: e.detail.items,
    }));
  }

  function handleDndFinalize(e) {
    basicOperationsStore.update(state => ({
      ...state,
      dropped: e.detail.items,
    }));
    console.log(e.detail.items);
  }
</script>

<section
  use:dndzone={{ items, flipDurationMs }}
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
    color: white;
    background-color: darkgrey;
    width: 60px;
    height: 60px;
    margin-left: 45px;
  }
</style>
