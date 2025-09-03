function afterTwoSeconds() {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      // resolve("Done!");
      reject("Failed!");
    }, 2000);
  });
}

async function caller() {
  try {
    let result = await afterTwoSeconds();
    console.log(result);
  } catch (error) {
    console.log(error);
  }
  console.log("Finished!");
}

caller();


// Output:
// Failed!
// Finished!