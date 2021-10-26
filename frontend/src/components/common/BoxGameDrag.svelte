<script>
  export let item;
</script>

<div class="box-container" style="background-color: var(--{item.color});">
  {item.name}
</div>

<style>
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
</style>
