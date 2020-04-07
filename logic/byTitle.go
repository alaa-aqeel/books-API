package logic
import (
	"../books"
	"strings"
	)


func FindByTitleKeyword(keyword string) books.Books {
	var foundBooks books.Books
	for i:=0;i<len(books.BOOKS.Books);i++  {
		book := books.BOOKS.Books[i]
		if strings.Contains(strings.ToUpper(book.Title),strings.ToUpper(keyword)) {
			foundBooks.Books = append(foundBooks.Books, book)
		}
	}
	return foundBooks

}
