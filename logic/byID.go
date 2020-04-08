package logic

import (
	"../books"
	"fmt"
	"strconv"
)

func FindByID(id string) (books.Books,bool) {
	var foundBook books.Books
	failed := false
	ID,err := strconv.ParseInt(id, 10, 64)
	if err!=nil{
		fmt.Println(err)
		failed = true
	} else if ID<=0{
		failed = true
	}else{
		for i:=0;i<len(books.BOOKS.Books);i++{
			if ID == books.BOOKS.Books[i].Id{
				foundBook.Books = append(foundBook.Books,books.BOOKS.Books[i])
				break
			}
		}
	}
	return foundBook,failed
}