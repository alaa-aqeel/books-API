package logic

import (
	"../data"
	"../template"
	"fmt"
	"strconv"
	"strings"
)

func UpdateBook(id string,book template.Book) (template.Book,string){
	var foundBook template.Book

	ID,err := strconv.ParseInt(id, 10, 64)

	if err!=nil{
		fmt.Println(err)
		return foundBook, "Book ID should be an integer"

	}else if ID > int64(0) && ID <= int64(len(data.BOOKS.Books)) && data.BOOKS.Books[ID-1] != (template.Book{}){
		if strings.TrimSpace(book.Title) != "" {
			data.BOOKS.Books[ID-1].Title = book.Title
	
		}
		if strings.TrimSpace(book.Author) != "" {
			data.BOOKS.Books[ID-1].Author = book.Author
	
		}
		if strings.TrimSpace(book.Language) != "" {
			data.BOOKS.Books[ID-1].Language = book.Language
	
		}
		if strings.TrimSpace(book.Link) != "" {
			data.BOOKS.Books[ID-1].Link = book.Link
	
		}
		if book.Pages < int64(0){
			return foundBook, "Book should have at least one page"
		}else if book.Pages > int64(0){
			data.BOOKS.Books[ID-1].Pages = book.Pages
		}
		
		return data.BOOKS.Books[ID-1], ""
		
	}else{
		return foundBook, "No book exist for the ID"
		
	}
}