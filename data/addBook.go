package data

import "fmt"
import "../template"
func AddBook(book template.Book)  {
	book.Id = int64(len(BOOKS.Books)+1)
	BOOKS.Books = append(BOOKS.Books,book)
	fmt.Println("New Book with ID:",book.Id,"added")

}
