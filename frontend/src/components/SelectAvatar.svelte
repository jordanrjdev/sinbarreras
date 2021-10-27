<script>
  import { Link } from 'svelte-routing';
  import { fade } from 'svelte/transition';
  import { userStore, avatarStore, loginStore } from '../store/store';
  import Avatar from '../components/common/Avatar.svelte';
  import { onMount } from 'svelte';
  $: isSelect = false;
  const setAvatar = avatar => {
    userStore.update(state => ({
      ...state,
      avatar,
    }));
    isSelect = true;
  };
  const selectAvatarAgain = () => {
    userStore.update(state => ({
      ...state,
      avatar: null,
    }));
    isSelect = false;
  };
  const isLogin = async () => {
    loginStore.update(state => ({
      ...state,
      isFirstLogin: false,
    }));
    let username = $userStore.userName.toUpperCase();
    let avatar_id = $userStore.avatar.avatar_id;
    const data = {
      username,
      avatar_id,
    };
    console.log(data);
    const res = await fetch('https://sinbarreras.herokuapp.com/user', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });
    userStore.update(state => ({
      ...state,
      avatar: $userStore.avatar.url,
    }));
  };
  onMount(async () => {
    const res = await fetch('https://sinbarreras.herokuapp.com/avatars');
    const avatars = await res.json();
    avatarStore.set({ ...$avatarStore, avatars });
  });
</script>

<main transition:fade>
  {#if !isSelect}
    <h1>Hola <span>{$userStore.userName}</span></h1>
    <p>Escoge un animalito</p>
    <div class="avatars-container">
      {#each $avatarStore.avatars as avatar}
        <Avatar>
          <img src={avatar.url} alt="avatar" on:click={setAvatar(avatar)} />
        </Avatar>
      {/each}
    </div>
  {:else}
    <div class="login-cart-container">
      <h1>{$userStore.userName}</h1>
      <img src={$userStore.avatar.url} alt="avatar" />
      <div class="login-cart__buttons">
        <button on:click={selectAvatarAgain}>Escoge otro</button>
        <Link to="/" on:click={isLogin}>ACEPTAR</Link>
      </div>
    </div>
  {/if}
</main>

<style>
  main {
    position: relative;
    width: 100%;
    height: 100%;
  }
  h1 {
    font-size: 7rem;
    color: var(--orange);
    margin: 0 0 4rem;
    padding-top: 4rem;
  }
  h1 > span:nth-child(1) {
    color: var(--purple);
  }
  p {
    color: var(--black);
  }
  h1,
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
  .login-cart-container {
    width: 50%;
    height: 80%;
    border-radius: 4rem;
    box-shadow: 0 0 4rem -2rem var(--shadow);
    display: flex;
    flex-direction: column;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    align-items: center;
    justify-content: space-around;
  }
  .login-cart-container > img {
    height: 12rem;
  }
  .login-cart__buttons {
    display: flex;
    justify-content: space-around;
    align-items: center;
    width: 100%;
  }
  .login-cart__buttons > button {
    cursor: pointer;
    border: none;
    border-radius: 1rem;
    background-color: var(--yellow);
    color: var(--white);
    padding: 1rem 2rem;
  }
</style>
