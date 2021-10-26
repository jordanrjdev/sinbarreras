<script>
  import { onMount } from 'svelte';
  import { navigate } from 'svelte-routing';
  import { userStore } from '../store/store';
  onMount(() => {
    if (!$userStore.userName) {
      navigate('/login', { replace: true });
    }
    setTimeout(() => {
      navigate('/', { replace: true });
    }, 5000);
  });
</script>

<main>
  <span>GANASTE!!!</span>
  <p>Felicitaciones <span>{$userStore.userName}</span></p>
  <img src="/images/medal.svg" alt="medalla" />
</main>

<style>
  main {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: space-around;
    font-size: 6rem;
  }
  main > span {
    color: var(--pink);
  }
  main > p {
    color: var(--black);
  }
  main > p > span {
    color: var(--orange);
  }
</style>
