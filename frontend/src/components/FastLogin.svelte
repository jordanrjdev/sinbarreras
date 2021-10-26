<script>
  import { navigate } from 'svelte-routing';
  import SelectAvatar from '../components/SelectAvatar.svelte';
  import { userStore, loginStore } from './../store/store';
  let userName;
  const Register = () => {
    userStore.update(state => ({
      ...state,
      userName,
    }));
    loginStore.update(state => ({
      ...state,
      isFirstLogin: true,
    }));
  };
  const Login = user => {
    userStore.update(state => ({
      ...state,
      userName: user.username,
      avatar: user.url,
      score: user.score,
    }));
    navigate('/', { replace: true });
  };
</script>

{#if !$loginStore.isFirstLogin}
  <main>
    <div class="register-container">
      <h1>APRENDE <span>Y</span> <br /> <span>DIVIERTE</span></h1>
      <form on:submit|preventDefault={Register}>
        <label for="inputName"> ¿Cómo te llamas? </label>
        <input id="inputName" bind:value={userName} />
      </form>
    </div>
    <div class="login-container">
      {#each $loginStore.users as user}
        <div on:click={Login(user)}>
          <span>{user.username}</span>
          <img src={user.url} alt="avatar" />
        </div>
      {/each}
    </div>
  </main>
{:else}
  <SelectAvatar />
{/if}

<style>
  main {
    height: 100%;
    width: 100%;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: center;
  }
  .register-container {
    height: 18rem;
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: space-around;
  }
  .login-container {
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: space-around;
    flex-wrap: wrap;
    overflow-y: scroll;
  }
  .login-container::-webkit-scrollbar {
    width: 0.25rem;
    background-color: var(--white);
  }
  .login-container::-webkit-scrollbar-thumb {
    border-radius: 0.25rem;
    background-color: var(--orange);
  }
  h1 {
    font-size: 5rem;
    color: var(--orange);
  }
  h1 > span:nth-child(1) {
    color: var(--purple);
  }
  h1 > br + span {
    color: var(--yellow);
  }
  form > label {
    font-size: 3rem;
  }
  form > input {
    height: 4rem;
    font-size: 2rem;
    border: none;
    margin-top: 2rem;
    border-radius: 1rem;
    transition: box-shadow 1s ease;
  }
  form > input:hover {
    box-shadow: 0 0.5rem 1rem var(--shadow);
  }
  form > label,
  form > input {
    color: var(--black);
  }
  .login-container div {
    height: 12rem;
    width: 12rem;
    border-radius: 2rem;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: space-around;
    box-shadow: 0 0 2rem -1rem var(--shadow);
    cursor: pointer;
    margin: 2rem;
    transition: transform 0.5s ease;
  }
  .login-container div:hover {
    transform: scale(1.05);
  }
  .login-container div > span {
    font-size: 2rem;
    color: var(--purple);
  }
  .login-container div > img {
    height: 6rem;
  }
</style>
