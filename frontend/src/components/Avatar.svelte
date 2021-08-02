<div class="container" on:click><slot /></div>

<style>
  .container {
    display: flex;
    justify-content: center;
    align-items: center;
    margin-right: 1rem;
    padding: 1rem 2rem;
    cursor: pointer;
    transition: transform 0.25s ease;
  }
  .container:hover {
    transform: scale(1.1);
  }
</style>
