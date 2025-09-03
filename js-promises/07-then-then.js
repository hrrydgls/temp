function wait () {
  return new Promise((resolve, Rejected) => {
    resolve("Done")
  })
}

wait().then(result => result + ' + sth else')
  .then(result2 => console.log(result2))

// Each .then() can return a value or a Promise.
// If it returns a value → next .then() gets that value.
// If it returns a Promise → next .then() waits for it to resolve.