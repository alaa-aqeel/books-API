package logic

import (
	"../books"
	"fmt"
	"strconv"
)

func UpdateBook(id string,aBook books.Book) (books.Books,bool){
	failed := false
	var result books.Books
	ID,err := strconv.ParseInt(id, 10, 64)
	if err!=nil{
		fmt.Println(err)
		failed = true
	}else if ID<=0{
		failed = true
	}else{
				foundBook:=books.UpdateABook(aBook,ID)
				result.Books = append(result.Books,foundBook)
	}
	return result,failed
}
