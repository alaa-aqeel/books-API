package logic

import (
	"../data"
	"../template"
	"fmt"
	"strconv"
)

func UpdateBook(id string,aBook template.Book) (template.Books,bool){
	failed := false
	var result template.Books
	ID,err := strconv.ParseInt(id, 10, 64)
	if err!=nil{
		fmt.Println(err)
		failed = true
	}else if ID<=0 || ID > int64(len(data.BOOKS.Books)){
		failed = true
	}else{
				foundBook:= data.UpdateABook(aBook,ID)
				result.Books = append(result.Books,foundBook)
	}
	return result,failed
}
