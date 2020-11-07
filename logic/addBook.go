package logic

import (
	"../data"
	"strings"
	"../template"
)

func AddABook(newBook template.Book) (template.Book,string) {
	var book template.Book
	if strings.TrimSpace(newBook.Title) == "" {
		return book,"Book title should not be empty"

	}else if strings.TrimSpace(newBook.Author) == "" {
		return book,"Book author should not be empty"

	}else if strings.TrimSpace(newBook.Language) == "" {
		return book,"Book language should not be empty"

	}else if newBook.Pages <= 0 {
		return book,"Book should have at least one page"

	}else {
		newBook.Id = int64(len(data.BOOKS.Books)+1)
		data.BOOKS.Books = append(data.BOOKS.Books,newBook)
		
		return data.BOOKS.Books[len(data.BOOKS.Books)-1],""
	}

}
