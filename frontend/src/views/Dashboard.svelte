<script>
  import { navigate } from 'svelte-routing';
  import { onMount } from 'svelte';
  import { userStore } from '../store/store';
  import NavBar from '../components/NavBar.svelte';
  import Header from '../components/Header.svelte';
  import Score from '../components/common/Score.svelte';
  onMount(() => {
    if (!$userStore.userName) {
      navigate('/login', { replace: true });
    }
  });
</script>

<NavBar />
<main>
  <Header><div class="score"><Score /></div></Header>
  <h2>FUNSIBA <span>beta</span></h2>
</main>

<style>
  main {
    height: 100%;
    width: calc(100% - var(--navbar-width));
    position: relative;
    margin-left: var(--navbar-width);
  }
  main > h2 {
    font-size: 8rem;
    text-align: center;
    color: var(--orange);
  }
  main > h2 > span {
    color: var(--yellow);
  }
  .score {
    position: absolute;
    right: 5%;
  }
</style>
