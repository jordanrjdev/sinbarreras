<script>
  import { loginStore } from '../store/store';
  import FirstLogin from '../components/FirstLogin.svelte';
  import FastLogin from '../components/FastLogin.svelte';
  import SelectAvatar from '../components/SelectAvatar.svelte';
  import { onMount } from 'svelte';

  onMount(async () => {
    const res = await fetch('https://sinbarreras.herokuapp.com/users');
    // const res = await fetch('http://localhost:4000/users');
    const users = await res.json();
    loginStore.update(state => ({
      ...state,
      users,
    }));
  });
</script>

{#if $loginStore.users.length === 0}
  {#if !$loginStore.isFirstLogin}
    <FirstLogin />
  {:else}
    <SelectAvatar />
  {/if}
{:else}
  <FastLogin />
{/if}

<style>
</style>
