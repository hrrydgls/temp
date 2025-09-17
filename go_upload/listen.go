package go_upload

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Listen() {
	http.HandleFunc("/", home)
	http.HandleFunc("/upload", upload)

	fmt.Println("Listenning to port 8000...")

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		panic(err)
	}
}
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Server is up!")
	fmt.Fprintf(w, "Home is working!")
}

func upload(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != http.MethodPost {
		http.Error(w, "This route only supports POST request!", http.StatusMethodNotAllowed)
		return
	}

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	dst_name := "./files/" + handler.Filename

	dst, err := os.Create(dst_name)

	if err != nil {
		panic(err)
	}

	_, err = io.Copy(dst, file)

	if err != nil {
		panic(err)
	}

	fmt.Printf("File uploaded to %s", dst_name)
}
