package logic

import (
	"fmt"
	"strconv"
	"../books"
)

func DeleteABook(id string) (bool,bool) {
	failed := false
	found := false
	ID,err := strconv.ParseInt(id, 10, 64)
	if err!=nil{
		fmt.Println(err)
		failed = true
	}else if ID<=0{
		failed = true
	}else{
		found = books.DeleteBook(ID)
	}
	return failed,found
}