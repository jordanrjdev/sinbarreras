<script>
  export let backgroundColor;
  export let word = '';
</script>

<div style="background-color: var(--{backgroundColor});">
  <span>{word || ''}</span>
</div>

<style>
  div {
    width: 80%;
    height: 10%;
    border-radius: 1rem;
    display: flex;
    justify-content: center;
    align-items: center;
  }
  span {
    color: var(--white);
  }
</style>
