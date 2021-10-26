<script>
  import { flip } from 'svelte/animate';
  import { dndzone } from 'svelte-dnd-action';
  export let items;
  export let color;
  const flipDurationMs = 300;
  function handleDndConsider(e) {
    items = e.detail.items;
  }
  function handleDndFinalize(e) {
    items = e.detail.items;
  }
</script>

<section
  use:dndzone={{ items, flipDurationMs }}
  on:consider={handleDndConsider}
  on:finalize={handleDndFinalize}
>
  {#each items as item (item.id)}
    <div
      class="coloritem"
      animate:flip={{ duration: flipDurationMs }}
      style="backgrounds-color:{color}"
    >
      {item.name}
    </div>
  {/each}
</section>

<style>
  section {
    border: #999 1px solid;
    background-color: #999;
    width: 80px;
    height: 70px;
    margin-left: 45px;
    margin-bottom: 45px;
  }
  div {
    /* width: 50%;
    padding: 0.2em;
    border: 1px solid blue;
    margin: 0.15em 0; */

    color: white;
    width: 80px;
    height: 75px;
    margin: -4px 0 4px 4px;
    text-align: center;
    justify-content: center;
    cursor: move;
    margin-bottom: 40px;
    margin-left: 0px;
  }

  .coloritem {
    background-color: #56cf9f;
  }
</style>
