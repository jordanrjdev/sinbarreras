<script>
  import { onMount } from 'svelte';
  import { navigate } from 'svelte-routing';
  import { userStore, subjectStore } from '../store/store';
  import Subject from './common/Subject.svelte';
  const logOut = () => {
    userStore.update(state => ({
      ...state,
      userName: null,
      avatar: null,
      score: 0,
    }));
    navigate('/login', { replace: true });
  };
  onMount(async () => {
    const res = await fetch('https://sinbarreras.herokuapp.com/subjects');
    // const res = await fetch('http://localhost:4000/subjects');
    const subjects = await res.json();
    subjectStore.update(state => ({
      ...state,
      subjects,
    }));
  });
</script>

<nav>
  <div class="greeting-container">
    <h3>Hola <span>{$userStore.userName || 'Usuario'}</span></h3>
  </div>
  <div class="subject-container">
    {#each $subjectStore.subjects as subject}
      <Subject
        {subject}
        on:click={subjectStore.update(state => ({
          ...state,
          subject,
        }))}
      />
    {/each}
  </div>
  <img src={$userStore.avatar || 'images/avatar-default.svg'} alt="avatar" />
  <span class="log-out" on:click={logOut}>SALIR</span>
</nav>

<style>
  nav {
    width: var(--navbar-width);
    height: 100%;
    background-color: var(--purple);
    left: 0;
  }
  .greeting-container {
    padding: 1rem 0;
    position: relative;
    text-align: center;
  }

  nav,
  .greeting-container ::after {
    position: absolute;
  }
  .greeting-container ::after {
    content: '';
    left: 50%;
    bottom: 0;
    border: 0.25rem solid var(--white);
    width: 90%;
    transform: translateX(-50%);
  }
  .greeting-container > h3 {
    font-size: 4rem;
    text-align: center;
    color: var(--white);
  }
  .greeting-container > h3 > span {
    font-size: 3rem;
    color: var(--yellow);
  }
  nav,
  .greeting-container,
  .subject-container {
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  .greeting-container,
  .subject-container {
    width: 100%;
  }
  .log-out {
    font-size: 3rem;
    color: var(--white);
    position: absolute;
    bottom: 2.5%;
    cursor: pointer;
    transition: font-size 0.5s ease;
  }
  .log-out:hover {
    font-size: 4rem;
  }
  img[alt='avatar'] {
    position: absolute;
    width: 10rem;
    bottom: 15%;
    transition: width 0.25s ease;
  }
  img[alt='avatar']:hover {
    width: 10.1rem;
  }
</style>
