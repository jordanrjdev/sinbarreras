<script>
  import { themeStore, userStore } from '../../store/store';
</script>

<div class="score-container">
  <span>{$userStore.score}</span>
  <img src={$themeStore.scoreImage || 'images/score-default.svg'} alt="scoreImage" />
</div>

<style>
  .score-container {
    width: 10rem;
    height: 4rem;
    background-color: var(--cyan);
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0 1rem;
    box-shadow: 0 4px 4px rgba(0, 0, 0, 0.25);
  }
  .score-container > span {
    font-size: 2.5rem;
    text-align: center;
    color: var(--white);
  }
  img[alt='scoreImage'] {
    width: 3rem;
  }
</style>
