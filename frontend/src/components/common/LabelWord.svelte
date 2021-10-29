<script>
  export let word = '';
</script>

<div style="background-color: var(--yellow);">
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
    font-size: 2rem;
  }
</style>
