package logic

import (
	"fmt"
	"strconv"
	"strings"
	"../data"
	"../template"
)

func DeleteABook(id string) (bool,string) {

	ID,err := strconv.ParseInt(strings.TrimSpace(id), 10, 64)
	var emptyBook template.Book

	if err!=nil{
		fmt.Println(err)
		return false, "Book ID should be an integer"

	}else if ID > int64(0) && ID <= int64(len(data.BOOKS.Books)) && (template.Book{}) != data.BOOKS.Books[ID-1]{
		data.BOOKS.Books[ID-1] = emptyBook
		return true, ""

	}else{
		return false,  "No book exist for the ID"
		
	}
	
}