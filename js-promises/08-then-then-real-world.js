function fetchUser(userId) {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      console.log(`User with id ${userId} fetched!`);
      resolve({
        id: userId,
        name: "Hamid Haghdoost",
        username: "hamid",
      });
    }, 2000);
  });
}

function fetchPosts(username) {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      console.log(`Posts fetched for @${username}`)
      resolve([
        {
          postId: 1,
          title: "First post",
        },
        {
          postId: 2,
          title: "Second post",
        },
      ]);
    }, 2000);
  });
}

function fetchUserPosts(userId) {
  return new Promise((resolve, reject) => {
    return fetchUser(userId).then((user) => {
      resolve(fetchPosts(user["username"]));
    });
  });
}

fetchUserPosts(1).then((posts) => console.log(posts));

// Output:
// User with id 1 fetched!
// Posts fetched for @hamid
// [
//   { postId: 1, title: 'First post' },
//   { postId: 2, title: 'Second post' }
// ]