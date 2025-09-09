import http from "http"
import { greet } from "hrrydgls-greet"

let server = http.createServer((req, res) => {
    res.writeHead(200, {"Content-Type": "text/plain"})
    res.end(greet("NodeJS World"))
})


server.listen(8888, () => {
    console.log("Listenning to port 8888!")
})