<script>
  import { memoryGameStore } from '../store/store';
  export let value;
  let backGroundColor;
  let flipClass;
  switch (value) {
    case 1:
      backGroundColor = 'purple';
      break;
    case 2:
      backGroundColor = 'pink';
      break;
    case 3:
      backGroundColor = 'cyan';
      break;
    case 4:
      backGroundColor = 'yellow';
      break;
    case 5:
      backGroundColor = 'orange';
      break;
    default:
      backGroundColor = 'black';
      break;
  }
  const voltear = e => {
    let number = Number.parseInt(e.target.parentElement.innerText);
    flipClass = 'flip';
    if ($memoryGameStore.number === 0) {
      memoryGameStore.update(state => ({
        ...state,
        number,
      }));
    } else {
      if ($memoryGameStore.number !== number) {
        memoryGameStore.update(state => ({
          ...state,
          flip: true,
        }));
        setTimeout(() => {
          flipClass = '';
        }, 2000);
      } else {
        let numbers = $memoryGameStore.numbers;
        numbers.push(number);
        memoryGameStore.update(state => ({
          ...state,
          flip: false,
          numbers,
        }));
      }
      memoryGameStore.update(state => ({
        ...state,
        number: 0,
      }));
    }
  };
  $: !$memoryGameStore.numbers.includes(value) && $memoryGameStore.flip
    ? setTimeout(() => {
        flipClass = '';
      }, 2000)
    : '';
</script>

<div class="card-container" style="background-color: var(--{backGroundColor});">
  <div class="card-background {flipClass}" on:click={e => voltear(e)} />
  {value || 0}
</div>

<style>
  .card-container {
    width: 8rem;
    height: 8rem;
    cursor: pointer;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 5rem;
    color: var(--white);
    position: relative;
    box-shadow: 0 4px 8px 0 var(--shadow);
  }

  .card-background {
    width: 100%;
    height: 100%;
    background-color: var(--black);
    position: absolute;
    top: 0;
    left: 0;
    transition: transform 0.6s;
    transform-style: preserve-3d;
    animation: iniciar 5s;
  }

  .flip {
    z-index: -1;
  }

  @keyframes iniciar {
    20%,
    90% {
      transform: rotateY(180deg);
      z-index: -1;
    }

    0%,
    100% {
      transform: rotateY(0deg);
    }
  }
</style>
