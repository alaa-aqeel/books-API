package data

import (
	"fmt"
	"../template"
)

func UpdateABook(book template.Book,ID int64) template.Book{
	var foundIndex int
	for i := 0; i < len(BOOKS.Books); i++ {
		if ID == BOOKS.Books[i].Id {
			if book.Language != "" {
				BOOKS.Books[i].Language = book.Language
			}
			if book.Author != ""{
				BOOKS.Books[i].Author = book.Author
			}
			if book.Title != "" {
				BOOKS.Books[i].Title = book.Title
			}
			if book.Link != "" {
				BOOKS.Books[i].Link = book.Link
			}
			if book.Pages > 0 {
				BOOKS.Books[i].Pages = book.Pages
			}
			foundIndex = i
			fmt.Println("Book with ID:",BOOKS.Books[i].Id,"updated")
			break
		}
	}
	return BOOKS.Books[foundIndex]
}
