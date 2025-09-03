function wait(seconds) {
  return new Promise((resolve, reject) => {
    setTimeout(
      () => resolve(`Resolved after ${seconds} seconds...`),
      seconds * 1000
    );
  });
}

async function runQueue() {
  console.log(await wait(4));
  console.log(await wait(2));
}

runQueue();

// Output: 
// Resolved after 4 seconds...
// Resolved after 2 seconds...