function wait(seconds) {
  setTimeout(() => {
    console.log(`${seconds} seconds passed.`);
  }, seconds * 1000);
}

wait(5);
wait(2);


// Output:
// 2 seconds passed.
// 5 seconds passed.