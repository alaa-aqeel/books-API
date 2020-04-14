package logic

import (
	"../template"
	"strconv"
	"../data"
)

func FindByPageRange(min string,max string) (template.Books,int){
	var foundBooks template.Books
	var flag int = 0
	minPage,err1 := strconv.ParseInt(min, 10, 64)
	maxPage,err2 := strconv.ParseInt(max, 10, 64)
	if err1!=nil || err2!=nil{
		flag = 1
	}else if minPage<0 || maxPage<0{
		flag = 2
	}else if minPage>maxPage{
		flag = 3
	}else {
		for i := 0; i < len(data.BOOKS.Books); i++ {
			book := data.BOOKS.Books[i]
			if book.Pages>=minPage && book.Pages<=maxPage {
				foundBooks.Books = append(foundBooks.Books, book)
			}
		}
	}
	return foundBooks,flag
}
