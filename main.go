package main

import (
	"flag"
	"fmt"

	"github.com/hrrydgls/temp/echo"
	"github.com/hrrydgls/temp/go_channels"
	"github.com/hrrydgls/temp/go_upload"
	"github.com/hrrydgls/temp/goroutines"
)

type Book struct {
	Title  string
	Author string
}

func (book Book) Describe() string {
	return fmt.Sprintf("%s by %s", book.Title, book.Author)
}

func newBook(title, author string) Book {
	return Book{Title: title, Author: author}
}

func main() {

	command := flag.String("command", "", "The command you wanna run!")

	flag.Parse()

	switch *command {
	case "echo":
		echo.Index()
	case "upload":
		go_upload.Listen()
	case "go":
		goroutines.Go()
	case "channel":
		go_channels.ChannelRunner()
	default:
		goHere()
	}

}

func goHere() {
	book := newBook("Learn Golang", "Harry")
	fmt.Println(book.Describe())
}
