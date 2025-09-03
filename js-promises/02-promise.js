let promise = new Promise((resolve, reject) => {
  let success = true;

  if (success) {
    setTimeout(() => {
      resolve("Resolved");
    }, 2000)
  } else {
    reject("Rejected");
  }
});

promise
  .then((result) => console.log(result))
  .catch((error) => console.log(error))
  .finally(() => console.log("Done"));
