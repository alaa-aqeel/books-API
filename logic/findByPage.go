package logic

import (
	"../template"
	"strconv"
	"../data"
)

func FindByPageRange(min string,max string) (template.Books,string){

	var foundBooks template.Books
	foundBooks.Books = make([]template.Book, 0)
	
	minPage,err1 := strconv.ParseInt(min, 10, 64)
	maxPage,err2 := strconv.ParseInt(max, 10, 64)

	if err1!=nil || err2!=nil {
		return foundBooks,"Number of pages should be integer"

	}else if minPage<=0 || maxPage<=0 {
		return foundBooks,"Number of pages should be at least 1"

	}else if minPage>maxPage {
		return foundBooks,"Minimum pages should not exceed maximum pages"
		
	}else {
		for i := 0; i < len(data.BOOKS.Books); i++ {

			book := data.BOOKS.Books[i]

			if  book != (template.Book{}) && 
				book.Pages>=minPage && 
				book.Pages<=maxPage {
				foundBooks.Books = append(foundBooks.Books, book)
				
			}
		}
	}
	
	return foundBooks,""
}
