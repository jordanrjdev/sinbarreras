<script>
  import { Link } from 'svelte-routing';
  export let subject;
</script>

<div
  class="subject-container"
  style="background-color: var(--{subject.color});"
>
  <Link to="/games"><h4 on:click>{subject.name}</h4></Link>
</div>

<style>
  .subject-container {
    width: 96%;
    height: 4rem;
    display: flex;
    justify-content: center;
    align-items: center;
    margin-top: 1rem;
    cursor: pointer;
    transition: width 0.5s ease, height 0.5s ease;
  }
  .subject-container:hover {
    width: 100%;
    height: 4.5rem;
  }
  h4 {
    font-size: 2rem;
    color: var(--white);
  }
</style>
