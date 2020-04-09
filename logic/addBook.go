package logic

import "../books"

func AddABook(newBook books.Book) bool {
	failed := false
	if newBook.Pages>=0{
		books.AddBook(newBook)
	}else{
		failed = true
	}
	return failed
}
