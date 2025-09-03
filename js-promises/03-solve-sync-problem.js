// function wait(seconds) {
//   setTimeout(() => {
//     console.log(`${seconds} seconds passed.`);
//   }, seconds * 1000);
// }

// wait(5);
// wait(2);

function wait(seconds) {
  console.log(`wait with ${seconds} seconds called`);
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      console.log(`${seconds} seconds passed.`);
      resolve(`Resolve called after ${seconds} seconds...`);
    }, seconds * 1000);
  });
}

wait(5).then((result) => {
  console.log(`Result: ${result}`);
  wait(2).then((result) => {
    console.log(`Result: ${result}`);
  });
});

// Output: 
// 5 seconds passed.
// Result: Resolve called after 5 seconds...
// wait with 2 seconds called
// 2 seconds passed.
// Result: Resolve called after 2 seconds...
