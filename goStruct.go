package main

import "fmt"

/*
go结构体

	type struct_variable_type struct {
	   member definition
	   member definition
	   ...
	   member definition
	}
*/
type Book struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var book1 = Book{"lbytech", "lby", "all", 1}

	var book2 Book
	book2.title = "Go Program"
	book2.author = "www.golang.org"
	book2.subject = "Go Program"
	book2.book_id = 2

	printBook(book1)
	printBook(book2)

	var bookPtr *Book
	bookPtr = &book1
	printBookByPointer(bookPtr)
	printBookByPointer(&book2)
}

func printBook(book Book) {
	fmt.Print("id:", book.book_id)
	fmt.Print(",title:" + book.title)
	fmt.Print(",author:" + book.author)
	fmt.Println(",subject:" + book.subject)
}

func printBookByPointer(book *Book) {
	fmt.Print("id:", book.book_id)
	fmt.Print(",title:" + book.title)
	fmt.Print(",author:" + book.author)
	fmt.Println(",subject:" + book.subject)
}
