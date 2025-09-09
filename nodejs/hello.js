import http from "http"

let server = http.createServer((req, res) => {
    res.writeHead(200, {"Content-Type": "text/plain"})
    res.end("Welcome to my first server!")
})


server.listen(8888, () => {
    console.log("Listenning to port 8888!")
})