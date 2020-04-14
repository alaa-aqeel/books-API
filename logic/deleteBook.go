package logic

import (
	"fmt"
	"strconv"
	"../data"
)

func DeleteABook(id string) (bool,bool) {
	failed := false
	found := false
	ID,err := strconv.ParseInt(id, 10, 64)
	if err!=nil{
		fmt.Println(err)
		failed = true
	}else if ID<=0 || ID > int64(len(data.BOOKS.Books)){
		failed = true
	}else{
		found = data.DeleteBook(ID)
	}
	return failed,found
}