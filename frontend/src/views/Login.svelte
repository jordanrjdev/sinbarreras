<script lang="ts">
  import { fade } from 'svelte/transition';
  import { userStore } from '../store/userStore';
  $: login = false;
  $: userName = '';
  const Login = (e: Event) => {
    e.preventDefault();
    login = true;
    userStore.update(state => ({
      ...state,
      userName: userName,
    }));
  };
</script>

<main transition:fade>
  {#if !login}
    <h1>APRENDE <span>Y</span> <br /> <span>DIVIERTE</span></h1>
    <form on:submit={e => Login(e)}>
      <label for="inputName"> ¿Cómo te llamas? </label>
      <input type="text" id="inputName" bind:value={userName} />
    </form>
  {:else}
    <h1>Hola <span>{userName}</span></h1>
    <p>Escoge un avatar</p>
    <div class="avatars-container" />
    <input type="submit" value="ACEPTAR" disabled />
  {/if}
</main>

<style>
  h1 {
    font-size: 7rem;
    color: var(--orange);
    margin: 4rem 0;
  }
  h1 > span:nth-child(1) {
    color: var(--purple);
  }
  h1 > br + span {
    color: var(--yellow);
  }
  form,
  main {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
  }
  form {
    margin-top: 1rem;
  }
  form > label {
    font-size: 6rem;
  }
  form > input {
    width: 50%;
    height: 4rem;
    font-size: 2rem;
    border: none;
    box-shadow: 0 0.5rem 1rem var(--shadow);
    margin-top: 2rem;
  }
  form > label,
  form > input,
  p {
    color: var(--black);
  }
  h1,
  form > label,
  form > input,
  p {
    text-align: center;
  }
  p {
    font-size: 3rem;
  }
  .avatars-container {
    height: 16rem;
    margin: 2rem 0;
  }
  input[type='submit'] {
    height: 4rem;
    width: 20rem;
    border: none;
    border-radius: 2rem;
    letter-spacing: 0.5rem;
    font-size: 2rem;
    background-color: var(--yellow);
    color: var(--white);
    cursor: pointer;
  }
</style>
