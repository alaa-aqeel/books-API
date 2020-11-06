package logic

import (
	"../data"
	"../template"
	"fmt"
	"strconv"
	"strings"
)

// careful empty array return

func FindByID(id string) (template.Book,string) {
	var foundBook template.Book

	ID,err := strconv.ParseInt(strings.TrimSpace(id), 10, 64)

	if err!=nil{
		fmt.Println(err)
		return foundBook,"Book ID should be an integer"

	} else if ID<=int64(0) || ID>int64(len(data.BOOKS.Books)){
		return foundBook,"No book exist for the ID"

	}else if (template.Book{}) == data.BOOKS.Books[ID-1]{
		return foundBook,"No book exist for the ID"

	}else{
		return data.BOOKS.Books[ID-1],""
	}
}
