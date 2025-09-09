package main 

// import "fmt"

// type Book struct {
// 	Title string
// 	Author string
// }

// func (book Book) Describe () string {
// 	return fmt.Sprintf("%s by %s", book.Title, book.Author)
// }

// func newBook(title, author string) Book {
// 	return Book {Title: title, Author: author}
// }

// func main () {
// 	book := newBook("Learn Golang", "Harry")
// 	fmt.Println(book.Describe())
// }

import "github.com/hrrydgls/temp/echo" 

func main () {
	echo.Echo()
	echo.AnotherFunc()
}

