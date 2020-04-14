package data

import "fmt"

func DeleteBook(id int64) bool {
	found := false
	for i:=0;i< len(BOOKS.Books);i++{
		if id == BOOKS.Books[i].Id{
			fmt.Println("Book with ID:",BOOKS.Books[i].Id,"deleted")
			BOOKS.Books = append(BOOKS.Books[:i],BOOKS.Books[i+1:]...)
			found = true
			break
		}
	}
	return found
}
