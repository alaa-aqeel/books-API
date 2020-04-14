package logic

import "../data"
import "../template"

func AddABook(newBook template.Book) bool {
	failed := false
	if newBook.Pages>=0{
		data.AddBook(newBook)
	}else{
		failed = true
	}
	return failed
}
