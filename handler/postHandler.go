package handler
import (
	"../books"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func handleFprintf(err error,n int)  {
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("%d bytes written\n",n)
}

func HandlePost(writer http.ResponseWriter,r *http.Request)  {
	var newBook books.Book
	reqBody,err := ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		n,err := fmt.Fprintf(writer,"Error occurred")
		handleFprintf(err,n)
	}else{
		err := json.Unmarshal(reqBody,&newBook)
		if err!=nil{
			writer.WriteHeader(http.StatusBadRequest)
			n,err := fmt.Fprintf(writer,"Please enter information according to format")
			handleFprintf(err,n)
		}else {

			writer.WriteHeader(http.StatusCreated)
			books.AddBook(newBook)
			books.UpdateID()
			err = json.NewEncoder(writer).Encode(books.BOOKS.Books[len(books.BOOKS.Books)-1])
			if err!=nil{
				fmt.Println(err)
			}
		}

	}
}
