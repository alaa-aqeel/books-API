package logic

import (
	"../data"
	"../template"
	"fmt"
	"strconv"
)

func UpdateBook(id string,aBook template.Book) (template.Book,bool){
	failed := false
	var result template.Book
	ID,err := strconv.ParseInt(id, 10, 64)
	if err!=nil{
		fmt.Println(err)
		failed = true
	}else if ID<=0 || ID > int64(len(data.BOOKS.Books)){
		failed = true
	}else{
				foundBook:= data.UpdateABook(aBook,ID)
				result = foundBook
	}
	return result,failed
}
