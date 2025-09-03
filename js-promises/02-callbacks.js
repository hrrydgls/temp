function wait(seconds, callback) {
  setTimeout(() => {
    console.log(`${seconds} seconds passed.`);
    callback()
  }, seconds * 1000);
}

wait(5, function () {
    wait(2, () => {})
});