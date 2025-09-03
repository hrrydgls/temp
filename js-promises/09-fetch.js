fetch("https://jsonplaceholder.typicode.com/todos/1")
  .then(response => {
    console.log(response.status)
    console.log(response.statusText)
    console.log(response.ok)
    return response.json()
  })
  .then(data => console.log(data))
  .catch(err => console.log("Error:", err));

// Output:
// 200
// OK
// true
// { userId: 1, id: 1, title: 'delectus aut autem', completed: false }