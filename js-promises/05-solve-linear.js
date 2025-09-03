
function wait(seconds) {
  console.log(`Wait with ${seconds} seconds called`);
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      console.log(`${seconds} seconds passed.`);
      resolve(`Resolve called after ${seconds} seconds...`);
    }, seconds * 1000);
  });
}

async function callWaits()
{
    let result1 = await wait(5)
    console.log("Result 1: ", result1)
    
    let result2 = await wait(2)
    console.log("Result 2: ", result2)
}

callWaits()