<script lang="ts">
  import { fade } from 'svelte/transition';
  import { userStore } from '../store/userStore';
  import Avatar from '../components/Avatar.svelte';
  const images: string[] = [];
  $: login = false;
  $: userName = '';
  $: disabled = true;
  for (let i = 1; i <= 22; i++) {
    images.push(`images/${i}.svg`);
  }
  const Login = (e: Event) => {
    e.preventDefault();
    login = true;
    userStore.update(state => ({
      ...state,
      userName: userName,
    }));
  };
</script>

{#if !login}
  <main>
    <h1>APRENDE <span>Y</span> <br /> <span>DIVIERTE</span></h1>
    <form on:submit={e => Login(e)}>
      <label for="inputName"> ¿Cómo te llamas? </label>
      <input type="text" id="inputName" bind:value={userName} />
    </form>
  </main>
{:else}
  <main transition:fade>
    <h1>Hola <span>{userName}</span></h1>
    <p>Escoge un animalito</p>
    <div class="avatars-container">
      {#each images as src}
        <Avatar on:click={() => (disabled = false)}>
          <img {src} alt="" />
        </Avatar>
      {/each}
    </div>
    <button {disabled}> ACEPTAR </button>
  </main>
{/if}

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
    height: 4rem;
    font-size: 2rem;
    border: none;
    margin-top: 2rem;
    border-radius: 2rem;
    transition: box-shadow 1s ease;
  }
  form > input:hover {
    box-shadow: 0 0.5rem 1rem var(--shadow);
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
    width: 100%;
    margin: 2rem 0;
    display: flex;
    overflow-x: scroll;
    overflow-y: hidden;
  }
  .avatars-container::-webkit-scrollbar {
    height: 0.5rem;
    background-color: var(--white);
  }
  .avatars-container::-webkit-scrollbar-thumb {
    border-radius: 0.25rem;
    background-color: var(--orange);
  }
  button {
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
  button:disabled {
    background-color: rgb(224, 224, 224);
    color: var(--white);
    cursor: auto;
  }
</style>
