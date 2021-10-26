<script>
  import { Router, Route } from 'svelte-routing';
  import Login from './views/Login.svelte';
  import Dashboard from './views/Dashboard.svelte';
  import SubjectGames from './views/SubjectGames.svelte';
  import GrowingUpGame from './views/GrowingUpGame.svelte';
  import MemoryGame from './views/MemoryGame.svelte';
  import SyllabeGame from './views/SyllabeGame.svelte';
  import BasicOperationsGames from './views/BasicOperationsGames.svelte';
  import LivingBeigns from './views/LivingBeigns.svelte';
  import Congratulations from './views/Congratulations.svelte';
</script>

<Router>
  <Route path="/login"><Login /></Route>
  <Route path="/"><Dashboard /></Route>
  <Route path="/games"><SubjectGames /></Route>
  <Route path="/basicOperations"><BasicOperationsGames /></Route>
  <Route path="/memory"><MemoryGame /></Route>
  <Route path="/growingUp"><GrowingUpGame /></Route>
  <Route path="/syllabes"><SyllabeGame /></Route>
  <Route path="/beigns"><LivingBeigns /></Route>
  <Route path="/congratulations"><Congratulations /></Route>
</Router>

<style>
</style>
