package main

import "fmt"

type Book struct {
	Bookname   string
	BookAuthor string
	NoOfCopies int
}

type BindBook Book

func main() {
	// var oldBook Book
	var oldBook BindBook

	newBook1 := Book{}
	newBook1.Bookname = "Innovation Capital"
	newBook1.BookAuthor = "Chris"
	newBook1.NoOfCopies = 10

	oldBookIface := testingInterface(newBook1)
	oldBook = oldBookIface.(BindBook)
	fmt.Printf("%s", oldBook)

}

func testingInterface(inputIn interface{}) (OutputIn interface{}) {
	// var tmpBook Book

	// tmpBook = inputIn.(Book)
	// tmpBook.BookAuthor = "Thamizh"

	// return tmpBook
	return inputIn
}
