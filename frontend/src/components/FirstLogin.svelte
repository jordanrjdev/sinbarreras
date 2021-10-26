<script>
  import { userStore, loginStore } from '../store/store';
  $: userName = '';
  const Login = () => {
    userStore.update(state => ({
      ...state,
      userName,
    }));

    loginStore.update(state => ({
      ...state,
      isFirstLogin: true,
    }));
  };
</script>

<main>
  <h1>APRENDE <span>Y</span> <br /> <span>DIVIERTE</span></h1>
  <form on:submit|preventDefault={Login}>
    <label for="inputName"> ¿Cómo te llamas? </label>
    <input id="inputName" bind:value={userName} />
  </form>
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
  h1,
  form > label,
  form > input {
    text-align: center;
  }
</style>
