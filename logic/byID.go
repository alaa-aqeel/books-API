package logic

import (
	"../data"
	"../template"
	"fmt"
	"strconv"
)

func FindByID(id string) (template.Books,bool) {
	var foundBook template.Books
	failed := false
	ID,err := strconv.ParseInt(id, 10, 64)
	if err!=nil{
		fmt.Println(err)
		failed = true
	} else if ID<=0{
		failed = true
	}else{
		for i:=0;i<len(data.BOOKS.Books);i++{
			if ID == data.BOOKS.Books[i].Id{
				foundBook.Books = append(foundBook.Books, data.BOOKS.Books[i])
				break
			}
		}
	}
	return foundBook,failed
}