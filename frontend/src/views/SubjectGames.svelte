<script>
  import { navigate } from 'svelte-routing';
  import { Link } from 'svelte-routing';
  import { subjectStore, userStore } from '../store/store';
  import { onMount } from 'svelte';
  let games = [];
  switch ($subjectStore.subject.name) {
    case 'Matemáticas':
      games.push({ name: 'Operaciones Básicas', path: 'basicOperations' });
      games.push({ name: 'Memória Numérica', path: 'memory' });
      break;
    case 'Lenguaje':
      games.push({ name: 'Adivina Adivinador', path: 'growingUp' });
      games.push({ name: 'Sílabas', path: 'syllabes' });
      break;
    case 'Ciencias':
      games.push({ name: 'Seres Vivos o Inertes', path: 'beigns' });
      break;
    default:
      break;
  }
  onMount(() => {
    if (!$userStore.userName) {
      navigate('/login', { replace: true });
    }
  });
</script>

<main
  style="background-color: var(--{$subjectStore.subject.color || 'yellow'}) ;"
>
  <div class="header"><Link to="/">VOLVER</Link></div>
  <h1>JUEGA CON {$subjectStore.subject.name || 'Materia'}</h1>
  <div class="slice">
    {#each games as game}
      <Link to="/{game.path}"><div class="game">{game.name}</div></Link>
    {/each}
  </div>
</main>

<style>
  main {
    width: 100%;
    height: 100%;
  }
  .header {
    height: 8rem;
    display: flex;
    align-items: center;
    justify-content: end;
    padding-right: 5rem;
  }
  h1 {
    color: var(--white);
    font-size: 6rem;
    text-align: center;
    margin-bottom: 2rem;
    text-shadow: 0 0 3rem -1rem var(--shadow);
  }
  .slice {
    height: 50%;
    display: flex;
    justify-content: space-around;
    align-items: center;
  }

  .slice .game {
    height: 20rem;
    width: 40rem;
    background-color: var(--white);
    box-shadow: 0 0 3rem -2rem var(--shadow);
    display: flex;
    align-items: center;
    justify-content: center;
  }
</style>
