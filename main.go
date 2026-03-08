package main

import (
	"flag"
	"fmt"

	"github.com/hrrydgls/temp/echo"
	"github.com/hrrydgls/temp/go_channels"
	"github.com/hrrydgls/temp/go_upload"
	"github.com/hrrydgls/temp/godi"
	"github.com/hrrydgls/temp/goroutines"
	"github.com/hrrydgls/temp/go_pointers"
	"github.com/hrrydgls/temp/go_vars"
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

	option := flag.String("option", "", "Some commands need an option!")

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
	case "di":
		godi.Run()
	case "vars":
		vars.Run(*option)
	case "pointer":
		pointer.Run()
	default:
		goHere()
	}

}

func goHere() {
	fmt.Println("This is not a valid command but I suggest you to read this book:")
	book := newBook("Learn Golang", "Harry")
	fmt.Println(book.Describe())
}
