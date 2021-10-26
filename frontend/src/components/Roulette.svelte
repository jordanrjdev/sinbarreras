<script></script>

<div class="roulette-container" />

<style>
  .roulette-container {
    width: 25rem;
    height: 25rem;
    background-color: var(--pink);
    border-radius: 50%;
    position: relative;
  }
  .roulette-container::before {
    content: '';
    position: absolute;
    top: -2.5%;
    left: 50%;
    transform: translateX(-50%);
    width: 0;
    height: 0;
    border-style: solid;
    border-width: 80px 32px 0 32px;
    border-color: var(--black) transparent transparent transparent;
  }
</style>
