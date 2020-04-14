package logic
import (
	"../data"
	"strings"
	"../template"
)


func FindByAuthorKeyword(keyword string) template.Books {
	var foundBooks template.Books
	for i:=0;i<len(data.BOOKS.Books);i++  {
		book := data.BOOKS.Books[i]
		if strings.Contains(strings.ToUpper(book.Author),strings.ToUpper(keyword)) {
			foundBooks.Books = append(foundBooks.Books, book)
		}
	}
	return foundBooks

}
