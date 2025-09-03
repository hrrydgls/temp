<script>
  let count = 0;

  function increment() {
    count += 1;
  }

  function decrement() {
    count -= 1;
  }
</script>

<div>
  <h2>Counter: {count}</h2>
  <button on:click={increment}>+</button>
  <button on:click={decrement}>-</button>
</div>

<style>
  div {
    margin: 10px 0;
  }
  button {
    margin-right: 5px;
    padding: 5px 10px;
  }
</style>
